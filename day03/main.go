package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// lines := utils.ReadLines("input.txt") // line by line
	// grid := utils.ReadGrid("input.txt")   // map 2D [][]rune

	input := utils.ReadLines("day03/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2 Optimal: %v (took %v)\n", Part2Optimal(input), time.Since(start))
}

func Part1(input []string) int {
	var total int
	for i := 0; i < len(input); i++ {
		cur := input[i]
		subString := cur[0 : len(cur)-1]

		largestDigit, largestDigitIndex := utils.FindLargestDigit(subString)
		secondLargestDigit, _ := utils.FindLargestDigit(cur[largestDigitIndex+1:])
		total += largestDigit*10 + secondLargestDigit
	}
	return total
}

func Part2(input []string) int {
	nDigits := 12
	var total int

	for i := 0; i < len(input); i++ {
		currentTotal := utils.ToInt(findLargestNumberOf(input[i], nDigits))
		total += currentTotal
	}
	return total
}

func findLargestNumberOf(input string, nDigits int) string {
	if len(input) < nDigits {
		return "0"
	}

	var result string
	cur := 0
	for i := 0; i < nDigits; i++ {
		curRanges := input[cur : len(input)-nDigits+i+1]
		if len(curRanges) == 1 {
			result += input[cur:]
			i = nDigits
		} else {
			largest, index := utils.FindLargestDigit(curRanges)
			result += strconv.Itoa(largest)
			cur += index + 1
		}
	}
	return result
}

func Part2Optimal(input []string) int {
	nDigits := 12
	var total int

	for i := 0; i < len(input); i++ {
		currentTotal := utils.ToInt(findLargestNumberOfMonotonicStack(input[i], nDigits))
		total += currentTotal
	}
	return total
}

func findLargestNumberOfMonotonicStack(input string, nDigits int) string {
	if len(input) < nDigits {
		return "0"
	}

	dropQuota := len(input) - nDigits

	stack := make([]byte, 0, len(input))

	for i := 0; i < len(input); i++ {
		char := input[i]

		for len(stack) > 0 && dropQuota > 0 && char > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			dropQuota--
		}

		stack = append(stack, char)
	}

	if len(stack) > nDigits {
		stack = stack[:nDigits]
	}

	return string(stack)
}
