package encrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

// Sha1 加密
func Sha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	return fmt.Sprintf("%x", has)
}

// Sha256 加密
func Sha256(str string) string {
	w := sha256.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	return hex.EncodeToString(bw)
}

// Sha512 加密
func Sha512(str string) string {
	w := sha512.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	return hex.EncodeToString(bw)
}
