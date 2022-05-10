package fileutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsExist(t *testing.T) {
	cases := []string{"./", "./file.go", "./a.txt"}
	expected := []bool{true, true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsExist(cases[i])
		require.Equal(t, expected[i], actual)
	}
}

func TestCreateFile(t *testing.T) {
	f := "./text.txt"
	if CreateFile(f) {
		file, err := os.Open(f)
		require.Nil(t, err)
		require.Equal(t, f, file.Name())
	} else {
		t.FailNow()
	}
	os.Remove(f)
}

func TestIsDir(t *testing.T) {
	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsDir(cases[i])
		require.Equal(t, expected[i], actual)
	}
}

func TestRemoveFile(t *testing.T) {
	f := "./text.txt"
	if !IsExist(f) {
		CreateFile(f)
		err := RemoveFile(f)
		require.Nil(t, err)
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	destFile := "./text_copy.txt"

	err := CopyFile(srcFile, destFile)
	if err != nil {
		file, err := os.Open(destFile)
		require.Nil(t, err)
		require.Equal(t, destFile, file.Name())
	}
	os.Remove(srcFile)
	os.Remove(destFile)
}

func TestListFileNames(t *testing.T) {
	filesInPath, err := ListFileNames("./")
	require.Nil(t, err)

	expected := []string{"file.go", "file_test.go"}
	require.Equal(t, expected, filesInPath)
}

func TestReadFileToString(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	f.WriteString("hello world")

	content, _ := ReadFileToString(path)
	require.Equal(t, "hello world", content)

	os.Remove(path)
}

func TestClearFile(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	f.WriteString("hello world")

	err := ClearFile(path)
	require.Nil(t, err)

	content, _ := ReadFileToString(path)
	require.Equal(t, "", content)

	os.Remove(path)
}

func TestReadFileByLine(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()
	f.WriteString("hello\nworld")

	expected := []string{"hello", "world"}
	actual, _ := ReadFileByLine(path)
	require.Equal(t, expected, actual)

	os.Remove(path)
}

func TestZipAndUnZip(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	file, _ := os.OpenFile(srcFile, os.O_WRONLY|os.O_TRUNC, 0777)
	defer file.Close()
	file.WriteString("hello\nworld")

	zipFile := "./text.zip"
	err := Zip(srcFile, zipFile)
	require.Nil(t, err)

	unZipPath := "./unzip"
	err = UnZip(zipFile, unZipPath)
	require.Nil(t, err)

	unZipFile := "./unzip/text.txt"
	require.Equal(t, true, IsExist(unZipFile))

	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
}

func TestFileMode(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	mode, err := FileMode(srcFile)
	require.Nil(t, err)

	t.Log(mode)

	os.Remove(srcFile)
}

func TestIsLink(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	linkFile := "./text.link"
	if !IsExist(linkFile) {
		_ = os.Symlink(srcFile, linkFile)
	}
	require.Equal(t, true, IsLink(linkFile))

	require.Equal(t, false, IsLink("./file.go"))

	os.Remove(srcFile)
	os.Remove(linkFile)
}

func TestMiMeType(t *testing.T) {
	f, _ := os.Open("./file.go")
	require.Equal(t, "text/plain; charset=utf-8", MiMeType(f))
	require.Equal(t, "text/plain; charset=utf-8", MiMeType("./file.go"))
}
