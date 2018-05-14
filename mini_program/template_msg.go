package mini_program

import (
	"encoding/json"
	"fmt"

	"miniprogram/utils"
)

type TemplateKeyWord struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateData struct {
	Keyword1  TemplateKeyWord `json:"keyword1"`
	Keyword2  TemplateKeyWord `json:"keyword2"`
	Keyword3  TemplateKeyWord `json:"keyword3"`
	Keyword4  TemplateKeyWord `json:"keyword4"`
	Keyword5  TemplateKeyWord `json:"keyword5"`
	Keyword6  TemplateKeyWord `json:"keyword6"`
	Keyword7  TemplateKeyWord `json:"keyword7"`
	Keyword8  TemplateKeyWord `json:"keyword8"`
	Keyword9  TemplateKeyWord `json:"keyword9"`
	Keyword10 TemplateKeyWord `json:"keyword10"`
}

type SendTemplateMsgParam struct {
	ToUser     string       `json:"touser"`
	TemplateID string       `json:"template_id"`
	Page       string       `json:"page"`
	FormID     string       `json:"form_id"`
	Data       TemplateData `json:"data"`
}

type SendTemplateMsgRet struct {
	WechatError
}

func SendTemplateMsg(appID, appSecret string, param *SendTemplateMsgParam) (bool, error) {
	token, _, err := GetMiniProgramAccessToken(appID, appSecret, true)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf(sendTemplateMsg, token)
	var response []byte
	response, err = utils.PostJSON(url, param)
	if err != nil {
		return false, err
	}

	result := new(SendTemplateMsgRet)
	err = json.Unmarshal(response, result)
	if err != nil {
		return false, err
	}

	if result.Errcode != 0 {
		err := fmt.Errorf("send template msg error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return false, err
	}

	return true, nil
}
