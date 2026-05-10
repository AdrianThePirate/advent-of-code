package matrix

import (
	"testing"

	"github.com/AdrianThePirate/advent-of-code/pkg/array"
)

func TestDet1x1(t *testing.T) {
	m := array.Array2D[int]{
		array.Array[int]{7},
	}
	got, err := Det(m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 7 {
		t.Errorf("Det(1x1) = %d, want 7", got)
	}
}

func TestDet2x2(t *testing.T) {
	m := array.Array2D[int]{
		array.Array[int]{3, 8},
		array.Array[int]{4, 6},
	}
	got, err := Det(m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != -14 {
		t.Errorf("Det(2x2) = %d, want -14", got)
	}
}

func TestDet3x3(t *testing.T) {
	m := array.Array2D[int]{
		array.Array[int]{6, 1, 1},
		array.Array[int]{4, -2, 5},
		array.Array[int]{2, 8, 7},
	}
	got, err := Det(m)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != -306 {
		t.Errorf("Det(3x3) = %d, want -306", got)
	}
}

func TestDetEmpty(t *testing.T) {
	m := array.Array2D[int]{}
	_, err := Det(m)
	if err == nil {
		t.Error("Det(empty) should return an error")
	}
}

func TestDetNonSquare(t *testing.T) {
	m := array.Array2D[int]{
		array.Array[int]{1, 2, 3},
		array.Array[int]{4, 5, 6},
	}
	_, err := Det(m)
	if err == nil {
		t.Error("Det(non-square) should return an error")
	}
}
