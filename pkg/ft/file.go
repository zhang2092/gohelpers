package ft

import "os"

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if notExist := Exists(src); !notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
