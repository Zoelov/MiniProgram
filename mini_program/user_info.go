package mini_program

import (
	"encoding/json"
	"fmt"
	"strings"

	"miniprogram/utils"
)

type WechatUserInfo struct {
	WaterMark WaterMark `json:"watermark"`
	OpenID    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int32     `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
}

func DecryptedWechatUserInfo(encryptedData, iv, sessionKey, appID string) (*WechatUserInfo, error) {

	aesPlantText, err := utils.Decrypt(encryptedData, sessionKey, iv, true)
	if err != nil {
		return nil, fmt.Errorf("err occurs while decrypt data: %v", err)
	}

	info := new(WechatUserInfo)
	aesPlantText = strings.Replace(aesPlantText, "\a", "", -1)
	err = json.Unmarshal([]byte(aesPlantText), info)
	if err != nil {
		return nil, fmt.Errorf("err occurs while unmarshal decypted data to group info: %v", err)
	}

	if info.WaterMark.AppID != appID {
		return nil, fmt.Errorf("appId in decrypted data does not match, appID:%v", appID)
	}

	return info, nil
}
