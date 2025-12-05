package utils

import (
	"regexp"
	"strconv"
)

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

func ParsePattern(line string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {
		return nil
	}

	return matches[1:]
}
