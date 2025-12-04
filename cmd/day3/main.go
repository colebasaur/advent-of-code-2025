package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		fmt.Println("You suck, you can't even open a file")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	banks := []string{}
	for scanner.Scan() {
		bank := scanner.Text()
		banks = append(banks, bank)
	}
	joltages := []int{}
	maxDigits := 2
	for _, bank := range banks {
		joltage := toInt(bank[:maxDigits])
		joltageStr := bank[:maxDigits]
		for i := maxDigits; i < len(bank); i++ {
			opt1 := toInt(string(bank[x]) + string(bank[i]))
			opt2 := toInt(string(bank[y]) + string(bank[i]))

			maxOpt := max(opt1, opt2)
			if maxOpt <= joltage {
				continue
			}

			joltage = maxOpt

			if opt1 > opt2 {
				y = i
			} else {
				x = i
			}
		}
		joltages = append(joltages, joltage)
	}

	sum := 0
	for _, j := range joltages {
		sum += j
	}

	fmt.Println(sum)
	// fmt.Println(joltages)
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func getMinIndex(s string) int {
	mindex := 0
	for i := 1; i < len(s); i++ {
		if s[mindex] > s[i] {
			mindex = 1
		}
	}
	return mindex
}
