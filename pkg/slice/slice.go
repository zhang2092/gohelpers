package slice

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Contains slice中是否包含某个字符串
func Contains(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}

	return false
}

// RemoveDuplicatesFromSlice 去重
func RemoveDuplicatesFromSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// Shuffle 随机打乱
func Shuffle(array []string) []string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}

	return array
}

// ReverseSlice 反转
func ReverseSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}

	return a
}

// ConvertSliceToString 将 slice 转换为逗号分隔的字符串
func ConvertSliceToString(input []int) string {
	var output []string
	for _, i := range input {
		output = append(output, strconv.Itoa(i))
	}

	return strings.Join(output, ",")
}
