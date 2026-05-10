package registry

import "fmt"

var solvers = map[int]map[int]func(string){}

func Register(year, day int, fn func(string)) {
	if solvers[year] == nil {
		solvers[year] = make(map[int]func(string))
	}
	solvers[year][day] = fn
}

func Run(year, day int, variant string) {
	if yearMap, ok := solvers[year]; ok {
		if fn, ok := yearMap[day]; ok {
			fn(variant)
			return
		}
	}
	fmt.Printf("No solution found for year %d day %d\n", year, day)
}
