package day1

import (
	"fmt"
	"os"
)

func Run(variant string) {
	data, err := os.ReadFile("2015/day1/day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	floor, basement := getFloor(data)
	fmt.Printf("Floor: %d\nBasement Entered: %d\n", floor, basement)
}

func getFloor(data []byte) (int, int) {
	floor, base := 0, 0
	for i, r := range data {
		if r == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 && base == 0 {
			base = i + 1
		}
	}
	return floor, base
}
