package matrix

import (
	"advent/pkg/array"
	"fmt"
)

type Numeric interface {
	int | int32 | int64 | float32 | float64
}

func Det[T Numeric](a array.Array2D[T]) (T, error){
	y := len(a)
	var emt T
	if y < 2 { return emt, fmt.Errorf("not a valid matrix")}
	for i := range a { if len(a[i]) != y { return emt, fmt.Errorf("not a XnX matrix") } }

	switch y {
	case 2:
		return (a[0][0] * a[1][1]) - (a[0][1] * a[1][0]), nil
	default:
		return emt, fmt.Errorf("not implemented support")
	}
}