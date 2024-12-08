package main

import (
	"adventutils"
	"bufio"
	"fmt"
	"os"
)

var antPos map[rune][]adventutils.Vec2[int]
var maxX, maxY int

func main() {
	antPos = make(map[rune][]adventutils.Vec2[int])
	populateMap()
	
	fmt.Println("Result", len(findAnt()))
	fmt.Println("Result", len(fintAntReson()))
}

func findAnt() map[adventutils.Vec2[int]]int {
	var foundAnt map[adventutils.Vec2[int]]int
	foundAnt = make(map[adventutils.Vec2[int]]int)
	for _, poss := range antPos {
		for _, loc := range poss {
			for _, pos := range poss {
				if pos.Len(loc) == 0 { continue }
				antPos := pos.Add(pos.Sub(loc))
				if antPos.X >= 0 && antPos.X < maxX && antPos.Y >= 0 && antPos.Y < maxY { foundAnt[antPos]++ }
			}
		}
	}
	return foundAnt
}

func fintAntReson() map[adventutils.Vec2[int]]int {
	var foundAnt map[adventutils.Vec2[int]]int
	foundAnt = make(map[adventutils.Vec2[int]]int)
	for _, poss := range antPos {
		for _, loc := range poss {
			for _, pos := range poss {
				if pos.Len(loc) == 0 { continue }
				vector := pos.Sub(loc)
				var i int = 1
				for {
					antPos := loc.Add(vector.Mul(i))
					if antPos.X >= 0 && antPos.X < maxX && antPos.Y >= 0 && antPos.Y < maxY {
						 foundAnt[antPos]++ 
					} else {
						break
					}
					i++
				}
			}
		}
	}
	return foundAnt
}

func populateMap() {
	file, err := os.Open("tasks/day8_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		xLine := []rune(scanner.Text())
		if len(xLine) > maxX { maxX = len(xLine) }
		for j, val := range xLine {
			if val != '.' {
				antPos[val] = append(antPos[val], adventutils.Vec2[int]{X:j, Y:i})
			}
		}
		i++
	}
	maxY = i
}