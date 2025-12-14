package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetLines(day int) []string {
	filePath := fmt.Sprintf("input/day%d.txt", day)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
