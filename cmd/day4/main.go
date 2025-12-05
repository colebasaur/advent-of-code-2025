package main

import (
	"bufio"
	"fmt"
	"os"
)

type Row []bool
type Grid []Row
type Coord struct {
	x, y int
}

func (g Grid) IsAccessible(x, y int) bool {
	return g.AdjacentRecords(x, y) < 4
}

func (g Grid) isValidCoord(x, y int) bool {
	maxX := len(g)
	maxY := len(g[0])
	if x < 0 || x >= maxX {
		return false
	}
	if y < 0 || y >= maxY {
		return false
	}
	return true
}

func (g Grid) Up(x, y int) bool {
	x--
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) Down(x, y int) bool {
	x++
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) Left(x, y int) bool {
	y--
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) Right(x, y int) bool {
	y++
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) UpLeft(x, y int) bool {
	x--
	y--
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) UpRight(x, y int) bool {
	x--
	y++
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) DownLeft(x, y int) bool {
	x++
	y--
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) DownRight(x, y int) bool {
	x++
	y++
	if !g.isValidCoord(x, y) {
		return false
	}
	return g[x][y]
}

func (g Grid) AdjacentRecords(x, y int) int {
	count := 0
	if g.Up(x, y) {
		count++
	}
	if g.Down(x, y) {
		count++
	}
	if g.Left(x, y) {
		count++
	}
	if g.Right(x, y) {
		count++
	}
	if g.UpLeft(x, y) {
		count++
	}
	if g.UpRight(x, y) {
		count++
	}
	if g.DownRight(x, y) {
		count++
	}
	if g.DownLeft(x, y) {
		count++
	}
	return count
}

func (g Grid) RemoveAccessible() int {
	coords := []Coord{}
	for x, r := range g {
		for y, isPaperRoll := range r {
			if isPaperRoll && g.IsAccessible(x, y) {
				coords = append(coords, Coord{x: x, y: y})
			}
		}
	}
	for _, c := range coords {
		g[c.x][c.y] = false
	}
	return len(coords)
}

func (g Grid) Print() {
	for _, r := range g {
		for _, b := range r {
			if b {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Part1(g Grid) {
	count := 0
	for x, r := range g {
		for y, isPaperRoll := range r {
			if isPaperRoll && g.IsAccessible(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Part2(g Grid) {
	count := 0
	for removed := g.RemoveAccessible(); removed > 0; removed = g.RemoveAccessible() {
		count += removed
	}
	fmt.Println(count)
}

func main() {
	file, err := os.Open("./input/day4.txt")
	if err != nil {
		fmt.Println("You suck, you can't even open a file")
		panic(err)
	}
	defer file.Close()

	grid := Grid{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := Row{}
		line := scanner.Text()
		for _, item := range line {
			row = append(row, item == '@')
		}
		grid = append(grid, row)
	}

	Part1(grid)
	Part2(grid)
}
