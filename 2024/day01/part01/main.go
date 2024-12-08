package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var leftColumn, rightColumn []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)

		leftNum, _ := strconv.Atoi(tokens[0])
		rightNum, _ := strconv.Atoi(tokens[1])

		leftColumn = append(leftColumn, leftNum)
		rightColumn = append(rightColumn, rightNum)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	result := 0
	for index, leftNum := range leftColumn {
		result += abs(leftNum - rightColumn[index])
	}

	fmt.Println(result)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
