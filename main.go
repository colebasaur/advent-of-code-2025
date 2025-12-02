package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("You suck, you can't even open a file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := 50
	clicks := 0
	var c int

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("bro, do you even parse ints?")
		}

		if direction == "L" {
			cur, c = rotate(cur, -1*value)
		} else {
			cur, c = rotate(cur, value)

		}

		clicks += c
	}

	fmt.Println(clicks)
}

func rotate(cur, value int) (int, int) {
	clicks := 0
	increment := 1
	if value < 0 {
		increment = -1
	}
	absValue := max(value, -value)
	for range absValue {
		cur += increment
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
