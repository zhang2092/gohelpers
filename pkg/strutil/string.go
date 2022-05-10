// Package strutil 实现了一些函数来操作字符串
package strutil

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase 转换字符串到驼峰法(CamelCase)
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			res += string(vv)
		} else {
			res += Capitalize(v)
		}
	}

	return res
}

// Capitalize 将一个字符串的第一个字符转换为大写,其余的转换为小写
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	out := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			out[i] = unicode.ToUpper(v)
		} else {
			out[i] = unicode.ToLower(v)
		}
	}

	return string(out)
}

// UpperFirst 将字符串的第一个字符转换为大写
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst 将字符串的第一个字符转换为小写
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// PadEnd 如果字符串比尺寸短,则将其垫在右侧
// 填充字符如果超过大小,将被截断
func PadEnd(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return source + fill[0:size-len1]
}

// PadStart 如果字符串比尺寸短,则将其垫在左侧
//填充字符如果超过大小,将被截断
func PadStart(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return fill[0:size-len1] + source
}

// KebabCase 将字符串转为短横线隔开式(kebab-case)
func KebabCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var res []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			res = append(res, splitWords...)
		}
	}

	return strings.Join(res, "-")
}

// SnakeCase 将字符串转为蛇形命名(snake_case)
func SnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var res []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			res = append(res, splitWords...)
		}
	}

	return strings.Join(res, "_")
}

// Before 在字符首次出现的位置之前,在源字符串中创建子串
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast 在字符最后出现的位置之前,在源字符串中创建子串
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After 在字符首次出现的位置后,在源字符串中创建子串
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast 在字符最后出现的位置后,在源字符串中创建子串
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString 检查值的数据类型是否为字符串
func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// ReverseStr 返回字符顺序与给定字符串相反的字符串
func ReverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap 用另一个字符串包住一个字符串
func Wrap(str string, wrapWith string) string {
	if str == "" || wrapWith == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapWith)
	sb.WriteString(str)
	sb.WriteString(wrapWith)

	return sb.String()
}

// Unwrap 从另一个字符串中解开一个给定的字符串,将改变str值
func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}

	return str
}

// RemoveHTML 去除字符串中的 html, js
func RemoveHTML(str string) string {
	if len(str) > 0 {
		//删除脚本
		reg := regexp.MustCompile(`([\r\n])[\s]+`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`<script[^>]*?>.*?</script>`)
		str = reg.ReplaceAllString(str, "")
		//删除HTML
		reg = regexp.MustCompile(`<(.[^>]*)>`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`([\r\n])[\s]+`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`-->`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`<!--.*`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(quot|#34);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(amp|#38);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(lt|#60);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(gt|#62);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(nbsp|#160);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(iexcl|#161);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(cent|#162);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(pound|#163);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&(copy|#169);`)
		str = reg.ReplaceAllString(str, "")
		reg = regexp.MustCompile(`&#(\d+);`)
		str = reg.ReplaceAllString(str, "")

		str = strings.ReplaceAll(str, "<", "")
		str = strings.ReplaceAll(str, ">", "")
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, " ", "")
		str = strings.ReplaceAll(str, "　", "")

		return str
	}
	return ""
}
