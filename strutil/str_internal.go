// =============================================================================
/*!
 *  @file       str_internal.go
 *  @brief      str internal functions
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

type Pos int

const (
	PosBoth Pos = iota
	PosLeft
	PosRight
)

// splitIntoStrings splits string into string slices based on type of unicode character.
//
//	 example:
//		"GoLangProgramming" -> {"go", "lang", "programming"}
//		"12345!@#$%" -> {"12345"},
func splitIntoStrings(str string, upper bool) []string {
	// 预估切片容量为字符串长度的1/3，因为大多数情况下会分割成多个单词
	runes := make([][]rune, 0, len(str)/3)
	lastCharType, charType := 0, 0

	// split string into runes based on type of unicode character
	for _, r := range str {
		switch true {
		case unicode.IsLower(r):
			charType = 1
		case unicode.IsUpper(r):
			charType = 2
		case unicode.IsDigit(r):
			charType = 3
		default:
			charType = 4
		}

		if charType == lastCharType {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}

		lastCharType = charType
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			length := len(runes[i]) - 1
			temp := runes[i][length]
			runes[i+1] = append([]rune{temp}, runes[i+1]...)
			runes[i] = runes[i][:length]
		}
	}

	// filter all none letters and non digits
	result := make([]string, 0, len(runes))
	for _, rs := range runes {
		if len(rs) > 0 && (unicode.IsLetter(rs[0]) || unicode.IsDigit(rs[0])) {
			if upper {
				result = append(result, string(toUpperAll(rs)))
			} else {
				result = append(result, string(toLowerAll(rs)))
			}
		}
	}

	return result
}

// toUpperAll converts all runes to upper case.
func toUpperAll(rs []rune) []rune { return []rune(strings.ToUpper(string(rs))) }

// toLowerAll converts all runes to lower case.
func toLowerAll(rs []rune) []rune { return []rune(strings.ToLower(string(rs))) }

// padAtPos pads string at given position.
// pos: 0 - both, 1 - left, 2 - right
//
//	Example:
//		padAtPos("hello", 6, "*", 0) -> "hello*"
//		padAtPos("hello", 10, "abc", 0) -> "abhelloabc"
func padAtPos(str string, size int, pad string, pos Pos) string {
	if len(str) >= size {
		return str
	}

	if pad == "" {
		pad = " "
	}
	padSize := len(pad)

	// Calculate the number of padding characters needed.
	size = size - len(str)
	leftPadSize := 0
	if pos == 0 {
		leftPadSize = size / 2
	} else if pos == 1 {
		leftPadSize = size
	}
	rightPadSize := size - leftPadSize

	// Using strings.Builder to optimize string concatenation
	var builder strings.Builder
	builder.Grow(size + len(str))

	// Pad left
	for i := 0; i < leftPadSize; i++ {
		builder.WriteByte(pad[i%padSize])
	}

	// Write original string
	builder.WriteString(str)

	// Pad right
	for i := 0; i < rightPadSize; i++ {
		builder.WriteByte(pad[i%padSize])
	}

	return builder.String()
}
