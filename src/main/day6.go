package main

import (
	"adventutils"
	"bufio"
	"fmt"
	"os"
	"slices"
)

var board [][]rune
var obstructions_pos []adventutils.Vec2[int]
var guard_pos adventutils.Vec2[int]
var crossing = false

func main() {
	populate_array()
	part1()
	part2()
}

func part1() {
	var _, result = simulate()

	fmt.Println("Result", result)
}

func part2() {
	for y := range board {
		for x := range board[y] {
			if adventutils.IsInSet(board[y][x], []rune{'|','-','+', 'v'}) { obstructions_pos = append(obstructions_pos, adventutils.Vec2[int]{X: x, Y: y}) }
		}
	}

	var result int
	for _, pos := range obstructions_pos {
		populate_array()
		if pos == guard_pos { continue }
		board[pos.Y][pos.X] = 'O'
		val, _ := simulate() 
		if val { 
			result++
		}
	}

	fmt.Println("Result", result)
}

func simulate() (bool, int) {
	var walks map[adventutils.Vec2[int]][]rune
	walks = nil
	walks = make(map[adventutils.Vec2[int]][]rune)

	walks[guard_pos] = append(walks[guard_pos], board[guard_pos.Y][guard_pos.X])

	var turned = false
	for {
		switch board[guard_pos.Y][guard_pos.X]{
		case '^':
			if guard_pos.Y == 0 { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guard_pos.X, Y: guard_pos.Y-1}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '^') { return true, len(walks)+1 }
			if adventutils.IsInSet(next, []rune{'#', 'O'}) {
				board[guard_pos.Y][guard_pos.X] = '>'
				turned = true
			} else {
				walks[guard_pos] = append(walks[guard_pos], '^')
				move_guard(next_pos, '^', &turned)
			}
		case '<':
			if guard_pos.X == 0 { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guard_pos.X-1, Y: guard_pos.Y}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '<') { return true, len(walks)+1 }
			if adventutils.IsInSet(next, []rune{'#', 'O'}) {
				board[guard_pos.Y][guard_pos.X] = '^'
				turned = true
			} else {
				walks[guard_pos] = append(walks[guard_pos], '<')
				move_guard(next_pos, '<', &turned)
			}
		case '>':
			if guard_pos.X+1 == len(board[guard_pos.Y]) { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guard_pos.X+1, Y: guard_pos.Y}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '>') { return true, len(walks)+1 }
			if adventutils.IsInSet(next, []rune{'#', 'O'}) {
				board[guard_pos.Y][guard_pos.X] = 'v'
				turned = true
			} else {
				walks[guard_pos] = append(walks[guard_pos], '>')
				move_guard(next_pos, '>', &turned)
			}
		case 'v':
			if guard_pos.Y+1 == len(board) { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guard_pos.X, Y: guard_pos.Y+1}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], 'v') { return true, len(walks)+1 }
			if adventutils.IsInSet(next, []rune{'#', 'O'}) {
				board[guard_pos.Y][guard_pos.X] = '<'
				turned = true
			} else {
				walks[guard_pos] = append(walks[guard_pos], 'v')
				move_guard(next_pos, 'v', &turned)
			}
		}
	}
}

func move_guard(pos adventutils.Vec2[int], dir rune, turned *bool) {
	if *turned {
		board[guard_pos.Y][guard_pos.X] = '+'
		*turned = false
	} else if crossing {
		board[guard_pos.Y][guard_pos.X] = '+'
		crossing = false
	} else if adventutils.IsInSet(dir, []rune{'<', '>'}) {
		board[guard_pos.Y][guard_pos.X] = '-'
	} else {
		board[guard_pos.Y][guard_pos.X] = '|'
	}
	if adventutils.IsInSet(board[pos.Y][pos.X], []rune{'|','-'}) { crossing = true }
	board[pos.Y][pos.X] = dir
	guard_pos = pos
}

func populate_array() {
	file, err := os.Open("tasks/day6_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	board = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		board = append(board, []rune(scanner.Text()))
	}

	var guard_found = false
	for y, line := range board {
		for x, val := range line {
			if adventutils.IsInSet(val, []rune{'^', '<', '>', 'v'}) {
				guard_pos.X = x
				guard_pos.Y = y
				guard_found = true
				break
			}
			if guard_found { break }
		}
	}
}
