package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"time"

	"github.com/garyburd/redigo/redis"
)

var CachePrefix = "minicache-"

var CacheStore = RedisConn{Prefix: CachePrefix}

type RedisConn struct {
	pool              *redis.Pool
	server            string
	max_idle_conn     int
	conn_idle_timeout time.Duration
	Prefix            string
}

func (rc *RedisConn) keyWithPrefix(key string) string {
	buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, rc.Prefix, key)
	return buffer.String()
}

func (rc *RedisConn) do(command string, args ...interface{}) (reply interface{}, err error) {
	conn := rc.pool.Get()
	defer conn.Close()

	reply, err = conn.Do(command, args...)
	if err != nil {
		if err != redis.ErrNil {

		}
	}

	return
}

func (rc *RedisConn) Set(key string, value interface{}, expire int64) (bool, error) {
	var args []interface{}
	if expire == 0 {
		args = append(args, rc.keyWithPrefix(key), value)
	} else {
		args = append(args, rc.keyWithPrefix(key), value, "EX", expire)
	}

	r, err := redis.String(rc.do("SET", args...))
	var ret bool
	if r == "OK" {
		ret = true
	} else {
		ret = false
	}

	return ret, err
}

func (rc *RedisConn) GetString(key string) (string, error) {
	v, err := redis.String(rc.do("GET", rc.keyWithPrefix(key)))
	return v, err
}

func (rc *RedisConn) CacheStruct(key string, dest interface{}) (bool, error) {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	encoder.Encode(dest)

	v, err := redis.String(rc.do("SET", rc.keyWithPrefix(key), buff.String()))
	var ret bool
	if v == "OK" {
		ret = true
	} else {
		ret = false
	}

	return ret, err
}

func (rc *RedisConn) GetCacheStruct(key string, dest interface{}) error {
	v, err := redis.Bytes(rc.do("GET", rc.keyWithPrefix(key)))
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewReader(v))
	if err := decoder.Decode(dest); err != nil {
		return err
	}

	return nil
}

func (rc *RedisConn) Expire(key string, seconds int64) (bool, error) {
	r, err := redis.Bool(rc.do("EXPIRE", rc.keyWithPrefix(key), seconds))
	return r, err
}

func InitRedis(server, password string, storePtr *RedisConn) {
	dialFunc := func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", server)
		return c, err
	}

	pool := &redis.Pool{
		MaxIdle:     int(math.Pow(2, 30)),
		IdleTimeout: 1 * time.Second,
		Dial:        dialFunc,
	}

	storePtr.pool = pool
	storePtr.server = server
	storePtr.max_idle_conn = pool.MaxIdle
	storePtr.conn_idle_timeout = pool.IdleTimeout
}
