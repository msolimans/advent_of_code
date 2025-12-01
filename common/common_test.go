package aoc

import "testing"

func TestReadLines(t *testing.T) {
	lines, err := ReadLines("./test/sample.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedLines := []string{
		"First line",
		"Second line",
		"Third line",
	}

	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Expected line %d to be %q, got %q", i, expectedLines[i], line)
		}
	}
}
