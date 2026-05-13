package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

var wireValue map[string]uint16
var wireInstruction map[string][]string


func Run(variant string) {
	lines, err := input.FileToLines("2015/day7/day7.txt")
    if err != nil{
		fmt.Println(err)
		return
	}
	
    part1(lines)
	part2()
}

func part1(lines []string) {
	wireValue = map[string]uint16{}
	wireInstruction = map[string][]string{}

	for _, line := range lines{
		parts := strings.Fields(line)
		switch len(parts) {
		case 3:
			wireInstruction[parts[2]] = parts[:1]
		case 4:
			wireInstruction[parts[3]] = parts[:2]
		case 5:
			wireInstruction[parts[4]] = parts[:3]
		}
	}

	fmt.Printf("Part 1: Wire a: %v\n", solveWire("a"))
}

func part2() {
	overideVal := wireValue["a"]
	wireValue = map[string]uint16{"b": overideVal}

	fmt.Printf("Part 2: Wire a: %v\n", solveWire("a"))
}

func solveWire(wire string) uint16 {
	if val, err := strconv.ParseUint(wire, 10, 16); err == nil {
		return uint16(val)
	}
	if val, ok := wireValue[wire]; ok {
		return val
	}
	parts := wireInstruction[wire]
	switch len(parts) {
	case 1:
		val := solveWire(parts[0])
		wireValue[wire] = val
		return val
	case 2:
		val := solveWire(parts[1])
		val = ^val
		wireValue[wire] = val
		return val
	case 3:
		val1, val2 := solveWire(parts[0]), solveWire(parts[2])
		val := uint16(0)
		switch parts[1] {
		case "AND":
			val = val1 & val2
		case "OR":
			val = val1 | val2
		case "LSHIFT":
			val = val1 << val2
		case "RSHIFT":
			val = val1 >> val2
		}
		wireValue[wire] = val
		return val
	}
	panic(fmt.Sprintf("Error: wire '%q' failed to be solved\nFollowing instructions avilable: %v\nall instructions %v", wire, parts, wireInstruction))
}