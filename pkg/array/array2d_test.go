package array

import (
	"testing"

	"github.com/AdrianThePirate/advent-of-code/pkg/vector"
)

func TestGetPos(t *testing.T) {
	grid := Array2D[rune]{
		Array[rune]{'a', 'b', 'c'},
		Array[rune]{'d', 'e', 'f'},
	}

	got, err := grid.GetPos(vector.Vec2[int]{X: 1, Y: 0})
	if err != nil {
		t.Fatalf("GetPos returned unexpected error: %v", err)
	}
	if got != 'b' {
		t.Errorf("GetPos(1,0) = %c, want 'b'", got)
	}
}

func TestGetPosOutOfBounds(t *testing.T) {
	grid := Array2D[rune]{
		Array[rune]{'a', 'b'},
	}
	cases := []vector.Vec2[int]{
		{X: -1, Y: 0},
		{X: 0, Y: -1},
		{X: 2, Y: 0},
		{X: 0, Y: 1},
	}
	for _, pos := range cases {
		_, err := grid.GetPos(pos)
		if err == nil {
			t.Errorf("GetPos(%v) should return error for out-of-bounds", pos)
		}
	}
}

func TestSetPos(t *testing.T) {
	grid := Array2D[rune]{
		Array[rune]{'a', 'b', 'c'},
	}
	err := grid.SetPos(vector.Vec2[int]{X: 1, Y: 0}, 'z')
	if err != nil {
		t.Fatalf("SetPos returned unexpected error: %v", err)
	}
	if grid[0][1] != 'z' {
		t.Errorf("SetPos: grid[0][1] = %c, want 'z'", grid[0][1])
	}
}

func TestSetPosOutOfBounds(t *testing.T) {
	grid := Array2D[rune]{
		Array[rune]{'a', 'b'},
	}
	err := grid.SetPos(vector.Vec2[int]{X: 5, Y: 0}, 'z')
	if err == nil {
		t.Error("SetPos out-of-bounds should return error")
	}
}
