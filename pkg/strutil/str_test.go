package strutil

import (
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
