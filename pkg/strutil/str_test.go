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
