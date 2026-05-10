package day6

import (
	"fmt"
	"slices"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

type state struct {
	board    array.Array2D[rune]
	guardPos vector.Vec2[int]
	crossing bool
}

func Run() {
	s, err := populateArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	obstructionPos := part1(s)
	part2(obstructionPos)
}

func part1(s state) []vector.Vec2[int] {
	var obstructionPos []vector.Vec2[int]
	simulate(&s)
	for y := range s.board {
		for x := range s.board[y] {
			if slices.Contains([]rune{'|', '-', '+', 'v'}, s.board[y][x]) {
				obstructionPos = append(obstructionPos, vector.Vec2[int]{X: x, Y: y})
			}
		}
	}
	fmt.Println("Result", len(obstructionPos)+1)
	return obstructionPos
}

func part2(obstructionPos []vector.Vec2[int]) {
	var result int
	for _, pos := range obstructionPos {
		s, err := populateArray()
		if err != nil {
			fmt.Println(err)
			return
		}
		if pos == s.guardPos {
			continue
		}
		s.board[pos.Y][pos.X] = 'O'
		if looped, _ := simulate(&s); looped {
			result++
		}
	}
	fmt.Println("Result", result)
}

func simulate(s *state) (bool, int) {
	walks := make(map[vector.Vec2[int]][]rune)
	walks[s.guardPos] = append(walks[s.guardPos], s.board[s.guardPos.Y][s.guardPos.X])

	turned := false
	for {
		switch s.board[s.guardPos.Y][s.guardPos.X] {
		case '^':
			if s.guardPos.Y == 0 {
				return false, len(walks) + 1
			}
			next_pos := vector.Vec2[int]{X: s.guardPos.X, Y: s.guardPos.Y - 1}
			next := s.board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '^') {
				return true, len(walks) + 1
			}
			if slices.Contains([]rune{'#', 'O'}, next) {
				s.board[s.guardPos.Y][s.guardPos.X] = '>'
				turned = true
			} else {
				walks[s.guardPos] = append(walks[s.guardPos], '^')
				moveGuard(s, next_pos, '^', &turned)
			}
		case '<':
			if s.guardPos.X == 0 {
				return false, len(walks) + 1
			}
			next_pos := vector.Vec2[int]{X: s.guardPos.X - 1, Y: s.guardPos.Y}
			next := s.board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '<') {
				return true, len(walks) + 1
			}
			if slices.Contains([]rune{'#', 'O'}, next) {
				s.board[s.guardPos.Y][s.guardPos.X] = '^'
				turned = true
			} else {
				walks[s.guardPos] = append(walks[s.guardPos], '<')
				moveGuard(s, next_pos, '<', &turned)
			}
		case '>':
			if s.guardPos.X+1 == len(s.board[s.guardPos.Y]) {
				return false, len(walks) + 1
			}
			next_pos := vector.Vec2[int]{X: s.guardPos.X + 1, Y: s.guardPos.Y}
			next := s.board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], '>') {
				return true, len(walks) + 1
			}
			if slices.Contains([]rune{'#', 'O'}, next) {
				s.board[s.guardPos.Y][s.guardPos.X] = 'v'
				turned = true
			} else {
				walks[s.guardPos] = append(walks[s.guardPos], '>')
				moveGuard(s, next_pos, '>', &turned)
			}
		case 'v':
			if s.guardPos.Y+1 == len(s.board) {
				return false, len(walks) + 1
			}
			next_pos := vector.Vec2[int]{X: s.guardPos.X, Y: s.guardPos.Y + 1}
			next := s.board[next_pos.Y][next_pos.X]
			if slices.Contains(walks[next_pos], 'v') {
				return true, len(walks) + 1
			}
			if slices.Contains([]rune{'#', 'O'}, next) {
				s.board[s.guardPos.Y][s.guardPos.X] = '<'
				turned = true
			} else {
				walks[s.guardPos] = append(walks[s.guardPos], 'v')
				moveGuard(s, next_pos, 'v', &turned)
			}
		}
	}
}

func moveGuard(s *state, pos vector.Vec2[int], dir rune, turned *bool) {
	if *turned {
		s.board[s.guardPos.Y][s.guardPos.X] = '+'
		*turned = false
	} else if s.crossing {
		s.board[s.guardPos.Y][s.guardPos.X] = '+'
		s.crossing = false
	} else if slices.Contains([]rune{'<', '>'}, dir) {
		s.board[s.guardPos.Y][s.guardPos.X] = '-'
	} else {
		s.board[s.guardPos.Y][s.guardPos.X] = '|'
	}
	if slices.Contains([]rune{'|', '-'}, s.board[pos.Y][pos.X]) {
		s.crossing = true
	}
	s.board[pos.Y][pos.X] = dir
	s.guardPos = pos
}

func populateArray() (state, error) {
	board, err := input.FileToArray2D[rune]("2024/days/day6/day6_sample.txt")
	if err != nil {
		return state{}, err
	}

	var s state
	s.board = board

	for y, line := range s.board {
		for x, val := range line {
			if slices.Contains([]rune{'^', '<', '>', 'v'}, val) {
				s.guardPos.X = x
				s.guardPos.Y = y
			}
		}
	}
	return s, nil
}
