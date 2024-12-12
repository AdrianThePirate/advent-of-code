package input

import (
	"advent/pkg/array"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileToArray2D[T any](path string) (array.Array2D[T], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result array.Array2D[T]

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row[]T
		for _, r := range line {
			var value T
			switch any(new(T)).(type) {
			case *rune:
				value = any(r).(T)
			case *int:
				if r < '0' || r > '9' {
					return nil, fmt.Errorf("invalid character '%c' for int conversion", r)
				}
				value = any(r - '0').(T)
			default:
				return nil, fmt.Errorf("unsupported type")
			}
			row = append(row, value)
		}
		result = append(result, row)
	}

	return result, nil
}

func FileToArray[T any](path string) (array.Array[T], error) {
	var arr array.Array[T]
	file, err := os.Open(path)
	if err != nil {
		return array.Array[T]{}, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	switch any(new(T)).(type) {
	case int:
		for scanner.Scan() {
			line := scanner.Text()
			var parts []string
			if strings.Contains(line, " ") {
				parts = strings.Fields(line)
			} else {
				parts = strSplitStrChar(line)
			}
			for _, r := range parts {
				num, err := strconv.Atoi(r)
				if err != nil {
					return array.Array[T]{}, err
				}
				arr = append(arr, any(num).(T))
			}
		}
	default:
		return array.Array[T]{}, fmt.Errorf("unsupported type")
	}

	return arr, nil
}

func strSplitStrChar(str string) []string {
	var result []string
	for _, c := range str {
		result = append(result, string(c))
	}
	return result
}