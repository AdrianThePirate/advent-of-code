package input

import (
	"bufio"
	"fmt"
	"os"
)

func Array2D(path string, targetType interface{}) (interface{}, error) {
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