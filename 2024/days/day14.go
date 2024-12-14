package main

import (
	"advent/pkg/array"
	"advent/pkg/vector"
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("2024/tasks/day14.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr array.Array2D[vector.Vec2[int]]
	for scanner.Scan() {
		var px, py, vx, vy int
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		p, v := vector.Vec2[int]{X: px, Y: py}, vector.Vec2[int]{X: vx, Y: vy}
		arr = append(arr, array.Array[vector.Vec2[int]]{p, v})
	}

	// sample rules
	// fmt.Println(calcPos(arr, 11, 7, 100))
	val, _ := calcPos(arr, 101, 103, 100)
	fmt.Println(val)

	fmt.Println(findTree(arr, 10, 101, 103))
}

func calcPos(arr array.Array2D[vector.Vec2[int]], mW int, mH int, step int) (int, array.Array[vector.Vec2[int]]) {
	var tl, tr, bl, br int
	var found array.Array[vector.Vec2[int]]
	for _, val := range arr {
		tV := val[1].Mul(step)
		p := val[0].Add(tV)
		p, err := p.Modulo(vector.Vec2[int]{X: mW, Y: mH})
		if err != nil {
			return 0, nil
		}
		if p.X < 0 { p.X += mW }
		if p.Y < 0 { p.Y += mH }
		if p.Y < mH/2{
			if p.X < mW/2 { tl++ } else if p.X > mW/2 { tr++ }
		} else if p.Y > mH/2 {
			if p.X < mW/2 { bl++ } else if p.X > mW/2 { br++ }
		}
		found = append(found, p)
	}
	return tl*tr*bl*br, found
}

func findTree(arr array.Array2D[vector.Vec2[int]], h int, mW int, mH int) int {
	var i = 0
	for {
		if i > 10000 { fmt.Println("hit limit"); break }
		_, f := calcPos(arr, mW, mH, i)
		for _, p := range f{
			var c = 0
			test := p
			for {
				if c >= h { fmt.Println("hit hight", c, f); return i }
				if slices.Contains(f, test.Down()) { c++; test = test.Down() } else { break }
			}
		}
		i++
	}

	return 0
}