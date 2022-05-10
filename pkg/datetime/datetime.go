// Package datetime 实现了一些格式化日期和时间的功能
// Note:
// 1. `format` param in FormatTimeToStr function should be as flow:
//"yyyy-mm-dd hh:mm:ss"
//"yyyy-mm-dd hh:mm"
//"yyyy-mm-dd hh"
//"yyyy-mm-dd"
//"yyyy-mm"
//"mm-dd"
//"dd-mm-yy hh:mm:ss"
//"yyyy/mm/dd hh:mm:ss"
//"yyyy/mm/dd hh:mm"
//"yyyy/mm/dd hh"
//"yyyy/mm/dd"
//"yyyy/mm"
//"mm/dd"
//"dd/mm/yy hh:mm:ss"
//"yyyy"
//"mm"
//"hh:mm:ss"
//"mm:ss"
package datetime

import (
	"fmt"
	"time"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15:04",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyy":                "2006",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"mm:ss":               "04:05",
	}
}

// AddMinute 加或减分钟的时间
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour 加或减小时的时间
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay 加或减天的时间
func AddDay(t time.Time, day int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(day))
}

// GetNowDate 返回当前日期的格式yyy-mm-dd
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime 返回当前时间的格式 hh-mm-ss
func GetNowTime() string {
	return time.Now().Format("15:04:05")
}

// GetNowDateTime 返回当前日期时间的格式yyy-mm-dd hh-mm-ss
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetZeroHourTimestamp 返回零时的时间戳 (时间戳为00:00)
func GetZeroHourTimestamp() int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.UTC().Unix() - 8*3600
}

// GetNightTimestamp 返回零时的时间戳 (时间戳为23:59)
func GetNightTimestamp() int64 {
	return GetZeroHourTimestamp() + 86400 - 1
}

// FormatTimeToStr 将时间转换为字符串
func FormatTimeToStr(t time.Time, format string) string {
	return t.Format(timeFormat[format])
}

// FormatStrToTime 将字符串转换为时间
func FormatStrToTime(str, format string) (time.Time, error) {
	v, ok := timeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not found", format)
	}

	return time.Parse(v, str)
}

// BeginOfMinute 返回起始分钟的时间
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute 返回结束时的时间
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour 返回开始时间 一天中的时间
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour 返回结束时间 一天中的时间
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay 返回开始时间 一天中的时间
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay 一天中的返回结束时间
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek 回归周，从周日开始的一周
func BeginOfWeek(t time.Time) time.Time {
	y, m, d := t.AddDate(0, 0, 0-int(BeginOfDay(t).Weekday())).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfWeek 周末返程时间，周末有周六
func EndOfWeek(t time.Time) time.Time {
	y, m, d := BeginOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond).Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfMonth 月初返回
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 月末返回
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear 年初回报
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 年终回报
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
