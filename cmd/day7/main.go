package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid [][]string

func (g Grid) Print() {
	for _, chars := range g {
		for _, char := range chars {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func Part1(g Grid) {
	activeIndexes := map[int]bool{}
	count := 0
	start := strings.Index(strings.Join(g[0], ""), "S")
	activeIndexes[start] = true
	for _, row := range g {
		for j, active := range activeIndexes {
			if !active {
				continue
			}

			switch row[j] {
			case ".":
				// nothing happends
				continue
			case "^":
				count++
				left, right := j-1, j+1
				activeIndexes[j] = false
				if left >= 0 && !activeIndexes[left] {
					activeIndexes[left] = true
				}
				if right < len(row) && !activeIndexes[right] {
					activeIndexes[right] = true
				}
			}
		}
	}

	fmt.Println(count)
}

func Part2(g Grid) {
	activeIndexes := map[int]int{}
	start := strings.Index(strings.Join(g[0], ""), "S")
	activeIndexes[start] = 1
	for _, row := range g {
		for j, totalPaths := range activeIndexes {
			switch row[j] {
			case ".":
				// nothing happends
				continue
			case "^":
				left, right := j-1, j+1
				// delete
				delete(activeIndexes, j)
				// if left on the board
				if left >= 0 {
					// if left has nothing, set to ttoal
					if activeIndexes[left] == 0 {
						activeIndexes[left] = totalPaths
					} else {
						activeIndexes[left] += totalPaths
					}
				}
				// if right on the board
				if right < len(row) {
					if activeIndexes[right] == 0 {
						activeIndexes[right] = totalPaths
					} else {
						activeIndexes[right] += totalPaths
					}
				}
			}
		}
	}

	sum := 0
	for _, v := range activeIndexes {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	file, err := os.Open("./input/day7.txt")
	if err != nil {
		panic("just open the file correctly please")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := Grid{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}

	Part1(grid)
	Part2(grid)
}
