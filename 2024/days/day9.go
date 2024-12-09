package main

import (
	"advent/adventutils"
	"fmt"
	"os"
)

func main() {
	disk, err := os.ReadFile("2024/tasks/day9_sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	exp := expand(disk)
	exp2 := make([]int, len(exp))
	copy(exp2, exp)
	format(&exp)
	fmt.Println(checksum(&exp))
	fileFormat(&exp2)
	fmt.Println(checksum(&exp2))
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

func fileFormat(disk *[]int) {
	p, si, ei := -1, -1, -1
	for i := len(*disk) - 1; i >= 0; i--{
		if p == -1 { p = (*disk)[i]; si = i }
		if i != si && p != (*disk)[i] { ei = i+1 }
		if ei != -1 {
			s := si - ei + 1
			ni := -1
			for j, n := range *disk {
				if n == -1 && ni == -1 { ni = j }
				if ni != -1 && j-ni == s { 
					for k := 0; k < s; k++{
						*disk = adventutils.MoveIndex(*disk, si-k, ni+k)
						*disk = adventutils.MoveIndex(*disk, ni+k+1, si-k)
					}
					ei, si, p = -1, i, (*disk)[i]
					break
				}
				if n != -1 { ni = -1 }
				if j >= ei { ei, si, p = -1, i, (*disk)[i]; break; }
			}
		}
	}
}

func checksum(disk *[]int) int{
	sum := 0
	for i, n := range *disk{
		if n == -1 { continue }
		sum += i * n
	}
	return sum
}