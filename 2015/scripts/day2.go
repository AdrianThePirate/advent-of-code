package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	list, err := getList("2015/tasks/day2.txt")
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
			if n < sm { sm = n }
		}
		lg := l
		for _, n := range []int{w, h} {
			if n > lg { lg = n }
		}

		wrap += t*2 + f*2 + s*2 + sm
		ribbon += a + l*2 + h*2 + w*2 - 2*lg
	}

	fmt.Printf("Sqr.ft. Wrap: %d\nFt. Ribbon: %d\n", wrap, ribbon)
}

func getList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		list = append(list, s.Text())
	}
	return list, nil
}