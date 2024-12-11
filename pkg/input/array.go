package input

import (
	"advent/pkg/array"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileToArray2D(path string, targetType interface{}) (interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var intResult [][]int
	var runeResult [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch targetType.(type) {
		case rune:
			runeResult = append(runeResult, []rune(scanner.Text()))
			break
		case int:
			line := scanner.Text()
			var row []int
			for _, r := range line {
				row = append(row, int(r - '0'))
			}
			intResult = append(intResult, row)
		}
	}

	switch targetType.(type) {
	case rune:
		return runeResult, nil
	case int:
		return intResult, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func FileToArray[T any](path string) (array.Array[T], error) {
	var arr array.Array[T]
	file, err := os.Open(path)
	if err != nil {
		return array.Array[T]{}, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	var test T
	switch any(test).(type) {
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
				arr.Append(any(num).(T))
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