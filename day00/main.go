package main

import (
	"adventofcode2025/utils"
	"fmt"
	"time"
)

func main() {
	// lines := utils.ReadLines("day00/input.txt") // line by line
	// grid := utils.ReadGrid("input.txt")   // map 2D [][]rune

	input := utils.ReadLines("day00/example-input.txt")
	// input := utils.ReadLines("day00/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part1(input []string) int {
	return 0
}

func Part2(input []string) int {
	return 0
}
