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
	if y < 1 { return emt, fmt.Errorf("not a valid matrix")}
	for i := range a { if len(a[i]) != y { return emt, fmt.Errorf("not a XnX matrix") } }

	switch {
	case y == 1:
		return a[0][0], nil
	case y == 2:
		return (a[0][0] * a[1][1]) - (a[0][1] * a[1][0]), nil
	case y > 2:
		var result T
		for i := 0; i < y; i++{
			var b array.Array2D[T]
			for j := 1; j < y; j++{
				var row array.Array[T]
				for k := 0; k < y; k++{
					if k == i { continue }
					row = append(row, a[j][k])
				}
				b = append(b, row)
			}
			val, err := Det(b)
			if err != nil { 
				return emt, err
			}
			if i % 2 == 0{
				result += (a[0][i] * val) 
			} else {
				result -= (a[0][i] * val)
			}
		}
		return result, nil
	default:
		return emt, fmt.Errorf("not implemented support")
	}
}