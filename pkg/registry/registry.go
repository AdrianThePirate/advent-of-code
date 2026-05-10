package registry

import "fmt"

var solvers = map[int]map[int]func(){}

func Register(year, day int, fn func()) {
	if solvers[year] == nil {
		solvers[year] = make(map[int]func())
	}
	solvers[year][day] = fn
}

func Run(year, day int) {
	if yearMap, ok := solvers[year]; ok {
		if fn, ok := yearMap[day]; ok {
			fn()
			return
		}
	}
	fmt.Printf("No solution found for year %d day %d\n", year, day)
}
