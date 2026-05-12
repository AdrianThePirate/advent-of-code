package year2015

import (
	"github.com/AdrianThePirate/advent-of-code/pkg/registry"

	d1 "github.com/AdrianThePirate/advent-of-code/2015/day1"
	d2 "github.com/AdrianThePirate/advent-of-code/2015/day2"
	d3 "github.com/AdrianThePirate/advent-of-code/2015/day3"
	d4 "github.com/AdrianThePirate/advent-of-code/2015/day4"
	d5 "github.com/AdrianThePirate/advent-of-code/2015/day5"
	d6 "github.com/AdrianThePirate/advent-of-code/2015/day6"
)

func init() {
	registry.Register(2015, 1, d1.Run)
	registry.Register(2015, 2, d2.Run)
	registry.Register(2015, 3, d3.Run)
	registry.Register(2015, 4, d4.Run)
	registry.Register(2015, 5, d5.Run)
	registry.Register(2015, 6, d6.Run)
}
