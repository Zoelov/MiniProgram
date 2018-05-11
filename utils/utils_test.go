package utils

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	server := fmt.Sprintf("%v:%v", "127.0.0.1", "6379")
	InitRedis(server, "", &CacheStore)

	mc := &CacheStore
	key := "testKey"
	value := 123
	_, err := mc.Set(key, value, 20)

	if err == nil {
		t.Log("set pass")
	} else {
		t.Error("set failed")
		fmt.Println(err)
	}

	v, err := mc.GetString(key)
	if err == nil && v == fmt.Sprintf("%v", value) {
		t.Log("get string pass")
	} else {
		t.Error("get string failed")
	}

	type testStruct struct {
		A int
		B string
	}

	d := new(testStruct)
	d.A = 2
	d.B = "test"

	_, err = mc.CacheStruct("keystruct", d)
	if err == nil {
		t.Log("cache struct pass")
	} else {
		t.Error("cache struct failed")
	}

	var getd testStruct
	err = mc.GetCacheStruct("keystruct", &getd)
	if err == nil && getd.A == 2 && getd.B == "test" {
		t.Log("get struct pass")
	} else {
		t.Error("get struct failed")
	}

}
