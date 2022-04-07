package conver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 1, 64)
}

func SliceIntToString(num []int64) string {
	result := ""

	if len(num) > 0 {
		for _, item := range num {
			result = fmt.Sprintf("%s,%d", result, item)
		}
		if !strings.HasSuffix(result, ",") {
			result = fmt.Sprintf("%s,", result)
		}
	}

	return result
}

func SliceStringToString(num []string) string {
	result := ""

	if len(num) > 0 {
		for _, item := range num {
			result = fmt.Sprintf("%s,%s", result, item)
		}
		if !strings.HasSuffix(result, ",") {
			result = fmt.Sprintf("%s,", result)
		}
	}

	return result
}

func StringToSlice(num string) []string {
	var result []string

	if len(num) > 0 {
		result = strings.Split(num, ",")
	}

	return result
}

func StringToSliceInt(num string) []int64 {
	var result []int64

	if len(num) > 0 {
		array := strings.Split(num, ",")
		if len(array) > 0 {
			for _, item := range array {
				i, err := strconv.Atoi(item)
				if err == nil {
					result = append(result, int64(i))
				}
			}
		}
	}

	return result
}

func MapToJson(param map[string]interface{}) string {
	dataType, err := json.Marshal(param)
	if err != nil {
		return ""
	}

	dataString := string(dataType)
	return dataString
}

func MapSliceStringToJson(param map[string][]string) string {
	dataType, err := json.Marshal(param)
	if err != nil {
		return ""
	}

	dataString := string(dataType)
	return dataString
}
