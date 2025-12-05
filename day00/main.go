package main

import (
	"adventofcode2025/utils"
	"fmt"
	"time"
)

func main() {
	// Setup input path
	input := utils.ReadInput("input.txt")

	// Measure Part 1
	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	// Measure Part 2
	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part1(input []string) int {
	// TODO: Solve Part 1
	return 0
}

func Part2(input []string) int {
	// TODO: Solve Part 2
	return 0
}
