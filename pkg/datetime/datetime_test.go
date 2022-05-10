package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddDay(t *testing.T) {
	now := time.Now()
	after2Days := AddDay(now, 2)
	diff1 := after2Days.Sub(now)
	require.Equal(t, float64(48), diff1.Hours())

	before2Days := AddDay(now, -2)
	diff2 := before2Days.Sub(now)
	require.Equal(t, float64(-48), diff2.Hours())
}

func TestAddHour(t *testing.T) {
	now := time.Now()
	after2Hours := AddHour(now, 2)
	diff1 := after2Hours.Sub(now)
	require.Equal(t, float64(2), diff1.Hours())

	before2Hours := AddHour(now, -2)
	diff2 := before2Hours.Sub(now)
	require.Equal(t, float64(-2), diff2.Hours())
}

func TestAddMinute(t *testing.T) {
	now := time.Now()
	after2Minutes := AddMinute(now, 2)
	diff1 := after2Minutes.Sub(now)
	require.Equal(t, float64(2), diff1.Minutes())

	before2Minutes := AddMinute(now, -2)
	diff2 := before2Minutes.Sub(now)
	require.Equal(t, float64(-2), diff2.Minutes())
}

func TestGetNowDate(t *testing.T) {
	expected := time.Now().Format("2006-01-02")
	require.Equal(t, expected, GetNowDate())
}

func TestGetNotTime(t *testing.T) {
	expected := time.Now().Format("15:04:05")
	require.Equal(t, expected, GetNowTime())
}

func TestGetNowDateTime(t *testing.T) {
	expected := time.Now().Format("2006-01-02 15:04:05")
	require.Equal(t, expected, GetNowDateTime())
}

func TestFormatTimeToStr(t *testing.T) {
	datetime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"hh:mm:ss", "yyyy/mm"}

	expected := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"16:04:08", "2021/01"}

	for i := 0; i < len(cases); i++ {
		actual := FormatTimeToStr(datetime, cases[i])
		require.Equal(t, expected[i], actual)

	}
}

func TestFormatStrToTime(t *testing.T) {
	formats := []string{
		"2006-01-02 15:04:05", "2006-01-02",
		"02-01-06 15:04:05", "2006/01/02 15:04:05",
		"2006/01"}
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"yyyy/mm"}

	datetimeStr := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"2021/01"}

	for i := 0; i < len(cases); i++ {
		actual, err := FormatStrToTime(datetimeStr[i], cases[i])
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := time.Parse(formats[i], datetimeStr[i])
		require.Equal(t, expected, actual)
	}
}

func TestBeginOfMinute(t *testing.T) {
	expected := time.Date(2022, 2, 15, 15, 48, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMinute(td)

	require.Equal(t, expected, actual)
}

func TestEndOfMinute(t *testing.T) {
	expected := time.Date(2022, 2, 15, 15, 48, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMinute(td)

	require.Equal(t, expected, actual)
}

func TestBeginOfHour(t *testing.T) {
	expected := time.Date(2022, 2, 15, 15, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfHour(td)

	require.Equal(t, expected, actual)
}

func TestEndOfHour(t *testing.T) {
	expected := time.Date(2022, 2, 15, 15, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfHour(td)

	require.Equal(t, expected, actual)
}

func TestBeginOfDay(t *testing.T) {
	expected := time.Date(2022, 2, 15, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfDay(td)

	require.Equal(t, expected, actual)
}

func TestEndOfDay(t *testing.T) {
	expected := time.Date(2022, 2, 15, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfDay(td)

	require.Equal(t, expected, actual)
}

func TestBeginOfWeek(t *testing.T) {
	expected := time.Date(2022, 2, 13, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfWeek(td)

	require.Equal(t, expected, actual)
}

func TestEndOfWeek(t *testing.T) {
	expected := time.Date(2022, 2, 19, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfWeek(td)

	require.Equal(t, expected, actual)
}

func TestBeginOfMonth(t *testing.T) {
	expected := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMonth(td)

	require.Equal(t, expected, actual)
}

func TestEndOfMonth(t *testing.T) {
	expected := time.Date(2022, 2, 28, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMonth(td)

	require.Equal(t, expected, actual)
}

func TestBeginOfYear(t *testing.T) {
	expected := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfYear(td)

	require.Equal(t, expected, actual)
}

func TestEndOfYear(t *testing.T) {
	expected := time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfYear(td)

	require.Equal(t, expected, actual)
}
