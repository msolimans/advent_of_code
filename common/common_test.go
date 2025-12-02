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

func TestReadCommaSeparatedLine(t *testing.T) {
	values, err := ReadCommaSeparatedLine("./test/comma.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedValues := []string{"apple", "banana", "cherry"}

	if len(values) != len(expectedValues) {
		t.Fatalf("Expected %d values, got %d", len(expectedValues), len(values))
	}

	for i, value := range values {
		if value != expectedValues[i] {
			t.Errorf("Expected value %d to be %q, got %q", i, expectedValues[i], value)
		}
	}
}
