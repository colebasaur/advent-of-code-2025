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
	for _, bank := range banks {
		joltStr := largestNDigitSubsequence(bank, 12)
		joltages = append(joltages, toInt(joltStr))
	}

	sum := 0
	for _, j := range joltages {
		sum += j
	}

	fmt.Println(sum)
	// fmt.Println(joltages)
}

func largestNDigitSubsequence(s string, n int) string {
	removeLeft := len(s) - n
	sequence := []byte{}

	// 415131 (n = 3) (drop 3)
	// ------
	// [4], [41], [5] (drop 1 ... dropped 4, 1)
	// [51], [53] (drop 0 ... dropped 1)
	// dropped is 0, return 531 as they are last n remaining digits
	for i := range len(s) {
		char := s[i]

		// while we can still drop more characters
		// and the last character in the sequence is less than the current character
		for removeLeft > 0 && len(sequence) > 0 && sequence[len(sequence)-1] < char {
			// drop the last character from the sequence
			sequence = sequence[:len(sequence)-1]
			removeLeft--
		}

		sequence = append(sequence, char)
	}

	// only pull first n characters if we didn't drop enough
	return string(sequence[:n])
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
