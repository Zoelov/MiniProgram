package mini_program

type MiniProgramConfig interface {
	AppID() string
	AppSecret() string
}

type WechatError struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
