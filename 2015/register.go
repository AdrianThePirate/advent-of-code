package year2015

import (
	"github.com/AdrianThePirate/advent-of-code/pkg/registry"

	d1 "github.com/AdrianThePirate/advent-of-code/2015/days/day1"
	d2 "github.com/AdrianThePirate/advent-of-code/2015/days/day2"
	d3 "github.com/AdrianThePirate/advent-of-code/2015/days/day3"
)

func init() {
	registry.Register(2015, 1, d1.Run)
	registry.Register(2015, 2, d2.Run)
	registry.Register(2015, 3, d3.Run)
}
