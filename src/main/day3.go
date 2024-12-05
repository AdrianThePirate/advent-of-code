package main

import (
	"fmt"
	"os"
	"regexp"
)

var mem string

func main() {
	data, err := os.ReadFile("tasks/day3_sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	mem = string(data)

	part1()
	part2()
}

func part1() {
	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	hits := reg.FindAllStringSubmatch(mem,-1)
	var result, num1, num2 int 

	for _, val := range hits {
		_, err := fmt.Sscanf(val[0], "mul(%d,%d)", &num1, &num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		} 

		result += num1 * num2
	}

	fmt.Println("Result:", result)
}

func part2() {
	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	hits := reg.FindAllStringSubmatch(mem,-1)
	var enabled = true
	var result, num1, num2 int

	for _, val := range hits {
		str := val[0]
		if str == "do()" { enabled = true; continue }
		if str == "don't()" { enabled = false; continue }
		if enabled {
			_, err := fmt.Sscanf(str, "mul(%d,%d)", &num1, &num2)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			} 

			result += num1 * num2
		}
	}
	fmt.Println("Result:", result)
}