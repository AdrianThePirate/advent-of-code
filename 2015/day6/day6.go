package day6

import (
	"fmt"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

func Run(variant string) {
	lines, err := input.FileToLines("2015/day6/day6.txt")
	if err != nil{
		fmt.Println(err)
		return
	}

    part1(lines)
	part2(lines)
}

func part1(lines []string) {
	grid := make(array.Array2D[bool], 1000)
	for i := range grid {
		grid[i] = make(array.Array[bool], 1000)
	}

	for _, line := range lines{
		parts := strings.Fields(line)
		coordA := parts[len(parts)-3]
		coordB := parts[len(parts)-1]
		action := parts[0]
		if action == "turn" {
			action = parts[1]
		}

		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Sscanf(coordA, "%d,%d", &x1, &y1)
		fmt.Sscanf(coordB, "%d,%d", &x2, &y2)
		
		for x := x1; x <= x2; x++{
			for y := y1; y <= y2; y++{
				target := vector.Vec2[int]{X: x, Y: y}
				switch action{
				case "toggle":
					flip, err := grid.GetPos(target)
					if err != nil {
						fmt.Println(err)
						return
					}
					err = grid.SetPos(target, !flip)
					if err != nil {
						fmt.Println(err)
						return
					}
				case "on":
					err := grid.SetPos(target, true)
					if err != nil {
						fmt.Println(err)
						return
					}
				case "off":
					err := grid.SetPos(target, false)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
	
	lit := 0
	for i := range grid{
		for _, val := range grid[i]{
			if val {lit++}
		}
	}
	fmt.Printf("There are %d lit lights\n", lit)
}

func part2(lines []string) {
	grid := make(array.Array2D[int], 1000)
	for i := range grid {
		grid[i] = make(array.Array[int], 1000)
	}

	for _, line := range lines{
		parts := strings.Fields(line)
		coordA := parts[len(parts)-3]
		coordB := parts[len(parts)-1]
		action := parts[0]
		if action == "turn" {
			action = parts[1]
		}

		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Sscanf(coordA, "%d,%d", &x1, &y1)
		fmt.Sscanf(coordB, "%d,%d", &x2, &y2)
		
		for x := x1; x <= x2; x++{
			for y := y1; y <= y2; y++{
				target := vector.Vec2[int]{X: x, Y: y}
				base, err := grid.GetPos(target)
					if err != nil {
						fmt.Println(err)
						return
					}
				switch action{
				case "toggle":
					err = grid.SetPos(target, base+2)
					if err != nil {
						fmt.Println(err)
						return
					}
				case "on":
					err := grid.SetPos(target, base+1)
					if err != nil {
						fmt.Println(err)
						return
					}
				case "off":
					if base == 0 { continue }
					err := grid.SetPos(target, base-1)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
	
	bright := 0
	for i := range grid{
		for _, val := range grid[i]{
			bright += val
		}
	}
	fmt.Printf("Brightnes level %d\n", bright)
}