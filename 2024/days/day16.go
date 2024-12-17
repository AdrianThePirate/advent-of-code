package main

import (
	"advent/pkg/array"
	"advent/pkg/vector"
	"bufio"
	"fmt"
	"os"
)

type node struct {
	pos position
	g int
	parent *node
}

type position struct {
	pos vector.Vec2[int]
	deg int
}

func main() {
	walls := make(map[vector.Vec2[int]]bool)
	var goal vector.Vec2[int]
	var rain position
	
	file, err := os.Open("2024/tasks/day16_sample1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var y int
	for scanner.Scan() {
		for x, r := range scanner.Text(){
			switch r {
			case '#': walls[vector.Vec2[int]{X: x, Y: y}] = true
			case 'S': rain = position{pos: vector.Vec2[int]{X: x, Y: y}, deg: 1}
			case 'E': goal = vector.Vec2[int]{X: x, Y: y}
			}
		}
		y++
	}
	cost, seats := aStarAlg(rain, goal, walls)
	fmt.Printf("Cost: %d Seats: %d\n", cost, seats)
}

func aStarAlg(rain position, goal vector.Vec2[int], walls map[vector.Vec2[int]]bool) (int, int) {
	var openSet array.Array[node]
	openSet = append(openSet, node{pos: rain})

	closedSet := make(map[position]bool)

	var paths array.Array[node]
	var maxG int

	for len(openSet) > 0 {
		current := lowest(&openSet)

		if current.pos.pos == goal {
			paths = append(paths, current)
			maxG = current.g
			//println(backtrace(cameFrom, &current))
		}

		closedSet[current.pos] = true

		var options array.Array[position]
		switch current.pos.deg {
			case 0: options = append(options, position{pos: current.pos.pos.Up(), deg: current.pos.deg})
			case 1: options = append(options, position{pos: current.pos.pos.Right(), deg: current.pos.deg})
			case 2: options = append(options, position{pos: current.pos.pos.Down(), deg: current.pos.deg})
			case 3: options = append(options, position{pos: current.pos.pos.Left(), deg: current.pos.deg})
		}
		options = append(options, position{pos: current.pos.pos, deg: (current.pos.deg-1+4)%4})
		options = append(options, position{pos: current.pos.pos, deg: (current.pos.deg+1+4)%4})

		for _, neighPos := range options{
			if closedSet[neighPos] { continue }

			if walls[neighPos.pos] { continue }

			tentG := current.g+1
			if neighPos.deg != current.pos.deg { tentG = current.g+1000 }
			neigh := node{
				pos: neighPos,
				g: tentG,
				parent: &current,
			}

			if maxG != 0 && tentG > maxG { continue }

			skip := false
			for _, n := range openSet {
				if n.pos == neighPos && tentG > n.g {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			openSet = append(openSet, neigh)
		}
	} 

	return maxG, backtrace(paths)
}

func backtrace(arr array.Array[node]) int{
	seats := make(map[vector.Vec2[int]]bool)
	for _, n := range arr {
		for {
			seats[n.pos.pos] = true
			if n.parent != nil { n = *n.parent } else { break }
		}
	}
	return len(seats)
}

func lowest(arr *array.Array[node]) node {
	low, lowI:= (*arr)[0].g, 0
	for i, v := range *arr {
		if v.g < low { low = v.g; lowI = i}
	}
	n := (*arr)[lowI]
	arr.RemoveIndex(lowI)
	return n
}