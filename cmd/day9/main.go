/*
Example:
	7,1
	11,1
	11,7
	9,7
	9,5
	2,5
	2,3
	7,3
*/

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"sort"

	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type Bounds struct {
	MinX, MaxX int
	MinY, MaxY int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Coord) Area(other Coord) int {
	length := int(math.Abs(float64(c.X-other.X)) + 1)
	height := int(math.Abs(float64(c.Y-other.Y)) + 1)
	return length * height
}

func GetCoords() ([]Coord, Bounds) {
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		panic("..... you suck")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := []Coord{}
	bounds := Bounds{}
	for scanner.Scan() {
		xy := strings.Split(scanner.Text(), ",")
		xStr, yStr := xy[0], xy[1]
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		if x < bounds.MinX || bounds.MinX == 0 {
			bounds.MinX = x
		}
		if x > bounds.MaxX {
			bounds.MaxX = x
		}
		if y < bounds.MinY || bounds.MinY == 0 {
			bounds.MinY = y
		}
		if y > bounds.MaxY {
			bounds.MaxY = y
		}
		coord := Coord{X: x, Y: y}
		coords = append(coords, coord)
	}

	return coords, bounds
}

func Part1(coords []Coord) {
	maxArea := 0
	for i := range len(coords) {
		for j := i + 1; j < len(coords); j++ {
			area := coords[i].Area(coords[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea)
}

func IsInPolygon(c Coord, polygon []Coord) bool {
	inside := false
	for i := range len(polygon) {
		j := (i + 1) % len(polygon)

		a := polygon[i]
		b := polygon[j]

		// Check if point is a vertex
		if (c.X == a.X) && (c.Y == a.Y) || (c.X == b.X) && (c.Y == b.Y) {
			return true
		}

		// Check if point is on horizontal edge
		if a.Y == b.Y && c.Y == a.Y {
			if c.X >= min(a.X, b.X) && c.X <= max(a.X, b.X) {
				return true
			}
		}

		// Check if point is on vertical edge
		if a.X == b.X && c.X == a.X {
			if c.Y >= min(a.Y, b.Y) && c.Y <= max(a.Y, b.Y) {
				return true
			}
		}

		// Original ray casting logic
		// https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule
		if (a.Y > c.Y) != (b.Y > c.Y) {
			slope := (c.X-a.X)*(b.Y-a.Y) - (b.X-a.X)*(c.Y-a.Y)
			if slope == 0 {
				return true
			}
			if (slope < 0) != (b.Y < a.Y) {
				inside = !inside
			}
		}
	}
	return inside
}

func CheckAllInsidePolygon(a, b Coord, polygon []Coord) bool {
	fmt.Printf("Checking area between %s and %s\n", a.String(), b.String())
	minX, maxX := min(a.X, b.X), max(a.X, b.X)
	minY, maxY := min(a.Y, b.Y), max(a.Y, b.Y)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			c := Coord{X: x, Y: y}
			if !IsInPolygon(c, polygon) {
				return false
			}
		}
	}
	return true
}

func Part2(coords []Coord) {
	maxBoundedArea := 0
	percentComplete := 0
	for i := range len(coords) {
		newPercentComplete := (i * 100) / len(coords)
		for j := i + 1; j < len(coords); j++ {

			a := coords[i]
			b := coords[j]
			area := a.Area(b)
			if area <= maxBoundedArea {
				continue
			}
			if CheckAllInsidePolygon(a, b, coords) {
				if area > maxBoundedArea {
					maxBoundedArea = area
				}
			}
		}
		// log every 5 percent
		if newPercentComplete >= percentComplete+5 {
			percentComplete = newPercentComplete
			fmt.Printf("Part 2: %d%% complete\n", percentComplete)
		}
	}
	fmt.Println(maxBoundedArea)
}

func main() {
	// get those coords
	coords, _ := GetCoords()

	for _, c := range coords {
		fmt.Println(c.String())
	}
	fmt.Println("-----")

	sort.Slice(coords, func(i, j int) bool {
		return cmp.Or(cmp.Compare(coords[i].X, coords[j].X), cmp.Compare(coords[i].Y, coords[j].Y)) < 0
	})

	for _, c := range coords {
		fmt.Println(c.String())
	}
	fmt.Println("-----")

	sort.Slice(coords, func(i, j int) bool {
		return cmp.Or(cmp.Compare(coords[i].Y, coords[j].Y), cmp.Compare(coords[i].X, coords[j].X)) < 0
	})

	for _, c := range coords {
		fmt.Println(c.String())
	}

	// Part1(coords)
	// Part2(coords)
}
