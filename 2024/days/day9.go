package main

import (
	"advent/adventutils"
	"fmt"
	"os"
)

func main() {
	disk, err := os.ReadFile("2024/tasks/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	exp := expand(disk)
	exp2 := exp
	format(&exp)
	fmt.Println(checksum(&exp))
	fileformat(&exp2)
	fmt.PrintLn(checksum(&exp2))
}

func expand(disk []byte) []int {
	var p int
	var exp []int
	for i, r := range disk{
		n := int(r - '0')
		if i % 2 == 0 {
			for j := 0; j < n; j++  {
				exp = append(exp, p)
			}
			p++
		} else {
			for j := 0; j < n; j++  {
				exp = append(exp, -1)
			}
		}
	}

	return exp
}

func format(disk *[]int) {
	formatted := false
	for i := len(*disk) - 1; i >= 0; i--{
		if (*disk)[i] == -1 { continue }
		for j, n := range *disk {
			if j >= i { formatted = true; break}
			if n == -1 { *disk = adventutils.MoveIndex(*disk, i, j); *disk = adventutils.MoveIndex(*disk, j+1, i) ;break }
		}
		if formatted { break }
	}
}

func fileformat(disk *[]int) {

}

func checksum(disk *[]int) int{
	sum := 0
	for i, n := range *disk{
		if n == -1 { break }
		sum += i * n
	}
	return sum
}