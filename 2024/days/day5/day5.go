package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run() {
	rules, rulesReversed, pages := importData()
	incorrect := part1(rules, pages)
	part2(rules, rulesReversed, incorrect)
}

func part1(rules map[int][]int, pages [][]int) [][]int {
	var result int
	var incorrect [][]int

	for i := range pages {
		var valid = true
		var forbidden []int
		for _, val := range pages[i] {
			if slices.Contains(forbidden, val) {
				valid = false
				incorrect = append(incorrect, pages[i])
				break
			}
			if rules[val] != nil {
				forbidden = append(forbidden, rules[val]...)
			}
		}
		if valid {
			result += pages[i][len(pages[i])/2]
		}
	}

	fmt.Println("Result:", result)
	return incorrect
}

func part2(rules map[int][]int, rulesReversed map[int][]int, incorrect [][]int) {
	var result int

	for i := range incorrect {
		var forbidden []int
		for j, val := range incorrect[i] {
			if slices.Contains(forbidden, val) {
				for z, lav := range incorrect[i] {
					if slices.Contains(rulesReversed[val], lav) {
						row := array.Array[int](incorrect[i])
						row.MoveIndex(j, z)
						incorrect[i] = row
						break
					}
				}
			}
			if rules[val] != nil {
				forbidden = append(forbidden, rules[val]...)
			}
		}
		result += incorrect[i][len(incorrect[i])/2]
	}

	fmt.Println("Result:", result)
}

func importData() (map[int][]int, map[int][]int, [][]int) {
	lines, err := input.FileToLines("2024/days/day5/day5_sample.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}

	rules := make(map[int][]int)
	rulesReversed := make(map[int][]int)
	var pages [][]int

	rulemode := true
	for _, line := range lines {
		if line == "" {
			rulemode = false
			continue
		}
		if rulemode {
			var num1, num2 int
			fmt.Sscanf(line, "%d|%d", &num1, &num2)
			rules[num2] = append(rules[num2], num1)
			rulesReversed[num1] = append(rulesReversed[num1], num2)
		} else {
			parts := strings.Split(line, ",")
			var section []int
			for _, val := range parts {
				page, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				section = append(section, page)
			}
			pages = append(pages, section)
		}
	}

	return rules, rulesReversed, pages
}
