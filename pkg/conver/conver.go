package conver

import "strconv"

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 1, 64)
}
