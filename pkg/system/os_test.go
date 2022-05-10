package system

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOsDetection(t *testing.T) {
	osType, _, _ := ExecCommand("echo $OSTYPE")
	if strings.Index(osType, "linux") != -1 {
		require.Equal(t, true, IsLinux())
	}
	if strings.Index(osType, "darwin") != -1 {
		require.Equal(t, true, IsMac())
	}
}

func TestOsEnvOperation(t *testing.T) {
	envNotExist := GetOsEnv("foo")
	require.Equal(t, "", envNotExist)

	SetOsEnv("foo", "foo_value")
	envExist := GetOsEnv("foo")
	require.Equal(t, "foo_value", envExist)

	require.Equal(t, true, CompareOsEnv("foo", "foo_value"))
	require.Equal(t, false, CompareOsEnv("foo", "abc"))
	require.Equal(t, false, CompareOsEnv("abc", "abc"))
	require.Equal(t, false, CompareOsEnv("abc", "abc"))

	err := RemoveOsEnv("foo")
	if err != nil {
		t.Fail()
	}
	require.Equal(t, false, CompareOsEnv("foo", "foo_value"))
}

func TestExecCommand(t *testing.T) {
	out, errout, err := ExecCommand("ls")
	t.Log("std out: ", out)
	t.Log("std err: ", errout)
	require.Nil(t, err)

	out, errout, err = ExecCommand("abc")
	t.Log("std out: ", out)
	t.Log("std err: ", errout)
	if err != nil {
		t.Logf("error: %v\n", err)
	}

	if !IsWindows() {
		require.NotNil(t, err)
	}
}
