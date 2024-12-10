package main

import (
	"advent/pkg/input"
	"advent/pkg/vector"
	"fmt"
)

func main() {
	arr, err := input.Array2D("2024/tasks/day10_sample.txt", 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	if trails, ok := arr.([][]int); ok {
		var result1, result2 int
		for y := range trails {
			for x, val := range trails[y] {
				if val == 0 { 
					result := findPath(&trails, vector.Vec2[int]{X: x, Y: y})
					result1 += len(result)
					for _, val := range result {
						result2 += val
					}
				}
			}
		}
		fmt.Println(result1)
		fmt.Println(result2)
	}

}

func findPath(arr *[][]int, pos vector.Vec2[int]) map[vector.Vec2[int]]int {
	if (*arr)[pos.Y][pos.X] == 9 { return map[vector.Vec2[int]]int{pos: 1} }
	result := make(map[vector.Vec2[int]]int)
	find := (*arr)[pos.Y][pos.X] + 1
	for i := -1; i <= 1; i++ {
		if pos.Y+i < 0 { continue }
		if pos.Y+i >= len(*arr) { continue }
		for j :=-1; j <= 1; j++{
			if i == 1 && j == 1 { continue }
			if i == -1 && j == -1 { continue }
			if i == 1 && j == -1 { continue }
			if i == -1 && j == 1 { continue }
			if pos.X+j < 0 { continue }
			if pos.X+j >= len(*arr) { continue }
			if (*arr)[pos.Y+i][pos.X+j] == find { joinMap(result, findPath(arr, vector.Vec2[int]{X: pos.X+j, Y: pos.Y+i})) }
		}
	}
	return result
}

func joinMap(map1 map[vector.Vec2[int]]int, map2 map[vector.Vec2[int]]int) map[vector.Vec2[int]]int {
	for key, val := range map2 {
		if _, exists := map1[key]; exists {
			map1[key] += val
		} else {
			map1[key] = val
		}
	}
	return map1
}