package strutil

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsAlpha(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsAlpha(""))

	// Only Letters
	assert.True(IsAlpha("HelloWorld"))

	// Contains Numbers
	assert.False(IsAlpha("Hello123"))

	// Contains Special Characters
	assert.False(IsAlpha("Hello!"))

	// Letters and Numbers
	assert.False(IsAlpha("Hello123World"))

	// Letters and Special Characters
	assert.False(IsAlpha("Hello!World"))

	// Only Numbers
	assert.False(IsAlpha("123"))
}

func TestIsAllUpper(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsAllUpper(""))

	// All Uppercase
	assert.True(IsAllUpper("HELLO"))

	// Mixed Case
	assert.False(IsAllUpper("HelloWorld"))

	// All Lowercase
	assert.False(IsAllUpper("hello"))

	// Numbers and Symbols
	assert.False(IsAllUpper("123!@#"))

	// Uppercase and Symbols
	assert.False(IsAllUpper("HELLO!"))

	// Only Numbers
	assert.False(IsAllUpper("123"))
}

func TestIsAllLower(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsAllLower(""))

	// All Lowercase
	assert.True(IsAllLower("hello"))

	// Mixed Case
	assert.False(IsAllLower("HelloWorld"))

	// All Uppercase
	assert.False(IsAllLower("HELLO"))

	// Numbers and Symbols
	assert.False(IsAllLower("123!@#"))

	// Lowercase and Symbols
	assert.False(IsAllLower("hello!"))

	// Only Numbers
	assert.False(IsAllLower("123"))
}

func TestIsASCII(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.True(IsASCII(""))

	// ASCII Only
	assert.True(IsASCII("Hello, world!"))

	// Extended ASCII
	assert.False(IsASCII("Hello, world! \u0080"))

	// Unicode Characters
	assert.False(IsASCII("こんにちは世界"))

	// Mixed ASCII and Unicode
	assert.False(IsASCII("Hello, 世界!"))
}

func TestIsPrint(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.True(IsPrint(""))

	// Printable ASCII
	assert.True(IsPrint("Hello, world!"))

	// Non-Printable ASCII
	assert.False(IsPrint("Hello, \u0007world!"))

	// Printable Unicode
	assert.True(IsPrint("こんにちは世界"))

	// Non-Printable Unicode
	assert.False(IsPrint("こんにちは\u0007世界"))

	// Newline
	assert.True(IsPrint("Hello\nWorld"))

	// Tab
	assert.True(IsPrint("Hello\tWorld"))

	// Carriage Return
	assert.True(IsPrint("Hello\rWorld"))

	// Backtick
	assert.True(IsPrint("Hello`World"))
}

func TestContainUpper(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(ContainUpper(""))

	// All Uppercase
	assert.True(ContainUpper("HELLO"))

	// Mixed Case
	assert.True(ContainUpper("HelloWorld"))

	// All Lowercase
	assert.False(ContainUpper("hello"))

	// Numbers and Symbols
	assert.False(ContainUpper("123!@#"))

	// Uppercase and Symbols
	assert.True(ContainUpper("HELLO!"))

	// Only Numbers
	assert.False(ContainUpper("123"))

	// Only Symbols
	assert.False(ContainUpper("!@#"))
}

func TestContainLower(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(ContainLower(""))

	// All Lowercase
	assert.True(ContainLower("hello"))

	// Mixed Case
	assert.True(ContainLower("HelloWorld"))

	// All Uppercase
	assert.False(ContainLower("HELLO"))

	// Numbers and Symbols
	assert.False(ContainLower("123!@#"))

	// Lowercase and Symbols
	assert.True(ContainLower("hello!"))

	// Only Numbers
	assert.False(ContainLower("123"))

	// Only Symbols
	assert.False(ContainLower("!@#"))
}

func TestContainLetter_string(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(ContainLetter(""))

	// String With Letters
	assert.True(ContainLetter("Hello"))

	// String With Numbers
	assert.False(ContainLetter("12345"))

	// String With Special Characters
	assert.False(ContainLetter("!@#$%"))

	// String With Letters and Numbers
	assert.True(ContainLetter("Hello123"))

	// String With Letters and Special Characters
	assert.True(ContainLetter("Hello!"))

	// Uincode String
	assert.False(ContainLetter("こんにちは世界"))
}

func TestContainLetter_vb(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Bytes With Letters
	assert.True(ContainLetter([]byte("Hello")))

	// Bytes With Numbers
	assert.False(ContainLetter([]byte("12345")))

	// Bytes With Special Characters
	assert.False(ContainLetter([]byte("!@#$%")))

	// Bytes With Letters and Numbers
	assert.True(ContainLetter([]byte("Hello123")))

	// Bytes With Letters and Special Characters
	assert.True(ContainLetter([]byte("Hello!")))

	// Uincode String
	assert.False(ContainLetter([]byte("こんにちは世界")))
}

func TestContainNumber_string(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(ContainNumber(""))

	// String With Letters
	assert.False(ContainNumber("Hello"))

	// String With Numbers
	assert.True(ContainNumber("12345"))

	// String With Special Characters
	assert.False(ContainNumber("!@#$%"))

	// String With Letters and Numbers
	assert.True(ContainNumber("Hello123"))

	// String With Letters and Special Characters
	assert.False(ContainNumber("Hello!"))

	// Uincode String
	assert.False(ContainNumber("こんにちは世界"))
}

func TestContainNumber_vb(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Bytes With Letters
	assert.False(ContainNumber([]byte("Hello")))

	// Bytes With Numbers
	assert.True(ContainNumber([]byte("12345")))

	// Bytes With Special Characters
	assert.False(ContainNumber([]byte("!@#$%")))

	// Bytes With Letters and Numbers
	assert.True(ContainNumber([]byte("Hello123")))

	// Bytes With Letters and Special Characters
	assert.False(ContainNumber([]byte("Hello!")))

	// Uincode String
	assert.False(ContainNumber([]byte("こんにちは世界")))
}

func TestIsJSON(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsJSON(""))

	// Valid JSON Object
	assert.True(IsJSON(`{"name":"John", "age":30}`))

	// Valid JSON Array
	assert.True(IsJSON(`{"a":[1, 2, 3, 4]}`))

	// Invalid JSON
	assert.False(IsJSON(`{"name":"John", "age":30,}`))

	// Not JSON
	assert.False(IsJSON("Hello, world!"))

	// JSON with Spaces
	assert.True(IsJSON(` { "name" : "John" , "age" : 30 } `))

	// JSON with Newlines
	assert.True(IsJSON(`{
		"name": "John",
		"age": 30
	}`))

	// JSON with Tabs
	assert.True(IsJSON(`{
		"name": "John",
		"age": 30
	}`))

	// JSON with Carriage Returns
	assert.True(IsJSON(`{
		"name": "John",
		"age": 30
	}`))
}

func TestIsFloatStr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsFloatStr(""))

	// Valid Float
	assert.True(IsFloatStr("123.45"))

	// Invalid Float
	assert.False(IsFloatStr("123.45.67"))

	// Integer as Float
	assert.True(IsFloatStr("123"))

	// Non-numeric
	assert.False(IsFloatStr("abc"))

	// Scientific Notation
	assert.True(IsFloatStr("1.23e4"))

	// Negative Float
	assert.True(IsFloatStr("-123.45"))

	// Leading Zeros
	assert.True(IsFloatStr("0.123"))

	// Trailing Zeros
	assert.True(IsFloatStr("123.4500"))

	// Leading Trailing Zeros
	assert.True(IsFloatStr("0.1234500"))

	// Trailing Scientific Notation
	assert.True(IsFloatStr("123.45e-2"))
}

func TestIsIntStr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsIntStr(""))

	// Valid Integer
	assert.True(IsIntStr("123"))

	// Invalid Integer
	assert.False(IsIntStr("123.45"))

	// Negative Integer
	assert.True(IsIntStr("-123"))

	// Leading Zeros
	assert.True(IsIntStr("00123"))

	// Non-numeric
	assert.False(IsIntStr("abc"))

	// Scientific Notation
	assert.False(IsIntStr("1.23e4"))

	// Trailing Scientific Notation
	assert.False(IsIntStr("123.45e-2"))

	// Negative Number
	assert.False(IsIntStr("-123.45"))

	// Leading Zeros
	assert.True(IsIntStr("00123"))

	// Alpha Numeric
	assert.False(IsIntStr("a123"))
}

func TestIsNumberStr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsNumberStr(""))

	// Valid Integer
	assert.True(IsNumberStr("123"))

	// Valid Float
	assert.True(IsNumberStr("123.45"))

	// Invalid Number
	assert.False(IsNumberStr("123.45.67"))

	// Non-numeric
	assert.False(IsNumberStr("abc"))

	// Scientific Notation
	assert.True(IsNumberStr("1.23e4"))

	// Trailing Scientific Notation
	assert.True(IsNumberStr("123.45e-2"))

	// Negative Number
	assert.True(IsNumberStr("-123"))

	// Leading Zeros
	assert.True(IsNumberStr("00123"))

	// Alpha Numeric
	assert.False(IsNumberStr("a123"))
}

func TestIsIP(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsIP(""))

	// Valid IPv4
	assert.True(IsIP("192.168.1.1"))

	// Valid IPv6
	assert.True(IsIP("2001:db8::68"))

	// Invalid IP
	assert.False(IsIP("256.256.256.256"))

	// Non-IP String
	assert.False(IsIP("not an IP"))

	// Mixed IPv4 and IPv6
	assert.True(IsIP("2001:db8::68:192.168.1.1"))

	// IPv4 with Port
	assert.False(IsIP("192.168.1.1:80"))
}

func TestIsIPV4(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsIPV4(""))

	// Valid IPv4
	assert.True(IsIPV4("192.168.1.1"))

	// Valid IPv6
	assert.False(IsIPV4("2001:db8::68"))

	// Invalid IP
	assert.False(IsIPV4("256.256.256.256"))

	// Non-IP String
	assert.False(IsIPV4("not an IP"))

	// IPv4 with Port
	assert.False(IsIPV4("192.168.1.1:80"))

	// IPv6 with Port
	assert.False(IsIPV4("2001:db8::68:80"))
}

func TestIsIPV6(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsIPV6(""))

	// Valid IPv4
	assert.False(IsIPV6("192.168.1.1"))

	// Valid IPv6
	assert.True(IsIPV6("2001:db8::68"))

	// Invalid IP
	assert.False(IsIPV6("256.256.256.256"))

	// Non-IP String
	assert.False(IsIPV6("not an IP"))
}

func TestIsPort(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsPort(""))

	// Valid Port
	assert.True(IsPort("8080"))

	// Invalid Port
	assert.False(IsPort("65537"))
	assert.False(IsPort("-1"))

	// Non-numeric
	assert.False(IsPort("abc"))

	// Leading Zero
	assert.True(IsPort("08080"))
}

func TestIsURL(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsURL(""))

	// Too Short
	assert.False(IsURL("a"))

	// Too Long
	assert.False(IsURL(strings.Repeat("a", 2084)))

	// Starts with Dot
	assert.False(IsURL(".com"))

	// Invalid URL
	assert.False(IsURL("example"))

	// Valid URL
	assert.True(IsURL("http://kyden.us.kg"))

	// Host starts with Dot
	assert.False(IsURL("http://.kyden.us.kg"))

	// No Host, No Path
	assert.False(IsURL("http://"))

	// No Host, Path without Dot
	assert.False(IsURL("http:///path"))

	// Query String
	assert.True(IsURL("http://kyden.us.kg/?query=1"))

	// Fragment Identifier
	assert.True(IsURL("http://kyden.us.kg/#fragment"))
}

func TestIsDNS(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsDNS(""))

	// Valid DNS
	assert.True(IsDNS("kyden.us.kg"))
	assert.True(IsDNS("kyden1234.us.kg"))
	assert.True(IsDNS("kyden-top.us.kg"))
	assert.True(IsDNS("xn--fros-gra-6qa.com")) // Punycode encoded "frös-grä.com"

	// Invalid DNS
	assert.False(IsDNS("kyden.us.kg."))
	assert.False(IsDNS("kyden.us.kg.."))
	assert.False(IsDNS("kyden.us.kg."))
	assert.False(IsDNS("kyden.us.kg-"))
	assert.False(IsDNS("kyden.us.kg_"))
	assert.False(IsDNS("kyden.us.kg$"))

	assert.False(IsDNS("-kyden.us.kg"))
	assert.False(IsDNS("kyden-.us.kg"))

	// FIXME
	// assert.False(IsDNS("ky--den.us.kg"))
	// assert.False(IsDNS(strings.Repeat("a", 63) + ".com"))
	// assert.False(IsDNS(strings.Repeat("a", 64) + ".com"))
}

func TestIsEmail(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsEmail(""))

	// Valid Email
	assert.True(IsEmail("kytedance@gmail.com"))
	assert.True(IsEmail("kytedance+123@gmail.com"))
	assert.True(IsEmail("kytedance_123@gmail.com"))
	assert.True(IsEmail("kytedance.123@gmail.com"))
	assert.True(IsEmail("kytedance.123.456@gmail.com"))
	assert.True(IsEmail("kytedance@sub.example.com"))

	// Invalid Email
	assert.False(IsEmail("kytedancegmail.com"))
	assert.False(IsEmail(".kytedance@gmail"))
	assert.False(IsEmail("kytedance@gmail."))
	assert.False(IsEmail("kytedance..@gmail"))
	assert.False(IsEmail("kytedance!@#$%^&*()@gmail"))

	// FIXME
	// assert.False(IsEmail("kytedance@gmail"))

	// FIXME
	// {"Invalid Email - Local Part Too Long", strings.Repeat("a", 65) + "@example.com", false},
	// {"Invalid Email - Domain Too Long", "test@" + strings.Repeat("a", 256) + ".com", false},
}

func TestIsChineseMobile(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsChineseMobile(""))

	// Valid Mobile Number
	assert.True(IsChineseMobile("13800138000"))
	assert.True(IsChineseMobile("13800138001"))
	assert.True(IsChineseMobile("13800138002"))
	assert.True(IsChineseMobile("13800138003"))
	assert.True(IsChineseMobile("13800138004"))
	assert.True(IsChineseMobile("13800138005"))
	assert.True(IsChineseMobile("13800138006"))
	assert.True(IsChineseMobile("13800138007"))
	assert.True(IsChineseMobile("13800138008"))
	assert.True(IsChineseMobile("13800138009"))

	// Invalid Mobile Number - Too Short
	assert.False(IsChineseMobile("1380013800"))

	// Invalid Mobile Number - Too Long
	assert.False(IsChineseMobile("138001380000"))

	// Invalid Mobile Number - Invalid Prefix
	assert.False(IsChineseMobile("23800138000"))

	// Invalid Mobile Number - Contains Letters
	assert.False(IsChineseMobile("138001A8000"))

	// Invalid Mobile Number - Contains Special Characters
	assert.False(IsChineseMobile("13800138000!"))
}

func TestIsChineseIDNum(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsChineseIDNum(""))

	// Valid ID
	assert.True(IsChineseIDNum("34052419800101001X"))
	assert.True(IsChineseIDNum("34052419800101001x"))

	// Invalid ID - Wrong Checksum
	assert.False(IsChineseIDNum("340524198001010011"))

	// Invalid ID - Invalid Province Code
	assert.False(IsChineseIDNum("99052419800101001X"))

	// Invalid ID - Invalid Birthday
	assert.False(IsChineseIDNum("34052420301301001X"))

	// Invalid ID - Future Birthday
	assert.False(IsChineseIDNum("340524" + fmt.Sprintf("%02d", time.Now().Year()+1) + "0101001X"))

	// Invalid ID - Too Short
	assert.False(IsChineseIDNum("34052419800101001"))

	// Invalid ID - Too Long
	assert.False(IsChineseIDNum("34052419800101001XX"))

	// Invalid ID - Non-numeric Characters
	assert.False(IsChineseIDNum("340524198A0101001X"))
}

func TestContainChinese(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(ContainChinese(""))

	// No Chinese Characters
	assert.False(ContainChinese("HelloWorld"))

	// Contains Chinese Characters
	assert.True(ContainChinese("Hello世界"))

	// Mixed Chinese and ASCII Characters
	assert.True(ContainChinese("你好Hello"))

	// Only Chinese Characters
	assert.True(ContainChinese("你好"))

	// Special Characters
	assert.False(ContainChinese("!@#$%^&*()"))
}

func TestIsChinesePhone(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsChinesePhone(""))

	// Valid Phone Number
	assert.True(IsChinesePhone("123-45678901"))
	assert.True(IsChinesePhone("1234-5678901"))

	// Invalid Phone Number - Too Short
	assert.False(IsChinesePhone("1380013800"))

	// Invalid Phone Number - Too Long
	assert.False(IsChinesePhone("138001380000"))

	// Invalid Phone Number - Invalid Prefix
	assert.False(IsChinesePhone("23800138000"))

	// Invalid Phone Number - Contains Letters
	assert.False(IsChinesePhone("138001A8000"))

	// Invalid Phone Number - Contains Special Characters
	assert.False(IsChinesePhone("13800138000!"))

	// FIXME
	// assert.True(IsChinesePhone("13800138001"))
}

func TestIsCreditCard(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsCreditCard(""))

	// Valid Credit Card
	assert.True(IsCreditCard("4111111111111111"))

	// Invalid Format
	assert.False(IsCreditCard("1234567890123456"))

	// Non-Digit Characters
	assert.False(IsCreditCard("1234-abcd-5678-efgh"))
}

func TestIsBase64(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsBase64(""))

	// Valid Base64
	assert.True(IsBase64("SGVsbG8gV29ybGQh"))

	// Invalid Base64
	assert.False(IsBase64("Hello World!"))

	// Non-Base64 Characters
	assert.False(IsBase64("Hello, 世界!"))

	// Invalid Length
	assert.False(IsBase64("Hello, 世界!"))
}

func TestIsEmptyString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.True(IsEmptyString(""))

	// Non-Empty String
	assert.False(IsEmptyString("Hello, World!"))
}

func TestIsRegexMatch(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.True(IsRegexMatch("", ""))
	assert.True(IsRegexMatch("", "^$"))

	// Match
	assert.True(IsRegexMatch("Hello, World!", "^Hello"))
	assert.True(IsRegexMatch("Hello, World!", "Hello.*"))
	assert.True(IsRegexMatch("Hello, World!", "World!$"))

	// Not Match
	assert.False(IsRegexMatch("Hello, World!", "^Goodbye"))

	// Regex Error
	assert.False(IsRegexMatch("Hello, World!", `\[`))
}

func TestIsStrongPassword(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsStrongPassword("", 8))

	// Less Than Length
	assert.False(IsStrongPassword("pass", 8))

	// No Number
	assert.False(IsStrongPassword("Password!", 8))

	// No Lowercase
	assert.False(IsStrongPassword("PASSWORD1", 8))

	// No Uppercase
	assert.False(IsStrongPassword("password1", 8))

	// No Special Char
	assert.False(IsStrongPassword("Password1", 8))

	// Strong Password
	assert.True(IsStrongPassword("Password1!", 8))

	// Strong Password with Length 6
	assert.True(IsStrongPassword("Pass1!", 6))

	// Strong Password with Length 10
	assert.True(IsStrongPassword("Password1!", 10))
}

func TestIsWeakPassword(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsWeakPassword(""))

	// Only letters
	assert.True(IsWeakPassword("abcdef"))

	// Only numbers
	assert.True(IsWeakPassword("123456"))

	// Letters and numbers
	assert.True(IsWeakPassword("abc123"))

	// Contains special character
	assert.False(IsWeakPassword("abc!123"))

	// Too short
	assert.True(IsWeakPassword("a"))

	// Strong password
	assert.False(IsWeakPassword("abcABC123!"))

	// Weak password
	assert.True(IsWeakPassword("abc123"))
}

func TestIsZeroValue(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// nil
	assert.True(IsZeroValue(nil))

	// empty string
	assert.True(IsZeroValue(""))

	// false boolean
	assert.True(IsZeroValue(false))
	assert.False(IsZeroValue(true))

	// int zero
	assert.True(IsZeroValue(0))
	assert.False(IsZeroValue(1))

	// int8 zero
	assert.True(IsZeroValue(int8(0)))
	assert.False(IsZeroValue(int8(1)))

	// int16 zero
	assert.True(IsZeroValue(int16(0)))
	assert.False(IsZeroValue(int16(1)))

	// int32 zero
	assert.True(IsZeroValue(int32(0)))
	assert.False(IsZeroValue(int32(1)))

	// int64 zero
	assert.True(IsZeroValue(int64(0)))
	assert.False(IsZeroValue(int64(1)))

	// uint zero
	assert.True(IsZeroValue(uint(0)))
	assert.False(IsZeroValue(uint(1)))

	// uint8 zero
	assert.True(IsZeroValue(uint8(0)))
	assert.False(IsZeroValue(uint8(1)))

	// uint16 zero
	assert.True(IsZeroValue(uint16(0)))
	assert.False(IsZeroValue(uint16(1)))

	// uint32 zero
	assert.True(IsZeroValue(uint32(0)))
	assert.False(IsZeroValue(uint32(1)))

	// uint64 zero
	assert.True(IsZeroValue(uint64(0)))
	assert.False(IsZeroValue(uint64(1)))

	// float32 zero
	assert.True(IsZeroValue(float32(0)))
	assert.False(IsZeroValue(float32(1)))

	// float64 zero
	assert.True(IsZeroValue(float64(0)))
	assert.False(IsZeroValue(float64(1)))

	// nil pointer
	assert.True(IsZeroValue((*int)(nil)))
	assert.True(IsZeroValue(new(int)))

	// nil interface
	assert.True(IsZeroValue(new(interface{})))

	// nil slice
	assert.True(IsZeroValue([]int(nil)))

	// complex type zero
	assert.True(IsZeroValue(complex(0, 0)))
	assert.False(IsZeroValue(complex(1, 1)))

	// nil channel
	assert.True(IsZeroValue(make(chan int)))
	// zero len channel
	assert.True(IsZeroValue(make(chan int, 0)))

	// nil function
	assert.True(IsZeroValue(nil))

	// nil map
	assert.True(IsZeroValue(make(map[int]any)))
	// zero len map
	assert.True(IsZeroValue(make(map[string]int)))
}

func TestIsGBK(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.True(IsGBK([]byte("")))

	// Single byte
	assert.True(IsGBK([]byte{0x41}))
	assert.False(IsGBK([]byte{0x80}))

	// GBK range limits
	assert.True(IsGBK([]byte("\x81\x40\xfe\xfe")))
	assert.False(IsGBK([]byte("\x81\x40\xfe\xfe\x80")))

	// Invalid GBK start
	assert.False(IsGBK([]byte("\x80\x40")))
	// Invalid GBK end
	assert.False(IsGBK([]byte("\xfe\x3f")))
	// Invalid GBK with F7
	assert.False(IsGBK([]byte("\x81\xf7")))

	// Valid GBK
	assert.True(IsGBK([]byte("你好")))
	assert.True(IsGBK([]byte("kyden")))
}

func TestIsFloat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsFloat(""))

	// Valid Float
	assert.True(IsFloat(1.0))
	assert.True(IsFloat(float32(1.0)))
	assert.True(IsFloat(float64(1.0)))

	// Not Float
	assert.False(IsFloat(1))
	assert.False(IsFloat("1.0"))
	assert.False(IsFloat("not a float"))
	assert.False(IsFloat(nil))
	assert.False(IsFloat(make(chan any)))
	assert.False(IsFloat(make(map[any]any)))
	assert.False(IsFloat(make([]any, 0)))
	assert.False(IsFloat(func() {}))
}

func TestIsInt(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Valid Int
	assert.True(IsInt(1))
	assert.True(IsInt(int8(1)))
	assert.True(IsInt(int16(1)))
	assert.True(IsInt(int32(1)))
	assert.True(IsInt(int64(1)))
	assert.True(IsInt(uint(1)))
	assert.True(IsInt(uint8(1)))
	assert.True(IsInt(uint16(1)))
	assert.True(IsInt(uint32(1)))
	assert.True(IsInt(uint64(1)))
	assert.True(IsInt(uintptr(1)))

	// Empty String
	assert.False(IsInt(""))

	// Invalid int
	assert.False(IsInt(float32(1.0)))
	assert.False(IsInt(float64(1.0)))
	assert.False(IsInt(1.1))               // float
	assert.False(IsInt("1"))               // string
	assert.False(IsInt("not an int"))      // string
	assert.False(IsInt(nil))               // nil
	assert.False(IsInt(make(chan any)))    // channel
	assert.False(IsInt(make(map[any]any))) // map
	assert.False(IsInt(make([]any, 0)))    // slice
	assert.False(IsInt(func() {}))         // func
}

func TestIsNumber(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsNumber(""))

	// Valid Number
	assert.True(IsNumber(1))
	assert.True(IsNumber(int8(1)))
	assert.True(IsNumber(int16(1)))
	assert.True(IsNumber(int32(1)))
	assert.True(IsNumber(int64(1)))
	assert.True(IsNumber(uint(1)))
	assert.True(IsNumber(uint8(1)))
	assert.True(IsNumber(uint16(1)))
	assert.True(IsNumber(uint32(1)))
	assert.True(IsNumber(uint64(1)))
	assert.True(IsNumber(uintptr(1)))
	assert.True(IsNumber(1.0))
	assert.True(IsNumber(float32(1.0)))
	assert.True(IsNumber(float64(1.0)))
	assert.True(IsNumber(1.1))

	// Invalid Number
	assert.False(IsNumber("1"))               // string
	assert.False(IsNumber("not a number"))    // string
	assert.False(IsNumber(nil))               // nil
	assert.False(IsNumber(make(chan any)))    // channel
	assert.False(IsNumber(make(map[any]any))) // map
	assert.False(IsNumber(make([]any, 0)))    // slice
	assert.False(IsNumber(func() {}))         // func
}

func TestIsBin(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsBin(""))

	// All zeros
	assert.True(IsBin("0000"))
	assert.True(IsBin("00000000"))

	// All ones
	assert.True(IsBin("1111"))
	assert.True(IsBin("11111111"))

	// Mixed zeros and ones
	assert.True(IsBin("101010"))
	assert.True(IsBin("1010101010101010"))

	// Invalid characters
	assert.False(IsBin("102010"))
	assert.False(IsBin("103010"))
}

func TestIsHex(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty String
	assert.False(IsHex(""))

	// Valid Hex
	assert.True(IsHex("1A3FBC90"))
	assert.True(IsHex("0x1A3FBC90"))

	// Invalid Hex
	assert.False(IsHex("1A3FBC9G"))
	assert.False(IsHex("0x1A3FBC9G"))
	assert.False(IsHex("1A3FBC9!"))
	assert.False(IsHex("0x1A3FBC9!"))

	// Invalid leading characters
	assert.False(IsHex("x1A3FBC90"))
	assert.False(IsHex("X1A3FBC90"))

	// Invalid trailing characters
	assert.False(IsHex("1A3FBC90x"))
	assert.False(IsHex("1A3FBC90X"))
	assert.False(IsHex("1A3FBC90!"))
	assert.False(IsHex("1A3FBC90#"))
}

func TestIsBase64URL(t *testing.T) {
	// Valid Base64URL strings
	validCases := []string{
		"aGVsbG8gd29ybGQ=",
		"YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo-", // "abcdefghijklmnopqrstuvwxyz" encoded
		"U29tZSBzdHJpbmc=",                     // "Some string" encoded
	}

	for _, v := range validCases {
		if !IsBase64URL(v) {
			t.Errorf("Expected true for %s, got false", v)
		}
	}

	// Invalid Base64URL strings
	invalidCases := []string{
		"not a base64 url string",
		// "Zm9vYmFy",   // FIXME Missing padding '='
		"Zm9vYmFya",
		"Zm9vYmFy=",  // Invalid padding '=' at the end
		"Zm9vYmFy?=", // Contains invalid character '?'
	}

	for _, v := range invalidCases {
		if IsBase64URL(v) {
			t.Errorf("Expected false for %s, got true", v)
		}
	}

	// Edge cases
	edgeCases := []string{
		// "", // FIXME empty string
		"=",
		"==",
		"===",
		"Zm9vYmFy?=", // Contains invalid character '?'
	}

	for _, v := range edgeCases {
		if IsBase64URL(v) {
			t.Errorf("Expected false for %s, got true", v)
		}
	}
}

func TestIsJWT(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsJWT(""))

	// Valid JWT
	assert.True(IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"))

	// Invalid JWT - incorrect segment count
	assert.False(IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"))

	// Invalid JWT - not Base64URL encoded
	assert.False(IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV/adQssw5c."))

	// Invalid JWT - Too short to be a JWT
	assert.False(IsJWT("a.b.c"))
}

func TestIsVisa(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsVisa(""))

	// Valid Visa card number
	assert.True(IsVisa("4111111111111111"))
	assert.True(IsVisa("4012888888881881"))

	// Invalid length
	assert.False(IsVisa("401288888888188"))

	// Invalid characters
	assert.False(IsVisa("401288888888188a"))

	// Invalid prefix
	assert.False(IsVisa("3412888888881881"))
}

func TestIsMasterCard(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsMasterCard(""))

	// Valid MasterCard card number
	assert.True(IsMasterCard("5555555555554444"))
	assert.True(IsMasterCard("5105105105105100"))

	// Invalid length
	assert.False(IsMasterCard("510510510510510"))
}

func TestIsAmericanExpress(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsAmericanExpress(""))

	// Valid American Express card number
	assert.True(IsAmericanExpress("378282246310005"))
	assert.True(IsAmericanExpress("371449635398431"))

	// Invalid length
	assert.False(IsAmericanExpress("3714496353984"))

	// Invalid characters
	assert.False(IsAmericanExpress("371449635398431a"))

	// Invalid prefix
	assert.False(IsAmericanExpress("5555555555554444"))
}

func TestIsUnionPay(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsUnionPay(""))

	// Valid UnionPay
	assert.True(IsUnionPay("6211514433542201"))

	// Invalid length
	assert.False(IsUnionPay("621151443354"))

	// Invalid characters
	assert.False(IsUnionPay("621151443354220a"))

	// Invalid prefix
	assert.False(IsUnionPay("5555555555554444"))
}

func TestIsChinaUnionPay(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty string
	assert.False(IsChinaUnionPay(""))

	// Valid China UnionPay
	assert.True(IsChinaUnionPay("6222021001122334"))
	assert.True(IsChinaUnionPay("62220210011223344"))
	assert.True(IsChinaUnionPay("622202100112233445"))
	assert.True(IsChinaUnionPay("6222021001122334456"))

	// Invalid length
	assert.False(IsChinaUnionPay("622202100112233"))
	assert.False(IsChinaUnionPay("62220210011223344567"))
	assert.False(IsChinaUnionPay("622202100112233445678"))
}
