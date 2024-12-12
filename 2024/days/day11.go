package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/tasks/day11.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	
	stones := make(map[string]int)
	for _, val := range parts {
		stones[val] += 1
	}
	
	for i := 0; i < 50; i++{
		fmt.Println((i+1)*5, blinks(&stones, 5), len(stones))
	}
	fmt.Println(blinks(&stones, 25))
	fmt.Println(blinks(&stones, 50))
}

func blinks(stones *map[string]int, times int) uint64 {
	for i := 0; i < times; i++ {
		temp := make(map[string]int)
		for key, val := range *stones {
			if key == "0" { 
				temp["1"] += val
			} else if len(key) % 2 == 0 {
				n1, n2 := key[:len(key)/2], key[len(key)/2:]
				temp[n1] += val
				n2 = strings.TrimLeft(n2, "0")
				if n2 != "" { 
					temp[n2] += val 
				} else {
					temp["0"] += val
				}
			} else {
				n, err := strconv.Atoi(key)
				if err != nil {
					fmt.Println(err)
				}
				temp[strconv.Itoa(n*2024)] += val 
			}
		}
		*stones = temp
	}

	var result uint64
	for _, val := range *stones{
		result += uint64(val)
	}

	return result
}
