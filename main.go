package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/AdrianThePirate/advent-of-code/2015"
	_ "github.com/AdrianThePirate/advent-of-code/2024"
	"github.com/AdrianThePirate/advent-of-code/pkg/registry"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <year> <day>")
		os.Exit(1)
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid year:", os.Args[1])
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid day:", os.Args[2])
		os.Exit(1)
	}

	registry.Run(year, day)
}
