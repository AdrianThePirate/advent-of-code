package main

import (
	"adventutils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reports []string

func main() {
	populate_list()
	part1()
	part2()
}

func part1() {
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
		safe := true
		for _, num := range numbers{
			if prev == 0 {
				prev = num
				continue
			} else if ascending == "" {
				if prev > num { ascending = "false" } else { ascending = "true" }
			}

			if adventutils.Absolute(prev - num) > 3 { safe = false; break }
			if ascending == "true" && prev > num { safe = false; break }
			if ascending == "false" && prev < num { safe = false; break }
			if prev == num { safe = false; break }
			prev = num
		}
		if safe { safe_count += 1 }
	}

	println("Safe count:", safe_count)
}

func part2() {
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
		for _, num := range numbers{
			if prev == 0 {
				prev = num
				continue
			} else if ascending == "" {
				if prev > num { ascending = "false" } else { ascending = "true" }
			}

			if adventutils.Absolute(prev - num) > 3 { err += 1 }
			if ascending == "true" && prev > num { err += 1 }
			if ascending == "false" && prev < num { err += 1 }
			if prev == num { err += 1 }

			prev = num
		}
		if err < 2 { safe_count += 1 }
	}

	println("Safe count:", safe_count)
}

func populate_list() {
	file, err := os.Open("tasks/day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reports = append(reports, scanner.Text())
	}
}

