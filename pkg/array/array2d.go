package array

import (
	"advent/pkg/vector"
	"fmt"
)

type Array2D[T any] []Array[T]

func (a *Array2D[T]) GetPos(pos vector.Vec2[int]) (T, error) {
	if pos.Y < 0 || pos.Y >= len(*a) || pos.X < 0 || pos.X >= len((*a)[pos.Y])  { 
		var zeroValue T
		return zeroValue, fmt.Errorf("out of bounce") 
	}
	
	return (*a)[pos.Y][pos.X], nil
}

func (a *Array2D[T]) SetPos(pos vector.Vec2[int], val T) error {
	if pos.Y < 0 || pos.Y >= len(*a) || pos.X < 0 || pos.X >= len((*a)[pos.Y])  { 
		return fmt.Errorf("out of bounce") 
	}
	
	(*a)[pos.Y][pos.X] = val
	return nil
}
