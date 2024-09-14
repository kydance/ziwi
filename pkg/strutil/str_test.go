package strutil

import (
	"bytes"
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
			t.Errorf("Pad(%q, %d, %q) = %q; expected %q", test.src, test.size, test.pad, result, test.result)
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
			t.Errorf("PadLeft(%q, %d, %q) = %q; expected %.sql", test.src, test.size, test.pad, result, test.result)
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
			t.Errorf("PadRight(%q, %d, %q) = %q; expected %q", test.src, test.size, test.pad, result, test.result)
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
			t.Errorf("Before(%q, %q) = %q; expected %q", test.input, test.sep, result, test.expected)
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
			t.Errorf("BeforeLast(%q, %q) = %q; expected %q", test.input, test.ch, result, test.expected)
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
			t.Errorf("AfterLast(%q, %q) = %q; expected %q", test.str, test.ch, result, test.expected)
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
			t.Errorf("UnWarp(%q, %q) = %q; expected %q", test.input, test.sWarp, result, test.expected)
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
			t.Errorf("SubString(%q, %d, %d) = %q; expected %q", test.src, test.begin, test.size, result, test.result)
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
			t.Errorf("HasPrefixAny(%q, %v) = %v; expected %v", test.str, test.prefixs, result, test.expected)
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
			t.Errorf("HasSuffixAny(%q, %v) = %v; expected %v", test.str, test.suffixs, result, test.expected)
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
			t.Errorf("IndexOffset(%q, %q, %d) = %d; expected %d", test.str, test.substr, test.offset, result, test.expected)
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
			t.Errorf("ReplaceWithMap(%q, %v) = %q; expected %q", test.str, test.replaceMap, result, test.expected)
		}
	}
}
