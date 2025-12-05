package utils

import (
	"strconv"
	"strings"
)

// ToInt converts a string to an int, panics on error (fail fast approach for AoC).
func ToInt(s string) int {
	val, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return val
}

// ToIntSlice converts a slice of strings to a slice of ints.
func ToIntSlice(input []string) []int {
	nums := make([]int, len(input))
	for i, s := range input {
		nums[i] = ToInt(s)
	}
	return nums
}
