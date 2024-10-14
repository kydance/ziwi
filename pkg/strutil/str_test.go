package strutil

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"GoLang", "Golang"},
		{"", ""},
		{"h", "H"},
		{"H", "H"},
		{"123", "123"},
		{"!@#", "!@#"},
		{"你好", "你好"},
		{"你好世界", "你好世界"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"hello", "hello"},
		{"Hello", "hello"},
		{"123", "123"},
		{"!@#", ""},
		{"", ""},
		{"foobar", "foobar"},
		{"&FOO:BAR$BAZ", "fooBarBaz"},
		{"fooBar", "fooBar"},
		{"FOObar", "foObar"},
		{"$foo%", "foo"},
		{"   $#$Foo   22    bar   ", "foo22Bar"},
		{"Foo-#1😄$_%^&*(1bar", "foo11Bar"},
	}

	for _, test := range tests {
		result := CamelCase(test.input)
		if result != test.expected {
			t.Errorf("CamelCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"GoLang", "GoLang"},
		{"", ""},
		{"h", "H"},
		{"H", "H"},
		{"123", "123"},
		{"!@#", "!@#"},
		{"你好", "你好"},
		{"你好世界", "你好世界"},
		{"你好，世界！", "你好，世界！"},
		{"中文编程", "中文编程"},
	}

	for _, test := range tests {
		result := UpperFirst(test.input)
		if result != test.expected {
			t.Errorf("UpperFirst(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"World", "world"},
		{"GoLang", "goLang"},
		{"", ""},
		{"h", "h"},
		{"H", "h"},
		{"123", "123"},
		{"!@#", "!@#"},
		{"你好", "你好"},
		{"你好世界", "你好世界"},
		{"你好，世界！", "你好，世界！"},
		{"中文编程", "中文编程"},
		{"HELLO WORLD", "hELLO WORLD"},
	}

	for _, test := range tests {
		result := LowerFirst(test.input)
		if result != test.expected {
			t.Errorf("LowerFirst(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestPad(t *testing.T) {
	tests := []struct {
		src    string
		size   int
		pad    string
		result string
	}{
		{"hello", 10, "!", "!!hello!!!"},
		{"world", 5, "*", "world"},
		{"Go", 4, "-", "-Go-"},
		{"", 3, "#", "###"},
		{"你好", 8, "$", "$你好$"},
	}

	for _, test := range tests {
		result := Pad(test.src, test.size, test.pad)
		if result != test.result {
			t.Errorf(
				"Pad(%q, %d, %q) = %q; expected %q",
				test.src,
				test.size,
				test.pad,
				result,
				test.result,
			)
		}
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		src    string
		size   int
		pad    string
		result string
	}{
		{"hello", 10, "!", "!!!!!hello"},
		{"world", 5, "*", "world"},
		{"Go", 4, "-", "--Go"},
		{"", 3, "#", "###"},
		{"你好", 8, "@", "@@你好"},
	}

	for _, test := range tests {
		result := PadLeft(test.src, test.size, test.pad)
		if result != test.result {
			t.Errorf(
				"PadLeft(%q, %d, %q) = %q; expected %.sql",
				test.src,
				test.size,
				test.pad,
				result,
				test.result,
			)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		src    string
		size   int
		pad    string
		result string
	}{
		{"hello", 10, "!", "hello!!!!!"},
		{"world", 5, "*", "world"},
		{"Go", 4, "-", "Go--"},
		{"", 3, "#", "###"},
		{"你好", 8, "~", "你好~~"},
	}

	for _, test := range tests {
		result := PadRight(test.src, test.size, test.pad)
		if result != test.result {
			t.Errorf(
				"PadRight(%q, %d, %q) = %q; expected %q",
				test.src,
				test.size,
				test.pad,
				result,
				test.result,
			)
		}
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello-world"},
		{"Hello World", "hello-world"},
		{"hello-world", "hello-world"},
		{"hello_world", "hello-world"},

		{"hello", "hello"},
		{"Hello", "hello"},

		{"123", "123"},
		{"!@#", ""},
		{"", ""},

		{"foobar", "foobar"},
		{"&FOO:BAR$BAZ", "foo-bar-baz"},
		{"fooBar", "foo-bar"},
		{"FOObar", "fo-obar"},
		{"$foo%", "foo"},
		{"   $#$Foo   22    bar   ", "foo-22-bar"},
		{"Foo-#1😄$_%^&*(1bar", "foo-1-1-bar"},

		{"你好世界", "你好世界"},
		{"中文编程", "中文编程"},
	}

	for _, test := range tests {
		result := KebabCase(test.input)
		if result != test.expected {
			t.Errorf("KebabCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestUpperKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO-WORLD"},
		{"Hello World", "HELLO-WORLD"},
		{"hello-world", "HELLO-WORLD"},
		{"hello_world", "HELLO-WORLD"},

		{"hello", "HELLO"},
		{"Hello", "HELLO"},
		{"123", "123"},

		{"!@#", ""},

		{"你好世界", "你好世界"},
		{"中文编程", "中文编程"},

		{"Foo-Bar-Baz", "FOO-BAR-BAZ"},
		{"fooBarBaz", "FOO-BAR-BAZ"},
	}

	for _, test := range tests {
		result := UpperKebabCase(test.input)
		if result != test.expected {
			t.Errorf("UpperKebabCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "hello_world"},
		{"&FOO:BAR$BAZ", "foo_bar_baz"},
		{"Hello World", "hello_world"},
		{"123", "123"},
		{"!@#", ""},
		{"你好世界", "你好世界"},
		{"你好，世界！", "你好，世界！"},
	}

	for _, test := range tests {
		result := SnakeCase(test.input)
		if result != test.expected {
			t.Errorf("SnakeCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestUpperSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HELLO_WORLD"},
		{"&FOO:BAR$BAZ", "FOO_BAR_BAZ"},
		{"Hello World", "HELLO_WORLD"},
		{"123", "123"},
		{"!@#", ""},
		{"你好世界", "你好世界"},
		{"你好，世界！", "你好，世界！"},
	}

	for _, test := range tests {
		result := UpperSnakeCase(test.input)
		if result != test.expected {
			t.Errorf("UpperSnakeCase(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected string
	}{
		{"hello,world", ",", "hello"},
		{"GoLang is fun", " ", "GoLang"},
		{"no-separator", "-", "no"},
		{"", ",", ""},

		{"你好,世界", ",", "你好"},
		{"中文编程,测试", ",", "中文编程"},

		{"Hello-kyden", "kyden", "Hello-"},
		{"Hello-kyden", "", "Hello-kyden"},
	}

	for _, test := range tests {
		result := Before(test.input, test.sep)
		if result != test.expected {
			t.Errorf(
				"Before(%q, %q) = %q; expected %q",
				test.input,
				test.sep,
				result,
				test.expected,
			)
		}
	}
}

func TestBeforeLast(t *testing.T) {
	tests := []struct {
		input    string
		ch       string
		expected string
	}{
		{"hello world", " ", "hello"},
		{"hello", "l", "hel"},
		{"hello", "o", "hell"},

		{"hello", "z", "hello"},
		{"", "a", ""},
		{"hello", "", "hello"},

		{"abcabc", "c", "abcab"},
		{"abcabc", "b", "abca"},
		{"abcabc", "a", "abc"},

		{"abcabc", "d", "abcabc"},
	}

	for _, test := range tests {
		result := BeforeLast(test.input, test.ch)
		if result != test.expected {
			t.Errorf(
				"BeforeLast(%q, %q) = %q; expected %q",
				test.input,
				test.ch,
				result,
				test.expected,
			)
		}
	}
}

func TestAfter(t *testing.T) {
	tests := []struct {
		str      string
		ch       string
		expected string
	}{
		{"hello world", " ", "world"},
		{"hello world", "o", " world"},
		{"hello world", "z", "hello world"},

		{"", " ", ""},
		{"hello", "", "hello"},
		{"", "", ""},
	}

	for _, test := range tests {
		result := After(test.str, test.ch)
		if result != test.expected {
			t.Errorf("After(%q, %q) = %q; expected %q", test.str, test.ch, result, test.expected)
		}
	}
}

func TestAfterLast(t *testing.T) {
	tests := []struct {
		str      string
		ch       string
		expected string
	}{
		{"hello world", " ", "world"},
		{"hello world", "o", "rld"},
		{"hello world", "z", "hello world"},

		{"", " ", ""},
		{"hello", "", "hello"},

		{"hello world hello", "o", ""},
	}

	for _, test := range tests {
		result := AfterLast(test.str, test.ch)
		if result != test.expected {
			t.Errorf(
				"AfterLast(%q, %q) = %q; expected %q",
				test.str,
				test.ch,
				result,
				test.expected,
			)
		}
	}
}

func TestIsString(t *testing.T) {
	tests := []struct {
		input    any
		expected bool
	}{
		{nil, false},
		{"hello", true},
		{123, false},
		{true, false},
		{nil, false},
		{[]int{}, false},
		{map[string]int{}, false},
		{"你好", true},
		{complex(1, 2), false},
	}

	for _, test := range tests {
		result := IsString(test.input)
		if result != test.expected {
			t.Errorf("IsString(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"GoLang", "gnaLoG"},

		{"", ""},
		{"h", "h"},
		{"H", "H"},

		{"123", "321"},
		{"!@#", "#@!"},
		{"你好", "好你"},
		{"你好世界", "界世好你"},
		{"你好，世界！", "！界世，好你"},
		{"中文编程", "程编文中"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestWarp(t *testing.T) {
	tests := []struct {
		str      string
		sWarp    string
		expected string
	}{
		{"hello", "!", "!hello!"},
		{"world", "*", "*world*"},
		{"Go", "-", "-Go-"},
		{"", "#", ""},
		{"你好", "@", "@你好@"},
		{"hello", "", "hello"},
		{"", "!", ""},
	}

	for _, test := range tests {
		result := Warp(test.str, test.sWarp)
		if result != test.expected {
			t.Errorf("Warp(%q, %q) = %q; expected %q", test.str, test.sWarp, result, test.expected)
		}
	}
}

func TestUnWarp(t *testing.T) {
	tests := []struct {
		input    string
		sWarp    string
		expected string
	}{
		{"hello", "", "hello"},
		{"", "world", ""},

		{"hello", "world", "hello"},
		{"helloworld", "world", "helloworld"},
		{"worldhello", "world", "worldhello"},

		{"helloworldworld", "world", "helloworldworld"},
		{"abcabc", "abc", ""},
		{"abcdefabc", "abc", "def"},
	}

	for _, test := range tests {
		result := UnWarp(test.input, test.sWarp)
		if result != test.expected {
			t.Errorf(
				"UnWarp(%q, %q) = %q; expected %q",
				test.input,
				test.sWarp,
				result,
				test.expected,
			)
		}
	}
}

func TestSubString(t *testing.T) {
	tests := []struct {
		src    string
		begin  int
		size   int
		result string
	}{
		{"hello", 0, 2, "he"},
		{"hello", 1, 3, "ell"},
		{"hello", 4, 1, "o"},
		{"hello", 5, 1, ""},
		{"hello", -1, 2, "o"},
		{"hello", -5, 2, "he"},
		{"hello", 0, 10, "hello"},
		{"你好世界", 1, 2, "好世"},
		{"你好世界", -2, 2, "世界"},
		{"你好世界", 3, 0, ""},
		{"你好世界", 0, -1, ""},
		{"\x00你好世界\x00", 2, 3, "好世界"},
	}

	for _, test := range tests {
		result := SubString(test.src, test.begin, test.size)
		if result != test.result {
			t.Errorf(
				"SubString(%q, %d, %d) = %q; expected %q",
				test.src,
				test.begin,
				test.size,
				result,
				test.result,
			)
		}
	}
}

func TestRemoveNonPrintable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "Hello, World!"},
		{"こんにちは", "こんにちは"},

		{"\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0A\x0B\x0C\x0D\x0E\x0F", ""},
		{"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F", ""},
		{"\x7F", ""},

		{"Hello\x00World", "HelloWorld"},
		{"\u00A0\u00A1\u00A2", "\u00A1\u00A2"},
	}

	for _, test := range tests {
		result := RemoveNonPrintable(test.input)
		if result != test.expected {
			t.Errorf("RemoveNonPrintable(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
	}{
		{"hello", []byte("hello")},
		{"world", []byte("world")},
		{"", []byte("")},
		{"你好", []byte("你好")},
	}

	for _, test := range tests {
		result := StringToBytes(test.input)
		if !bytes.Equal(result, test.expected) {
			t.Errorf("StringToBytes(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestBytesToString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte("hello"), "hello"},
		{[]byte("world"), "world"},
		{[]byte(""), ""},
		{[]byte("你好"), "你好"},
	}

	for _, test := range tests {
		result := BytesToString(test.input)
		if result != test.expected {
			t.Errorf("BytesToString(%v) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{" ", true},
		{"\t", true},
		{"\n", true},
		{"\r", true},
		{" \t\n\r", true},
		{"hello", false},
		{"你好", false},
		{"123", false},
		{"!@#", false},
		{"你好 世界", false},
	}

	for _, test := range tests {
		result := IsSpace(test.input)
		if result != test.expected {
			t.Errorf("IsSpace(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsNotSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{" ", false},
		{"\t", false},
		{"\n", false},
		{"hello", true},
		{"你好", true},
		{" 你好 ", true},
		{"\t你好\n", true},
		{"", false},
	}

	for _, test := range tests {
		result := IsNotSpace(test.input)
		if result != test.expected {
			t.Errorf("IsNotSpace(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestHasPrefixAny(t *testing.T) {
	tests := []struct {
		str      string
		prefixs  []string
		expected bool
	}{
		{"hello", []string{"he", "ha"}, true},
		{"world", []string{"wo", "wa"}, true},
		{"GoLang", []string{"Go", "GoL"}, true},

		{"", []string{"a", "b"}, false},
		{"h", []string{}, false},

		{"H", []string{"h"}, false},
		{"123", []string{"12", "123"}, true},
		{"!@#", []string{"!@", "!#"}, true},
		{"你好", []string{"你", "好"}, true},
		{"你好世界", []string{"你好", "好世"}, true},
	}

	for _, test := range tests {
		result := HasPrefixAny(test.str, test.prefixs...)
		if result != test.expected {
			t.Errorf(
				"HasPrefixAny(%q, %v) = %v; expected %v",
				test.str,
				test.prefixs,
				result,
				test.expected,
			)
		}
	}
}

func TestHasSuffixAny(t *testing.T) {
	tests := []struct {
		str      string
		suffixs  []string
		expected bool
	}{
		{"hello", []string{"lo", "he"}, true},
		{"world", []string{"ld", "wo"}, true},
		{"GoLang", []string{"g", "Lang"}, true},

		{"", []string{"a", "b"}, false},
		{"h", []string{}, false},

		{"H", []string{"h"}, false},
		{"123", []string{"23", "123"}, true},
		{"!@#", []string{"#@", "!#"}, false},
		{"你好", []string{"好", "你"}, true},
		{"你好世界", []string{"世界", "好世"}, true},
	}

	for _, test := range tests {
		result := HasSuffixAny(test.str, test.suffixs...)
		if result != test.expected {
			t.Errorf(
				"HasSuffixAny(%q, %v) = %v; expected %v",
				test.str,
				test.suffixs,
				result,
				test.expected,
			)
		}
	}
}

func TestIndexOffset(t *testing.T) {
	tests := []struct {
		str      string
		substr   string
		offset   int
		expected int
	}{
		{"hello world", "world", 6, 6},
		{"hello world", "o", 7, 7},
		{"hello world", "l", 3, 3},
		{"hello world", "z", 0, -1},
		{"hello world", "o", 11, -1},
		{"hello world", "", 0, 0},
		{"", "a", 0, -1},
		{"hello world", "world", -1, -1},
		{"hello world", "world", 11, -1},
	}

	for _, test := range tests {
		result := IndexOffset(test.str, test.substr, test.offset)
		if result != test.expected {
			t.Errorf(
				"IndexOffset(%q, %q, %d) = %d; expected %d",
				test.str,
				test.substr,
				test.offset,
				result,
				test.expected,
			)
		}
	}
}

func TestReplaceWithMap(t *testing.T) {
	tests := []struct {
		str        string
		replaceMap map[string]string
		expected   string
	}{
		{"hello world", map[string]string{"hello": "hi"}, "hi world"},
		{"Hello World", map[string]string{"World": "Go"}, "Hello Go"},
		{"hello-world", map[string]string{"-": "_"}, "hello_world"},
		{"hello_world", map[string]string{"_": "-"}, "hello-world"},

		{"", map[string]string{"a": "b"}, ""},

		{"abc", map[string]string{"a": "A", "b": "B", "c": "C"}, "ABC"},
		{"你好世界", map[string]string{"你好": "Hello"}, "Hello世界"},
		{"你好，世界！", map[string]string{"！": "?"}, "你好，世界?"},
		{"你好，世界！", map[string]string{"！": "？"}, "你好，世界？"},
	}

	for _, test := range tests {
		result := ReplaceWithMap(test.str, test.replaceMap)
		if result != test.expected {
			t.Errorf(
				"ReplaceWithMap(%q, %v) = %q; expected %q",
				test.str,
				test.replaceMap,
				result,
				test.expected,
			)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"  hello world  ", "", "hello world"},
		{"\t\nhello world\t\n", "", "hello world"},
		{"abcHello World!cba", "abc!", "Hello World"},
		{"abcHello World!cba", "abc!dEF", "Hello Worl"},
		{"你好 世界！", "", "你好 世界！"},
		{"你好 世界！", "你好", "世界！"},
	}

	for _, test := range tests {
		result := Trim(test.input, test.cutset)
		if result != test.expected {
			t.Errorf(
				"Trim(%q, %q) = %q; expected %q",
				test.input,
				test.cutset,
				result,
				test.expected,
			)
		}
	}
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		input     string
		delimeter string
		cutset    string
		expected  []string
	}{
		{"hello, world, go", ",", "", []string{"hello", "world", "go"}},
		{"  hello, world, go  ", ",", "", []string{"hello", "world", "go"}},
		{"a,b,cHello,d,eWorld,f", ",", "abc", []string{"Hello", "d", "eWorld", "f"}},
		{"你好,世界,Go", ",", "", []string{"你好", "世界", "Go"}},
		{"你好,世界,Go", ",", "你好", []string{"世界", "Go"}},
	}

	for _, test := range tests {
		result := SplitAndTrim(test.input, test.delimeter, test.cutset)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf(
				"SplitAndTrim(%q, %q, %q) = %v; expected %v",
				test.input,
				test.delimeter,
				test.cutset,
				result,
				test.expected,
			)
		}
	}
}

func TestHideString(t *testing.T) {
	tests := []struct {
		src      string
		beg      int
		end      int
		hideChar string
		expected string
	}{
		{"hello world", 3, 7, "*", "hel****orld"},
		{"1234567890", 1, 9, "X", "1XXXXXXXX0"},
		{"abcdef", 0, 6, "-", "------"},

		{"abc", 1, 2, "", "abc"},
		{"abc", 3, 2, "*", "abc"},
		{"abc", -1, 2, "*", "abc"},
		{"abc", 2, -1, "*", "abc"},
		{"abc", 4, 5, "*", "abc"},
		{"", 0, 0, "*", ""},
	}

	for _, test := range tests {
		result := HideString(test.src, test.beg, test.end, test.hideChar)
		if result != test.expected {
			t.Errorf(
				"HideString(%q, %d, %d, %q) = %q; expected %q",
				test.src,
				test.beg,
				test.end,
				test.hideChar,
				result,
				test.expected,
			)
		}
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		src     string
		substrs []string
		want    bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"hello", "planet"}, false},
		{"", []string{}, true},

		{"hello", []string{"hello", "h"}, true},
		{"hello", []string{"hell", "o"}, true},
		{"hello", []string{"he", "llo"}, true},
		{"hello", []string{"he", "l", "lo"}, true},
		{"hello", []string{"he", "ll", "o"}, true},

		{"hello", []string{"h", "e", "l", "l", "o"}, true},
		{"hello", []string{"h", "e", "l", "l", "o", "x"}, false},
	}

	for _, test := range tests {
		got := ContainsAll(test.src, test.substrs)
		if got != test.want {
			t.Errorf("ContainsAll(%q, %v) = %v; want %v", test.src, test.substrs, got, test.want)
		}
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		src     string
		substrs []string
		want    bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"planet", "mars"}, false},
		{"hello world", []string{"planet", "mars", "hello"}, true},
		{"hello world", []string{"planet", "mars", "jupiter"}, false},

		{"", []string{}, false},
		{"hello", []string{"hell", "o"}, true},
		{"hello", []string{"he", "llo"}, true},
		{"hello", []string{"he", "ll", "o"}, true},
		{"hello", []string{"he", "l", "lo"}, true},
		{"hello", []string{"he", "ll", "o", "x"}, true},
	}

	for _, test := range tests {
		got := ContainsAny(test.src, test.substrs)
		if got != test.want {
			t.Errorf("ContainsAny(%q, %v) = %v; want %v", test.src, test.substrs, got, test.want)
		}
	}
}

func TestRemoveWhiteSpace(t *testing.T) {
	tests := []struct {
		input    string
		rmAll    bool
		expected string
	}{
		{"hello world", false, "hello world"},
		{"hello world", true, "helloworld"},

		{"  hello   world  ", false, "hello world"},
		{"  hello   world  ", true, "helloworld"},

		{"no-whitespace", false, "no-whitespace"},
		{"no-whitespace", true, "no-whitespace"},

		{"  \t\n  ", false, ""},
		{"  \t\n  ", true, ""},

		{"中文 编程", false, "中文 编程"},
		{"中文 编程", true, "中文编程"},
	}

	for _, test := range tests {
		result := RemoveWhiteSpace(test.input, test.rmAll)
		if result != test.expected {
			t.Errorf(
				"RemoveWhiteSpace(%q, %v) = %q; expected %q",
				test.input,
				test.rmAll,
				result,
				test.expected,
			)
		}
	}
}

func TestSubInBetween(t *testing.T) {
	tests := []struct {
		str      string
		beg      string
		end      string
		expected string
	}{
		{"startmiddleend", "start", "end", "middle"},
		{"startmiddle", "start", "end", ""},
		{"middlestartend", "start", "end", ""},
		{"startendmiddle", "start", "end", ""},

		{"startmiddleendextra", "start", "end", "middle"},
		{"", "start", "end", ""},

		{"start", "start", "end", ""},
		{"end", "start", "end", ""},
	}

	for _, test := range tests {
		result := SubInBetween(test.str, test.beg, test.end)
		if result != test.expected {
			t.Errorf(
				"SubInBetween(%q, %q, %q) = %q; expected %q",
				test.str,
				test.beg,
				test.end,
				result,
				test.expected,
			)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		str1        string
		str2        string
		expected    int
		expectedErr error
	}{
		{"", "", 0, nil},
		{"abc", "abc", 0, nil},
		{"abc", "abd", 1, nil},
		{"abc", "abcd", -1, errors.New("the length of two strings must be equal")},

		{"你好世界", "你好世界", 0, nil},
		{"你好世界", "你好世界啊", -1, errors.New("the length of two strings must be equal")},
	}

	for _, test := range tests {
		result, err := HammingDistance(test.str1, test.str2)
		if err != nil && test.expectedErr == nil {
			t.Errorf(
				"HammingDistance(%q, %q) returned unexpected error: %v",
				test.str1,
				test.str2,
				err,
			)
		} else if err == nil && test.expectedErr != nil {
			t.Errorf("HammingDistance(%q, %q) did not return expected error: %v", test.str1, test.str2, test.expectedErr)
		} else if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("HammingDistance(%q, %q) returned wrong error message: got %q, want %q", test.str1, test.str2, err.Error(), test.expectedErr.Error())
		} else if result != test.expected {
			t.Errorf("HammingDistance(%q, %q) = %d; want %d", test.str1, test.str2, result, test.expected)
		}
	}
}

func TestShuffle(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"helloworld"},
		{"12345!@#$%"},
		{"hellokyden"},

		{""}, // 空字符串测试用例
	}

	for _, tc := range testCases {
		shuffled, err := Shuffle(tc.input)
		if tc.input == "" {
			if err == nil {
				t.Errorf("Shuffle(%q) = %q, expected an error", tc.input, shuffled)
			}
			continue
		}

		if shuffled == tc.input {
			t.Errorf("Shuffle(%q) = %q, expected a different string", tc.input, shuffled)
		}
		if len(shuffled) != len(tc.input) {
			t.Errorf("Shuffle(%q) length = %d, expected %d", tc.input, len(shuffled), len(tc.input))
		}

		// 可选：检查是否每个字符都在洗牌后的字符串中
		for _, char := range tc.input {
			if !strings.Contains(shuffled, string(char)) {
				t.Errorf("Shuffle(%q) missing character %q", tc.input, char)
			}
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		input    string
		shift    int
		expected string
	}{
		{"hello", 2, "lohel"},
		{"hello", -2, "llohe"},
		{"hello", 0, "hello"},
		{"hello", 5, "hello"},
		{"hello", -5, "hello"},

		{"", 3, ""},
		{"a", 1, "a"},
		{"ab", 1, "ba"},
		{"ab", -1, "ba"},
	}

	for _, test := range tests {
		result := Rotate(test.input, test.shift)
		if result != test.expected {
			t.Errorf(
				"Rotate(%q, %d) = %q; expected %q",
				test.input,
				test.shift,
				result,
				test.expected,
			)
		}
	}
}

func TestRegexMatchAllGroups(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		pattern  string
		expected [][]string
	}{
		{
			name:     "基本测试",
			str:      "abc123",
			pattern:  `(\w)(\w)(\w)`,
			expected: [][]string{{"abc", "a", "b", "c"}, {"123", "1", "2", "3"}},
		},
		{
			name:     "无匹配测试",
			str:      "abc123",
			pattern:  `(\d\d)`,
			expected: [][]string{{"12", "12"}},
		},
		{
			name:    "全局匹配测试",
			str:     "abc123abc",
			pattern: `(\w)(\w)(\w)`,
			expected: [][]string{
				{"abc", "a", "b", "c"},
				{"123", "1", "2", "3"},
				{"abc", "a", "b", "c"},
			},
		},
		{
			name:    "复杂模式测试",
			str:     "The quick brown fox jumps over the lazy dog.",
			pattern: `(\w+)\s+(\w+)`,
			expected: [][]string{
				{"The quick", "The", "quick"},
				{"brown fox", "brown", "fox"},
				{"jumps over", "jumps", "over"},
				{"the lazy", "the", "lazy"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RegexMatchAllGroups(tt.str, tt.pattern)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf(
					"RegexMatchAllGroups(%q, %q) = %v; want %v",
					tt.str,
					tt.pattern,
					actual,
					tt.expected,
				)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		strings  []string
		expected string
	}{
		{"Empty input", 0, []string{}, ""},
		{"Single string", 0, []string{"hello"}, "hello"},
		{"Multiple strings", 0, []string{"hello", "world"}, "helloworld"},
		{"Specified length", 10, []string{"hello", "world"}, "helloworld"},
		{"Length exceeds", 5, []string{"hello", "world"}, "helloworld"},
		{"Empty strings", 0, []string{"", "", ""}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Concat(tt.length, tt.strings...)
			if actual != tt.expected {
				t.Errorf(
					"Concat(%d, %v) = %q; expected %q",
					tt.length,
					tt.strings,
					actual,
					tt.expected,
				)
			}
		})
	}
}

func TestEllipsis(t *testing.T) {
	tests := []struct {
		str      string
		size     int
		expected string
	}{
		{"Hello, world!", 5, "Hello..."},
		{"你好，世界！", 2, "你好..."},
		{"", 5, ""},
		{"Short", 10, "Short"},
		{"Longer text to be truncated", 20, "Longer text to be tr..."},
		{"", 0, ""},
	}

	for _, test := range tests {
		result := Ellipsis(test.str, test.size)
		if result != test.expected {
			t.Errorf(
				"Ellipsis(%q, %d) = %q; expected %q",
				test.str,
				test.size,
				result,
				test.expected,
			)
		}
	}
}
