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
			t.Errorf(
				"splitIntoStrings(%q, %v) = %v; expected %v",
				test.input,
				test.upper,
				result,
				test.expected,
			)
		}
	}
}

func TestPadAtPosEdgeCases(t *testing.T) {
	tests := []struct {
		str  string
		size int
		pad  string
		pos  int
		want string
	}{
		{"", 5, "-", 0, "-----"},
		{"", 5, "-", 1, "-----"},
		{"", 5, "-", 2, "-----"},

		{"hello", 3, "", 0, "hello"},
		{"hello", 3, "", 1, "hello"},
		{"hello", 3, "", 2, "hello"},

		{"hello", 6, "*", 0, "hello*"},
		{"hello", 6, "*", 1, "*hello"},
		{"hello", 6, "*", 2, "hello*"},

		{"hello", 10, "abc", 0, "abhelloabc"},
		{"hello", 10, "abc", 1, "abcabhello"},
		{"hello", 10, "abc", 2, "helloabcab"},
	}

	for _, test := range tests {
		result := padAtPos(test.str, test.size, test.pad, test.pos)
		if result != test.want {
			t.Errorf(
				"padAtPos(%q, %d, %q, %d) = %q; expected %q",
				test.str,
				test.size,
				test.pad,
				test.pos,
				result,
				test.want,
			)
		}
	}
}
