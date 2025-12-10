package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	input := utils.ReadLines("day08/input.txt")

	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part2(input []string) int {
	points := parseInput(input)
	distances := generateSortedDistances(points)
	parent := initUnionFind(len(points))

	numComponents := len(points)

	for _, d := range distances {
		root1 := findRoot(parent, d.idx1)
		root2 := findRoot(parent, d.idx2)

		if root1 != root2 {
			parent[root2] = root1
			numComponents--

			if numComponents == 1 {
				return points[d.idx1].x * points[d.idx2].x
			}
		}
	}

	return 0
}

func parseInput(input []string) []Point {
	points := make([]Point, 0)
	for _, line := range input {
		coords := utils.ExtractInts(line)
		if len(coords) >= 3 {
			points = append(points, Point{coords[0], coords[1], coords[2]})
		}
	}
	return points
}

func generateSortedDistances(points []Point) []Distance {
	distances := make([]Distance, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := calculateDistanceByPoint(points[i], points[j])
			distances = append(distances, Distance{points[i], points[j], i, j, dist})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	return distances
}

func initUnionFind(size int) []int {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return parent
}

func Part1(input []string) int {
	points := make([]Point, 0)
	for _, line := range input {
		coords := utils.ExtractInts(line)
		if len(coords) >= 3 {
			points = append(points, Point{coords[0], coords[1], coords[2]})
		}
	}

	distances := make([]Distance, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := calculateDistanceByPoint(points[i], points[j])
			distances = append(distances, Distance{points[i], points[j], i, j, dist})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}

	limit := 1000
	if len(distances) < limit {
		limit = len(distances)
	}

	for i := 0; i < limit; i++ {
		p1Idx := distances[i].idx1
		p2Idx := distances[i].idx2

		root1 := findRoot(parent, p1Idx)
		root2 := findRoot(parent, p2Idx)

		if root1 != root2 {
			parent[root2] = root1
		}
	}

	circuitSizes := make(map[int]int)
	for i := range points {
		root := findRoot(parent, i)
		circuitSizes[root]++
	}

	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	result := 1
	count := 0
	for _, s := range sizes {
		result *= s
		count++
		if count == 3 {
			break
		}
	}

	return result
}

func findRoot(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = findRoot(parent, parent[i])
	return parent[i]
}

func calculateDistance(x1, x2, y1, y2, z1, z2 int) float64 {
	dist := pow2(x1-x2) + pow2(y1-y2) + pow2(z1-z2)
	return math.Sqrt(float64(dist))
}

func calculateDistanceByPoint(point1, point2 Point) float64 {
	return calculateDistance(point1.x, point2.x, point1.y, point2.y, point1.z, point2.z)
}

func pow2(x int) int {
	return int(math.Pow(float64(x), float64(2)))
}

type Point struct {
	x, y, z int
}

type Distance struct {
	point1, point2 Point
	idx1, idx2     int
	distance       float64
}
