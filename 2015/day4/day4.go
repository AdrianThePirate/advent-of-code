package day4

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run(variant string) {
	key, err := input.FileToString("2015/day4/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(key)
}

func part1(key string) {
	iterate := 1
	found5, found6 := 0, 0
	for {
		test := key + strconv.Itoa(iterate)
		hash := md5.Sum([]byte(test))
		hex := fmt.Sprintf("%x", hash)
		if hex[:5] == "00000" && found5 == 0 {
			found5 = iterate
		}
		if hex[:6] == "000000" && found6 == 0{
			found6 = iterate
		}
		if found5 != 0 && found6 != 0 {
			break
		}
		iterate++
	}

	fmt.Printf("5 0's: %d\n6 0's: %d\n", found5, found6)
}