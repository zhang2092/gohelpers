package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomInt(n int) string {
	var letters = []byte("0123456789")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RandomLevel() string {
	levels := []string{"DEBUG", "INFO", "ERROR", "WARNING", "FAIL"}
	n := len(levels)
	return levels[rand.Intn(n)]
}

func GenerateFilename() (string, error) {
	filename, err := gonanoid.Nanoid()
	if err != nil {
		return "", fmt.Errorf("could not generate avatar filename: %v", err)
	}
	return filename, nil
}
