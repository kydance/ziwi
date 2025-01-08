package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUpperAll(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]rune("HELLO"), toUpperAll([]rune("hello")))
	assert.Equal([]rune("WORLD"), toUpperAll([]rune("world")))
	assert.Equal([]rune("GOLANG"), toUpperAll([]rune("GoLang")))
	assert.Equal([]rune(""), toUpperAll([]rune("")))
}

func TestToLowerAll(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]rune("hello"), toLowerAll([]rune("HELLO")))
	assert.Equal([]rune("world"), toLowerAll([]rune("WORLD")))
	assert.Equal([]rune("golang"), toLowerAll([]rune("GoLang")))
	assert.Equal([]rune(""), toLowerAll([]rune("")))
}

func TestSplitIntoStrings(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]string{"HELLO", "WORLD", "123"}, splitIntoStrings("HelloWorld123", true))
	assert.Equal([]string{"hello", "world", "123"}, splitIntoStrings("HelloWorld123", false))

	assert.Equal([]string{"GO", "LANG", "PROGRAMMING"}, splitIntoStrings("GoLangProgramming", true))
	assert.Equal([]string{"go", "lang", "programming"}, splitIntoStrings("GoLangProgramming", false))

	assert.Equal([]string{"12345"}, splitIntoStrings("12345!@#$%", true))
	assert.Equal([]string{"12345"}, splitIntoStrings("12345!@#$%", false))

	assert.Equal([]string{}, splitIntoStrings("", true))
	assert.Equal([]string{}, splitIntoStrings("", false))
}

func TestPadAtPosEdgeCases(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("-----", padAtPos("", 5, "-", PosBoth))
	assert.Equal("-----", padAtPos("", 5, "-", PosLeft))
	assert.Equal("-----", padAtPos("", 5, "-", PosRight))

	assert.Equal("hello", padAtPos("hello", 3, "", PosBoth))
	assert.Equal("hello", padAtPos("hello", 3, "", PosLeft))
	assert.Equal("hello", padAtPos("hello", 3, "", PosRight))

	assert.Equal("hello*", padAtPos("hello", 6, "*", PosBoth))
	assert.Equal("*hello", padAtPos("hello", 6, "*", PosLeft))
	assert.Equal("hello*", padAtPos("hello", 6, "*", PosRight))

	assert.Equal("abhelloabc", padAtPos("hello", 10, "abc", PosBoth))
	assert.Equal("abcabhello", padAtPos("hello", 10, "abc", PosLeft))
	assert.Equal("helloabcab", padAtPos("hello", 10, "abc", PosRight))
}
