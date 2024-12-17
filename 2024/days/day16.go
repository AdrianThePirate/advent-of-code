package main

import (
	"advent/pkg/array"
	"advent/pkg/vector"
	"bufio"
	"fmt"
	"os"
	"slices"
)

type node struct {
	pos position
	g, h, f int
	parent *node
}

type position struct {
	pos vector.Vec2[int]
	deg int
}

func main() {
	var walls array.Array[vector.Vec2[int]]
	var goal vector.Vec2[int]
	var rain position
	
	file, err := os.Open("2024/tasks/day16.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var y int
	for scanner.Scan() {
		for x, r := range scanner.Text(){
			switch r {
			case '#': walls = append(walls, vector.Vec2[int]{X: x, Y: y})
			case 'S': rain = position{pos: vector.Vec2[int]{X: x, Y: y}, deg: 1}
			case 'E': goal = vector.Vec2[int]{X: x, Y: y}
			}
		}
		y++
	}
	fmt.Println(aStarAlg(rain, goal, walls))
}

func aStarAlg(rain position, goal vector.Vec2[int], walls array.Array[vector.Vec2[int]]) int {
	var openSet array.Array[node]
	openSet = append(openSet, node{pos: rain})

	closedSet := make(map[position]bool)
	cameFrom := make(map[position]*node)

	for len(openSet) > 0 {
		current := lowest(&openSet)

		if current.pos.pos == goal {
			//println(backtrace(cameFrom, &current))
			return current.g
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

			if slices.Contains(walls, neighPos.pos) {	continue }

			tentG := current.g+1
			if neighPos.deg != current.pos.deg { tentG = current.g+1000 }
			neigh := node{
				pos: neighPos,
				g: tentG,
				h: int(neighPos.pos.DistanceTo(goal)),
				f: tentG + int(neighPos.pos.DistanceTo(goal)),
				parent: &current,
			}

			skip := false
			for _, n := range openSet {
				if n.pos == neighPos && tentG >= n.g {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			cameFrom[neighPos] = &current
			openSet = append(openSet, neigh)
		}
	} 

	return 0
}

func backtrace(cf map[position]*node, s *node) int{
	var l int
	for n := s; n != nil; n = cf[n.pos] {
		fmt.Println(n.pos, n.g)
		l++
	}
	return l
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