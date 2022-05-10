package strutil

import "strings"

// splitWordsToLower 将一个字符串按大写字母分割成若干个字符串
func splitWordsToLower(s string) []string {
	var res []string

	upperIndexes := upperIndex(s)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if s != "" {
			res = append(res, s)
		}
		return res
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			res = append(res, strings.ToLower(s[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			res = append(res, strings.ToLower(s[upperIndexes[i]:]))
		}
	}
	return res
}

// upperIndex 获得一个int slice,其元素是一个字符串的所有大写字母索引
func upperIndex(s string) []int {
	var res []int
	for i := 0; i < len(s); i++ {
		if 64 < s[i] && s[i] < 91 {
			res = append(res, i)
		}
	}
	if len(s) > 0 && res != nil && res[0] != 0 {
		res = append([]int{0}, res...)
	}

	return res
}
