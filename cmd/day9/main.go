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
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func GetCoords() []Coord {
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		panic("..... you suck")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := []Coord{}
	for scanner.Scan() {
		xy := strings.Split(scanner.Text(), ",")
		xStr, yStr := xy[0], xy[1]
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		coord := Coord{X: x, Y: y}
		coords = append(coords, coord)
	}

	return coords
}

func main() {
	coords := GetCoords()
	for _, c := range coords {
		fmt.Println(c)
	}
}
