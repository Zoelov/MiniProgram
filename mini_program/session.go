// mini program code to session
// cache session in redis
// return a uuid as key instead of session info

package mini_program

import (
	"encoding/json"
	"fmt"

	"miniprogram/utils"

	uuid "github.com/satori/go.uuid"
)

const (
	cacheKeyForSession = "mini:program:sessioin:key:v2:%v"
)

type WechatUserSessionData struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
type MiniProgramSession struct {
	WechatError
	WechatUserSessionData
}

func GetMiniProgramSessionInfo(code, appID, appSecret string) (*MiniProgramSession, error) {
	url := fmt.Sprintf(sessionURL, appID, appSecret, code)
	var response []byte
	response, err := utils.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	result := new(MiniProgramSession)
	err = json.Unmarshal(response, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		err := fmt.Errorf("GetMiniProgramSessionInfo error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return nil, err
	}

	return result, nil
}

func CacheWechatUserSessionData(userSession WechatUserSessionData) (string, error) {
	ns := uuid.NewV1()
	key := uuid.NewV3(ns, userSession.OpenID).String()
	cacheKey := fmt.Sprintf(cacheKeyForSession, key)
	mc := utils.CacheStore

	_, err := mc.CacheStruct(cacheKey, userSession)
	if err != nil {
		return "", err
	}
	_, err = mc.Expire(cacheKey, utils.ONE_WEEK)
	if err != nil {
		return "", err
	}

	return key, nil
}

func DeleteWechatUserSession(key string) error {
	mc := utils.CacheStore
	cacheKey := fmt.Sprintf(cacheKeyForSession, key)
	_, err := mc.Delete(cacheKey)
	return err
}

func GetWechatUserSessionData(key string) (*WechatUserSessionData, error) {
	mc := utils.CacheStore
	cacheKey := fmt.Sprintf(cacheKeyForSession, key)
	ret := new(WechatUserSessionData)
	err := mc.GetCacheStruct(cacheKey, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
