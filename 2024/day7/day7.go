package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/cmd"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run(variant string) {
	path := cmd.InputPath("2024/day7/day7", variant)
	cali := populateMap(path)

	taskRun(cali, []rune{'*', '+'})
	taskRun(cali, []rune{'*', '+', '|'})
}

func taskRun(cali map[int][]int, operations []rune) {
	var result int
	for key, vals := range cali {
		opts := oprOpts(operations, len(vals)-1)
		if solvable(key, opts, vals) {
			result += key
		}
	}
	fmt.Println("Result:", result)
}

func oprOpts(oprs []rune, x int) [][]rune {
	if x == 0 {
		return [][]rune{{}}
	}

	subset := oprOpts(oprs, x-1)

	var opts [][]rune
	for _, line := range subset {
		for _, opr := range oprs {
			newLine := append([]rune{}, line...)
			newLine = append(newLine, opr)
			opts = append(opts, newLine)
		}
	}

	return opts
}

func solvable(goal int, opts [][]rune, nums []int) bool {
	for _, line := range opts {
		var sum = nums[0]
		for i, r := range line {
			switch r {
			case '*':
				sum *= nums[i+1]
			case '+':
				sum += nums[i+1]
			case '|':
				val, err := strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(nums[i+1]))
				if err != nil {
					fmt.Println(err)
				}
				sum = val
			}
		}
		if sum == goal {
			return true
		}
	}
	return false
}

func populateMap(path string) map[int][]int {
	lines, err := input.FileToLines(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	cali := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		key, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}

		parts = strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, num := range parts {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}
			cali[key] = append(cali[key], val)
		}
	}
	return cali
}
