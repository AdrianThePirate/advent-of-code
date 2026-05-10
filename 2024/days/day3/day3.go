package day3

import (
	"fmt"
	"regexp"

	"github.com/AdrianThePirate/advent-of-code/pkg/cmd"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run(variant string) {
	path := cmd.InputPath("2024/days/day3/day3", variant)
	mem, err := input.FileToString(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	part1(mem)
	part2(mem)
}

func part1(mem string) {
	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	hits := reg.FindAllStringSubmatch(mem, -1)
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

func part2(mem string) {
	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	hits := reg.FindAllStringSubmatch(mem, -1)
	var enabled = true
	var result, num1, num2 int

	for _, val := range hits {
		str := val[0]
		if str == "do()" {
			enabled = true
			continue
		}
		if str == "don't()" {
			enabled = false
			continue
		}
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
