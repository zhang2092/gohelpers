package strutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCamelCase(t *testing.T) {
	require.Equal(t, "fooBar", CamelCase("foo_bar"))
	require.Equal(t, "fooBar", CamelCase("Foo-Bar"))
	require.Equal(t, "fooBar", CamelCase("Foo&bar"))
	require.Equal(t, "fooBar", CamelCase("foo bar"))

	require.NotEqual(t, "FooBar", CamelCase("foo_bar"))
}

func TestCapitalize(t *testing.T) {
	require.Equal(t, "Foo", Capitalize("foo"))
	require.Equal(t, "Foo", Capitalize("Foo"))
	require.Equal(t, "Foo", Capitalize("Foo"))

	require.NotEqual(t, "foo", Capitalize("Foo"))
}

func TestKebabCase(t *testing.T) {
	require.Equal(t, "foo-bar", KebabCase("Foo Bar-"))
	require.Equal(t, "foo-bar", KebabCase("foo_Bar"))
	require.Equal(t, "foo-bar", KebabCase("fooBar"))
	require.Equal(t, "f-o-o-b-a-r", KebabCase("__FOO_BAR__"))

	require.NotEqual(t, "foo_bar", KebabCase("fooBar"))
}

func TestSnakeCase(t *testing.T) {
	require.Equal(t, "foo_bar", SnakeCase("Foo Bar-"))
	require.Equal(t, "foo_bar", SnakeCase("foo_Bar"))
	require.Equal(t, "foo_bar", SnakeCase("fooBar"))
	require.Equal(t, "f_o_o_b_a_r", SnakeCase("__FOO_BAR__"))
	require.Equal(t, "a_bbc_s_a_b_b_c", SnakeCase("aBbc-s$@a&%_B.B^C"))

	require.NotEqual(t, "foo-bar", SnakeCase("foo_Bar"))
}

func TestUpperFirst(t *testing.T) {
	require.Equal(t, "Foo", UpperFirst("foo"))
	require.Equal(t, "BAR", UpperFirst("bAR"))
	require.Equal(t, "FOo", UpperFirst("FOo"))
	require.Equal(t, "FOo大", UpperFirst("fOo大"))

	require.NotEqual(t, "Bar", UpperFirst("BAR"))
}

func TestLowerFirst(t *testing.T) {
	require.Equal(t, "foo", LowerFirst("foo"))
	require.Equal(t, "bAR", LowerFirst("BAR"))
	require.Equal(t, "fOo", LowerFirst("FOo"))
	require.Equal(t, "fOo大", LowerFirst("FOo大"))

	require.NotEqual(t, "Bar", LowerFirst("BAR"))
}

func TestPadEnd(t *testing.T) {
	require.Equal(t, "a", PadEnd("a", 1, "b"))
	require.Equal(t, "ab", PadEnd("a", 2, "b"))
	require.Equal(t, "abcdmn", PadEnd("abcd", 6, "mno"))
	require.Equal(t, "abcdmm", PadEnd("abcd", 6, "m"))
	require.Equal(t, "abcaba", PadEnd("abc", 6, "ab"))

	require.NotEqual(t, "ba", PadEnd("a", 2, "b"))
}

func TestPadStart(t *testing.T) {
	require.Equal(t, "a", PadStart("a", 1, "b"))
	require.Equal(t, "ba", PadStart("a", 2, "b"))
	require.Equal(t, "mnabcd", PadStart("abcd", 6, "mno"))
	require.Equal(t, "mmabcd", PadStart("abcd", 6, "m"))
	require.Equal(t, "abaabc", PadStart("abc", 6, "ab"))

	require.NotEqual(t, "ab", PadStart("a", 2, "b"))
}

func TestBefore(t *testing.T) {
	require.Equal(t, "lancet", Before("lancet", ""))
	require.Equal(t, "github.com", Before("github.com/test/lancet", "/"))
	require.Equal(t, "github.com/", Before("github.com/test/lancet", "test"))
}

func TestBeforeLast(t *testing.T) {
	require.Equal(t, "lancet", BeforeLast("lancet", ""))
	require.Equal(t, "github.com/test", BeforeLast("github.com/test/lancet", "/"))
	require.Equal(t, "github.com/test/", BeforeLast("github.com/test/test/lancet", "test"))

	require.NotEqual(t, "github.com/", BeforeLast("github.com/test/test/lancet", "test"))
}

func TestAfter(t *testing.T) {
	require.Equal(t, "lancet", After("lancet", ""))
	require.Equal(t, "test/lancet", After("github.com/test/lancet", "/"))
	require.Equal(t, "/lancet", After("github.com/test/lancet", "test"))
}

func TestAfterLast(t *testing.T) {
	require.Equal(t, "lancet", AfterLast("lancet", ""))
	require.Equal(t, "lancet", AfterLast("github.com/test/lancet", "/"))
	require.Equal(t, "/lancet", AfterLast("github.com/test/lancet", "test"))
	require.Equal(t, "/lancet", AfterLast("github.com/test/test/lancet", "test"))

	require.NotEqual(t, "/test/lancet", AfterLast("github.com/test/test/lancet", "test"))
}

func TestIsString(t *testing.T) {
	require.Equal(t, true, IsString("lancet"))
	require.Equal(t, true, IsString(""))
	require.Equal(t, false, IsString(1))
	require.Equal(t, false, IsString(true))
	require.Equal(t, false, IsString([]string{}))
}

func TestReverseStr(t *testing.T) {
	require.Equal(t, "cba", ReverseStr("abc"))
	require.Equal(t, "54321", ReverseStr("12345"))
}

func TestWrap(t *testing.T) {
	require.Equal(t, "ab", Wrap("ab", ""))
	require.Equal(t, "", Wrap("", "*"))
	require.Equal(t, "*ab*", Wrap("ab", "*"))
	require.Equal(t, "\"ab\"", Wrap("ab", "\""))
	require.Equal(t, "'ab'", Wrap("ab", "'"))
}

func TestUnwrap(t *testing.T) {
	require.Equal(t, "", Unwrap("", "*"))
	require.Equal(t, "ab", Unwrap("ab", ""))
	require.Equal(t, "ab", Unwrap("ab", "*"))
	require.Equal(t, "*ab*", Unwrap("**ab**", "*"))
	require.Equal(t, "ab", Unwrap("**ab**", "**"))
	require.Equal(t, "ab", Unwrap("\"ab\"", "\""))
	require.Equal(t, "*ab", Unwrap("*ab", "*"))
	require.Equal(t, "ab*", Unwrap("ab*", "*"))
	require.Equal(t, "*", Unwrap("***", "*"))

	require.Equal(t, "", Unwrap("**", "*"))
	require.Equal(t, "***", Unwrap("***", "**"))
	require.Equal(t, "**", Unwrap("**", "**"))
}
