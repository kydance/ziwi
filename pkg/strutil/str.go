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
