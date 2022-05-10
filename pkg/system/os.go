// Package system 包含一些关于操作系统、运行时间、shell命令的功能
package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

// IsWindows 检查当前操作系统是否为Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux 检查当前操作系统是否为Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac 检查当前操作系统是否为Macos
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv 获取由键命名的环境变量的值
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv 设置由键命名的环境变量的值
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv 删除一个环境变量
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv 获取由键值命名的环境,并将其与compareEnv进行比较
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand 使用shell /bin/bash -c来执行命令
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = string(errout.Bytes())
	}
	stdout = string(out.Bytes())

	return
}
