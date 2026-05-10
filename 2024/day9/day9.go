package day9

import (
	"fmt"
	"os"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/cmd"
)

func Run(variant string) {
	path := cmd.InputPath("2024/day9/day9", variant)
	disk, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	exp := expand(disk)
	exp2 := make(array.Array[int], len(exp))
	copy(exp2, exp)
	format(&exp)
	fmt.Println(checksum(&exp))
	fileFormat(&exp2)
	fmt.Println(checksum(&exp2))
}

func expand(disk []byte) array.Array[int] {
	var p int
	var exp array.Array[int]
	for i, r := range disk {
		n := int(r - '0')
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				exp = append(exp, p)
			}
			p++
		} else {
			for j := 0; j < n; j++ {
				exp = append(exp, -1)
			}
		}
	}
	return exp
}

func format(disk *array.Array[int]) {
	formatted := false
	for i := len(*disk) -1; i >= 0; i-- {
		if (*disk)[i] == -1 {
			continue
		}
		for j, n := range *disk {
			if j >= i {
				formatted = true
				break
			}
			if n == -1 {
				disk.MoveIndex(i, j)
				disk.MoveIndex(j+1, i)
				break
			}
		}
		if formatted {
			break
		}
	}
}

func fileFormat(disk *array.Array[int]) {
	p, si, ei := -1, -1, -1
	for i := len(*disk) - 1; i >= 0; i-- {
		if p == -1 {
			p = (*disk)[i]
			si = i
		}
		if i != si && p != (*disk)[i] {
			ei = i + 1
		}
		if ei != -1 {
			s := si - ei + 1
			ni := -1
			for j, n := range *disk {
				if n == -1 && ni == -1 {
					ni = j
				}
				if ni != -1 && j-ni == s {
					for k := 0; k < s; k++ {
						disk.MoveIndex(si-k, ni+k)
						disk.MoveIndex(ni+k+1, si-k)
					}
					ei, si, p = -1, i, (*disk)[i]
					break
				}
				if n != -1 {
					ni = -1
				}
				if j >= ei {
					ei, si, p = -1, i, (*disk)[i]
					break
				}
			}
		}
	}
}

func checksum(disk *array.Array[int]) int {
	sum := 0
	for i, n := range *disk {
		if n == -1 {
			continue
		}
		sum += i * n
	}
	return sum
}
