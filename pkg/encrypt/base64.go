package encrypt

import "encoding/base64"

// Base64Encrypt Base64加密
func Base64Encrypt(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decrypt Base64解密
func Base64Decrypt(str string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes), err
}
