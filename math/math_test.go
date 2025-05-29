package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponent(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal(0, Exponent(0, -1))
	assert.Equal(1, Exponent(0, 0))
	assert.Equal(0, Exponent(0, 10))

	assert.Equal(1, Exponent(1, 0))
	assert.Equal(1, Exponent(1, 10))

	assert.Equal(0, Exponent(2, -1))
	assert.Equal(8, Exponent(2, 3))
	assert.Equal(1024, Exponent(2, 10))
	assert.Equal(1, Exponent(2, 0))
}

func TestFactorial(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal(uint(1), Factorial(uint(0)))
	assert.Equal(uint(1), Factorial(uint(1)))
	assert.Equal(uint(2*1), Factorial(uint(2)))
	assert.Equal(uint(3*2*1), Factorial(uint(3)))
	assert.Equal(uint(4*3*2*1), Factorial(uint(4)))
	assert.Equal(uint(5*4*3*2*1), Factorial(uint(5)))
	assert.Equal(uint(6*5*4*3*2*1), Factorial(uint(6)))
	assert.Equal(uint(7*6*5*4*3*2*1), Factorial(uint(7)))
	assert.Equal(uint(8*7*6*5*4*3*2*1), Factorial(uint(8)))
	assert.Equal(uint(9*8*7*6*5*4*3*2*1), Factorial(uint(9)))
	assert.Equal(uint(10*9*8*7*6*5*4*3*2*1), Factorial(uint(10)))
	assert.Equal(uint(11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(11)))
	assert.Equal(uint(12*11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(12)))
	assert.Equal(uint(13*12*11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(13)))
	assert.Equal(uint(14*13*12*11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(14)))
	assert.Equal(uint(15*14*13*12*11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(15)))
	assert.Equal(uint(16*15*14*13*12*11*10*9*8*7*6*5*4*3*2*1), Factorial(uint(16)))
}

func TestRoundToFloat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Round float64
	assert.Equal(float64(123.46), RoundToFloat(123.456789, 2))

	// Round int
	assert.Equal(float64(123.00), RoundToFloat(123, 2))

	// Round negative float64
	assert.Equal(float64(-123.46), RoundToFloat(-123.456789, 2))

	// Round float64 with zero decimals
	assert.Equal(float64(123.0), RoundToFloat(123.456789, 0))

	// Round float64 with 5 decimals
	assert.Equal(float64(123.45679), RoundToFloat(123.456789, 5))

	// Round float64 with negative decimals
	assert.Equal(float64(120.0), RoundToFloat(123.456789, -1))
	assert.Equal(float64(100.0), RoundToFloat(123.456789, -2))
}

func TestRoundToString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Round float64 to 2 decimals
	assert.Equal("123.46", RoundToString(123.456789, 2))

	// Round int to 2 decimals
	assert.Equal("123.00", RoundToString(123, 2))

	// Round negative float64 to 2 decimals
	assert.Equal("-123.46", RoundToString(-123.456789, 2))

	// Round float64 to 0 decimals
	assert.Equal("123", RoundToString(123.456789, 0))

	// Round float64 to 5 decimals
	assert.Equal("123.45679", RoundToString(123.456789, 5))
}

func TestPercent(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Both values are zero
	assert.Equal(float64(0), Percent(0, 0, 2))

	// Value is zero
	assert.Equal(float64(0), Percent(0, 100, 2))

	// Total is zero
	assert.Equal(float64(0), Percent(50, 0, 2))

	// Normal case
	assert.Equal(float64(50.00), Percent(50, 100, 2))

	// Decimal case
	assert.Equal(float64(33.333), Percent(1, 3, 3))

	// Negative values
	assert.Equal(float64(-50.00), Percent(-50, 100, 2))
}

func TestTruncRound(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Round float64 to 2 decimals
	assert.Equal(float64(123.45), TruncRound(123.456789, 2))

	// Round int to 2 decimals
	assert.Equal(123, TruncRound(123, 2))

	// Round negative float64 to 2 decimals
	assert.Equal(float64(-123.45), TruncRound(-123.456789, 2))

	// Round float64 to 0 decimals
	assert.Equal(float64(123.0), TruncRound(123.456789, 0))

	// Round float64 to 5 decimals
	assert.Equal(float64(123.45678), TruncRound(123.456789, 5))

	// Round float64 to 8 decimals
	assert.Equal(float64(123.45678901), TruncRound(123.4567890123456789, 8))

	// Round with n = 0 for integer
	assert.Equal(123, TruncRound(123, 0))

	// Round with n < 0
	assert.Equal(float64(123.0), TruncRound(123.456789, -1))
}

func TestFloorToFloat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Rount float64 with specified decimals
	assert.Equal(float64(123.45), FloorToFloat(123.456789, 2))
	assert.Equal(float64(123.0), FloorToFloat(123.456789, 0))
	assert.Equal(float64(123.45678), FloorToFloat(123.456789, 5))

	// Rount int to 2 decimals
	assert.Equal(123.00, FloorToFloat(123, 2))

	// Rount negative float64 to 2 decimals
	assert.Equal(-123.46, FloorToFloat(-123.456789, 2))
}

func TestFloorToStr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Round float64 with specified decimals
	assert.Equal("123.45", FloorToStr(123.456789, 2))
	assert.Equal("123", FloorToStr(123.456789, 0))
	assert.Equal("123.45678", FloorToStr(123.456789, 5))

	// Round negative
	assert.Equal("-124", FloorToStr(-123.456789, 0))
	assert.Equal("-123.46", FloorToStr(-123.456789, 2))

	// Round int
	assert.Equal("123.00", FloorToStr(123, 2))
}

func TestCeilToFloat(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Rount float64 with specified decimals
	assert.Equal(float64(123.46), CeilToFloat(123.456789, 2))
	assert.Equal(float64(124.0), CeilToFloat(123.456789, 0))
	assert.Equal(float64(123.45679), CeilToFloat(123.456789, 5))

	// Rount int to 2 decimals
	assert.Equal(123.00, CeilToFloat(123, 2))

	// Rount negative float64 to 2 decimals
	assert.Equal(-123.45, CeilToFloat(-123.456789, 2))
}

func TestCeilToStr(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Round float64 with specified decimals
	assert.Equal("123.46", CeilToStr(123.456789, 2))
	assert.Equal("124", CeilToStr(123.456789, 0))
	assert.Equal("123.45679", CeilToStr(123.456789, 5))

	// Round negative
	assert.Equal("-123.45", CeilToStr(-123.456789, 2))
	assert.Equal("-123", CeilToStr(-123.456789, 0))

	// Round int
	assert.Equal("123.00", CeilToStr(123, 2))
}

func TestSum(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Integer
	assert.Equal(15, Sum([]int{1, 2, 3, 4, 5}...))
	assert.Equal(1, Sum([]int{1}...))
	assert.Equal(3, Sum([]int{1, 2}...))

	// Float64
	assert.Equal(16.5, Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5}...))
	assert.Equal(float32(1.1), Sum([]float32{1.1}...))

	// nil
	assert.Equal(0, Sum([]int{}...))
	assert.Equal(float64(0.0), Sum([]float64{}...))
}

func TestAverage(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Integer
	assert.Equal(3, Average([]int{1, 2, 3, 4, 5}...))
	assert.Equal(1, Average([]int{1}...))
	assert.Equal(1, Average([]int{1, 2}...))

	// Float64
	assert.Equal(3.1, Average([]float64{1.0, 2.0, 3.5, 4.0, 5.0}...))
	assert.Equal(float32(1.1), Average([]float32{1.1}...))

	// nil
	assert.Equal(0, Average([]int{}...))
	assert.Equal(float64(0.0), Average([]float64{}...))
}

func TestGCD(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Zero
	assert.Equal(0, GCD(0, 0))
	assert.Equal(1, GCD(0, 1))
	assert.Equal(1, GCD(1, 0))

	// Positive
	assert.Equal(5, GCD(15, 5))
	assert.Equal(25, GCD(100, 25))
	assert.Equal(6, GCD(12, 18))
	assert.Equal(12, GCD(12, 24))
	assert.Equal(5, GCD(10, 25))
	assert.Equal(1, GCD(7, 13))
	assert.Equal(100, GCD(100, 1000))

	// Negative
	assert.Equal(5, GCD(-15, 5))
	assert.Equal(25, GCD(-100, 25))
}

func TestLCM(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Zero
	assert.Panics(func() { LCM(0, 0) })
	assert.Panics(func() { LCM(0, 1) })

	// Positive
	assert.Equal(15, LCM(3, 5))
	assert.Equal(100, LCM(10, 100))
	assert.Equal(36, LCM(12, 18))
	assert.Equal(24, LCM(12, 24))
	assert.Equal(50, LCM(10, 25))
	assert.Equal(91, LCM(7, 13))
	assert.Equal(1000, LCM(100, 1000))
}
