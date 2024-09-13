package strutil

import (
	"testing"
)

func TestToUpperAll(t *testing.T) {
	tests := []struct {
		input []rune
		want  []rune
	}{
		{[]rune("hello"), []rune("HELLO")},
		{[]rune("world"), []rune("WORLD")},
		{[]rune("GoLang"), []rune("GOLANG")},
		{[]rune(""), []rune("")},
	}

	for _, test := range tests {
		result := toUpperAll(test.input)
		if !equalRuneSlices(result, test.want) {
			t.Errorf("toUpperAll(%q) = %q; expected %q", test.input, result, test.want)
		}
	}
}

func TestToLowerAll(t *testing.T) {
	tests := []struct {
		input []rune
		want  []rune
	}{
		{[]rune("HELLO"), []rune("hello")},
		{[]rune("WORLD"), []rune("world")},
		{[]rune("golang"), []rune("golang")},
		{[]rune(""), []rune("")},
	}

	for _, test := range tests {
		result := toLowerAll(test.input)
		if !equalRuneSlices(result, test.want) {
			t.Errorf("toLowerAll(%q) = %q; expected %q", test.input, result, test.want)
		}
	}
}

func TestSplitIntoStrings(t *testing.T) {
	tests := []struct {
		input    string
		upper    bool
		expected []string
	}{
		{"HelloWorld123", true, []string{"HELLO", "WORLD", "123"}},
		{"HelloWorld123", false, []string{"hello", "world", "123"}},

		{"GoLangProgramming", true, []string{"GO", "LANG", "PROGRAMMING"}},
		{"GoLangProgramming", false, []string{"go", "lang", "programming"}},

		{"12345!@#$%", true, []string{"12345"}},
		{"12345!@#$%", false, []string{"12345"}},

		{"", true, []string{}},
		{"", false, []string{}},
	}

	for _, test := range tests {
		result := splitIntoStrings(test.input, test.upper)
		if !equalStringSlices(result, test.expected) {
			t.Errorf("splitIntoStrings(%q, %v) = %v; expected %v", test.input, test.upper, result, test.expected)
		}
	}
}
