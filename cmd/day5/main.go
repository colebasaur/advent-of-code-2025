package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

func SplitRange(s string) Range {
	bounds := strings.Split(s, "-")
	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])
	return Range{min: lower, max: upper}
}

func IsFresh(r Range, id int) bool {
	return id >= r.min && id <= r.max
}

func Part1(ranges []Range, ids []int) {
	fresh := []int{}
	for _, id := range ids {
		for _, r := range ranges {
			if IsFresh(r, id) {
				fresh = append(fresh, id)
				break
			}
		}
	}

	fmt.Println(len(fresh))
}

func main() {
	file, err := os.Open("./input/day5.txt")
	if err != nil {
		panic("Hey man... you gotta figure this out with this file opening")
	}
	scanner := bufio.NewScanner(file)

	ranges := []Range{}
	ids := []int{}

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			ranges = append(ranges, SplitRange(line))
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	Part1(ranges, ids)
}
