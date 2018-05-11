package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Decrypt -- 小程序解密
func Decrypt(toBeDecrypted, key, iv string, dataEncoded bool) (string, error) {
	var err error
	var cipherContent, aesKey, aesIV []byte

	if key == "" {
		return "", fmt.Errorf("no key given")
	}

	if toBeDecrypted == "" {
		return "", fmt.Errorf("nothing to be decrypted")
	}

	var decoder func(string) ([]byte, error)
	decoder = base64.StdEncoding.DecodeString

	aesKey, err = decoder(key)
	if err != nil {
		return "", err
	}

	if dataEncoded {
		cipherContent, err = decoder(toBeDecrypted)
		if err != nil {
			return "", err
		}
	} else {
		cipherContent = []byte(toBeDecrypted)
	}

	if len(cipherContent) < aes.BlockSize {
		return "", fmt.Errorf("ciphercontent too short")
	}

	// 获取初始向量 IV，情况有
	// 1. 若参数中 iv 不为 “”, 则使用 参数中 iv
	// 2. 若参数中 iv 为 “”, 则使用 cipherContent 的第一个 block 为 iv
	if iv == "" {
		aesIV = cipherContent[:aes.BlockSize]
		cipherContent = cipherContent[aes.BlockSize:]
	} else {
		aesIV, err = decoder(iv)
		if err != nil {
			return "", fmt.Errorf("could not decode iv via encode")
		}
	}

	if len(cipherContent)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphercontent is not a multiple of the block size")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	cbcDecrypter := cipher.NewCBCDecrypter(block, aesIV)
	plainContent := cipherContent
	cbcDecrypter.CryptBlocks(cipherContent, plainContent)

	var unpaddinger func([]byte) []byte
	unpaddinger = PKCS7UnPadding
	plainContent = unpaddinger(plainContent)

	return string(plainContent), nil

}

// PKCS7UnPadding return unpadding []Byte plantText
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	if unPadding < 1 || unPadding > 32 {
		unPadding = 0
	}
	return plantText[:(length - unPadding)]
}
