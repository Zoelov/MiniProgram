package mini_program

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// VerifySignature -- 小程序签名验证
func VerifySignature(signature, rawData, sessionKey string) bool {
	content := fmt.Sprintf("%v%v", rawData, sessionKey)

	h := sha1.New()
	h.Write([]byte(content))

	compareSignature := hex.EncodeToString(h.Sum(nil))

	return signature == compareSignature
}
