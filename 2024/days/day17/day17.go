package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
	"github.com/AdrianThePirate/advent-of-code/pkg/cmd"
)

type vm struct {
	a, b, c, ptr int
}

func Run(variant string) {
	path := cmd.InputPath("2024/days/day17/day17", variant)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var la, lb, lc int
	var str string
	var prg array.Array[int]
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register A: %d", &la)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register B: %d", &lb)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Register C: %d", &lc)
	scanner.Scan()
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Program: %s", &str)

	parts := strings.Split(str, ",")
	for _, r := range parts {
		i, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		prg = append(prg, i)
	}

	result, err := exec(&vm{a: la, b: lb, c: lc}, prg)
	if err != nil {
		fmt.Println(err)
	}

	var resstr string
	for _, i := range result {
		resstr += fmt.Sprintf("%d,", i)
	}
	if len(resstr) > 0 {
		resstr = resstr[:len(resstr)-1]
	}
	fmt.Println(resstr)

	for i := 1; ; i++ {
		result, err := exec(&vm{a: i, b: lb, c: lc}, prg)
		if err != nil {
			fmt.Println(err)
		}
		if i%100000 == 0 {
			fmt.Println(i)
		}

		var test string
		for _, v := range result {
			test += fmt.Sprintf("%d,", v)
		}
		if len(test) > 0 {
			test = test[:len(test)-1]
		}
		if test == str {
			fmt.Println(i)
			break
		}
	}
}

func exec(m *vm, prg array.Array[int]) (array.Array[int], error) {
	length := len(prg)
	var output array.Array[int]
	for {
		if m.ptr == length {
			break
		}
		opr := prg[m.ptr+1]
		switch prg[m.ptr] {
		case 0:
			if err := adv(m, opr); err != nil {
				return nil, err
			}
			m.ptr += 2
		case 1:
			bxl(m, opr)
			m.ptr += 2
		case 2:
			if err := bst(m, opr); err != nil {
				return nil, err
			}
			m.ptr += 2
		case 3:
			if !jnz(m, opr) {
				m.ptr += 2
			}
		case 4:
			bxc(m)
			m.ptr += 2
		case 5:
			v, err := out(m, opr)
			if err != nil {
				return nil, err
			}
			output = append(output, v)
			m.ptr += 2
		case 6:
			if err := bdv(m, opr); err != nil {
				return nil, err
			}
			m.ptr += 2
		case 7:
			if err := cdv(m, opr); err != nil {
				return nil, err
			}
			m.ptr += 2
		}
	}
	return output, nil
}

func combo(m *vm, v int) (int, error) {
	switch {
	case v < 4:
		return v, nil
	case v == 4:
		return m.a, nil
	case v == 5:
		return m.b, nil
	case v == 6:
		return m.c, nil
	default:
		return 0, fmt.Errorf("non valid combo operand")
	}
}

func adv(m *vm, opr int) error {
	v, err := combo(m, opr)
	if err != nil {
		return err
	}
	m.a = m.a / int(math.Pow(2, float64(v)))
	return nil
}

func bxl(m *vm, opr int) {
	m.b = m.b ^ opr
}

func bst(m *vm, opr int) error {
	v, err := combo(m, opr)
	if err != nil {
		return err
	}
	m.b = v % 8
	return nil
}

func jnz(m *vm, opr int) bool {
	if m.a == 0 {
		return false
	}
	m.ptr = opr
	return true
}

func bxc(m *vm) {
	m.b = m.b ^ m.c
}

func out(m *vm, opr int) (int, error) {
	v, err := combo(m, opr)
	if err != nil {
		return 0, err
	}
	return v % 8, nil
}

func bdv(m *vm, opr int) error {
	v, err := combo(m, opr)
	if err != nil {
		return err
	}
	m.b = m.a / int(math.Pow(2, float64(v)))
	return nil
}

func cdv(m *vm, opr int) error {
	v, err := combo(m, opr)
	if err != nil {
		return err
	}
	m.c = m.a / int(math.Pow(2, float64(v)))
	return nil
}
