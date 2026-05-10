package vector

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, expected Vec2[int]
	}{
		{Vec2[int]{1, 2}, Vec2[int]{3, 4}, Vec2[int]{4, 6}},
		{Vec2[int]{0, 0}, Vec2[int]{0, 0}, Vec2[int]{0, 0}},
		{Vec2[int]{-1, -2}, Vec2[int]{1, 2}, Vec2[int]{0, 0}},
	}
	for _, c := range cases {
		got := c.a.Add(c.b)
		if got != c.expected {
			t.Errorf("Add(%v, %v) = %v, want %v", c.a, c.b, got, c.expected)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		a, b, expected Vec2[int]
	}{
		{Vec2[int]{5, 7}, Vec2[int]{2, 3}, Vec2[int]{3, 4}},
		{Vec2[int]{0, 0}, Vec2[int]{1, 1}, Vec2[int]{-1, -1}},
	}
	for _, c := range cases {
		got := c.a.Sub(c.b)
		if got != c.expected {
			t.Errorf("Sub(%v, %v) = %v, want %v", c.a, c.b, got, c.expected)
		}
	}
}

func TestMulVec(t *testing.T) {
	a := Vec2[int]{2, 3}
	b := Vec2[int]{4, 5}
	got := a.Mul(b)
	expected := Vec2[int]{8, 15}
	if got != expected {
		t.Errorf("Mul(vec) = %v, want %v", got, expected)
	}
}

func TestMulScalar(t *testing.T) {
	a := Vec2[int]{2, 3}
	got := a.Mul(int(4))
	expected := Vec2[int]{8, 12}
	if got != expected {
		t.Errorf("Mul(scalar) = %v, want %v", got, expected)
	}
}

func TestDiv(t *testing.T) {
	a := Vec2[int]{8, 6}
	b := Vec2[int]{2, 3}
	got := a.Div(b)
	expected := Vec2[int]{4, 2}
	if got != expected {
		t.Errorf("Div(%v, %v) = %v, want %v", a, b, got, expected)
	}
}

func TestModulo(t *testing.T) {
	a := Vec2[int]{7, 9}
	b := Vec2[int]{3, 4}
	got, err := a.Modulo(b)
	expected := Vec2[int]{1, 1}
	if err != nil {
		t.Fatalf("Modulo returned unexpected error: %v", err)
	}
	if got != expected {
		t.Errorf("Modulo(%v, %v) = %v, want %v", a, b, got, expected)
	}
}

func TestModuloFloatError(t *testing.T) {
	a := Vec2[float64]{7.0, 9.0}
	_, err := a.Modulo(Vec2[float64]{3.0, 4.0})
	if err == nil {
		t.Error("Modulo on float64 should return an error")
	}
}

func TestDot(t *testing.T) {
	cases := []struct {
		a, b	Vec2[float64]
		expected	float64
	}{
		{Vec2[float64]{2,3},Vec2[float64]{4,5},23},
		{Vec2[float64]{4,5},Vec2[float64]{2,3},23},
		{Vec2[float64]{4,5},Vec2[float64]{-2,-3},-23},
		{Vec2[float64]{-4,-5},Vec2[float64]{2,3},-23},
		{Vec2[float64]{-4,-5},Vec2[float64]{-2,-3},23},
		{Vec2[float64]{-4,5},Vec2[float64]{2,3},7},
	}

	for _, c := range cases {
		got := c.a.Dot(c.b)
		if got != c.expected {
			t.Errorf("Dot(%v, %v) = %f, want %f", c.a, c.b, got, c.expected)
		}
	}
}

func TestMagn(t *testing.T) {
	cases := []struct {
		v        Vec2[float64]
		expected float64
	}{
		{Vec2[float64]{3, 4}, 5.0},
		{Vec2[float64]{-3, 4}, 5.0},
		{Vec2[float64]{3, -4}, 5.0},
		{Vec2[float64]{-3, -4}, 5.0},
		{Vec2[float64]{0, 0}, 0.0},
		{Vec2[float64]{1, 0}, 1.0},
	}
	for _, c := range cases {
		got := c.v.Magn()
		if got != c.expected {
			t.Errorf("Magn(%v) = %f, want %f", c.v, got, c.expected)
		}
	}
}

func TestDistanceTo(t *testing.T) {
	cases := []struct {
		a, b     Vec2[float64]
		expected float64
	}{
		{Vec2[float64]{0, 0}, Vec2[float64]{3, 4}, 5.0},
		{Vec2[float64]{3, 4}, Vec2[float64]{0, 0}, 5.0},   // reversed
		{Vec2[float64]{-3, 0}, Vec2[float64]{1, 0}, 4.0},  // along X, negative start
		{Vec2[float64]{0, -4}, Vec2[float64]{0, 1}, 5.0},  // along Y, negative start
		{Vec2[float64]{-3, -4}, Vec2[float64]{0, 0}, 5.0}, // from negative quadrant
		{Vec2[float64]{1, 1}, Vec2[float64]{1, 1}, 0.0},   // same point
	}
	for _, c := range cases {
		got := c.a.DistanceTo(c.b)
		if got != c.expected {
			t.Errorf("DistanceTo(%v, %v) = %f, want %f", c.a, c.b, got, c.expected)
		}
	}
}

func TestDirections(t *testing.T) {
	v := Vec2[int]{5, 5}
	cases := []struct {
		method   func() Vec2[int]
		expected Vec2[int]
	}{
		{v.Up, Vec2[int]{5, 4}},
		{v.Down, Vec2[int]{5, 6}},
		{v.Left, Vec2[int]{4, 5}},
		{v.Right, Vec2[int]{6, 5}},
	}
	for _, c := range cases {
		got := c.method()
		if got != c.expected {
			t.Errorf("got %v, want %v", got, c.expected)
		}
	}
}

func TestDirection(t *testing.T) {
	v := Vec2[int]{5, 5}
	cases := []struct {
		r        rune
		expected Vec2[int]
	}{
		{'>', Vec2[int]{6, 5}},
		{'<', Vec2[int]{4, 5}},
		{'^', Vec2[int]{5, 4}},
		{'v', Vec2[int]{5, 6}},
	}
	for _, c := range cases {
		got := v.Direction(c.r)
		if got != c.expected {
			t.Errorf("Direction('%c') = %v, want %v", c.r, got, c.expected)
		}
	}
}
