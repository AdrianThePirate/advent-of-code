package main

import (
	"advent/pkg/array"
	"advent/pkg/input"
	"advent/pkg/vector"
	"fmt"
	"slices"
)

func main() {
	farm, err := input.FileToArray2D[rune]("2024/tasks/day12_sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	normal, bulk := fencePrice(farm)
	fmt.Printf("Normal: %d\nBulk: %d", normal, bulk)
}

func fencePrice(farm array.Array2D[rune]) (int, int) {
	price := 0
	priceBulk := 0

	for y, l := range farm {
		for x, r := range l {
			var v []vector.Vec2[int]
			pos := vector.Vec2[int]{X: x, Y: y}
			if r != '.' { if area, border, err := sizeRegion(&farm, pos, &v); err == nil { price += area * border; priceBulk += area * findEdges(v) } }

		}
	}

	return price, priceBulk
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

func findEdges(visited []vector.Vec2[int]) int{
	corners := 0
		
	for _, val := range visited{
		upVisited := slices.Contains(visited, val.Up())
		leftVisited := slices.Contains(visited, val.Left())
		rightVisited := slices.Contains(visited, val.Right())
		downVisited := slices.Contains(visited, val.Down())

		if !upVisited && !leftVisited { corners++ }
		if !upVisited && !rightVisited { corners++ }
		if !downVisited && !leftVisited { corners++ }
		if !downVisited && !rightVisited { corners++ }

		if upVisited && leftVisited {
			if !slices.Contains(visited, val.Up().Left()) { corners++ }
		}
		if upVisited && rightVisited {
			if !slices.Contains(visited, val.Up().Right()) { corners++ }
		}
		if downVisited && leftVisited {
			if !slices.Contains(visited, val.Down().Left()) { corners++ }
		}
		if downVisited && rightVisited {
			if !slices.Contains(visited, val.Down().Right()) { corners++ }
		}
	}

	return corners
}