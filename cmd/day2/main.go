package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		fmt.Println("You suck, you can't even open a file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	ranges := strings.Split(input, ",")
	invalidIDs := []int{}
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		l, u := bounds[0], bounds[1]
		lower, _ := strconv.Atoi(l)
		upper, _ := strconv.Atoi(u)
		for i := lower; i <= upper; i++ {
			if !isValidIDPart2(i) {
				invalidIDs = append(invalidIDs, i)
			}
		}
	}

	sum := 0
	for _, id := range invalidIDs {
		sum += id
	}

	fmt.Println(invalidIDs)
	fmt.Println()
	fmt.Println(sum)
}

// Part 1
func isValidID(id int) bool {
	// Invalid id is one that is made only of some sequence
	// of digits repeated twice.
	strID := strconv.Itoa(id)

	// If not even length, it is valid
	if len(strID)%2 != 0 {
		return true
	}

	half := len(strID) / 2
	left, right := strID[:half], strID[half:]

	if left == right {
		return false
	}

	return true
}

// Part 2
func isValidIDPart2(id int) bool {
	strID := strconv.Itoa(id)

	// handle case where length == 1
	if len(strID) <= 1 {
		return true
	}

	// start at splitting in half and then work towards splitting to just 1 char
	for splitSize := len(strID) / 2; splitSize >= 1; splitSize-- {
		// need to split the string into equal parts to check
		// so continue if not possible for splitsize
		if len(strID)%splitSize != 0 {
			continue
		}

		// iterate through the equal parts and compare them to each other
		repeats := true
		for i := 0; i < len(strID)-splitSize; i += splitSize {
			left, right := strID[i:i+splitSize], strID[i+splitSize:i+splitSize*2]
			if left != right {
				repeats = false
				break
			}
		}
		// Return early if we found a repeating sequence
		if repeats {
			return false
		}
	}

	return true
}
