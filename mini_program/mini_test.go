package mini_program

import "testing"

var appID = "wx4f4bc4dec97d474b"
var appSecret = ""
var sessionKey = "tiihtNczf5v6AKRyjwEUhQ=="
var encryptedData = "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
var iv = "r7BXXKkLb8qrSNn05n0qiA=="

func TestAccessToken(t *testing.T) {
	_, _, err := GetMiniProgramAccessToken(appID, appSecret, false)
	if err == nil {
		t.Log("Pass")
	} else {
		t.Errorf("failed:%v", err.Error())
	}
}

func TestDecryptGroup(t *testing.T) {
	_, err := DecryptedWechatSharedGroup(encryptedData, iv, sessionKey, appID)
	if err == nil {
		t.Log("decrpted pass")
	} else {
		t.Error("decrpted failed")
	}
}
