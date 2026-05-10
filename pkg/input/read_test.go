package input

import (
	"os"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "read_test_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	return f.Name()
}

func TestFileToString(t *testing.T) {
	path := writeTempFile(t, "hello world")
	got, err := FileToString(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello world" {
		t.Errorf("FileToString = %q, want %q", got, "hello world")
	}
}

func TestFileToStringMissing(t *testing.T) {
	_, err := FileToString("nonexistent.txt")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestFileToLines(t *testing.T) {
	path := writeTempFile(t, "line1\nline2\nline3")
	got, err := FileToLines(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"line1", "line2", "line3"}
	if len(got) != len(expected) {
		t.Fatalf("FileToLines len = %d, want %d", len(got), len(expected))
	}
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("line %d = %q, want %q", i, got[i], v)
		}
	}
}

func TestFileToArrayInt(t *testing.T) {
	path := writeTempFile(t, "1 2 3\n4 5 6")
	got, err := FileToArray[int](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3, 4, 5, 6}
	if len(got) != len(expected) {
		t.Fatalf("FileToArray len = %d, want %d", len(got), len(expected))
	}
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("index %d = %d, want %d", i, got[i], v)
		}
	}
}

func TestFileToArray2DRune(t *testing.T) {
	path := writeTempFile(t, "abc\ndef")
	got, err := FileToArray2D[rune](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0][1] != 'b' {
		t.Errorf("grid[0][1] = %c, want 'b'", got[0][1])
	}
	if got[1][2] != 'f' {
		t.Errorf("grid[1][2] = %c, want 'f'", got[1][2])
	}
}

func TestFileToArray2DInt(t *testing.T) {
	path := writeTempFile(t, "123\n456")
	got, err := FileToArray2D[int](path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0][0] != 1 || got[1][2] != 6 {
		t.Errorf("unexpected values: got[0][0]=%d, got[1][2]=%d", got[0][0], got[1][2])
	}
}

func TestFileToArray2DIntInvalidChar(t *testing.T) {
	path := writeTempFile(t, "1a3")
	_, err := FileToArray2D[int](path)
	if err == nil {
		t.Error("expected error for non-digit character in int grid")
	}
}
