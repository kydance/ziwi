package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("helloWorld", CamelCase("hello world"))
	assert.Equal("helloWorld", CamelCase("Hello World"))

	assert.Equal("helloWorld", CamelCase("hello-world"))
	assert.Equal("helloWorld", CamelCase("hello_world"))

	assert.Equal("hello", CamelCase("hello"))
	assert.Equal("hello", CamelCase("Hello"))

	assert.Equal("123", CamelCase("123"))
	assert.Equal("", CamelCase("!@#"))
	assert.Equal("", CamelCase(""))

	assert.Equal("foobar", CamelCase("foobar"))
	assert.Equal("fooBarBaz", CamelCase("&FOO:BAR$BAZ"))
	assert.Equal("fooBar", CamelCase("fooBar"))
	assert.Equal("foObar", CamelCase("FOObar"))
	assert.Equal("foo", CamelCase("$foo%"))
	assert.Equal("foo22Bar", CamelCase("   $#$Foo   22    bar   "))
	assert.Equal("foo11Bar", CamelCase("Foo-#1😄$_%^&*(1bar"))
}

func TestCapitalize(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("Hello", Capitalize("hello"))
	assert.Equal("World", Capitalize("world"))
	assert.Equal("Golang", Capitalize("GoLang"))
	assert.Equal("", Capitalize(""))

	assert.Equal("H", Capitalize("h"))
	assert.Equal("H", Capitalize("H"))
	assert.Equal("123", Capitalize("123"))
	assert.Equal("!@#", Capitalize("!@#"))
	assert.Equal("你好", Capitalize("你好"))
	assert.Equal("你好世界", Capitalize("你好世界"))
}

func TestUpperFirst(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("Hello", UpperFirst("hello"))
	assert.Equal("World", UpperFirst("world"))
	assert.Equal("GoLang", UpperFirst("GoLang"))
	assert.Equal("", UpperFirst(""))

	assert.Equal("H", UpperFirst("h"))
	assert.Equal("H", UpperFirst("H"))
	assert.Equal("123", UpperFirst("123"))

	assert.Equal("!@#", UpperFirst("!@#"))
	assert.Equal("你好", UpperFirst("你好"))
	assert.Equal("你好世界", UpperFirst("你好世界"))
	assert.Equal("你好，世界！", UpperFirst("你好，世界！"))
	assert.Equal("中文编程", UpperFirst("中文编程"))
}

func TestLowerFirst(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello", LowerFirst("Hello"))
	assert.Equal("world", LowerFirst("World"))
	assert.Equal("goLang", LowerFirst("GoLang"))
	assert.Equal("", LowerFirst(""))

	assert.Equal("h", LowerFirst("H"))
	assert.Equal("h", LowerFirst("h"))
	assert.Equal("123", LowerFirst("123"))

	assert.Equal("!@#", LowerFirst("!@#"))

	assert.Equal("你好", LowerFirst("你好"))
	assert.Equal("你好世界", LowerFirst("你好世界"))
	assert.Equal("你好，世界！", LowerFirst("你好，世界！"))
	assert.Equal("中文编程", LowerFirst("中文编程"))

	assert.Equal("hELLO WORLD", LowerFirst("HELLO WORLD"))
}

func TestPad(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("!!hello!!!", Pad("hello", 10, "!"))
	assert.Equal("world", Pad("world", 5, "*"))
	assert.Equal("-Go-", Pad("Go", 4, "-"))
	assert.Equal("###", Pad("", 3, "#"))
	assert.Equal("$你好$", Pad("你好", 8, "$"))
}

func TestPadLeft(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("!!!!!hello", PadLeft("hello", 10, "!"))
	assert.Equal("world", PadLeft("world", 5, "*"))
	assert.Equal("--Go", PadLeft("Go", 4, "-"))
	assert.Equal("###", PadLeft("", 3, "#"))
	assert.Equal("@@你好", PadLeft("你好", 8, "@"))
}

func TestPadRight(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello!!!!!", PadRight("hello", 10, "!"))
	assert.Equal("world", PadRight("world", 5, "*"))
	assert.Equal("Go--", PadRight("Go", 4, "-"))
	assert.Equal("###", PadRight("", 3, "#"))
	assert.Equal("你好~~", PadRight("你好", 8, "~"))
}

func TestKebabCase(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello-world", KebabCase("hello world"))
	assert.Equal("hello-world", KebabCase("Hello World"))
	assert.Equal("hello-world", KebabCase("hello-world"))
	assert.Equal("hello-world", KebabCase("hello_world"))

	assert.Equal("hello", KebabCase("hello"))
	assert.Equal("hello", KebabCase("Hello"))

	assert.Equal("123", KebabCase("123"))
	assert.Equal("", KebabCase("!@#"))
	assert.Equal("", KebabCase(""))

	assert.Equal("foobar", KebabCase("foobar"))
	assert.Equal("foo-bar-baz", KebabCase("&FOO:BAR$BAZ"))
	assert.Equal("foo-bar", KebabCase("fooBar"))
	assert.Equal("fo-obar", KebabCase("FOObar"))
	assert.Equal("foo", KebabCase("$foo%"))
	assert.Equal("foo-22-bar", KebabCase("   $#$Foo   22    bar   "))
	assert.Equal("foo-1-1-bar", KebabCase("Foo-#1😄$_%^&*(1bar"))

	assert.Equal("你好世界", KebabCase("你好世界"))
	assert.Equal("中文编程", KebabCase("中文编程"))
}

func TestUpperKebabCase(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("HELLO-WORLD", UpperKebabCase("hello world"))
	assert.Equal("HELLO-WORLD", UpperKebabCase("Hello World"))
	assert.Equal("HELLO-WORLD", UpperKebabCase("hello-world"))
	assert.Equal("HELLO-WORLD", UpperKebabCase("hello_world"))

	assert.Equal("HELLO", UpperKebabCase("hello"))
	assert.Equal("HELLO", UpperKebabCase("Hello"))

	assert.Equal("123", UpperKebabCase("123"))
	assert.Equal("", UpperKebabCase("!@#"))
	assert.Equal("", UpperKebabCase(""))

	assert.Equal("FOO-BAR-BAZ", UpperKebabCase("&FOO:BAR$BAZ"))
	assert.Equal("FOO-BAR", UpperKebabCase("fooBar"))
	assert.Equal("FO-OBAR", UpperKebabCase("FOObar"))
	assert.Equal("FOO", UpperKebabCase("$foo%"))
	assert.Equal("FOO-22-BAR", UpperKebabCase("   $#$Foo   22    bar   "))
	assert.Equal("FOO-1-1-BAR", UpperKebabCase("Foo-#1😄$_%^&*(1bar"))

	assert.Equal("你好世界", UpperKebabCase("你好世界"))
	assert.Equal("中文编程", UpperKebabCase("中文编程"))
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello_world", SnakeCase("hello world"))
	assert.Equal("hello_world", SnakeCase("Hello World"))
	assert.Equal("hello_world", SnakeCase("hello-world"))
	assert.Equal("hello_world", SnakeCase("hello_world"))

	assert.Equal("hello", SnakeCase("hello"))
	assert.Equal("hello", SnakeCase("Hello"))

	assert.Equal("123", SnakeCase("123"))
	assert.Equal("", SnakeCase("!@#"))
	assert.Equal("", SnakeCase(""))

	assert.Equal("foobar", SnakeCase("foobar"))
	assert.Equal("foo_bar_baz", SnakeCase("&FOO:BAR$BAZ"))
	assert.Equal("foo_bar", SnakeCase("fooBar"))
	assert.Equal("fo_obar", SnakeCase("FOObar"))
	assert.Equal("foo", SnakeCase("$foo%"))
	assert.Equal("foo_22_bar", SnakeCase("   $#$Foo   22    bar   "))
	assert.Equal("foo_1_1_bar", SnakeCase("Foo-#1😄$_%^&*(1bar"))

	assert.Equal("你好世界", SnakeCase("你好世界"))
	assert.Equal("中文编程", SnakeCase("中文编程"))
}

func TestUpperSnakeCase(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("HELLO_WORLD", UpperSnakeCase("hello world"))
	assert.Equal("HELLO_WORLD", UpperSnakeCase("Hello World"))
	assert.Equal("HELLO_WORLD", UpperSnakeCase("hello-world"))
	assert.Equal("HELLO_WORLD", UpperSnakeCase("hello_world"))

	assert.Equal("HELLO", UpperSnakeCase("hello"))
	assert.Equal("HELLO", UpperSnakeCase("Hello"))

	assert.Equal("123", UpperSnakeCase("123"))
	assert.Equal("", UpperSnakeCase("!@#"))
	assert.Equal("", UpperSnakeCase(""))

	assert.Equal("FOO_BAR_BAZ", UpperSnakeCase("&FOO:BAR$BAZ"))
	assert.Equal("FOO_BAR", UpperSnakeCase("fooBar"))
	assert.Equal("FO_OBAR", UpperSnakeCase("FOObar"))
	assert.Equal("FOO", UpperSnakeCase("$foo%"))
	assert.Equal("FOO_22_BAR", UpperSnakeCase("   $#$Foo   22    bar   "))
	assert.Equal("FOO_1_1_BAR", UpperSnakeCase("Foo-#1😄$_%^&*(1bar"))

	assert.Equal("你好世界", UpperSnakeCase("你好世界"))
	assert.Equal("中文编程", UpperSnakeCase("中文编程"))
}

func TestBefore(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello", Before("hello world", " "))
	assert.Equal("he", Before("hello", "l"))
	assert.Equal("hell", Before("hello", "o"))

	assert.Equal("hello", Before("hello", "z"))
	assert.Equal("", Before("", "a"))
	assert.Equal("hello", Before("hello", ""))

	assert.Equal("Golang", Before("Golang is fun", " "))
	assert.Equal("no", Before("no-separator", "-"))

	assert.Equal("你好", Before("你好,世界！", ","))
	assert.Equal("你好", Before("你好，世界！", "，"))
	assert.Equal("中文编程", Before("中文编程,测试", ","))

	assert.Equal("Hello-", Before("Hello-kyden", "kyden"))
}

func TestBeforeLast(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello", BeforeLast("hello world", " "))
	assert.Equal("hel", BeforeLast("hello", "l"))
	assert.Equal("hell", BeforeLast("hello", "o"))

	assert.Equal("hello", BeforeLast("hello", "z"))
	assert.Equal("", BeforeLast("", "a"))
	assert.Equal("hello", BeforeLast("hello", ""))

	assert.Equal("Golang is", BeforeLast("Golang is fun", " "))
	assert.Equal("abcab", BeforeLast("abcabc", "c"))
}

func TestAfter(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("world", After("hello world", " "))
	assert.Equal("lo", After("hello", "l"))
	assert.Equal(" world", After("hello world", "o"))

	assert.Equal("hello", After("hello", "z"))
	assert.Equal("", After("", "a"))
	assert.Equal("hello", After("hello", ""))
}

func TestAfterLast(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("world", AfterLast("hello world", " "))
	assert.Equal("o", AfterLast("hello", "l"))
	assert.Equal("rld", AfterLast("hello world", "o"))

	assert.Equal("hello", AfterLast("hello", "z"))
	assert.Equal("", AfterLast("", "a"))
	assert.Equal("hello", AfterLast("hello", ""))
}

func TestIsString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.False(IsString(nil))
	assert.True(IsString("hello"))
	assert.False(IsString(123))
	assert.False(IsString(true))
	assert.False(IsString([]int{}))
	assert.False(IsString(map[string]int{}))
	assert.True(IsString("你好"))
	assert.False(IsString(complex(1, 2)))
}

func TestReverse(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello", Reverse("olleh"))
	assert.Equal("world", Reverse("dlrow"))
	assert.Equal("GoLang", Reverse("gnaLoG"))

	assert.Equal("", Reverse(""))
	assert.Equal("h", Reverse("h"))
	assert.Equal("H", Reverse("H"))

	assert.Equal("123", Reverse("321"))
	assert.Equal("!@#", Reverse("#@!"))
	assert.Equal("你好", Reverse("好你"))
	assert.Equal("你好世界", Reverse("界世好你"))
	assert.Equal("你好，世界!", Reverse("!界世，好你"))
	assert.Equal("中文编程", Reverse("程编文中"))
}

func TestWarp(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("!hello!", Warp("hello", "!"))
	assert.Equal("*world*", Warp("world", "*"))
	assert.Equal("-Go-", Warp("Go", "-"))
	assert.Equal("", Warp("", "#"))
	assert.Equal("@你好@", Warp("你好", "@"))
	assert.Equal("hello", Warp("hello", ""))
}

func TestUnWarp(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("", UnWarp("", "#"))
	assert.Equal("hello", UnWarp("hello", ""))

	assert.Equal("hello", UnWarp("!hello!", "!"))
	assert.Equal("world", UnWarp("*world*", "*"))
	assert.Equal("Go", UnWarp("-Go-", "-"))
	assert.Equal("你好", UnWarp("@你好@", "@"))

	assert.Equal("hello", UnWarp("hello", "world"))
	assert.Equal("helloworld", UnWarp("helloworld", "world"))
	assert.Equal("worldhello", UnWarp("worldhello", "world"))

	assert.Equal("def", UnWarp("abcdefabc", "abc"))
}

func TestSubString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("he", SubString("hello", 0, 2))
	assert.Equal("ello", SubString("hello", 1, 4))
	assert.Equal("o", SubString("hello", 4, 1))
	assert.Equal("", SubString("hello", 5, 1))
	assert.Equal("o", SubString("hello", -1, 2))
	assert.Equal("he", SubString("hello", -5, 2))
	assert.Equal("hello", SubString("hello", 0, 10))

	assert.Equal("好世", SubString("你好世界", 1, 2))
	assert.Equal("世界", SubString("你好世界", -2, 2))
	assert.Equal("", SubString("你好世界", 3, 0))
	assert.Equal("", SubString("你好世界", 0, -1))
	assert.Equal("好世界", SubString("\x00你好世界\x00", 2, 3))
}

func TestRemoveNonPrintable(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("Hello, World!", RemoveNonPrintable("Hello, World!"))
	assert.Equal("Hello, World!", RemoveNonPrintable("Hello, World!\x00"))
	assert.Equal("Hello, World!", RemoveNonPrintable("Hello, World!\x01"))
	assert.Equal("Hello, World!", RemoveNonPrintable("Hello, World!\x7F"))

	assert.Equal("こんにちは", RemoveNonPrintable("こんにちは"))
	assert.Equal("こんにちは", RemoveNonPrintable("こんにちは\x00"))

	assert.Equal("\u00A1\u00A2", RemoveNonPrintable("\u00A0\u00A1\u00A2"))
	assert.Equal("", RemoveNonPrintable("\x7F"))
}

func TestStringToBytes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]byte("hello"), StringToBytes("hello"))
	assert.Equal([]byte("world"), StringToBytes("world"))
	assert.Equal([]byte(nil), StringToBytes(""))
	assert.Equal([]byte("你好"), StringToBytes("你好"))
}

func TestBytesToString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello", BytesToString([]byte("hello")))
	assert.Equal("world", BytesToString([]byte("world")))
	assert.Equal("", BytesToString([]byte("")))
	assert.Equal("你好", BytesToString([]byte("你好")))
}

func TestIsSpace(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.True(IsSpace(""))
	assert.True(IsSpace(" "))
	assert.True(IsSpace("\t"))
	assert.True(IsSpace("\n"))
	assert.True(IsSpace("\r"))
	assert.True(IsSpace(" \t\n\r"))

	assert.False(IsSpace("hello"))
	assert.False(IsSpace("你好"))
	assert.False(IsSpace("123"))
	assert.False(IsSpace("!@#"))
	assert.False(IsSpace("你好 世界"))
}

func TestIsNotSpace(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.False(IsNotSpace(""))
	assert.False(IsNotSpace(" "))
	assert.False(IsNotSpace("\t"))
	assert.False(IsNotSpace("\n"))
	assert.False(IsNotSpace("\r"))
	assert.False(IsNotSpace(" \t\n\r"))

	assert.True(IsNotSpace("hello"))
	assert.True(IsNotSpace("你好"))
	assert.True(IsNotSpace("123"))
	assert.True(IsNotSpace("!@#"))
	assert.True(IsNotSpace("你好 世界"))
}

func TestHasPrefixAny(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.False(HasPrefixAny("", "a", "b"))
	assert.False(HasPrefixAny("H", "h"))

	assert.True(HasPrefixAny("hello", "he", "ha"))
	assert.True(HasPrefixAny("world", "wo", "wa"))
	assert.True(HasPrefixAny("GoLang", "Go", "GoL"))

	assert.True(HasPrefixAny("h", "", ""))
	assert.True(HasPrefixAny("123", "12", "123"))
	assert.True(HasPrefixAny("!@#", "!@", "!#"))
	assert.True(HasPrefixAny("你好", "你", "好"))
	assert.True(HasPrefixAny("你好世界", "你好", "好世"))
}

func TestHasSuffixAny(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.False(HasSuffixAny("", "a", "b"))
	assert.False(HasSuffixAny("H", "h"))
	assert.False(HasSuffixAny("!@#", "#@", "!#"))

	assert.True(HasSuffixAny("hello", "lo", "he"))
	assert.True(HasSuffixAny("world", "ld", "wo"))
	assert.True(HasSuffixAny("GoLang", "g", "Lang"))

	assert.True(HasSuffixAny("h", "", ""))
	assert.True(HasSuffixAny("123", "23", "123"))
	assert.True(HasSuffixAny("你好", "好", "你"))
	assert.True(HasSuffixAny("你好世界", "世界", "好世"))
}

func TestIndexOffset(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal(-1, IndexOffset("", "", 0))
	assert.Equal(-1, IndexOffset("", "a", 0))

	assert.Equal(6, IndexOffset("hello world", "world", 6))
	assert.Equal(7, IndexOffset("hello world", "o", 7))
	assert.Equal(3, IndexOffset("hello world", "l", 3))

	assert.Equal(-1, IndexOffset("hello world", "z", 0))
	assert.Equal(-1, IndexOffset("hello world", "o", 11))

	assert.Equal(0, IndexOffset("hello world", "", 0))
	assert.Equal(-1, IndexOffset("", "a", 0))
	assert.Equal(-1, IndexOffset("hello world", "world", -1))
	assert.Equal(-1, IndexOffset("hello world", "world", 11))
}

func TestReplaceWithMap(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello world", ReplaceWithMap("hello world", nil))
	assert.Equal("hello world", ReplaceWithMap("hello world", map[string]string{}))
	assert.Equal("hello world", ReplaceWithMap("hello world", map[string]string{"a": "b"}))

	assert.Equal("hi world", ReplaceWithMap("hello world", map[string]string{"hello": "hi"}))
	assert.Equal("Hi World", ReplaceWithMap("Hello World", map[string]string{"Hello": "Hi"}))
	assert.Equal("hi-world", ReplaceWithMap("hello-world", map[string]string{"hello": "hi"}))
	assert.Equal("hi_world", ReplaceWithMap("hello_world", map[string]string{"hello": "hi"}))

	assert.Equal("Hello世界", ReplaceWithMap("你好世界", map[string]string{"你好": "Hello"}))
}

func TestTrim(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello world", Trim("  hello world  ", ""))
	assert.Equal("hello world", Trim("\t\nhello world\t\n", ""))
	assert.Equal("Hello World", Trim("abcHello World!cba", "abc!"))
	assert.Equal("Hello Worl", Trim("abcHello World!cba", "abc!dEF"))
	assert.Equal("世界！", Trim("你好 世界！", "你好"))
	assert.Equal("！", Trim("你好 世界！", "你好 世界"))
}

func TestSplitAndTrim(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]string{"hello", "world", "go"}, SplitAndTrim("hello, world, go", ",", ""))
	assert.Equal([]string{"Hello", "d", "eWorld", "f"}, SplitAndTrim("a,b,cHello,d,eWorld,f", ",", "abc"))
	assert.Equal([]string{"你好", "世界", "Go"}, SplitAndTrim("你好,世界,Go", ",", ""))
	assert.Equal([]string{"世界", "Go"}, SplitAndTrim("你好,世界,Go", ",", "你好"))
}

func TestHideString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hel****orld", HideString("hello world", 3, 7, "*"))
	assert.Equal("1XXXXXXXX0", HideString("1234567890", 1, 9, "X"))
	assert.Equal("------", HideString("abcdef", 0, 6, "-"))

	assert.Equal("abc", HideString("abc", 1, 2, ""))
	assert.Equal("abc", HideString("abc", 3, 2, "*"))
	assert.Equal("abc", HideString("abc", -1, 2, "*"))
	assert.Equal("abc", HideString("abc", 2, -1, "*"))
	assert.Equal("abc", HideString("abc", 4, 5, "*"))
	assert.Equal("", HideString("", 0, 0, "*"))
}

func TestContainsAll(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.True(ContainsAll("hello world", []string{"hello", "world"}))
	assert.False(ContainsAll("hello world", []string{"hello", "planet"}))
	assert.True(ContainsAll("", []string{}))

	assert.True(ContainsAll("hello", []string{"hello", "h"}))
	assert.True(ContainsAll("hello", []string{"hell", "o"}))
	assert.True(ContainsAll("hello", []string{"he", "llo"}))
	assert.True(ContainsAll("hello", []string{"he", "l", "lo"}))
	assert.True(ContainsAll("hello", []string{"he", "ll", "o"}))

	assert.True(ContainsAll("hello", []string{"h", "e", "l", "l", "o"}))
	assert.False(ContainsAll("hello", []string{"h", "e", "l", "l", "o", "x"}))
}

func TestContainsAny(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.False(ContainsAny("", []string{}))
	assert.False(ContainsAny("hello world", []string{}))
	assert.False(ContainsAny("hello world", []string{"a", "b", "c"}))

	assert.True(ContainsAny("hello world", []string{"hello", "world"}))
	assert.True(ContainsAny("hello world", []string{"hello", "planet"}))

	assert.True(ContainsAny("hello", []string{"hello", "h"}))
	assert.True(ContainsAny("hello", []string{"hell", "o"}))
	assert.True(ContainsAny("hello", []string{"he", "llo"}))
	assert.True(ContainsAny("hello", []string{"he", "l", "lo"}))
	assert.True(ContainsAny("hello", []string{"he", "ll", "o"}))

	assert.True(ContainsAny("hello", []string{"h", "e", "l", "l", "o"}))
	assert.True(ContainsAny("hello", []string{"h", "e", "l", "l", "o", "x"}))
}

func TestRemoveWhiteSpace(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("hello world", RemoveWhiteSpace("hello world", false))
	assert.Equal("helloworld", RemoveWhiteSpace("hello world", true))

	assert.Equal("hello world", RemoveWhiteSpace("  hello   world  ", false))
	assert.Equal("helloworld", RemoveWhiteSpace("  hello   world  ", true))

	assert.Equal("no-whitespace", RemoveWhiteSpace("no-whitespace", false))
	assert.Equal("no-whitespace", RemoveWhiteSpace("no-whitespace", true))

	assert.Equal("", RemoveWhiteSpace("  \t\n  ", false))
	assert.Equal("", RemoveWhiteSpace("  \t\n  ", true))

	assert.Equal("中文 编程", RemoveWhiteSpace("中文 编程", false))
	assert.Equal("中文编程", RemoveWhiteSpace("中文 编程", true))
}

func TestSubInBetween(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("middle", SubInBetween("startmiddleend", "start", "end"))
	assert.Equal("", SubInBetween("startmiddle", "start", "end"))
	assert.Equal("", SubInBetween("middlestartend", "start", "end"))
	assert.Equal("middle", SubInBetween("startmiddleendextra", "start", "end"))

	assert.Equal("", SubInBetween("", "start", "end"))
}

func TestHammingDistance(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	dis, err := HammingDistance("", "")
	assert.Equal(0, dis)
	assert.Nil(err)

	dis, err = HammingDistance("abc", "abc")
	assert.Equal(0, dis)
	assert.Nil(err)

	dis, err = HammingDistance("abc", "abd")
	assert.Equal(1, dis)
	assert.Nil(err)

	dis, err = HammingDistance("abc", "abcd")
	assert.Equal(-1, dis)
	assert.NotNil(err)
	assert.Equal("the length of two strings must be equal", err.Error())

	dis, err = HammingDistance("你好世界", "你好世界")
	assert.Equal(0, dis)
	assert.Nil(err)
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	s, e := Shuffle("hello world")
	assert.NotEqual("hello world", s)
	assert.Nil(e)

	s, e = Shuffle("")
	assert.Equal("", s)
	assert.NotNil(e)

	s, e = Shuffle("你好世界")
	assert.NotEqual("你好世界", s)
	assert.Nil(e)
}

func TestRotate(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("lohel", Rotate("hello", 2))
	assert.Equal("llohe", Rotate("hello", -2))
	assert.Equal("hello", Rotate("hello", 0))
	assert.Equal("hello", Rotate("hello", 5))
	assert.Equal("hello", Rotate("hello", -5))

	assert.Equal("", Rotate("", 3))
	assert.Equal("a", Rotate("a", 1))
	assert.Equal("ba", Rotate("ab", 1))
	assert.Equal("ba", Rotate("ab", -1))
}

func TestRegexMatchAllGroups(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// 基本测试
	assert.Equal([][]string{
		{"abc", "a", "b", "c"}, {"123", "1", "2", "3"},
	}, RegexMatchAllGroups("abc123", `(\w)(\w)(\w)`))
	// 无匹配测试
	assert.Equal([][]string{{"12", "12"}},
		RegexMatchAllGroups("abc123", `(\d\d)`))
	// 全局匹配测试
	assert.Equal([][]string{
		{"abc", "a", "b", "c"},
		{"123", "1", "2", "3"},
		{"abc", "a", "b", "c"},
	}, RegexMatchAllGroups("abc123abc", `(\w)(\w)(\w)`))
	// 复杂模式测试
	assert.Equal(
		[][]string{
			{"The quick", "The", "quick"},
			{"brown fox", "brown", "fox"},
			{"jumps over", "jumps", "over"},
			{"the lazy", "the", "lazy"},
		},
		RegexMatchAllGroups("The quick brown fox jumps over the lazy dog.", `(\w+)\s+(\w+)`),
	)
}

func TestConcat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Empty input
	assert.Equal("", Concat(0, []string{}...))

	// Single string
	assert.Equal("hello", Concat(0, []string{"hello"}...))

	// Multiple strings
	assert.Equal("helloworld", Concat(0, []string{"hello", "world"}...))

	// Specified length
	assert.Equal("helloworld", Concat(10, []string{"hello", "world"}...))

	// Length exceeds
	assert.Equal("helloworld", Concat(5, []string{"hello", "world"}...))

	// Empty strings
	assert.Equal("", Concat(0, []string{"", "", ""}...))
}

func TestEllipsis(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("Hello...", Ellipsis("Hello, world!", 5))
	assert.Equal("你好...", Ellipsis("你好，世界！", 2))
	assert.Equal("", Ellipsis("", 5))
	assert.Equal("Short", Ellipsis("Short", 10))
	assert.Equal("Longer text to be tr...", Ellipsis("Longer text to be truncated", 20))
	assert.Equal("", Ellipsis("", 0))
}

func TestStrEscape(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// empty string
	assert.Equal("", StrEscape(""))

	// string with carriage return
	assert.Equal("hello\\rworld", StrEscape("hello\rworld"))

	// string with newline
	assert.Equal("hello\\nworld", StrEscape("hello\nworld"))

	// string with backslash
	assert.Equal("hello\\\\world", StrEscape("hello\\world"))

	// string with single quote
	assert.Equal("hello\\'world", StrEscape("hello'world"))

	// string with double quote
	assert.Equal("hello\\\"world", StrEscape("hello\"world"))

	// string with Ctrl+Z
	assert.Equal("hello\\Zworld", StrEscape("hello\032world"))

	// string with multiple special chars
	assert.Equal("hello\\r\\n\\\"\\'\\\\world\\Z", StrEscape("hello\r\n\"'\\world\032"))

	// normal string without special chars
	assert.Equal("hello world", StrEscape("hello world"))

	// {}
	assert.Equal("{}", StrEscape("{}"))

	// Json str
	assert.Equal("{\\\"key\\\": \\\"value\\\"}", StrEscape(`{"key": "value"}`))
}
