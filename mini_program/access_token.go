package mini_program

import (
	"encoding/json"
	"fmt"

	"miniprogram/utils"
)

type WechatAccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type MiniProgramAccessToken struct {
	WechatError
	WechatAccessTokenData
}

// GetMiniProgramAccessToken -- get token
func GetMiniProgramAccessToken(config MiniProgramConfig) (string, int64, error) {
	appID := config.AppID()
	appSecret := config.AppSecret()

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
		err := fmt.Errorf("GetMiniProgramSessionInfo error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return "", 0, err
	}

	return result.AccessToken, result.ExpiresIn, nil
}
