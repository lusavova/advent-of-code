package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	leftColumn := []int{}
	occurrences := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)

		leftNum, _ := strconv.Atoi(tokens[0])
		leftColumn = append(leftColumn, leftNum)

		rightNum, _ := strconv.Atoi(tokens[1])
		occurrences[rightNum]++
	}

	result := 0
	for _, num := range leftColumn {
		if _, exists := occurrences[num]; exists {
			result += occurrences[num] * num
		}
	}

	fmt.Println(result)
}
