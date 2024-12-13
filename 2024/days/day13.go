package main

import (
	"advent/pkg/array"
	"advent/pkg/matrix"
	"advent/pkg/vector"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("2024/tasks/day13.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr array.Array[struct{vec1 vector.Vec2[int64]; vec2 vector.Vec2[int64]; pos vector.Vec2[int64]}]
	var arr2 array.Array[struct{vec1 vector.Vec2[int64]; vec2 vector.Vec2[int64]; pos vector.Vec2[int64]}]
	for scanner.Scan() {
		if scanner.Text() == "" { continue }
		var x, y int64
		var vec1, vec2, pos, pos2 vector.Vec2[int64]
		fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d", &x, &y)
		vec1 = vector.Vec2[int64]{X: x, Y: y}
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &x, &y)
		vec2 = vector.Vec2[int64]{X: x, Y: y}
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &x, &y)
		pos = vector.Vec2[int64]{X: x, Y: y}
		pos2 = vector.Vec2[int64]{X: (x+10000000000000), Y: (y+10000000000000)}

		arr = append(arr, struct{vec1 vector.Vec2[int64]; vec2 vector.Vec2[int64]; pos vector.Vec2[int64]}{vec1: vec1, vec2:  vec2, pos: pos})
		arr2 = append(arr2, struct{vec1 vector.Vec2[int64]; vec2 vector.Vec2[int64]; pos vector.Vec2[int64]}{vec1: vec1, vec2:  vec2, pos: pos2})
	}

	tokens, err := findTokenPrice(arr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokens)

	tokens2, err := findTokenPrice(arr2)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(tokens2)
}

func findTokenPrice(arr array.Array[struct{vec1 vector.Vec2[int64]; vec2 vector.Vec2[int64]; pos vector.Vec2[int64]}]) (int64, error) {
	var price int64 = 0
	for _, val := range arr {
		a, err := matrix.Det(array.Array2D[int64]{{val.vec1.X, val.vec1.Y}, {val.vec2.X, val.vec2.Y}})
		if err != nil {
			return 0, err
		}
		if a == 0 { continue }
		x, err := matrix.Det(array.Array2D[int64]{{val.pos.X, val.vec2.X}, {val.pos.Y, val.vec2.Y}})
		if err != nil {
			return 0, err
		}
		y, err := matrix.Det(array.Array2D[int64]{{val.vec1.X, val.pos.X}, {val.vec1.Y, val.pos.Y}})
		if err != nil {
			return 0, err
		}
		if x % a != 0 { continue }
		if y % a != 0 { continue }

		price += x/a * 3
		price += y/a
	}

	return price, nil
}