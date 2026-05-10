package day8

import (
	"fmt"

	"github.com/AdrianThePirate/advent-of-code/pkg/cmd"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

func Run(variant string) {
	path := cmd.InputPath("2024/days/day8/day8", variant)
	antPos, maxX, maxY := populateMap(path)

	fmt.Println("Result", len(findAnt(antPos, maxX, maxY)))
	fmt.Println("Result", len(findAntReson(antPos, maxX, maxY)))
}

func findAnt(antPos map[rune][]vector.Vec2[int], maxX, maxY int) map[vector.Vec2[int]]int {
	foundAnt := make(map[vector.Vec2[int]]int)
	for _, poss := range antPos {
		for _, loc := range poss {
			for _, pos := range poss {
				if loc.Len(pos) == 0 {
					continue
				}
				antPos := pos.Add(pos.Sub(loc))
				if antPos.X >= 0 && antPos.X < maxX && antPos.Y >= 0 && antPos.Y < maxY {
					foundAnt[antPos]++
				}
			}
		}
	}
	return foundAnt
}

func findAntReson(antPos map[rune][]vector.Vec2[int], maxX, maxY int) map[vector.Vec2[int]]int {
	foundAnt := make(map[vector.Vec2[int]]int)
	for _, poss := range antPos {
		for _, loc := range poss {
			for _, pos := range poss {
				if loc.Len(pos) == 0 {
					continue
				}
				diff := pos.Sub(loc)
				for i := 1; ; i++ {
					antPos := loc.Add(diff.Mul(i))
					if antPos.X >= 0 && antPos.X < maxX && antPos.Y >= 0 && antPos.Y < maxY {
						foundAnt[antPos]++
					} else {
						break
					}
				}
			}
		}
	}
	return foundAnt
}

func populateMap(path string) (map[rune][]vector.Vec2[int], int, int) {
	lines, err := input.FileToLines(path)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0
	}

	antPos := make(map[rune][]vector.Vec2[int])
	var maxX, maxY int
	for y, line := range lines {
		xLine := []rune(line)
		if len(xLine) > maxX {
			maxX = len(xLine)
		}
		for x, val := range xLine {
			if val != '.' {
				antPos[val] = append(antPos[val], vector.Vec2[int]{X: x, Y: y})
			}
		}
		maxY++
	}
	return antPos, maxX, maxY
}
