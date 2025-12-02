package aoc

import (
	"bufio"
	"os"
)

// ReadLines reads a file and returns its lines as a slice of strings
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadCommaSeparatedLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		return splitCommaSeparated(line), nil
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func splitCommaSeparated(line string) []string {
	var result []string
	current := ""
	for _, char := range line {
		if char == ',' {
			result = append(result, current)
			current = ""
		} else {
			current += string(char)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
