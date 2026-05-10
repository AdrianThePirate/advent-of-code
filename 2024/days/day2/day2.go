package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/input"
	"github.com/AdrianThePirate/advent-of-code/pkg/math"
)

func Run() {
	reports, err := input.FileToLines("2024/days/day2/day2_sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	part1(reports)
	part2(reports)
}

func part1(reports []string) {
	var safeCount int

	for _, value := range reports {
		parts := strings.Split(value, " ")

		var numbers []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}

		var ascending string
		prev := 0
		safe := true
		for _, num := range numbers {
			if prev == 0 {
				prev = num
				continue
			} else if ascending == "" {
				if prev > num {
					ascending = "false"
				} else {
					ascending = "true"
				}
			}

			if math.Absolute(prev-num) > 3 {
				safe = false
				break
			}
			if ascending == "true" && prev > num {
				safe = false
				break
			}
			if ascending == "false" && prev < num {
				safe = false
				break
			}
			if prev == num {
				safe = false
				break
			}
			prev = num
		}
		if safe {
			safeCount += 1
		}
	}

	println("Safe count:", safeCount)
}

func part2(reports []string) {
	var safe_count int

	for _, value := range reports {
		parts := strings.Split(value, " ")

		var numbers []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}

		var ascending string
		prev := 0
		err := 0
		for _, num := range numbers {
			if prev == 0 {
				prev = num
				continue
			} else if ascending == "" {
				if prev > num {
					ascending = "false"
				} else {
					ascending = "true"
				}
			}

			if math.Absolute(prev-num) > 3 {
				err += 1
			}
			if ascending == "true" && prev > num {
				err += 1
			}
			if ascending == "false" && prev < num {
				err += 1
			}
			if prev == num {
				err += 1
			}

			prev = num
		}
		if err < 2 {
			safe_count += 1
		}
	}

	println("Safe count:", safe_count)
}
