package mini_program

import (
	"encoding/json"
	"fmt"

	"miniprogram/utils"
)

const (
	cacheKeyForAccessTokenByAppID = "mini:program:appID:%v:access:token:v1"
)

type WechatError struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WechatAccessTokenData -- token json data
type WechatAccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// MiniProgramAccessToken -- with error
type MiniProgramAccessToken struct {
	WechatError
	WechatAccessTokenData
}

// GetMiniProgramAccessToken -- get token and cache
func GetMiniProgramAccessToken(appID, appSecret string, cached bool) (string, int64, error) {
	mc := utils.CacheStore
	key := fmt.Sprintf(cacheKeyForAccessTokenByAppID, appID)
	if cached {
		if token, err := mc.GetString(key); err == nil {
			expired, err := mc.TTL(key)
			if err == nil {
				return token, expired, nil
			}
		}
	}

	url := fmt.Sprintf(accessTokenURL, appID, appSecret)
	var response []byte
	response, err := utils.HTTPGet(url)
	if err != nil {
		return "", 0, err
	}
	result := new(MiniProgramAccessToken)
	err = json.Unmarshal(response, result)
	if err != nil {
		return "", 0, err
	}
	if result.Errcode != 0 {
		err := fmt.Errorf("get access token error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return "", 0, err
	}

	if cached {
		mc.Set(key, result.AccessToken, result.ExpiresIn)
	}

	return result.AccessToken, result.ExpiresIn, nil
}
