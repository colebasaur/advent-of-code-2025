package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func MergeRanges(r1, r2 Range) Range {
	return Range{min: min(r1.min, r2.min), max: max(r1.max, r2.max)}
}

func Part2(ranges []Range) {
	// Need to converge Ranges to where they don't overlap
	// Then sum'em
	for {
		hasChanged := false

		newRanges := []Range{}
		// Keep a list of indexes combined
		combined := []int{}
		for i := 0; i < len(ranges); i++ {
			curRange := ranges[i]

			// We've already combined this one, move to the next
			if slices.Contains(combined, i) {
				continue
			}

			// We are checking the first against all, so no need to start j at 0
			for j := i + 1; j < len(ranges); j++ {
				compareRange := ranges[j]
				if IsFresh(compareRange, curRange.min) || IsFresh(compareRange, curRange.max) || IsFresh(curRange, compareRange.min) || IsFresh(curRange, compareRange.max) {
					curRange = MergeRanges(curRange, compareRange)
					combined = append(combined, j)
					hasChanged = true
				}
			}
			newRanges = append(newRanges, curRange)
		}

		ranges = newRanges
		if !hasChanged {
			break
		}
	}

	sum := 0
	for _, r := range ranges {
		sum += r.max - r.min + 1 // + 1 because inclusive range
	}
	fmt.Println(sum)
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
	Part2(ranges)
}
