package main

import (
	"advent/pkg/array"
	"advent/pkg/vector"
	"bufio"
	"fmt"
	"os"
)

func main() {
	wrh := make(map[vector.Vec2[int]]rune)
	var inst array.Array[rune]
	var rb vector.Vec2[int]

	file, err := os.Open("2024/tasks/day15_sample.txt")
	if err != nil { fmt.Println(err); return }
	defer file.Close()

	instSet := false
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan(){
		if scanner.Text() == "" { instSet = true; continue }
		if instSet { inst = append(inst, []rune(scanner.Text())...); continue }
		for x, r := range scanner.Text(){
			if r == '.' { continue }
			if r == '@' { rb = vector.Vec2[int]{X: x, Y: y}}
			wrh[vector.Vec2[int]{X: x, Y: y}] = r
		}
		y++
	}
	lwrh := wrhEnlarge(wrh)

	// var x, z int
	// for p, _ := range wrh {
	// 	if p.X > x { x = p.X }
	// 	if p.Y > z { z = p.Y }
	// }

	// for i := 0; i <= z; i++ {
	// 	for j := 0; j <= x; j++ {
	// 		r, ok := wrh[vector.Vec2[int]{X: j, Y: i}]
	// 		if ok { fmt.Printf("%c", r)} else { fmt.Print(".") }
	// 	}
	// 	fmt.Printf("\n")
	// } 
	// for p, _ := range lwrh {
	// 	if p.X > x { x = p.X }
	// 	if p.Y > z { z = p.Y }
	// }

	// for i := 0; i <= z; i++ {
	// 	for j := 0; j <= x; j++ {
	// 		r, ok := lwrh[vector.Vec2[int]{X: j, Y: i}]
	// 		if ok { fmt.Printf("%c", r)} else { fmt.Print(".") }
	// 	}
	// 	fmt.Printf("\n")
	// } 
	fmt.Println(simulate(wrh, inst, rb))
	rb = vector.Vec2[int]{X: rb.X*2, Y: rb.Y}
	fmt.Println(simulate(lwrh, inst, rb))
}

func simulate(wrh map[vector.Vec2[int]]rune, inst array.Array[rune], rb vector.Vec2[int]) int {
	for _, r := range inst {
		if move(&wrh, r, rb) { rb = rb.Direction(r) }

		// var x, y int
		// for p, _ := range wrh {
		// 	if p.X > x { x = p.X }
		// 	if p.Y > y { y = p.Y }
		// }

		// for i := 0; i <= y; i++ {
		// 	for j := 0; j <= x; j++ {
		// 		r, ok := wrh[vector.Vec2[int]{X: j, Y: i}]
		// 		if ok { fmt.Printf("%c", r)} else { fmt.Print(".") }
		// 	}
		// 	fmt.Printf("\n")
		// } 
	}

	var gps int
	for pos, r := range wrh{
		if r == 'O' { gps += pos.X + pos.Y * 100 }
		if r == '[' { gps += pos.X + pos.Y * 100 }
	}

	
	return gps
}

func move(wrh *map[vector.Vec2[int]]rune, dir rune, pos vector.Vec2[int]) bool {
	// r := (*wrh)[pos]
	r, r2 := (*wrh)[pos], '0'
	var sc vector.Vec2[int]
	if r == '#' { return false }	
	if r == '[' { sc = pos.Right() }
	if r == ']' { sc = pos.Left() }
	//np := pos.Direction(dir)
	np, np2 := pos.Direction(dir), vector.Vec2[int]{}
	if (sc != vector.Vec2[int]{}) { np2 = sc.Direction(dir); r2 = (*wrh)[sc]; }

	skip := false
	if !(r == '[' && dir == '>') && !(r == ']' && dir == '<') {
		if (r == '[' && dir == '<') || (r == ']' && dir == '>') { skip = true }
		_, ok := (*wrh)[np]
		if ok { if !(move(wrh, dir, np)) { return false; } }
	}
	if r2 != '0' && (!(r2 == '[' && dir == '>') && !(r2 == ']' && dir == '<')) {
		if !skip {  
			//fmt.Println(np2, dir, r2, sc)
			_, ok := (*wrh)[np2]
			if ok { if !(move(wrh, dir, np2)) { return false; } }
		}
	}

	if (sc != vector.Vec2[int]{}) {
		// fmt.Println(r, r2)
		delete(*wrh, sc)
		(*wrh)[np2] = r2
	}

	delete(*wrh, pos)
	(*wrh)[np] = r

	return true
}

func wrhEnlarge(wrh map[vector.Vec2[int]]rune) map[vector.Vec2[int]]rune {
	lwrh := make(map[vector.Vec2[int]]rune)
	for pos, r := range wrh {
		enPos := vector.Vec2[int]{X: pos.X*2, Y: pos.Y}
		switch r {
		case '#':
			lwrh[enPos] = r
			lwrh[enPos.Right()] = r
		case 'O':
			lwrh[enPos] = '['
			lwrh[enPos.Right()] = ']'
		case '@':
			lwrh[enPos] = r
		}
	}

	return lwrh
}