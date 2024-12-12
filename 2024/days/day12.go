package main

import (
	"advent/pkg/array"
	"advent/pkg/input"
	"advent/pkg/vector"
	"fmt"
	"slices"
)

func main() {
	farm, err := input.FileToArray2D[rune]("2024/tasks/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fencePrice(farm))
}

func fencePrice(farm array.Array2D[rune]) int{
	price := 0

	for y, l := range farm {
		for x, r := range l {
			var v []vector.Vec2[int]
			if r != '.' { if area, border, err := sizeRegion(&farm, vector.Vec2[int]{X: x, Y: y}, &v); err == nil { price += area * border } }

		}
	}

	return price
}

func sizeRegion(farm *array.Array2D[rune], pos vector.Vec2[int], visited *[]vector.Vec2[int]) (int, int, error) {
	if slices.Contains(*visited, pos) { return 0, 0, nil }
	*visited = append(*visited, pos)
	var area, border int
	t, err := farm.GetPos(pos)
	if err != nil {
		return 0, 0, err
	}

	area += 1
	if val, err := farm.GetPos(pos.Up()); err == nil && val == t { if a1, b1, err := sizeRegion(farm, pos.Up(), visited); err == nil { area += a1; border += b1 } } else if !slices.Contains(*visited, pos.Up()) { border += 1 }
	if val, err := farm.GetPos(pos.Down()); err == nil && val == t { if a1, b1, err := sizeRegion(farm, pos.Down(), visited); err == nil { area += a1; border += b1 } } else if !slices.Contains(*visited, pos.Down()) { border += 1 }
	if val, err := farm.GetPos(pos.Left()); err == nil && val == t { if a1, b1, err := sizeRegion(farm, pos.Left(), visited); err == nil { area += a1; border += b1 } } else if !slices.Contains(*visited, pos.Left()) { border += 1 }
	if val, err := farm.GetPos(pos.Right()); err == nil && val == t { if a1, b1, err := sizeRegion(farm, pos.Right(), visited); err == nil { area += a1; border += b1 } } else if !slices.Contains(*visited, pos.Right()) { border += 1 }

	farm.SetPos(pos, '.')

	return area, border, nil
}