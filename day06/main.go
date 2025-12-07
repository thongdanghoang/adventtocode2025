package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	input := utils.ReadLines("day06/puzzle-input.txt")
	input2 := utils.ReadGrid("day06/puzzle-input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input2), time.Since(start))
}

func Part1(input []string) int {
	arrays := make([][]string, len(RemoveEmptyElements(strings.Split(input[0], " "))))
	for _, line := range input {
		split := RemoveEmptyElements(strings.Split(line, " "))
		for j, col := range split {
			arrays[j] = append(arrays[j], col)
		}
	}
	result := 0
	for _, col := range arrays {
		numbers := col[0 : len(col)-1]
		operator := col[len(col)-1]
		val := totalBig(numbers, operator)
		result += int(val.Int64())
	}
	return result
}

func RemoveEmptyElements(slice []string) []string {
	var newSlice []string
	for _, element := range slice {
		if element != "" {
			newSlice = append(newSlice, element)
		}
	}
	return newSlice
}

func totalBig(numbers []string, operator string) *big.Int {
	if len(numbers) == 0 {
		return big.NewInt(0)
	}

	res := big.NewInt(0)

	if operator == "*" {
		res.SetInt64(1)
		for _, number := range numbers {
			if number != "" {
				n := new(big.Int)
				n.SetString(number, 10)
				res.Mul(res, n)
			}
		}
	} else if operator == "+" {
		for _, number := range numbers {
			if number != "" {
				n := new(big.Int)
				n.SetString(number, 10)
				res.Add(res, n)
			}
		}
	}
	return res
}

func Part2(grid [][]rune) string {
	if len(grid) == 0 {
		return "0"
	}

	rows := len(grid)
	cols := len(grid[0])

	grandTotal := big.NewInt(0)

	blockStart := 0

	for c := 0; c <= cols; c++ {
		isEmpty := true
		if c < cols {
			for r := 0; r < rows; r++ {
				if grid[r][c] != ' ' {
					isEmpty = false
					break
				}
			}
		}

		if isEmpty {
			if c > blockStart {
				blockResult := solveBlock(grid, blockStart, c-1)
				grandTotal.Add(grandTotal, blockResult)
			}
			blockStart = c + 1
		}
	}

	return grandTotal.String()
}

func solveBlock(grid [][]rune, startCol, endCol int) *big.Int {
	rows := len(grid)

	operator := ""
	for c := startCol; c <= endCol; c++ {
		char := grid[rows-1][c]
		if char == '+' || char == '*' {
			operator = string(char)
			break
		}
	}

	var numbers []string

	for c := endCol; c >= startCol; c-- {
		numStr := ""
		for r := 0; r < rows-1; r++ {
			char := grid[r][c]
			if char != ' ' {
				numStr += string(char)
			}
		}

		if numStr != "" {
			numbers = append(numbers, numStr)
		}
	}

	return totalBig(numbers, operator)
}
