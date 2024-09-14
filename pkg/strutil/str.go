// =============================================================================
/*!
 *  @file       str.go
 *  @brief      strutil implements some functions to manipulate string.
 *  @author     kydenlu
 *  @date       2024.09
 *  @note
 */
// =============================================================================

package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase converts string to camelCase string.
// Non letters and numbers will be ignored.
//
//	Example:
//		"!@#" -> ""
//		"&FOO:BAR$BAZ" -> "fooBarBaz"
//		"hello_world" -> "helloWorld"
func CamelCase(str string) string {
	var sbr strings.Builder

	strs := splitIntoStrings(str, false)
	for idx, str := range strs {
		if idx == 0 {
			sbr.WriteString(strings.ToLower(str))
		} else {
			sbr.WriteString(Capitalize(str))
		}
	}

	return sbr.String()
}

// Capitalize converts first character of string to upper case and the remaining to lower case.
//
//	Example:
//		"hello" -> "Hello"
//		"GoLang" -> "Golang"
func Capitalize(str string) string {
	rs := make([]rune, 0, len(str))
	for idx, val := range str {
		if idx == 0 {
			rs = append(rs, unicode.ToUpper(val))
			continue
		}

		rs = append(rs, unicode.ToLower(val))
	}

	return string(rs)
}

// UpperFirst converts first character of string to upper case.
//
//	Example:
//		"hello" -> "Hello"
func UpperFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[size:]
}

// LowerFirst converts first character of string to lower case.
//
//	Example:
//		"HELLO WORLD" -> "hELLO WORLD"
func LowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToLower(r)) + str[size:]
}

// Pad pads the src to the left side and the right side with pad string until the src is size long.
func Pad(src string, size int, pad string) string {
	return padAtPos(src, size, pad, 0)
}

// PadLeft pads the src to the left side with pad string until the src is size long.
func PadLeft(src string, size int, pad string) string {
	return padAtPos(src, size, pad, 1)
}

// PadRight pads the src to the right side with pad string until the src is size long.
func PadRight(src string, size int, pad string) string {
	return padAtPos(src, size, pad, 2)
}

// KebabCase converts string to kebab-case string, non letters and numbers will be ignored.
//
//	Example:
//		"hello_world" -> "hello-world"
//		"&FOO:BAR$BAZ" -> "foo-bar-baz"
func KebabCase(str string) string {
	vs := splitIntoStrings(str, false)
	return strings.Join(vs, "-")
}

// UpperKebabCase converts string to upper KEBAB-CASE string,
// non letters and numbers will be ignored.
//
//	Example:
//		"Hello World" -> "HELLO-WORLD"
//		"!@#" -> ""
func UpperKebabCase(str string) string {
	return strings.ToUpper(KebabCase(str))
}

// SnakeCase converts string to snake_case string, non letters and numbers will be ignored.
//
//	Example:
//		"hello_world" -> "hello_world"
//		"&FOO:BAR$BAZ" -> "foo_bar_baz"
func SnakeCase(str string) string {
	vs := splitIntoStrings(str, false)
	return strings.Join(vs, "_")
}

// UpperSnakeCase converts string to upper SNAKE_CASE string,
// non letters and numbers will be ignored.
//
//	Example:
//		"Hello World" -> "HELLO_WORLD"
//		"!" -> ""
func UpperSnakeCase(str string) string {
	return strings.ToUpper(SnakeCase(str))
}

// Before returns the string before the first occurrence of ch.
//
//	Example:
//		Before("no-separator", "-") -> "no"
func Before(str, ch string) string {
	if str == "" || ch == "" {
		return str
	}

	idx := strings.Index(str, ch)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// BeforeLast returns the string before the last occurrence of ch.
//
//	Example:
//		BeforeLast("abcabc", "c") -> "abcab"
func BeforeLast(str, ch string) string {
	if str == "" || ch == "" {
		return str
	}

	idx := strings.LastIndex(str, ch)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// After returns the substring after the first occurrence of a specified string in the source string
//
//	Example:
//		After("hello world", "o") -> " world"
func After(str, ch string) string {
	if str == "" || ch == "" {
		return str
	}

	idx := strings.Index(str, ch)
	if idx == -1 {
		return str
	}
	return str[idx+len(ch):]
}

// AfterLast returns the substring after the last occurrence of a specified string
// in the source string
//
//	Example:
//		AfterLast("hello world", "o") -> "rld"
func AfterLast(str, ch string) string {
	if str == "" || ch == "" {
		return str
	}

	idx := strings.LastIndex(str, ch)
	if idx == -1 {
		return str
	}
	return str[idx+len(ch):]
}
