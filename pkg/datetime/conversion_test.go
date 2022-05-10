package datetime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToUnix(t *testing.T) {
	tm1 := NewUnixNow()
	unixTimestamp := tm1.ToUnix()
	tm2 := NewUnix(unixTimestamp)

	require.Equal(t, tm1, tm2)
}

func TestToFormat(t *testing.T) {
	_, err := NewFormat("2022/03/18 17:04:05")
	require.NotNil(t, err)

	tm, err := NewFormat("2022-03-18 17:04:05")
	require.Nil(t, err)

	t.Log("ToFormat -> ", tm.ToFormat())
}

func TestToFormatForTpl(t *testing.T) {
	_, err := NewFormat("2022/03/18 17:04:05")
	require.NotNil(t, err)

	tm, err := NewFormat("2022-03-18 17:04:05")
	require.Nil(t, err)

	t.Log("ToFormatForTpl -> ", tm.ToFormatForTpl("2006/01/02 15:04:05"))
}

func TestToIso8601(t *testing.T) {
	_, err := NewISO8601("2022-03-18 17:04:05")
	require.NotNil(t, err)

	tm, err := NewISO8601("2006-01-02T15:04:05.999Z")
	require.Nil(t, err)

	t.Log("ToIso8601 -> ", tm.ToIso8601())
}
