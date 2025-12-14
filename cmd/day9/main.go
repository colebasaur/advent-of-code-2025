package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"aoc/utils"
)

type Grid struct {
	Data        [][]bool
	XCompressed map[int]int
	YCompressed map[int]int
}

func (g Grid) Print() {
	for y := range len(g.Data) {
		for x := range len(g.Data[y]) {
			value := g.Data[y][x]
			if value {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g Grid) Up(p Point) Point {
	return Point{x: p.x, y: p.y - 1}
}
func (g Grid) Down(p Point) Point {
	return Point{x: p.x, y: p.y + 1}
}
func (g Grid) Right(p Point) Point {
	return Point{x: p.x + 1, y: p.y}
}
func (g Grid) Left(p Point) Point {
	return Point{x: p.x - 1, y: p.y}
}

type Polygon []Point

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.y, p.x)
}

func area(p1, p2 Point) int {
	width := max(p1.x, p2.x) - min(p1.x, p2.x) + 1
	height := max(p1.y, p2.y) - min(p1.y, p2.y) + 1
	return width * height
}

func getInput() Polygon {
	lines := utils.GetLines(9)
	polygon := Polygon{}
	for _, line := range lines {
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		polygon = append(polygon, Point{x, y})
	}
	return polygon
}

func compressGrid(polygon Polygon) (Grid, Polygon) {
	xSet := map[int]bool{}
	ySet := map[int]bool{}
	for _, p := range polygon {
		xSet[p.x] = true
		ySet[p.y] = true
	}
	xCoords := []int{}
	yCoords := []int{}
	for x := range xSet {
		xCoords = append(xCoords, x)
	}
	for y := range ySet {
		yCoords = append(yCoords, y)
	}
	sort.Ints(xCoords)
	sort.Ints(yCoords)

	xCompressed := map[int]int{}
	yCompressed := map[int]int{}
	for i, x := range xCoords {
		xCompressed[x] = i
	}
	for i, y := range yCoords {
		yCompressed[y] = i
	}

	grid := Grid{XCompressed: xCompressed, YCompressed: yCompressed}
	grid.Data = make([][]bool, len(yCoords))
	for i := range grid.Data {
		grid.Data[i] = make([]bool, len(xCoords))
	}
	comprssedPolygon := Polygon{}

	for i := range len(polygon) {
		j := (i + 1) % len(polygon)
		p1 := polygon[i]
		p2 := polygon[j]

		px1, py1 := xCompressed[p1.x], yCompressed[p1.y]
		px2, py2 := xCompressed[p2.x], yCompressed[p2.y]

		cp1 := Point{x: px1, y: py1}
		cp2 := Point{x: px2, y: py2}
		if i == 0 {
			comprssedPolygon = append(comprssedPolygon, cp1)
		}
		comprssedPolygon = append(comprssedPolygon, cp2)

		if px1 == px2 {
			yStart, yEnd := min(py1, py2), max(py1, py2)
			for y := yStart; y <= yEnd; y++ {
				grid.Data[y][px1] = true
			}
		} else if py1 == py2 {
			xStart, xEnd := min(px1, px2), max(px1, px2)
			for x := xStart; x <= xEnd; x++ {
				grid.Data[py1][x] = true
			}
		}
	}
	return grid, comprssedPolygon
}

func floodFill(g Grid, p Point) {
	if p.y < 0 || p.y >= len(g.Data) || p.x < 0 || p.x >= len(g.Data[0]) {
		return
	}

	if g.Data[p.y][p.x] {
		return
	}

	g.Data[p.y][p.x] = true

	floodFill(g, g.Up(p))
	floodFill(g, g.Down(p))
	floodFill(g, g.Left(p))
	floodFill(g, g.Right(p))
}

func Part1(polygon Polygon) {
	maxArea := 0

	for i := range len(polygon) {
		for j := i + 1; j < len(polygon); j++ {
			p1, p2 := polygon[i], polygon[j]
			maxArea = max(maxArea, area(p1, p2))
		}
	}

	fmt.Println(maxArea)
}

func Part2(grid Grid, polygon Polygon) {
	maxArea := 0

	for i := range len(polygon) {
		for j := i + 1; j < len(polygon); j++ {
			p1, p2 := polygon[i], polygon[j]

			a := area(p1, p2)
			if a <= maxArea {
				continue
			}

			cx1, cy1 := grid.XCompressed[p1.x], grid.YCompressed[p1.y]
			cx2, cy2 := grid.XCompressed[p2.x], grid.YCompressed[p2.y]

			xStart, xEnd := min(cx1, cx2), max(cx1, cx2)
			yStart, yEnd := min(cy1, cy2), max(cy1, cy2)

			valid := true
			for cy := yStart; cy <= yEnd; cy++ {
				for cx := xStart; cx <= xEnd; cx++ {
					if !grid.Data[cy][cx] {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				maxArea = max(maxArea, a)
			}
		}
	}

	fmt.Println(maxArea)
}

func main() {
	timeStart := time.Now()
	polygon := getInput()
	grid, _ := compressGrid(polygon)
	// manually pick a point in the polygon to start flood fill
	floodFill(grid, Point{x: 121, y: 20})
	Part1(polygon)
	Part2(grid, polygon)
	timeEnd := time.Now()
	fmt.Println("Time taken:", timeEnd.Sub(timeStart))
}
