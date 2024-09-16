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
