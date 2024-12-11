package math

import "math"

func Absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GetDigitCount(n int) int {
	Absolute(n)
	if n == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(float64(n))) + 1)

}