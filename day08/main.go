package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"time"
)

func main() {
	// lines := utils.ReadLines("day08/input.txt") // line by line
	// grid := utils.ReadGrid("input.txt")   // map 2D [][]rune

	input := utils.ReadLines("day08/example-input.txt")
	// input := utils.ReadLines("day08/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part1(input []string) int {
	remainingPoints := make([]Point, 0)
	for _, line := range input {
		coords := utils.ExtractInts(line)
		if len(coords) >= 3 {
			remainingPoints = append(remainingPoints, Point{coords[0], coords[1], coords[2]})
		}
	}

	distances := make([]Distance, 0)
	for i, point := range remainingPoints {
		for _, other := range remainingPoints[i+1:] {
			distances = append(distances, Distance{point, other, calculateDistanceByPoint(point, other)})
		}
	}

	// sort distances
	Sort(distances, func(a, b Distance) bool {
		return a.distance > b.distance
	})

	total := float64(0)
	for _, distance := range distances {
		// if remainingPoints contains both distance.point1 and point 2
		if Contains(remainingPoints, distance.point1) && Contains(remainingPoints, distance.point2) {
			total += distance.distance
			remainingPoints = Remove(remainingPoints, distance.point1)
			remainingPoints = Remove(remainingPoints, distance.point2)
		}
		if len(remainingPoints) == 0 {
			break
		}
	}
	return int(total)
}

func Remove(points []Point, point Point) []Point {
	for i, p := range points {
		if p == point {
			return append(points[:i], points[i+1:]...)
		}
	}
	return points
}

func Contains(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func Sort(distances []Distance, less func(a, b Distance) bool) {
	for i := 0; i < len(distances); i++ {
		for j := i + 1; j < len(distances); j++ {
			if less(distances[i], distances[j]) {
				distances[i], distances[j] = distances[j], distances[i]
			}
		}
	}
}

func calculateDistance(x1, x2, y1, y2, z1, z2 int) float64 {
	distance := pow2(x1-x2) + pow2(y1-y2) + pow2(z1-z2)
	return math.Sqrt(float64(distance))
}

func calculateDistanceByPoint(point1, point2 Point) float64 {
	return calculateDistance(point1.x, point2.x, point1.y, point2.y, point1.z, point2.z)
}

func pow2(x int) int {
	return int(math.Pow(float64(x), float64(2)))
}

func Part2(input []string) int {
	return 0
}

type Point struct {
	x, y, z int
}

type Distance struct {
	point1, point2 Point
	distance       float64
}

type Circuit struct {
	points []Point
}
