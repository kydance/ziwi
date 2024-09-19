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
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Exponent calculates x^n
func Exponent(x, n int64) int64 {
	if n == 0 {
		return 1
	}

	t := Exponent(x, n/2)
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}

// Factorial calculates `x!`
func Factorial(x uint) uint {
	var f uint = 1
	for ; x > 1; x-- {
		f *= x
	}
	return f
}

// RoundToFloat rounds a number to the specified number of decimal places.
func RoundToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	tmp := math.Pow(10.0, float64(n))
	x *= T(tmp)
	return math.Round(float64(x)) / tmp
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

	// Multiply the input value by the calculated power of 10.
	x *= T(temp)

	// Round the multiplied value to the nearest integer.
	r := math.Round(float64(x))

	// Format the rounded value back to a string with the specified number of decimal places.
	return strconv.FormatFloat(r/temp, 'f', n, 64)
}

// Percent calculates the percentage of value to total.
func Percent(val, total float64, n int) float64 {
	if total == 0 {
		return float64(0)
	}

	tmp := val / total * 100
	return RoundToFloat(tmp, n)
}

// TruncRound rounds a floating-point number to a specified number of decimal places.
// It uses string manipulation to achieve the rounding and then converts the result
// to the original type.
//
// Parameters:
//
//	x: The number to be rounded.
//	n: The number of decimal places to round to.
//
// Returns:
//
//	T: The rounded number of the same type as the input.
func TruncRound[T constraints.Float | constraints.Integer](x T, n int) T {
	// Convert the number to a string with n+1 decimal places to handle rounding
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", float64(x))
	// Split the string into integer and fractional parts
	temp := strings.Split(floatStr, ".")

	var newFloat string
	// If there is no fractional part or the number of decimal places is
	// greater than or equal to the length of the fractional part,
	// use the original string representation
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		// Otherwise, truncate the fractional part to n decimal places
		newFloat = temp[0] + "." + temp[1][:n]
	}

	// Parse the new string representation back to a float64
	result, _ := strconv.ParseFloat(newFloat, 64)
	// Convert the result back to the original type and return it
	return T(result)
}

// FloorToFloat round down to n decimal places.
func FloorToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	return math.Floor(float64(x)) / temp
}

// FloorToStr round down to n decimal places.
func FloorToStr[T constraints.Float | constraints.Integer](x T, n int) string {
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	r := math.Floor(float64(x))
	return strconv.FormatFloat(r/temp, 'f', n, 64)
}

// CeilToFloat round up to n decimal places.
func CeilToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	return math.Ceil(float64(x)) / temp
}

// CeilToStr round up to n decimal places.
func CeilToStr[T constraints.Float | constraints.Integer](x T, n int) string {
	temp := math.Pow(10.0, float64(n))
	x *= T(temp)
	r := math.Ceil(float64(x))
	return strconv.FormatFloat(r/temp, 'f', n, 64)
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
func AngleToRadian(angle float64) float64 {
	return angle * (math.Pi / 180)
}

// RadianToAngle converts radian value to angle value.
func RadianToAngle(radian float64) float64 {
	return radian * (180 / math.Pi)
}

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
