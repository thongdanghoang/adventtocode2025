package utils

import (
	"regexp"
	"strconv"
	"unicode"
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

func FindLargestDigit(s string) (int, int) {
	maxDigit := -1
	index := -1

	for i, r := range s {
		if unicode.IsDigit(r) {
			digit := int(r - '0')
			if digit > maxDigit {
				maxDigit = digit
				index = i
			}
		}
	}

	return maxDigit, index
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
