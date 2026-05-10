package array

import "testing"

func TestInsertIndex(t *testing.T) {
	a := Array[int]{1, 2, 3}
	a.InsertIndex(99, 1)
	expected := Array[int]{1, 99, 2, 3}
	for i, v := range expected {
		if a[i] != v {
			t.Errorf("InsertIndex: index %d = %d, want %d", i, a[i], v)
		}
	}
}

func TestRemoveIndex(t *testing.T) {
	a := Array[int]{1, 2, 3, 4}
	a.RemoveIndex(1)
	expected := Array[int]{1, 3, 4}
	for i, v := range expected {
		if a[i] != v {
			t.Errorf("RemoveIndex: index %d = %d, want %d", i, a[i], v)
		}
	}
}

func TestMoveIndex(t *testing.T) {
	a := Array[int]{1, 2, 3, 4}
	a.MoveIndex(0, 2)
	expected := Array[int]{2, 3, 1, 4}
	for i, v := range expected {
		if a[i] != v {
			t.Errorf("MoveIndex: index %d = %d, want %d", i, a[i], v)
		}
	}
}

func TestPop(t *testing.T) {
	a := Array[int]{10, 20, 30}
	got := a.Pop()
	if got != 10 {
		t.Errorf("Pop() = %d, want 10", got)
	}
	if len(a) != 2 {
		t.Errorf("len after Pop = %d, want 2", len(a))
	}
}

func TestPopEmpty(t *testing.T) {
	a := Array[int]{}
	got := a.Pop()
	if got != 0 {
		t.Errorf("Pop() on empty = %d, want 0", got)
	}
}
