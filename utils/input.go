package utils

import (
	"os"
	"strconv"
	"strings"
)

// ReadInput reads the entire file and returns a slice of strings (lines).
// It trims the trailing newline of the file but keeps internal empty lines.
func ReadInput(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err) // Panic is fine for AoC scripts if input is missing
	}
	// Normalize Windows line endings just in case
	s := strings.ReplaceAll(string(content), "\r\n", "\n")
	// Trim trailing newline only
	s = strings.TrimSuffix(s, "\n")
	return strings.Split(s, "\n")
}

func ReadRaw(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(b), "\r\n", "\n")
}

func ReadLines(path string) []string {
	content := ReadRaw(path)
	content = strings.TrimSuffix(content, "\n")
	return strings.Split(content, "\n")
}

func ReadGrid(path string) [][]rune {
	lines := ReadLines(path)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func ReadMatrixInt(path string, sep string) [][]int {
	lines := ReadLines(path)
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, sep)
		var row []int
		for _, p := range parts {
			if p == "" {
				continue
			}
			val, _ := strconv.Atoi(p)
			row = append(row, val)
		}
		matrix[i] = row
	}
	return matrix
}

func ReadInts(path string) []int {
	lines := ReadLines(path)
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		nums = append(nums, ToInt(line))
	}
	return nums
}
