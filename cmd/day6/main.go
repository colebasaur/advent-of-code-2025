package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
)

type Grid struct {
	Data    [][]string
	Columns []int
}

func (g *Grid) AddRow(row []string) {
	g.Data = append(g.Data, row)
}

func (g *Grid) AddColumn(i int) {
	g.Columns = append(g.Columns, i)
}

func (g *Grid) Print() {
	for _, r := range g.Data {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (g *Grid) GetColumnOp(i int) string {
	opRow := len(g.Data) - 1
	opCol := 0
	if i > 0 {
		opCol = g.Columns[i-1] + 1
	}
	return g.Data[opRow][opCol]
}

func (g *Grid) GetHumanColumnNumbers(i int) []int {
	lenRows := len(g.Data) - 1

	colStart := 0
	colEnd := g.Columns[i]
	if i > 0 {
		colStart = g.Columns[i-1] + 1
	}

	columnNumbers := []int{}
	for x := range lenRows {
		numStr := ""
		for y := colStart; y < colEnd; y++ {
			item := g.Data[x][y]
			if item != " " {
				numStr += item
			}
		}
		numInt, _ := strconv.Atoi(numStr)
		columnNumbers = append(columnNumbers, numInt)
	}

	return columnNumbers
}

// Only difference between this and GetHumanColumnNumbers is the order of x and y loops
func (g *Grid) GetColumnNumbers(i int) []int {
	lenRows := len(g.Data) - 1

	colStart := 0
	colEnd := g.Columns[i]
	if i > 0 {
		colStart = g.Columns[i-1] + 1
	}

	columnNumbers := []int{}
	// This is the only difference, just swap x and y loops
	for y := colStart; y < colEnd; y++ {
		numStr := ""
		for x := range lenRows {
			item := g.Data[x][y]
			if item != " " {
				numStr += item
			}
		}
		numInt, _ := strconv.Atoi(numStr)
		columnNumbers = append(columnNumbers, numInt)
	}

	return columnNumbers
}

func Operate(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	default:
		err := fmt.Errorf("Unsupported Op %s", op)
		panic(err)
	}
}

func Part1(grid Grid) {
	answers := []int{}
	for i := range len(grid.Columns) {
		numbers := grid.GetHumanColumnNumbers(i)
		operation := grid.GetColumnOp(i)
		result := map[string]int{"+": 0, "*": 1}[operation]
		for _, n := range numbers {
			result = Operate(operation, result, n)
		}
		answers = append(answers, result)
	}

	sum := 0
	for _, a := range answers {
		sum += a
	}
	fmt.Println(sum)
}

// Only difference between this and Part1 is not getting human column numbers
func Part2(grid Grid) {
	answers := []int{}
	for i := range len(grid.Columns) {
		numbers := grid.GetColumnNumbers(i)
		operation := grid.GetColumnOp(i)
		result := map[string]int{"+": 0, "*": 1}[operation]
		for _, n := range numbers {
			result = Operate(operation, result, n)
		}
		answers = append(answers, result)
	}

	sum := 0
	for _, a := range answers {
		sum += a
	}
	fmt.Println(sum)
}

func main() {
	file, err := os.Open("./input/day6.txt")
	if err != nil {
		panic("seriously?")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := Grid{}
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid.AddRow(row)
	}

	for y := 0; y < len(grid.Data[0]); y++ {
		isColumn := true
		for x := 0; x < len(grid.Data); x++ {
			if grid.Data[x][y] != " " {
				isColumn = false
				break
			}
		}
		if isColumn {
			grid.AddColumn(y)
		}
	}
	// Add the last "column end" index
	grid.AddColumn(len(grid.Data[0]))

	Part1(grid)
	Part2(grid)
}
