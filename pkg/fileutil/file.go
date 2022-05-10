// Package fileutil 实现了一些文件操作的基本功能
package fileutil

import (
	"archive/zip"
	"bufio"
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// IsExist 检查一个文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateFile 在路径中创建一个文件
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// IsDir 检查路径是否为目录
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// RemoveFile 删除路径中的文件
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile 复制src文件到dest文件
func CopyFile(srcFilePath string, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//ClearFile 在路径文件中写入空字符串
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

//ReadFileToString 返回文件内容的字符串
func ReadFileToString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine 逐行读取文件
func ReadFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		res = append(res, l)
	}

	return res, nil
}

// ListFileNames 返回路径中的所有文件名
func ListFileNames(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	res := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			res = append(res, fs[i].Name())
		}
	}

	return res, nil
}

// Zip 创建压缩文件, 参数`filePath`可以是一个单独的文件或一个目录
func Zip(filePath string, destPath string) error {
	zipFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(filePath)+"/")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

// UnZip 解压缩文件并将其保存到destPath
func UnZip(zipFile string, destPath string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		path := filepath.Join(destPath, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// IsLink 检查一个文件是否是符号链接
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// FileMode 返回文件的模式和权限
func FileMode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// MiMeType 返回文件的mime类型
// 参数 `file`应该是字符串（文件路径）或*os.File
func MiMeType(file any) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}
	return mediatype
}

// IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if notExist := IsExist(src); !notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
