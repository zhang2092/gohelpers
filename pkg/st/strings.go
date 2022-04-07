package st

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

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

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
