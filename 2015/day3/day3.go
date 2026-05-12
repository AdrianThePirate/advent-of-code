package day3

import (
	"fmt"
	"os"

	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

func Run(variant string) {
	data, err := os.ReadFile("2015/day3/day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Santa visited:", len(allVisits(data, false)))
	fmt.Println("Santa & Robo-santa visited:", len(allVisits(data, true)))
}

func allVisits(data []byte, robotMode bool) map[vector.Vec2[int]]int {
	visited := make(map[vector.Vec2[int]]int)
	santa := vector.Vec2[int]{X: 0, Y: 0}
	robot := vector.Vec2[int]{X: 0, Y: 0}
	visited[santa]++
	if robotMode {
		visited[robot]++
	}

	robotTurn := false
	for _, r := range data {
		if robotMode && robotTurn {
			robot = robot.Direction(rune(r))
			visited[robot]++
			robotTurn = false
		} else {
			santa = santa.Direction(rune(r))
			visited[santa]++
			robotTurn = true
		}
	}
	return visited
}