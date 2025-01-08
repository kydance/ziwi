// =============================================================================
/*!
 *  @file       math.go
 *  @brief      Package math implements some functions for math calculating.
 *  @author     kydenlu
 *  @date       2024.09
 *  @note
 */
// =============================================================================

package math

import (
	"math"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Exponent calculates base^exp. Returns 0 for negative exponents.
// For positive exponents, it uses the binary exponentiation algorithm.
func Exponent[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](base, exp T) T {
	// Handle special cases
	switch {
	case exp < 0:
		return 0
	case exp == 0:
		return 1
	case exp == 1:
		return base
	case base == 0:
		return 0
	case base == 1:
		return 1
	}

	result := T(1)
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

// Factorial calculates `x!`
func Factorial[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](x T) T {
	if x == 0 {
		return 1
	}

	ret := x
	for x--; x > 1; x-- {
		ret *= x
	}
	return ret
}

// RoundToFloat rounds a number to the specified number of decimal places.
//
// Parameters:
//
//	x: The number to be rounded.
//	n: The number of decimal places to round to.
//
// Returns:
//
//	A float64 representation of the rounded number.
//
// Example:
//
//	RoundToFloat(3.14159, 2) -> Output: 3.14
func RoundToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	return math.Round(float64(x)) / temp
}

// RoundToString rounds a numeric value to a specified number of decimal places and
// returns it as a string. It supports both floating point and integer types.
//
// Parameters:
//
//	x: The numeric value to be rounded.
//	n: The number of decimal places to round to.
//
// Returns:
//
//	A string representation of the rounded numeric value.
func RoundToString[T constraints.Float | constraints.Integer](x T, n int) string {
	// Calculate the power of 10 to multiply the input value by, to shift the decimal point.
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	// Format the rounded value back to a string with the specified number of decimal places.
	return strconv.FormatFloat(math.Round(float64(x))/temp, 'f', n, 64)
}

// Percent calculates the percentage of value to total.
//
// Parameters:
//   - val: The value to calculate the percentage of. Can be negative.
//   - total: The total value to calculate the percentage against. Can be negative.
//   - n: The number of decimal places to round to. Must be non-negative.
//
// Returns:
//   - A float64 representation of the percentage value.
//   - Returns 0 if total is 0 to prevent division by zero.
//   - For negative values, the sign of the result follows mathematical rules:
//     (-val/total) and (val/-total) will return negative percentage.
//
// Examples:
//
//	Percent(25, 100, 2) returns 25.00
//	Percent(-25, 100, 2) returns -25.00
//	Percent(1, 3, 2) returns 33.33
func Percent[T constraints.Float | constraints.Integer](val, total T, n int) float64 {
	if total == 0 {
		return 0
	}

	if n < 0 {
		n = 0
	}

	return RoundToFloat((float64(val)/float64(total))*100, n)
}

// TruncRound rounds a floating-point number to a specified number of decimal places
// by truncating any additional decimal places.
//
// Parameters:
//   - x: The number to round (can be float or integer)
//   - n: The number of decimal places to keep (must be non-negative)
//
// Returns:
//   - The rounded number of the same type as the input
//
// Examples:
//
//	TruncRound(3.14159, 2) returns 3.14
//	TruncRound(3.14159, 0) returns 3.0
//	TruncRound(-3.14159, 1) returns -3.1
func TruncRound[T constraints.Float | constraints.Integer](x T, n int) T {
	if n < 0 {
		n = 0
	}

	// Convert to float64 for calculation
	f := float64(x)

	// Calculate the multiplier (10^n)
	multiplier := math.Pow10(n)

	// Shift decimal point right, truncate decimal part, then shift back
	return T(math.Trunc(f*multiplier) / multiplier)
}

// FloorToFloat round down to n decimal places.
//
// Parameters:
//   - x: The number to round (can be float or integer)
//   - n: The number of decimal places to keep (must be non-negative)
//
// Returns:
//   - The rounded number of the same type as the input
func FloorToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	if n < 0 {
		n = 0
	}

	multiplier := math.Pow10(n)
	return math.Floor(float64(x)*multiplier) / multiplier
}

// FloorToStr round down to n decimal places.
func FloorToStr[T constraints.Float | constraints.Integer](x T, n int) string {
	if n < 0 {
		n = 0
	}

	return strconv.FormatFloat(FloorToFloat(x, n), 'f', n, 64)
}

// CeilToFloat round up to n decimal places.
func CeilToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	if n < 0 {
		n = 0
	}

	multiplier := math.Pow10(n)
	return math.Ceil(float64(x)*multiplier) / multiplier
}

// CeilToStr round up to n decimal places.
func CeilToStr[T constraints.Float | constraints.Integer](x T, n int) string {
	if n < 0 {
		n = 0
	}

	return strconv.FormatFloat(CeilToFloat(x, n), 'f', n, 64)
}

// Sum returns sum of all passed values.
func Sum[T constraints.Float | constraints.Integer](vs ...T) T {
	var sum T
	for _, v := range vs {
		sum += v
	}
	return sum
}

// Average returns the average of all passed values.
func Average[T constraints.Float | constraints.Integer](vs ...T) T {
	if len(vs) == 0 {
		return 0
	}

	return Sum(vs...) / T(len(vs))
}

// Range returns a slice of numbers from start to start+count step by step.
func Range[T constraints.Integer | constraints.Float](start T, count int) []T {
	size := count
	if count < 0 {
		size = -count
	}

	ret := make([]T, size)
	for i, j := 0, start; i < size; i, j = i+1, j+1 {
		ret[i] = j
	}

	return ret
}

// RangeWithStep creates a slice of numbers from start to end with specified step.
func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T {
	ret := []T{}

	if start >= end || step == 0 {
		return ret
	}

	for i := start; i < end; i += step {
		ret = append(ret, i)
	}

	return ret
}

// AngleToRadian converts angle value to radian value.
func AngleToRadian(angle float64) float64 { return angle * (math.Pi / 180) }

// RadianToAngle converts radian value to angle value.
func RadianToAngle(radian float64) float64 { return radian * (180 / math.Pi) }

// IsPrime checks if value is primer number.
func IsPrime(val int) bool {
	if val < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(val))); i++ {
		if val%i == 0 {
			return false
		}
	}

	return true
}

// GCD returns greatest common divisor (GCD) of passed integers.
func GCD[T constraints.Integer](integers ...T) T {
	ret := integers[0]

	for k := range integers {
		ret = gcd(ret, integers[k])

		if ret == 1 {
			return 1
		}
	}

	return ret
}

func gcd[T constraints.Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// LCM returns least common multiple (LCM) of passed integers.
func LCM[T constraints.Integer](integers ...T) T {
	ret := integers[0]

	for i := range integers {
		ret = lcm(integers[i], ret)
	}

	return ret
}

func lcm[T constraints.Integer](a, b T) T {
	if a == 0 || b == 0 {
		panic("LCM: provide non zero integers only.")
	}

	return a * b / gcd(a, b)
}
