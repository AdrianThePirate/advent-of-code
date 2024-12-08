package main

import (
	"advent/adventutils"
	"bufio"
	"fmt"
	"os"
	"slices"
)

var board [][]rune
var obstructions_pos []adventutils.Vec2[int]
var guardPos adventutils.Vec2[int]
var crossing = false

func main() {
	populateArray()
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
			if slices.Contains([]rune{'|','-','+', 'v'}, board[y][x]) { obstructions_pos = append(obstructions_pos, adventutils.Vec2[int]{X: x, Y: y}) }
		}
	}

	var result int
	for _, pos := range obstructions_pos {
		populateArray()
		if pos == guardPos { continue }
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

	walks[guardPos] = append(walks[guardPos], board[guardPos.Y][guardPos.X])

	var turned = false
	for {
		switch board[guardPos.Y][guardPos.X]{
		case '^':
			if guardPos.Y == 0 { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guardPos.X, Y: guardPos.Y-1}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '^') { return true, len(walks)+1 }
			if slices.Contains([]rune{'#', 'O'}, next) {
				board[guardPos.Y][guardPos.X] = '>'
				turned = true
			} else {
				walks[guardPos] = append(walks[guardPos], '^')
				moveGuard(next_pos, '^', &turned)
			}
		case '<':
			if guardPos.X == 0 { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guardPos.X-1, Y: guardPos.Y}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '<') { return true, len(walks)+1 }
			if slices.Contains([]rune{'#', 'O'}, next) {
				board[guardPos.Y][guardPos.X] = '^'
				turned = true
			} else {
				walks[guardPos] = append(walks[guardPos], '<')
				moveGuard(next_pos, '<', &turned)
			}
		case '>':
			if guardPos.X+1 == len(board[guardPos.Y]) { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guardPos.X+1, Y: guardPos.Y}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '>') { return true, len(walks)+1 }
			if slices.Contains([]rune{'#', 'O'}, next) {
				board[guardPos.Y][guardPos.X] = 'v'
				turned = true
			} else {
				walks[guardPos] = append(walks[guardPos], '>')
				moveGuard(next_pos, '>', &turned)
			}
		case 'v':
			if guardPos.Y+1 == len(board) { return false, len(walks)+1 }
			next_pos := adventutils.Vec2[int]{X: guardPos.X, Y: guardPos.Y+1}
			next := board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], 'v') { return true, len(walks)+1 }
			if slices.Contains([]rune{'#', 'O'}, next) {
				board[guardPos.Y][guardPos.X] = '<'
				turned = true
			} else {
				walks[guardPos] = append(walks[guardPos], 'v')
				moveGuard(next_pos, 'v', &turned)
			}
		}
	}
}

func moveGuard(pos adventutils.Vec2[int], dir rune, turned *bool) {
	if *turned {
		board[guardPos.Y][guardPos.X] = '+'
		*turned = false
	} else if crossing {
		board[guardPos.Y][guardPos.X] = '+'
		crossing = false
	} else if slices.Contains([]rune{'<', '>'}, dir) {
		board[guardPos.Y][guardPos.X] = '-'
	} else {
		board[guardPos.Y][guardPos.X] = '|'
	}
	if slices.Contains([]rune{'|','-'}, board[pos.Y][pos.X]) { crossing = true }
	board[pos.Y][pos.X] = dir
	guardPos = pos
}

func populateArray() {
	file, err := os.Open("2024/tasks/day6_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	board = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		board = append(board, []rune(scanner.Text()))
	}

	var guardFound = false
	for y, line := range board {
		for x, val := range line {
			if slices.Contains([]rune{'^', '<', '>', 'v'}, val) {
				guardPos.X = x
				guardPos.Y = y
				guardFound = true
				break
			}
			if guardFound { break }
		}
	}
}
