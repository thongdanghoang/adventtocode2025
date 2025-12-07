package utils

import (
	"regexp"
	"strconv"
)

// ExtractInts extracts all integers from a string.
func ExtractInts(line string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)

	res := make([]int, 0, len(matches))
	for _, match := range matches {
		val, _ := strconv.Atoi(match)
		res = append(res, val)
	}
	return res
}

// ParsePattern parses a line using a regex pattern.
func ParsePattern(line string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {
		return nil
	}

	return matches[1:]
}
