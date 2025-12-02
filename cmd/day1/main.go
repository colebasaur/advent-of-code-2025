package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		fmt.Println("You suck, you can't even open a file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := 50
	clicks := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("bro, do you even parse ints?")
			panic(err)
		}

		step := map[string]int{"R": 1, "L": -1}[direction]
		var c int
		cur, c = rotate(cur, value, step)
		clicks += c
	}

	fmt.Println(clicks)
}

func rotate(cur, value, step int) (int, int) {
	clicks := 0
	absValue := max(value, -value)
	for range absValue {
		cur += step
		if cur > 99 {
			cur = 0
		}
		if cur < 0 {
			cur = 99
		}
		if cur == 0 {
			clicks++
		}
	}
	return cur, clicks
}
