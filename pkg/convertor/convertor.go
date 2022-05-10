// Package convertor 实现了一些函数来转换数据
package convertor

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ToBool 将字符串转换为布尔值
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToBytes 将接口转换为字节数
func ToBytes(value any) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ToChar 将字符串转换为char slice
func ToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ToString 将值转换为字符串
func ToString(value any) string {
	res := ""
	if value == nil {
		return res
	}

	v := reflect.ValueOf(value)

	switch value.(type) {
	case float32, float64:
		res = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		return res
	case int, int8, int16, int32, int64:
		res = strconv.FormatInt(v.Int(), 10)
		return res
	case uint, uint8, uint16, uint32, uint64:
		res = strconv.FormatUint(v.Uint(), 10)
		return res
	case string:
		res = v.String()
		return res
	case []byte:
		res = string(v.Bytes())
		return res
	default:
		newValue, _ := json.Marshal(value)
		res = string(newValue)
		return res
	}
}

// ToJson 将值转换为有效的json字符串
func ToJson(value any) (string, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// ToFloat 将数值转换为float64,如果输入的不是float,则返回0.0和错误
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	res := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = float64(v.Int())
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = float64(v.Uint())
		return res, nil
	case float32, float64:
		res = v.Float()
		return res, nil
	case string:
		res, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			res = 0.0
		}
		return res, err
	default:
		return res, err
	}
}

// ToInt 将数值转换为int64,如果输入的不是数字格式,则返回0和错误
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var res int64
	err := fmt.Errorf("ToInt: invalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = v.Int()
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = int64(v.Uint())
		return res, nil
	case float32, float64:
		res = int64(v.Float())
		return res, nil
	case string:
		res, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			res = 0
		}
		return res, err
	default:
		return res, err
	}
}

// StructToMap 将结构体转换为Map,只转换导出的结构体字段
// Map key的指定与结构字段标签`json`值相同
func StructToMap(value any) (map[string]any, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", value)
	}

	res := make(map[string]any)

	fieldNum := t.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			//res[name] = v.Field(i).Interface()
			res[tag] = v.Field(i).Interface()
		}
	}

	return res, nil
}

// ColorHexToRGB 将十六进制颜色转换为RGB颜色
func ColorHexToRGB(colorHex string) (red, green, blue int) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF
}

// ColorRGBToHex 将RGB颜色转换为十六进制颜色
func ColorRGBToHex(red, green, blue int) string {
	r := strconv.FormatInt(int64(red), 16)
	g := strconv.FormatInt(int64(green), 16)
	b := strconv.FormatInt(int64(blue), 16)

	if len(r) == 1 {
		r = "0" + r
	}
	if len(g) == 1 {
		g = "0" + g
	}
	if len(b) == 1 {
		b = "0" + b
	}

	return "#" + r + g + b
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
