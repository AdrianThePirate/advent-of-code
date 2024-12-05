package main

import (
	"bufio"
	"fmt"
	"os"
)

var board [][]rune

func main() {
	populate_array()
	part1()
	part2()
}

func part1() {
	var count int
	for i := range board {
		for j := range board[i] {
			var words []string
			if board[i][j] == 'X' || board[i][j] == 'S' {
				if i+3 < len(board) { words = append(words, string(board[i][j]) + string(board[i+1][j]) + string(board[i+2][j]) + string(board[i+3][j])) }
				if j+3 < len(board[i]) { words = append(words, string(board[i][j]) + string(board[i][j+1]) + string(board[i][j+2]) + string(board[i][j+3])) }
				if i+3 < len(board) && j+3 < len(board[i]) { words = append(words, string(board[i][j]) + string(board[i+1][j+1]) + string(board[i+2][j+2]) + string(board[i+3][j+3])) }
				if i+3 < len(board) && j-3 >= 0 { words = append(words, string(board[i][j]) + string(board[i+1][j-1]) + string(board[i+2][j-2]) + string(board[i+3][j-3])) }
			}
			for _, word := range words {
				if word == "XMAS" || word == "SAMX" { count++ }
			}
		}
	}

	fmt.Println("Result:", count)
}

func part2() {
	var count int
	for i := range board {
		for j := range board[i] {
			var word1, word2 string
			if board[i][j] == 'A' {
				if 0 < i && i < len(board)-1 && 0 < j && j < len(board[i])-1 {
					word1 = string(board[i-1][j-1]) + string(board[i+1][j+1])
					word2 = string(board[i+1][j-1]) + string(board[i-1][j+1])
				}
				if (word1 == "SM" || word1 == "MS") && (word2 == "SM" || word2 == "MS") { count++ }
			}
		}
	}

	fmt.Println("Result:", count)
}

func populate_array() {
	file, err := os.Open("tasks/day4_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}
}