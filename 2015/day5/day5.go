package day5

import (
	"fmt"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run(variant string) {
	lines, err := input.FileToArray2D[rune]("2015/day5/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

    part1(lines)
	part2(lines)
}

func part1(lines array.Array2D[rune]) {
	niceLines := 0
	bandWords := map[string]struct{}{"ab": {}, "cd": {}, "pq": {}, "xy": {}}
	volwels := map[rune]struct{}{'a': {}, 'i': {}, 'u': {}, 'e': {}, 'o': {}}
	for line := range lines {
		volwelCount, is2InRow, isIlligeal := 0, false, false
		for index, letter := range lines[line]{
			if _, ok := volwels[letter]; ok {volwelCount++}
			if index == len(lines[line])-1 { break }
			if letter == lines[line][index+1] {is2InRow = true}
			if _, ok := bandWords[fmt.Sprintf("%s%s", string(letter), string(lines[line][index+1]))]; ok {isIlligeal = true; break}
		}
		if isIlligeal || volwelCount < 3 || !is2InRow { continue } else { niceLines++ }
	}
	fmt.Printf("Nice lines part 1: %d\n", niceLines)
}

func part2(lines array.Array2D[rune]) {
	niceLines := 0
	for line := range lines {
		niceRepeat, validPair := false, false
		pairs := map[string]int{}

		for index, letter := range lines[line]{
			if index == len(lines[line])-1 { break }
			pair := fmt.Sprintf("%s%s", string(letter), string(lines[line][index+1]));
			if loc, ok := pairs[pair]; ok && loc != index-1 {
				validPair = true
			}else if !ok{
				pairs[pair] = index
			}
			if index == len(lines[line])-2 { continue }
			if letter == lines[line][index+2] { niceRepeat = true }
		}
		if niceRepeat && validPair { niceLines++ }
	}

	fmt.Printf("Nice lines part 2: %d\n", niceLines)
}