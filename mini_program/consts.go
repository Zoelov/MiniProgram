package mini_program

const (
	sessionURL      = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
	accessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
	qrCodeURL       = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%v"
	sendTemplateMsg = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s"
)
