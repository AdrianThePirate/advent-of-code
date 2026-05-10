package cmd

import "testing"

func TestInputPath(t *testing.T) {
	cases := []struct {
		base, variant, expected string
	}{
		{"day01/input", "", "day01/input.txt"},
		{"day01/input", "sample", "day01/input_sample.txt"},
		{"day01/input", "small", "day01/input_small.txt"},
	}

	for _, c := range cases {
		got := InputPath(c.base, c.variant)
		if got != c.expected {
			t.Errorf("InputPath(%q, %q) = %q, want %q", c.base, c.variant, got, c.expected)
		}
	}
}
