package mini_program

import (
	"encoding/json"
	"fmt"
	"strings"

	"miniprogram/utils"
)

type WaterMark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type WechatSharedGroup struct {
	WaterMark WaterMark `json:"watermark"`
	OpenGID   string    `json:"openGId"`
}

func DecryptedWechatSharedGroup(encryptedData, iv, sessionKey, appID string) (*WechatSharedGroup, error) {

	aesPlantText, err := utils.Decrypt(encryptedData, sessionKey, iv, true)
	if err != nil {
		return nil, fmt.Errorf("err occurs while decrypt data: %v", err)
	}

	groupInfo := new(WechatSharedGroup)
	aesPlantText = strings.Replace(aesPlantText, "\a", "", -1)
	err = json.Unmarshal([]byte(aesPlantText), groupInfo)
	if err != nil {
		return nil, fmt.Errorf("err occurs while unmarshal decypted data to group info: %v", err)
	}

	if groupInfo.WaterMark.AppID != appID {
		return nil, fmt.Errorf("appId in decrypted data does not match, appID:%v", appID)
	}

	return groupInfo, nil
}
