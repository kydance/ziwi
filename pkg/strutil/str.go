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
	"unsafe"
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

// IsString checks if the val data type is string or not.
func IsString(val any) bool {
	_, ok := val.(string)
	return ok
}

// Reverse reverses the string.
func Reverse(str string) string {
	vr := []rune(str)

	for i, j := 0, len(vr)-1; i < len(vr)/2; i, j = i+1, j-1 {
		vr[i], vr[j] = vr[j], vr[i]
	}

	return string(vr)
}

// Warp wraps the string with the given string.
//
//	Example:
//		Warp("hello", "*") -> "*hello*"
func Warp(str, sWarp string) string {
	if str == "" || sWarp == "" {
		return str
	}

	var sb strings.Builder

	sb.WriteString(sWarp)
	sb.WriteString(str)
	sb.WriteString(sWarp)

	return sb.String()
}

// UnWarp unwraps the string with the given string.
//
//	Example:
//		UnWarp("*hello*", "*") -> "hello"
//		UnWarp("abcdefabc", "abc") -> "def"
func UnWarp(str, sWarp string) string {
	if str == "" || sWarp == "" {
		return str
	}

	if !strings.HasPrefix(str, sWarp) || !strings.HasSuffix(str, sWarp) {
		return str
	}

	return str[len(sWarp) : len(str)-len(sWarp)]
}

// SubString returns a substring of the specified size from begin.
func SubString(src string, begin int, size int) string {
	vr := []rune(src)

	if begin < 0 {
		begin += len(vr)
		if begin < 0 {
			begin = 0
		}
	}
	if begin > len(vr) || size <= 0 {
		return ""
	}

	if size > len(vr)-begin {
		size = len(vr) - begin
	}

	str := string(vr[begin : begin+size])
	return strings.Replace(str, "\x00", "", -1)
}

// RemoveNonPrintable removes all non-printable characters from the string.
func RemoveNonPrintable(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
}

// StringToBytes converts the string to byte slice without memory alloc.
func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// BytesToString converts the byte slice to string without memory alloc.
func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// IsSpace checks if the string is whitespace, empty or not.
//
//	Example:
//		"" -> true
//		" \t\n\r" -> true
func IsSpace(str string) bool {
	if len(str) == 0 {
		return true
	}

	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

// IsNotSpace checks if the string is not whitespace, empty or not.
func IsNotSpace(str string) bool {
	return !IsSpace(str)
}

// HasPrefixAny checks if the string has any of the given prefixes.
func HasPrefixAny(str string, prefixs ...string) bool {
	if len(str) == 0 || len(prefixs) == 0 {
		return false
	}

	for _, prefix := range prefixs {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}

	return false
}

// HasSuffixAny checks if the string has any of the given suffixes.
func HasSuffixAny(str string, suffixs ...string) bool {
	if len(str) == 0 || len(suffixs) == 0 {
		return false
	}

	for _, suffix := range suffixs {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}

	return false
}

// IndexOffset return the index of the first occurrence of the substring
// in the string after offset, or -1 if not found.
func IndexOffset(str, substr string, offset int) int {
	if offset > len(str)-1 || offset < 0 {
		return -1
	}

	return strings.Index(str[offset:], substr) + offset
}

// ReplaceWithMap returns a copy of `str`,
// which is replaced by a map in unordered way, case-sensitively.
func ReplaceWithMap(str string, replaceMap map[string]string) string {
	for k, v := range replaceMap {
		str = strings.ReplaceAll(str, k, v)
	}
	return str
}

// Trim strips whitespace (or other characters) from beginning and end of the string.
//
//	Example:
//		Trim("abcHello World!cba", "abc!dEF") -> "Hello Worl"

func Trim(str string, cutset ...string) string {
	// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
	DefaultTrimChars := string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})

	if len(cutset) > 0 {
		DefaultTrimChars += cutset[0]
	}

	return strings.Trim(str, DefaultTrimChars)
}

// SplitAndTrim splits the string by the delimeter and trims the result.
//
//	Example:
//		SplitAndTrim("a,b,cHello,d,eWorld,f", ",", "abc") -> []string{"Hello", "d", "eWorld", "f"}
func SplitAndTrim(str, delimeter string, cutset ...string) (vStr []string) {
	temp := strings.Split(str, delimeter)

	for _, v := range temp {
		v := Trim(v, cutset...)
		if v != "" {
			vStr = append(vStr, v)
		}
	}

	return
}

// HideString hides the string between beg and end with the hideChar.
//
//	Example:
//		HideString({"hello world", 3, 7, "*") -> "hel****orld"
func HideString(src string, beg, end int, hideChar string) string {
	srcSize := len(src)

	if beg > srcSize-1 || beg < 0 || end < 0 || beg > end {
		return src
	}

	if end > srcSize {
		end = srcSize
	}

	if hideChar == "" {
		return src
	}

	return src[:beg] + strings.Repeat(hideChar, end-beg) + src[end:]
}
