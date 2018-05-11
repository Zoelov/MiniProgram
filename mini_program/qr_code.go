package mini_program

import (
	"fmt"

	"miniprogram/utils"
)

type Color struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
}

type MiniProgramQRCodeParam struct {
	Scene     string `json:"scene"`
	Page      string `json:"page"`
	Width     int    `json:"width"`
	AutoColor bool   `json:"auto_color"`
	LineColor Color  `json:"line_color"`
}

// GetMiniProgramQRCode -- 获取永久小程序码
func GetMiniProgramQRCode(param *MiniProgramQRCodeParam, appID, appSecret string) ([]byte, error) {
	token, _, err := GetMiniProgramAccessToken(appID, appSecret, true)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(qrCodeURL, token)
	var response []byte
	response, err = utils.PostJSON(url, param)
	if err != nil {
		return nil, err
	}

	return response, nil
}
