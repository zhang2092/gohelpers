package datetime

import "time"

type theTime struct {
	unix int64
}

// NewUnixNow 返回当前时间的unix时间戳
func NewUnixNow() *theTime {
	return &theTime{unix: time.Now().Unix()}
}

// NewUnix 返回指定时间的unix时间戳
func NewUnix(unix int64) *theTime {
	return &theTime{unix: unix}
}

// NewFormat 返回指定时间字符串的unix时间戳，t应该是 "yyyy-mm-dd hh:mm:ss"
func NewFormat(t string) (*theTime, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc := time.FixedZone("CST", 8*3600)
	tt, err := time.ParseInLocation(timeLayout, t, loc)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: tt.Unix()}, nil
}

// NewISO8601 返回指定的iso8601时间字符串的unix时间戳
func NewISO8601(iso8601 string) (*theTime, error) {
	t, err := time.ParseInLocation(time.RFC3339, iso8601, time.UTC)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: t.Unix()}, nil
}

// ToUnix 返回unix时间戳
func (t *theTime) ToUnix() int64 {
	return t.unix
}

// ToFormat 返回unix时间的时间字符串'yyyy-mm-dd hh:mm:ss'
func (t *theTime) ToFormat() string {
	return time.Unix(t.unix, 0).Format("2006-01-02 15:04:05")
}

// ToFormatForTpl 返回指定格式的时间字符串 tpl
func (t *theTime) ToFormatForTpl(tpl string) string {
	return time.Unix(t.unix, 0).Format(tpl)
}

// ToIso8601 返回 iso8601 时间字符串
func (t *theTime) ToIso8601() string {
	return time.Unix(t.unix, 0).Format(time.RFC3339)
}
