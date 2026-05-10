package day2

import (
	"fmt"

	"github.com/AdrianThePirate/advent-of-code/pkg/input"
)

func Run(variant string) {
	list, err := input.FileToLines("2015/day2/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	wrap, ribbon := 0, 0
	for _, line := range list {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		t, f, s, a := l*w, l*h, w*h, l*w*h
		sm := t
		for _, n := range []int{f, s} {
			if n < sm {
				sm = n
			}
		}
		lg := l
		for _, n := range []int{w, h} {
			if n > lg {
				lg = n
			}
		}

		wrap += t*2 + f*2 + s*2 + sm
		ribbon += a + l*2 + h*2 + w*2 - 2*lg
	}

	fmt.Printf("Sqr.ft. Wrap: %d\nFt. Ribbon: %d\n", wrap, ribbon)
}
