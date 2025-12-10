package main

import (
	"adventofcode2025/utils"
	"fmt"
	"time"
)

func main() {
	// lines := utils.ReadLines("day00/input.txt") // line by line
	// grid := utils.ReadGrid("input.txt")   // map 2D [][]rune

	input := utils.ReadLines("day07/input.txt")
	// input := utils.ReadLines("day00/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part2(input []string) int {
	if len(input) == 0 {
		return 0
	}

	height := len(input)
	width := len(input[0])
	grid := make(map[Point]rune)
	var startPos Point

	for y, line := range input {
		for x, char := range line {
			p := Point{x, y}
			grid[p] = char
			if char == 'S' {
				startPos = p
			}
		}
	}

	timelines := make(map[int]uint64)
	timelines[startPos.x] = 1

	totalExited := uint64(0)

	for y := startPos.y; y < height; y++ {
		nextTimelines := make(map[int]uint64)

		for x, count := range timelines {
			targetY := y + 1

			processBeam(x, targetY, count, width, height, grid, nextTimelines, &totalExited)
		}
		timelines = nextTimelines
	}

	return int(totalExited)
}

func processBeam(x, y int, count uint64, width, height int, grid map[Point]rune, nextTimelines map[int]uint64, totalExited *uint64) {
	if y >= height {
		*totalExited += count
		return
	}

	if x < 0 || x >= width {
		return
	}

	char := grid[Point{x, y}]

	if char == '^' {
		processBeam(x-1, y, count, width, height, grid, nextTimelines, totalExited)
		processBeam(x+1, y, count, width, height, grid, nextTimelines, totalExited)
	} else {
		nextTimelines[x] += count
	}
}

type Point struct {
	x, y int
}

func Part1(input []string) int {
	if len(input) == 0 {
		return 0
	}

	height := len(input)
	width := len(input[0])
	grid := make(map[Point]rune)
	var startPos Point

	// 1. Parse Input
	for y, line := range input {
		for x, char := range line {
			p := Point{x, y}
			grid[p] = char
			if char == 'S' {
				startPos = p
			}
		}
	}

	queue := []Point{startPos}

	visited := make(map[Point]bool)
	visited[startPos] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		next := Point{curr.x, curr.y + 1}

		if next.y >= height {
			continue
		}

		if visited[next] {
			continue
		}

		visited[next] = true

		char := grid[next]

		if char == '.' || char == 'S' {
			queue = append(queue, next)
		} else if char == '^' {
			left := Point{next.x - 1, next.y}
			right := Point{next.x + 1, next.y}

			if left.x >= 0 && left.x < width {
				if !visited[left] {
					visited[left] = true
					queue = append(queue, left)
				}
			}

			if right.x >= 0 && right.x < width {
				if !visited[right] {
					visited[right] = true
					queue = append(queue, right)
				}
			}
		}
	}

	splitCount := 0
	for p := range visited {
		if grid[p] == '^' {
			splitCount++
		}
	}

	return splitCount
}
