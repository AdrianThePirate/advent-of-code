package array

type Array[T any] []T

func (a *Array[T]) InsertIndex(value T, index int) {
	*a = append((*a)[:index], append([]T{value}, (*a)[index:]...)...)
}

func (a *Array[T]) RemoveIndex(index int) {
	*a = append((*a)[:index], (*a)[index+1:]...)
}

func (a *Array[T]) MoveIndex(srcIdx int, destIdx int) {
	value := (*a)[srcIdx]
	a.RemoveIndex(srcIdx)
	if srcIdx < destIdx {
		destIdx--
	}
	a.InsertIndex(value, destIdx)
}

func (a *Array[T]) Pop() T {
	var emt T
	if len(*a) == 0 {
		return emt
	}
	value := (*a)[0]
	a.RemoveIndex(0)
	return value
}