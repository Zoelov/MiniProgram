package mini_program

import "testing"

type TestConfig struct {
}

func (tc *TestConfig) AppID() string {
	return "wx45b8b9d700c63c00"
}

func (tc *TestConfig) AppSecret() string {
	return ""
}

var d TestConfig

func TestAccessToken(t *testing.T) {
	_, err := GetMiniProgramAccessToken(&d)
	if err == nil {
		t.Log("Pass")
	} else {
		t.Errorf("failed:%v", err.Error())
	}

}
