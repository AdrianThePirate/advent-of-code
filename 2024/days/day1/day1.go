package day1

import (
	"fmt"
	"sort"

	"github.com/AdrianThePirate/advent-of-code/pkg/input"
	"github.com/AdrianThePirate/advent-of-code/pkg/math"
)

func Run() {
	left, right := populateList()
	part1(left, right)
	part2(left, right)
}

func part1(left, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	var result int
	for index, value := range left {
		result += math.Absolute(value - right[index])
	}

	fmt.Println("Result part 1:", result)
}

func part2(left, right []int) {
	right_map := repeatMap(right)

	var result int
	for _, value := range left {
		result += (value * right_map[value])
	}

	fmt.Println("Result part 2:", result)
}

func populateList() ([]int, []int) {
	lines, err := input.FileToLines("2024/days/day1/day1_sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}

	var left, right []int
	for _, line := range lines {
		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d   %d", &num1, &num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right
}

func repeatMap(n []int) map[int]int {
	mapped := map[int]int{}
	for _, value := range n {
		_, exists := mapped[value]
		if exists {
			mapped[value] += 1
		} else {
			mapped[value] = 1
		}
	}
	return mapped
}
