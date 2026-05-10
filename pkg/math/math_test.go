package math

import (
	"testing"
)

func TestAbsolute(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-1, 1},
	}

	for _, c := range cases {
		got := Absolute(c.input)
		if got != c.expected {
			t.Errorf("Absolute(%d) = %d, want %d", c.input, got, c.expected)
		}
	}
}

func TestGetDigitCount(t *testing.T) {
	cases := []struct {
		input	 int
		expected int
	}{
		{0, 1},
		{432, 3},
		{3, 1},
		{-2, 1},
		{-3285, 4},
	}

	for _, c := range cases {
		got := GetDigitCount(c.input)
		if got != c.expected {
			t.Errorf("GetDigitCount(%d) = %d, want %d", c.input, got, c.expected)
		}
	}
}