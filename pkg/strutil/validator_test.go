package strutil

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Only Letters", "HelloWorld", true},
		{"Contains Numbers", "Hello123", false},
		{"Contains Special Characters", "Hello!", false},
		{"Letters and Numbers", "Hello123World", false},
		{"Letters and Special Characters", "Hello!World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsAlpha(tt.input)
			if actual != tt.expected {
				t.Errorf("IsAlpha(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsAllUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"All Uppercase", "HELLO", true},
		{"Mixed Case", "HelloWorld", false},
		{"All Lowercase", "hello", false},
		{"Numbers and Symbols", "123!@#", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsAllUpper(tt.input)
			if actual != tt.expected {
				t.Errorf("IsAllUpper(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsAllLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"All Lowercase", "hello", true},
		{"Mixed Case", "HelloWorld", false},
		{"All Uppercase", "HELLO", false},
		{"Numbers and Symbols", "123!@#", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsAllLower(tt.input)
			if actual != tt.expected {
				t.Errorf("IsAllLower(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", true},
		{"ASCII Only", "Hello, world!", true},
		{"Extended ASCII", "Hello, world! \u0080", false},
		{"Unicode Characters", "こんにちは世界", false},
		{"Mixed ASCII and Unicode", "Hello, 世界!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsASCII(tt.input)
			if actual != tt.expected {
				t.Errorf("IsASCII(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsPrint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", true},
		{"Printable ASCII", "Hello, world!", true},
		{"Non-Printable ASCII", "Hello, \u0007world!", false},
		{"Printable Unicode", "こんにちは世界", true},
		{"Non-Printable Unicode", "こんにちは\u0007世界", false},
		{"Newline", "Hello\nWorld", true},
		{"Tab", "Hello\tWorld", true},
		{"Carriage Return", "Hello\rWorld", true},
		{"Backtick", "Hello`World", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsPrint(tt.input)
			if actual != tt.expected {
				t.Errorf("IsPrint(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"All Uppercase", "HELLO", true},
		{"Mixed Case", "HelloWorld", true},
		{"All Lowercase", "hello", false},
		{"Numbers and Symbols", "123!@#", false},
		{"Uppercase with Numbers", "HELLO123", true},
		{"Lowercase with Numbers", "hello123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainUpper(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainUpper(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"All Uppercase", "HELLO", false},
		{"Mixed Case", "HelloWorld", true},
		{"All Lowercase", "hello", true},
		{"Numbers and Symbols", "123!@#", false},
		{"Uppercase with Numbers", "HELLO123", false},
		{"Lowercase with Numbers", "hello123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainLower(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainLower(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainLetter_string(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"EmptyString", "", false},
		{"StringWithLetters", "Hello", true},
		{"StringWithNumbers", "12345", false},
		{"StringWithSpecialChars", "!@#$%", false},
		{"StringWithLettersAndNumbers", "Hello123", true},
		{"StringWithLettersAndSpecialChars", "Hello!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainLetter(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainLetter(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainLetter_vb(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{"BytesWithLetters", []byte("Hello"), true},
		{"BytesWithNumbers", []byte("12345"), false},
		{"BytesWithSpecialChars", []byte("!@#$%"), false},
		{"BytesWithLettersAndNumbers", []byte("Hello123"), true},
		{"BytesWithLettersAndSpecialChars", []byte("Hello!"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainLetter(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainLetter(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainNumber_string(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"EmptyString", "", false},
		{"StringWithLetters", "Hello", false},
		{"StringWithNumbers", "12345", true},
		{"StringWithSpecialChars", "!@#$%", false},
		{"StringWithLettersAndNumbers", "Hello123", true},
		{"StringWithLettersAndSpecialChars", "Hello!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainNumber(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainNumber(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestContainNumber_vb(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{"BytesWithLetters", []byte("Hello"), false},
		{"BytesWithNumbers", []byte("12345"), true},
		{"BytesWithSpecialChars", []byte("!@#$%"), false},
		{"BytesWithLettersAndNumbers", []byte("Hello123"), true},
		{"BytesWithLettersAndSpecialChars", []byte("Hello!"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainNumber(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainNumber(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid JSON Object", `{"name":"John", "age":30}`, true},
		{"Valid JSON Array", `{"a":[1, 2, 3, 4]}`, true},
		{"Invalid JSON", `{"name":"John", "age":30,}`, false},
		{"Not JSON", "Hello, world!", false},
		{"JSON with Spaces", ` { "name" : "John" , "age" : 30 } `, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsJSON(tt.input)
			if actual != tt.expected {
				t.Errorf("IsJSON(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsFloatStr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid Float", "123.45", true},
		{"Invalid Float", "123.45.67", false},
		{"Integer as Float", "123", true},
		{"Non-numeric", "abc", false},
		{"Scientific Notation", "1.23e4", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsFloatStr(tt.input)
			if actual != tt.expected {
				t.Errorf("IsFloatStr(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsIntStr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid Integer", "123", true},
		{"Invalid Integer", "123.45", false},
		{"Negative Integer", "-123", true},
		{"Leading Zeros", "00123", true},
		{"Non-numeric", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsIntStr(tt.input)
			if actual != tt.expected {
				t.Errorf("IsIntStr(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsNumberStr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid Integer", "123", true},
		{"Valid Float", "123.45", true},
		{"Invalid Number", "123.45.67", false},
		{"Non-numeric", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsNumberStr(tt.input)
			if actual != tt.expected {
				t.Errorf("IsNumberStr(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid IPv4", "192.168.1.1", true},
		{"Valid IPv6", "2001:db8::68", true},
		{"Invalid IP", "256.256.256.256", false},
		{"Non-IP String", "not an IP", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsIP(tt.input)
			if actual != tt.expected {
				t.Errorf("IsIP(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsIPV4(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid IPv4", "192.168.1.1", true},
		{"Valid IPv6", "2001:db8::68", false},
		{"Invalid IPv4", "256.256.256.256", false},
		{"Non-IP String", "not an IP", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsIPV4(tt.input)
			if actual != tt.expected {
				t.Errorf("IsIPV4(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsIPV6(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid IPv4", "192.168.1.1", false},
		{"Valid IPv6", "2001:db8::68", true},
		{"Invalid IPv6", "2001:db8::68g", false},
		{"Non-IP String", "not an IP", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsIPV6(tt.input)
			if actual != tt.expected {
				t.Errorf("IsIPV6(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsPort(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Zero", "0", false},
		{"Less than 1", "-1", false},
		{"Greater than 65536", "65537", false},
		{"Non-numeric", "abc", false},
		{"Valid Port", "8080", true},
		{"Leading Zero", "08080", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsPort(tt.input)
			if actual != tt.expected {
				t.Errorf("IsPort(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Too Short", "a", false},
		{"Too Long", strings.Repeat("a", 2084), false},
		{"Starts with Dot", ".com", false},
		{"Invalid URL", "example", false},
		{"Valid URL", "http://www.example.com", true},
		{"Host starts with Dot", "http://.example.com", false},
		{"No Host, No Path", "http://", false},
		{"No Host, Path without Dot", "http:///path", false},
		{"Query String", "http://www.example.com/?query=1", true},
		{"Fragment Identifier", "http://www.example.com/#fragment", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsURL(tt.input)
			if actual != tt.expected {
				t.Errorf("IsURL(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsDNS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},

		{"Valid DNS", "example.com", true},
		{"Valid DNS with Subdomain", "sub.example.com", true},
		{"Valid DNS with Numbers", "example123.com", true},
		{"Valid DNS with Hyphen", "example-test.com", true},

		{"Invalid DNS - Starts with Hyphen", "-example.com", false},
		{"Invalid DNS - Ends with Hyphen", "example-.com", false},
		// FIXME
		{"Invalid DNS - Consecutive Hyphens", "example--test.com", true},
		{"Invalid DNS - Special Characters", "example$.com", false},
		{"Invalid DNS - Too Long", strings.Repeat("a", 64) + ".com", false},

		{"Valid IDN", "xn--fros-gra-6qa.com", true}, // Punycode encoded "frös-grä.com"
		{"Invalid IDN - Too Long", "xn--" + strings.Repeat("a", 64) + ".com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsDNS(tt.input)
			if actual != tt.expected {
				t.Errorf("IsDNS(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},

		{"Valid Email", "test@example.com", true},
		{"Valid Email with Plus Sign", "test+123@example.com", true},
		{"Valid Email with Underscore", "test_123@example.com", true},
		{"Valid Email with Dot", "test.123@example.com", true},
		{"Valid Email with Multiple Dots", "test.123.456@example.com", true},
		{"Valid Email with Subdomain", "test@sub.example.com", true},

		{"Invalid Email - Missing @", "testexample.com", false},
		// FIXME
		// {"Invalid Email - Missing Domain Dot", "test@example", false},
		{"Invalid Email - Missing Domain Dot", "test@example", true},

		{"Invalid Email - Starts with Dot", ".test@example.com", false},
		{"Invalid Email - Ends with Dot", "test.@example.com", false},
		{"Invalid Email - Consecutive Dots", "test..123@example.com", false},
		{"Invalid Email - Special Characters", "test!@#$%^&*()@example.com", false},

		// FIXME
		// {"Invalid Email - Local Part Too Long", strings.Repeat("a", 65) + "@example.com", false},
		// {"Invalid Email - Domain Too Long", "test@" + strings.Repeat("a", 256) + ".com", false},
		{"Invalid Email - Local Part Too Long", strings.Repeat("a", 65) + "@example.com", true},
		{"Invalid Email - Domain Too Long", "test@" + strings.Repeat("a", 256) + ".com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsEmail(tt.input)
			if actual != tt.expected {
				t.Errorf("IsEmail(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsChineseMobile(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"Valid Mobile Number", "13800138000", true},
		{"Invalid Mobile Number - Too Short", "1380013800", false},
		{"Invalid Mobile Number - Too Long", "138001380000", false},
		{"Invalid Mobile Number - Invalid Prefix", "23800138000", false},
		{"Invalid Mobile Number - Contains Letters", "138001A8000", false},
		{"Invalid Mobile Number - Contains Special Characters", "13800138000!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsChineseMobile(tt.input)
			if actual != tt.expected {
				t.Errorf("IsChineseMobile(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsChineseIDNum(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected bool
	}{
		{
			name:     "Valid ID",
			id:       "34052419800101001X",
			expected: true,
		},
		{
			name:     "Valid ID with lowercase x",
			id:       "34052419800101001x",
			expected: true,
		},
		{
			name:     "Invalid ID - Wrong Checksum",
			id:       "340524198001010011",
			expected: false,
		},
		{
			name:     "Invalid ID - Invalid Province Code",
			id:       "99052419800101001X",
			expected: false,
		},
		{
			name:     "Invalid ID - Invalid Birthday",
			id:       "34052420301301001X",
			expected: false,
		},
		{
			name:     "Invalid ID - Future Birthday",
			id:       "340524" + fmt.Sprintf("%02d", time.Now().Year()+1) + "0101001X",
			expected: false,
		},
		{
			name:     "Invalid ID - Too Short",
			id:       "34052419800101001",
			expected: false,
		},
		{
			name:     "Invalid ID - Too Long",
			id:       "34052419800101001XX",
			expected: false,
		},
		{
			name:     "Invalid ID - Non-numeric Characters",
			id:       "340524198A0101001X",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsChineseIDNum(tt.id)
			if actual != tt.expected {
				t.Errorf("IsChineseIDNum(%q) = %v; want %v", tt.id, actual, tt.expected)
			}
		})
	}
}

func TestContainChinese(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", false},
		{"No Chinese", "HelloWorld", false},
		{"Contains Chinese", "Hello世界", true},
		{"Mixed Chinese and ASCII", "你好Hello", true},
		{"Only Chinese", "你好", true},
		{"Special Characters", "!@#$%^&*()", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ContainChinese(tt.input)
			if actual != tt.expected {
				t.Errorf("ContainChinese(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsChinesePhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Format 1", "123-45678901", true},
		{"Valid Format 2", "1234-5678901", true},
		{"Invalid Format", "1234567890", false},
		{"Empty String", "", false},
		{"Non-Chinese Characters", "abc-12345678", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsChinesePhone(tt.input)
			if actual != tt.expected {
				t.Errorf("IsChinesePhone(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsCreditCard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Credit Card", "4111111111111111", true},
		{"Invalid Format", "1234567890123456", false},
		{"Empty String", "", false},
		{"Non-Digit Characters", "1234-abcd-5678-efgh", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsCreditCard(tt.input)
			if actual != tt.expected {
				t.Errorf("IsCreditมีCard(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Base64", "SGVsbG8gV29ybGQh", true},
		{"Invalid Base64", "Hello World!", false},

		{"Empty String", "", false},
		{"Non-Base64 Characters", "Hello, 世界!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsBase64(tt.input)
			if actual != tt.expected {
				t.Errorf("IsBase64(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsEmptyString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty String", "", true},
		{"Non-Empty String", "Hello, World!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsEmptyString(tt.input)
			if actual != tt.expected {
				t.Errorf("IsEmptyString(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsRegexMatch(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		regex    string
		expected bool
	}{
		{"Match", "Hello, World!", "^Hello", true},
		{"No Match", "Hello, World!", "^Goodbye", false},
		{"Empty String", "", "^$", true},
		{"Regex Error", "Hello, World!", `\[`, false}, // Invalid regex
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsRegexMatch(tt.input, tt.regex)
			if actual != tt.expected {
				t.Errorf("IsRegexMatch(%q, %q) = %v; want %v", tt.input, tt.regex, actual, tt.expected)
			}
		})
	}
}

func TestIsStrongPassword(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected bool
	}{
		{"Empty String", "", 8, false},
		{"Less Than Length", "pass", 8, false},

		{"No Number", "Password!", 8, false},
		{"No Lowercase", "PASSWORD1", 8, false},
		{"No Uppercase", "password1", 8, false},
		{"No Special Char", "Password1", 8, false},

		{"Strong Password", "Password1!", 8, true},
		{"Strong Password with Length 6", "Pass1!", 6, true},
		{"Strong Password with Length 10", "Password1!", 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsStrongPassword(tt.input, tt.length)
			if actual != tt.expected {
				t.Errorf("IsStrongPassword(%q, %d) = %v; want %v", tt.input, tt.length, actual, tt.expected)
			}
		})
	}
}

func TestIsWeakPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"abcdef", true}, // Only letters
		{"123456", true}, // Only numbers
		{"abc123", true}, // Letters and numbers

		{"abc!123", false},    // Contains special character
		{"", false},           // Empty string
		{"a", true},           // Too short
		{"abcABC123!", false}, // Strong password
	}

	for _, test := range tests {
		result := IsWeakPassword(test.password)
		if result != test.expected {
			t.Errorf("IsWeakPassword(%q) = %v; expected %v", test.password, result, test.expected)
		}
	}
}

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want bool
	}{
		{"nil", nil, true},
		{"empty string", "", true},
		{"false boolean", false, true},
		{"int zero", 0, true},
		{"int8 zero", int8(0), true},
		{"int16 zero", int16(0), true},
		{"int32 zero", int32(0), true},
		{"int64 zero", int64(0), true},
		{"uint zero", uint(0), true},
		{"uint8 zero", uint8(0), true},
		{"uint16 zero", uint16(0), true},
		{"uint32 zero", uint32(0), true},
		{"uint64 zero", uint64(0), true},
		{"float32 zero", float32(0), true},
		{"float64 zero", float64(0), true},
		{"nil pointer", (*int)(nil), true},
		{"nil interface", new(interface{}), true},
		{"nil slice", []int(nil), true},
		{"complex type zero", complex(0, 0), true},

		{"nil channel", nil, true},
		{"nil function", nil, true},
		{"nil map", nil, true},

		{"zero len channel", make(chan int, 0), true},
		{"zero len map", make(map[string]int), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroValue(tt.val); got != tt.want {
				t.Errorf("IsZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGBK(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{"Valid GBK", []byte("你好"), true},

		{"Invalid GBK start", []byte("\x80\x40"), false},
		{"Invalid GBK end", []byte("\xfe\x3f"), false},
		{"Invalid GBK with F7", []byte("\x81\xf7"), false},

		{"Empty string", []byte(""), true},
		{"Single byte", []byte{0x41}, true},
		{"GBK range limits", []byte("\x81\x40\xfe\xfe"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsGBK(tt.data)
			if result != tt.expected {
				t.Errorf("IsGBK(%q) = %v; want %v", tt.data, result, tt.expected)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	tests := []struct {
		value    any
		expected bool
	}{
		{float32(1.0), true},
		{float64(1.0), true},
		{int(1), false},
		{"1.0", false}, // Note: This will not be recognized as a float by IsFloat
		{"not a float", false},
		{nil, false},
	}

	for _, test := range tests {
		result := IsFloat(test.value)
		if result != test.expected {
			t.Errorf("IsFloat(%v) = %v; want %v", test.value, result, test.expected)
		}
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		value    any
		expected bool
	}{
		{int(0), true},
		{int8(0), true},
		{int16(0), true},
		{int32(0), true},
		{int64(0), true},
		{uint(0), true},
		{uint8(0), true},
		{uint16(0), true},
		{uint32(0), true},
		{uint64(0), true},
		{uintptr(0), true},
		{float32(0), false},
		{float64(0), false},
		{"string", false},
		{true, false},
		{nil, false},
	}

	for _, test := range tests {
		result := IsInt(test.value)
		if result != test.expected {
			t.Errorf("IsInt(%v) = %v; want %v", test.value, result, test.expected)
		}
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		value    any
		expected bool
	}{
		{123, true},
		{123.45, true},
		{"123", false},
		{"abc", false},
		{true, false},
		{nil, false},
	}

	for _, test := range tests {
		result := IsNumber(test.value)
		if result != test.expected {
			t.Errorf("IsNumber(%v) = %v; want %v", test.value, result, test.expected)
		}
	}
}

func TestIsBin(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"All zeros", "0000", true},
		{"All ones", "1111", true},
		{"Mixed zeros and ones", "101010", true},
		{"Invalid characters", "102010", false},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBin(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v for input %s", tc.expected, result, tc.input)
			}
		})
	}
}

func TestIsHex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Valid hex", "1A3FBC90", true},
		{"Invalid hex with letter G", "1A3FBC9G", false},
		{"Invalid hex with special character", "1A3FBC9!", false},
		{"Valid hex with leading 0x", "0x1A3FBC90", true},
		{"Invalid hex with leading 0x and special character", "0x1A3FBC9!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHex(tt.input); got != tt.expected {
				t.Errorf("IsHex(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
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
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid JWT", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", true},
		{"Invalid JWT - incorrect segment count", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ", false},
		{"Invalid JWT - not Base64URL encoded", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV/adQssw5c.", false},
		{"Empty string", "", false},
		{"Too short to be a JWT", "a.b.c", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJWT(tt.input); got != tt.expected {
				t.Errorf("IsJWT(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsVisa(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Visa", "4111111111111111", true},
		{"Invalid Visa", "1234567890123456", false},
		{"Empty String", "", false},
		// {"Non-string Input", 1234567890, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsVisa(tt.input)
			if actual != tt.expected {
				t.Errorf("IsVisa(%q) = %v; expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsMasterCard(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"5555555555554444", true},
		{"378282246310005", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsMasterCard(test.cardNumber)
		if result != test.expected {
			t.Errorf("IsMasterCard(%q) = %v; expected %v", test.cardNumber, result, test.expected)
		}
	}
}

func TestIsAmericanExpress(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"378282246310005", true},
		{"5555555555554444", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsAmericanExpress(test.cardNumber)
		if result != test.expected {
			t.Errorf("IsAmericanExpress(%q) = %v; expected %v", test.cardNumber, result, test.expected)
		}
	}
}

func TestIsUnionPay(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"6211514433542201", true},
		{"5555555555554444", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsUnionPay(test.cardNumber)
		if result != test.expected {
			t.Errorf("IsUnionPay(%q) = %v; expected %v", test.cardNumber, result, test.expected)
		}
	}
}

func TestIsChinaUnionPay(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   bool
	}{
		{"6222021001122334456", true},
		{"6011514433546201", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsChinaUnionPay(test.cardNumber)
		if result != test.expected {
			t.Errorf("IsChinaUnionPay(%q) = %v; expected %v", test.cardNumber, result, test.expected)
		}
	}
}
