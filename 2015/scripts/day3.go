package main

import (
	"advent/pkg/vector"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("2015/tasks/day3.txt")
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
	if robotMode { visited[robot]++ }

	robotTurn := false
	for _, r := range data {
		if robotMode && robotTurn{
			moveDir(r, &robot)
			visited[robot]++
			robotTurn = false
		} else {
			moveDir(r, &santa)
			visited[santa]++
			robotTurn = true
		}
	}
	return visited
}

func moveDir(r byte, pos *vector.Vec2[int]) {
	switch r{
		case '^':
			pos.Y -= 1	
		case 'v':
			pos.Y += 1
		case '<':
			pos.X -= 1
		case '>':
			pos.X += 1
		}
}