// =============================================================================
/*!
 *  @file       validator.go
 *  @brief      Package validator implements some validate function for string.
 *  @author     kydenlu
 *  @date       2024.09
 *  @note
 */
// =============================================================================

package strutil

import (
	"encoding/json"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	alphaMatcher       = regexp.MustCompile(`^[a-zA-Z]+$`)
	letterRegexMatcher = regexp.MustCompile(`[a-zA-Z]`)

	numberRegexMatcher = regexp.MustCompile(`\d`)
	intStrMatcher      = regexp.MustCompile(`^[\+-]?\d+$`)

	urlMatcher = regexp.MustCompile(
		`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`, //nolint:lll
	)

	// FIXME Consecutive Hyphens
	dnsMatcher = regexp.MustCompile(
		`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

	// emailMatcher = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	emailMatcher = regexp.MustCompile(
		`^[a-z0-9][a-z0-9._%+\-]+[a-z0-9_%+\-]@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	chineseMobileMatcher = regexp.MustCompile(
		`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
	chineseIDMatcher    = regexp.MustCompile(`^(\d{17})([0-9]|X|x)$`)
	chineseMatcher      = regexp.MustCompile("[\u4e00-\u9fa5]")
	chinesePhoneMatcher = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}|\d{4}-\d{8}`)

	creditCardMatcher = regexp.MustCompile(
		`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`, //nolint:lll
	)

	base64Matcher = regexp.MustCompile(
		`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)

	// FIXME Missing padding '=' / empty string
	base64URLMatcher = regexp.MustCompile(
		`^([A-Za-z0-9_-]{4})*([A-Za-z0-9_-]{2}(==)?|[A-Za-z0-9_-]{3}=?)?$`)

	binMatcher = regexp.MustCompile(`^(0b)?[01]+$`)
	hexMatcher = regexp.MustCompile(`^(#|0x|0X)?[0-9a-fA-F]+$`)

	visaMatcher            = regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`)
	masterCardMatcher      = regexp.MustCompile(`^5[1-5][0-9]{14}$`)
	americanExpressMatcher = regexp.MustCompile(`^3[47][0-9]{13}$`)
	unionPay               = regexp.MustCompile(`^62[0-5]\d{13,16}$`)
	chinaUnionPay          = regexp.MustCompile(`^62[0-9]{14,17}$`)
)

// IsAlpha checks if the string contains only letters (a-zA-Z).
func IsAlpha(str string) bool { return alphaMatcher.MatchString(str) }

// IsAllUpper checks if the string is all upper letters (A-Z).
func IsAllUpper(str string) bool {
	for _, rune := range str {
		if !unicode.IsUpper(rune) {
			return false
		}
	}

	return str != ""
}

// IsAllLower checks if the string is all lower letters (a-z).
func IsAllLower(str string) bool {
	for _, rune := range str {
		if !unicode.IsLower(rune) {
			return false
		}
	}

	return str != ""
}

// IsACII checks if the string is all ASCII characters.
func IsASCII(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// IsPrint checks if the string contains only printable characters.
func IsPrint(str string) bool {
	for _, r := range str {
		if !unicode.IsPrint(r) {
			if r == '\n' || r == '\t' || r == '\r' || r == '`' {
				continue
			}
			return false
		}
	}

	return true
}

// ContainUpper checks if the string contains any upper letters (A-Z).
func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}

	return false
}

// ContainLower checks if the string contains any lower letters (a-z).
func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}

	return false
}

// ContainLetter checks if the string contains any letters (a-zA-Z).
func ContainLetter[T string | []byte](dat T) bool {
	switch v := any(dat).(type) {
	case string:
		return letterRegexMatcher.MatchString(v)
	case []byte:
		return letterRegexMatcher.Match(v)
	default:
		return false
	}
}

// ContainNumber checks if the string contains any numbers (0-9).
func ContainNumber[T string | []byte](dat T) bool {
	switch v := any(dat).(type) {
	case string:
		return numberRegexMatcher.MatchString(v)
	case []byte:
		return numberRegexMatcher.Match(v)
	default:
		return false
	}
}

// IsJSON checks if the string is a valid JSON string.
func IsJSON(str string) bool { return json.Unmarshal([]byte(str), &struct{}{}) == nil }

// IsBin check if the string is a valid binary value or not.
func IsBin(dat string) bool { return binMatcher.MatchString(dat) }

// IsHex check if the string is a valid hexadecimal value or not.
func IsHex(dat string) bool { return hexMatcher.MatchString(dat) }

// IsFloatStr check if the string can convert to a float.
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

// IsFloat check if value is a float (float32, float64) or not.
func IsFloat(v any) bool {
	switch v.(type) {
	case float32, float64:
		return true
	}
	return false
}

// IsIntStr check if the string can convert to a integer.
func IsIntStr(str string) bool { return intStrMatcher.MatchString(str) }

// IsInt check if value is a integer (int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64, uintptr) or not.
func IsInt(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		uintptr:
		return true
	}
	return false
}

// IsNumberStr check if the string can convert to a number.
func IsNumberStr(str string) bool { return IsIntStr(str) || IsFloatStr(str) }

// IsNumber check if value is a number (integer, float) or not.
func IsNumber(v any) bool { return IsInt(v) || IsFloat(v) }

// IsIP check if the string is a valid IP address.
// e.g 127.0.0.1, 2001:db8::68, 2001:db8::68:192.168.1.1
func IsIP(str string) bool { return net.ParseIP(str) != nil }

// IsIPV4 check if the string is a valid IPv4 address.
func IsIPV4(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}

	return ip.To4() != nil
}

// IsIPV6 check if the string is a valid IPv6 address.
func IsIPV6(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}

	return ip.To4() == nil && len(ip) == net.IPv6len
}

// IsPort check if the string is a valid port number [1 ~ 65536).
func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil &&
		i > 0 && i < 65536 {
		return true
	}

	return false
}

// IsURL check if the string is a valid URL.
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return urlMatcher.MatchString(str)
}

// IsBase64URL check if the string is a valid RUL-safe Base64 encoded string.
func IsBase64URL(v string) bool { return base64URLMatcher.MatchString(v) }

// IsDNS check if the string is a valid DNS name.
func IsDNS(str string) bool { return dnsMatcher.MatchString(str) }

// IsEmail check if the string is a valid email address.
func IsEmail(str string) bool {
	// FIXME patten of email
	emailMatcher.MatchString(str)

	_, err := mail.ParseAddress(str)
	return err == nil
}

// IsChineseMobile check if the string is a valid Chinese mobile phone number.
func IsChineseMobile(str string) bool { return chineseMobileMatcher.MatchString(str) }

// IsChineseIDNum check if the string is a valid Chinese ID card number.
func IsChineseIDNum(id string) bool {
	var (
		// Identity card formula
		factor = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		// ID verification bit
		verifyStr = [...]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
		// Starting year of ID card
		birthStartYear = 1900
		// Province code
		provinceKv = map[string]struct{}{
			"11": {},
			"12": {},
			"13": {},
			"14": {},
			"15": {},
			"21": {},
			"22": {},
			"23": {},
			"31": {},
			"32": {},
			"33": {},
			"34": {},
			"35": {},
			"36": {},
			"37": {},
			"41": {},
			"42": {},
			"43": {},
			"44": {},
			"45": {},
			"46": {},
			"50": {},
			"51": {},
			"52": {},
			"53": {},
			"54": {},
			"61": {},
			"62": {},
			"63": {},
			"64": {},
			"65": {},
			//"71": {},
			//"81": {},
			//"82": {},
		}
	)
	// All characters should be numbers, and the last digit can be either x or X
	if !chineseIDMatcher.MatchString(id) {
		return false
	}

	// Verify province codes and complete all province codes according to GB/T2260
	_, ok := provinceKv[id[0:2]]
	if !ok {
		return false
	}

	// Verify birthday, must be greater than birthStartYear and less than the current year
	birthStr := fmt.Sprintf("%s-%s-%s", id[6:10], id[10:12], id[12:14])
	birthday, err := time.Parse("2006-01-02", birthStr)
	if err != nil || birthday.After(time.Now()) || birthday.Year() < birthStartYear {
		return false
	}

	// Verification code
	sum := 0
	for i, c := range id[:17] {
		v, _ := strconv.Atoi(string(c))
		sum += v * factor[i]
	}

	return verifyStr[sum%11] == strings.ToUpper(id[17:18])
}

// ContainChinese check if the string contains Chinese characters.
func ContainChinese(str string) bool { return chineseMatcher.MatchString(str) }

// IsChinesePhone check if the string is chinese phone number.
// Valid chinese phone is xxx-xxxxxxxx or xxxx-xxxxxxx.
func IsChinesePhone(phone string) bool { return chinesePhoneMatcher.MatchString(phone) }

// IsCreditCard check if the string is credit card.
func IsCreditCard(creditCart string) bool { return creditCardMatcher.MatchString(creditCart) }

// IsBase64 check if the string is base64 string.
func IsBase64(base64 string) bool { return base64Matcher.MatchString(base64) }

// IsEmptyString check if the string is empty.
func IsEmptyString(str string) bool { return len(str) == 0 }

// IsRegexMatch check if the string match the regexp.
func IsRegexMatch(str, regex string) bool { return regexp.MustCompile(regex).MatchString(str) }

// IsStrongPassword check if the string is strong password,
// if len(password) is less than the length param, return false.
//
//	Strong password: alpha(lower+upper) + number + special chars(!@#$%^&*()?><).
func IsStrongPassword(str string, length int) bool {
	if len(str) < length {
		return false
	}

	var num, lower, upper, special bool
	for _, r := range str {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return num && lower && upper && special
}

// IsWeakPassword check if the string is weak password.
//
//	Weak password: only letter or only number or letter+number.
func IsWeakPassword(str string) bool {
	// return !IsStrongPassword(str, 8)

	var num, letter, special bool
	for _, r := range str {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return (num || letter) && !special
}

// IsJWT check if a string is a valid JSON Web Token (JWT).
func IsJWT(str string) bool {
	strs := strings.Split(str, ".")
	if len(strs) != 3 {
		return false
	}

	for _, s := range strs {
		if !IsBase64URL(s) {
			return false
		}
	}

	return true
}

// IsZeroValue check if the value is zero value.
func IsZeroValue(val any) bool {
	if val == nil {
		return true
	}

	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Func:
		return rv.IsNil()
	case reflect.Chan, reflect.Map:
		return rv.Len() == 0 || rv.IsNil()
	}

	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}

// IsGBK check if data encoding is gbk
// Note: this function is implemented by whether double bytes fall within the encoding range of gbk,
// while each byte of utf-8 encoding format falls within the encoding range of gbk.
// Therefore, utf8.valid() should be called first to check whether it is not utf-8 encoding,
// and then call IsGBK() to check gbk encoding. like below
/**
	data := []byte("你好")
	if utf8.Valid(data) {
		fmt.Println("data encoding is utf-8")
	}else if(IsGBK(data)) {
		fmt.Println("data encoding is GBK")
	}
	fmt.Println("data encoding is unknown")
**/
func IsGBK(data []byte) bool {
	for i := 0; i < len(data); i++ {
		// Check for single-byte ASCII characters (0x00-0x7F)
		if data[i] <= 0x7F {
			continue
		}

		// Check for the first byte of a double-byte character (0x81-0xFE)
		if data[i] >= 0x81 && data[i] <= 0xFE {
			// Ensure there is a next byte to form a valid double-byte character
			if i+1 < len(data) {
				// Check the second byte (0x40-0xFE)
				if data[i+1] >= 0x40 && data[i+1] <= 0xFE && data[i+1] != 0xF7 {
					i++ // Move to the next byte after the valid double-byte character
					continue
				}
			}
			return false
		}
		return false
	}
	return true
}

// IsVisa check if a string is a valid Visa card number.
func IsVisa(str string) bool { return visaMatcher.MatchString(str) }

// IsMasterCard check if a give string is a valid master card nubmer or not.
func IsMasterCard(v string) bool { return masterCardMatcher.MatchString(v) }

// IsAmericanExpress check if a give string is a valid american expression card nubmer or not.
func IsAmericanExpress(v string) bool { return americanExpressMatcher.MatchString(v) }

// IsUnionPay check if a give string is a valid union pay nubmer or not.
func IsUnionPay(v string) bool { return unionPay.MatchString(v) }

// IsChinaUnionPay check if a give string is a valid china union pay nubmer or not.
func IsChinaUnionPay(v string) bool { return chinaUnionPay.MatchString(v) }
