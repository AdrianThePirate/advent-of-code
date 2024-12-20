package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cali map[int][]int

func main() {
	cali = make(map[int][]int)
	populateMap()

	//part 1
	operations := []rune{'*', '+'}
	taskRun(operations)

	//part 2
	operations = []rune{'*', '+', '|'}
	taskRun(operations)
}

func taskRun(operations []rune) {
	var result int
	for key, vals := range cali{
		opts := oprOpts(operations, len(vals)-1)
		if solvable(key, opts, vals) { result += key }
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
		var sum = nums [0]
		for i, r := range line {
			switch r {
			case '*':
				sum *= nums[i+1]
			case '+':
				sum += nums[i+1]
			case '|':
				val, err :=  strconv.Atoi(strconv.Itoa(sum)+strconv.Itoa(nums[i+1]))
				if err != nil {
					fmt.Println(err)
				}
				sum = val
			}
		}
		if sum == goal { return true }
	}
	return false
}

func populateMap() {
	file, err := os.Open("2024/tasks/day7_sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner :=  bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":") 
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
}