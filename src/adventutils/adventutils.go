package adventutils

import "math"

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

func (v1 Vec2[T]) Sub(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vec2[T]) Add(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vec2[T]) Mul(arg interface{}) Vec2[T] {
	switch v2 := arg.(type) {
	case Vec2[T]:
		return Vec2[T]{
			X: v1.X * v2.X,
			Y: v1.Y * v2.Y,
		}
	case T:
		return Vec2[T]{
			X: v1.X * v2,
			Y: v1.Y * v2,
		}
	default:
		panic("unsupported type")
	}
}

func (v1 Vec2[T]) Div(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X / v2.X,
		Y: v1.Y / v2.Y,
	}
}


func (v1 Vec2[T]) Len(arg interface{}) float64 {
	switch v2 := arg.(type){
	case nil:
		return math.Sqrt(float64(v1.Dot(v1)))
	case Vec2[T]:
		if v1 == v2 { return 0 }
		return math.Sqrt(float64(v1.Dot(v2)))
	default:
		panic("unsupported type")
	}
}

func (v1 Vec2[T]) Dot(v2 Vec2[T]) T{
	return (v1.X * v2.X) + (v1.Y * v2.Y)
} 