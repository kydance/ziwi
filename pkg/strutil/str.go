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
