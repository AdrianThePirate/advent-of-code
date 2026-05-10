package day8

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

func Run() {
	antPos, maxX, maxY := populateMap()

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
				vector := pos.Sub(loc)
				for i := 1; ; i++ {
					antPos := loc.Add(vector.Mul(i))
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

func populateMap() (map[rune][]vector.Vec2[int], int, int) {
	file, err := os.Open("2024/days/day8/day8_sample.txt")
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0
	}
	defer file.Close()

	antPos := make(map[rune][]vector.Vec2[int])
	var maxX, maxY int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xLine := []rune(scanner.Text())
		if len(xLine) > maxX {
			maxX = len(xLine)
		}
		for j, val := range xLine {
			if val != '.' {
				antPos[val] = append(antPos[val], vector.Vec2[int]{X: j, Y: maxY})
			}
		}
		maxY++
	}
	return antPos, maxX, maxY
}
