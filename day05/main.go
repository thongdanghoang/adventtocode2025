package main

import (
	"adventofcode2025/utils"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	input := utils.ReadInput("day05/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part1(input []string) int {
	var rangeStrs []string
	var ids []string
	splitIdx := 0

	for i, line := range input {
		if line == "" {
			splitIdx = i
			break
		}
	}
	rangeStrs = input[:splitIdx]
	ids = input[splitIdx+1:]

	ranges := parseRanges(rangeStrs)

	count := 0
	for _, idStr := range ids {
		id := utils.ToInt(idStr)
		if isFresh(id, ranges) {
			count++
		}
	}
	return count
}

func Part2(input []string) int {
	var rangeStrs []string
	for i, line := range input {
		if line == "" {
			rangeStrs = input[:i]
			break
		}
	}

	ranges := parseRanges(rangeStrs)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var merged []Range
	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
	}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		lastMerged := &merged[len(merged)-1]

		if current.Start <= lastMerged.End {
			if current.End > lastMerged.End {
				lastMerged.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	totalFresh := 0
	for _, r := range merged {
		totalFresh += r.End - r.Start + 1
	}

	return totalFresh
}

func parseRanges(lines []string) []Range {
	var ranges []Range
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			ranges = append(ranges, Range{
				Start: utils.ToInt(parts[0]),
				End:   utils.ToInt(parts[1]),
			})
		}
	}
	return ranges
}

func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}

type Range struct {
	Start int
	End   int
}
