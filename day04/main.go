package main

import (
	"adventofcode2025/utils"
	"fmt"
	"time"
)

func main() {
	grid := utils.ReadGrid("day04/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(grid), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(grid), time.Since(start))
}

func Part1(input [][]rune) int {
	return count(input)
}

func calculateRollsOfPaper(input [][]rune, x int, y int) bool {
	minX := utils.Max(0, x-1)
	maxX := utils.Min(x+1, len(input[0])-1)
	minY := utils.Max(0, y-1)
	maxY := utils.Min(y+1, len(input)-1)

	rollsOfPaper := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if i == x && j == y {
				continue
			}
			if input[i][j] == '@' {
				rollsOfPaper++
			}
		}
	}
	if rollsOfPaper < 4 {
		return true
	}
	return false
}

func count(input [][]rune) int {
	var count int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '.' {
				continue
			}
			if calculateRollsOfPaper(input, i, j) {
				count++
			}
		}
	}
	return count
}

func replace(input [][]rune) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '.' {
				continue
			}
			if calculateRollsOfPaper(input, i, j) {
				input[i][j] = 'x'
			}
		}
	}
}

func Part2(input [][]rune) int {
	var currentCount = count(input)
	replace(input)
	for count(input) > currentCount {
		currentCount = count(input)
		replace(input)
	}
	return count(input)
}
