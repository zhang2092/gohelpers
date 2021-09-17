package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

var DES_KEY = "fBEznwcv"

// DESEncrypt DES加密  key：8位
func DESEncrypt(origData, key []byte) string {
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//对明文先进行补码操作
	origData = PKCS5Padding(origData, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, key)
	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))
	//加密明文,加密后的数据放到数组中
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// DESDecrypt DES解密  key：8位
func DESDecrypt(data string, key []byte) string {
	//倒叙执行一遍加密方法
	//将字符串转换成字节数组
	crypted, _ := base64.StdEncoding.DecodeString(data)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, key)
	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)
	//去补码
	origData = PKCS5UnPadding(origData)

	return string(origData)
}

// PKCS5Padding 实现明文的补码
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 去除补码
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}

// TripleDESEncrypt 3DES加密 key:24位
func TripleDESEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// TripleDESDecrypt 3DES解密 key:24位
func TripleDESDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
