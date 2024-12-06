package main

import (
	"adventutils"
	"bufio"
	"fmt"
	"os"
)

var board [][]rune
var guard_pos adventutils.Vec2[int]

func main() {
	populate_array()
	part1()
}

func part1() {
	var result int = 1
	var exit = false
	for {
		switch board[guard_pos.Y][guard_pos.X]{
		case '^':
			if guard_pos.Y == 0 { exit = true; break; }
			next := board[guard_pos.Y-1][guard_pos.X]
			if next  == '#' {
				board[guard_pos.Y][guard_pos.X] = '>'
			} else {
				if next == '.' { result++ }
				move_guard(guard_pos.X, guard_pos.Y-1, '^')
			}
		case '<':
			if guard_pos.X == 0 { exit = true; break; }
			next := board[guard_pos.Y][guard_pos.X-1]
			if next == '#' {
				board[guard_pos.Y][guard_pos.X] = '^'
			} else {
				if next == '.' { result++ }
				move_guard(guard_pos.X-1, guard_pos.Y, '<')
			}
			
		case '>':
			if guard_pos.X+1 == len(board[guard_pos.Y]) { exit = true; break; }
			next := board[guard_pos.Y][guard_pos.X+1]
			if next == '#' {
				board[guard_pos.Y][guard_pos.X] = 'v'
			} else {
				if next == '.' { result++ }
				move_guard(guard_pos.X+1, guard_pos.Y, '>')
			}
		case 'v':
			if guard_pos.Y+1 == len(board) { exit = true; break; }
			next := board[guard_pos.Y+1][guard_pos.X]
			if next == '#' {
				board[guard_pos.Y][guard_pos.X] = '<'
			} else {
				if next == '.' { result++ }
				move_guard(guard_pos.X, guard_pos.Y+1, 'v')
			}
		}
		if exit { break }
	}

	fmt.Println("Result", result)
}

func move_guard(x int, y int, dir rune){
	board[guard_pos.Y][guard_pos.X] = 'B'
	board[y][x] = dir
	guard_pos.X = x
	guard_pos.Y = y
}

func populate_array() {
	file, err := os.Open("tasks/day6_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

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