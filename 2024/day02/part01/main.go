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

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		numbers := convertInputToNumbers(inputLine)
		if isSafeRow(numbers) {
			result++
		}
	}

	fmt.Println(result)
}

func isSafeRow(numbers []int) bool {
	isIncreasing := numbers[1]-numbers[0] > 0
	for i := 1; i < len(numbers); i++ {
		prev, current := numbers[i-1], numbers[i]

		if prev == current {
			return false
		}

		isCurrentIncreasing := current-prev > 0
		if isCurrentIncreasing != isIncreasing {
			return false
		}

		if absDiff(current, prev) > 3 {
			return false
		}
	}

	return true
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func convertInputToNumbers(inputLine string) []int {
	tokens := strings.Fields(inputLine)
	var numbers []int
	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		numbers = append(numbers, num)
	}
	return numbers
}
