package main

import (
	"advent/adventutils"
	"bufio"
	"fmt"
	"os"
	"sort"
)

var left, right []int

func main() {
	populateList()
	part1()
	part2()
}

func part1() {
	sort.Ints(left)
	sort.Ints(right)

	var result int
	for index, value := range left {
		result += adventutils.Absolute(value - right[index])
	}

	fmt.Println("Result part 1:", result)
}

func part2() {
	right_map := repeatMap(right)

	var result int
	for _, value := range left {
		result += (value * right_map[value])
	}

	fmt.Println("Result part 2:", result)
}

func populateList() {
	file, err := os.Open("2024/tasks/day1_sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num1, num2 int

		_, err := fmt.Sscanf(scanner.Text(), "%d   %d", &num1, &num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		left = append(left, num1)
		right = append(right, num2)
	}
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