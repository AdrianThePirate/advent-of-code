package array

type Array[T any] struct {
	a []T
}

func (a *Array[T]) InsertIndex(value T, index int) {
	a.a = append(a.a[:index], append([]T{value}, a.a[index:]...)...)
}

func (a *Array[T]) RemoveIndex(index int) {
	a.a = append(a.a[:index], a.a[index+1:]...)
}

func (a *Array[T]) MoveIndex(srcIdx int, destIdx int) {
	value := a.a[srcIdx]
	a.RemoveIndex(srcIdx)
	if srcIdx < destIdx {
		destIdx--
	}
	a.InsertIndex(value, destIdx)
}

func (a *Array[T]) Append(value T) {
	a.a = append(a.a, value)
}

func (a *Array[T]) Get() *[]T {
	return &a.a
}

func (a *Array[T]) GetIndex(index int) T {
	return a.a[index]
}

func (a *Array[T]) Change(index int, value T) {
	a.a[index] = value
}

func (a *Array[T]) Length() int {
	return len(a.a)
}