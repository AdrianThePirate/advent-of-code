package adventutils

func Absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func InsertIndex[T any](array []T, value T, index int) []T {
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func RemoveIndex[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func MoveIndex[T any](array []T, srcIdx int, destIdx int) []T {
	value := array[srcIdx]
	return InsertIndex(RemoveIndex(array, srcIdx), value, destIdx)
}

type Numeric interface {
	int | int32 | int64 | float32 | float64
}

type Vec2[T Numeric] struct {
	X, Y T
}

func IsInSet[T comparable](val T, set []T) bool {
	for _, v := range set {
		if val == v {
			return true
		}
	}
	return false
}