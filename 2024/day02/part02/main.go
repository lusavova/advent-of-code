package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	result := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		numbers := []int{}
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			numbers = append(numbers, num)
		}

		if isSafeRow(numbers) {
			result++
		} else {
			for i, _ := range numbers {
				newNums := append([]int(nil), numbers[:i]...)
				newNums = append(newNums, numbers[i+1:]...)
				if isSafeRow(newNums) {
					result++
					break
				}
			}
		}
	}

	fmt.Println(result)
}

func isSafeRow(numbers []int) bool {
	isIncreasing := numbers[1]-numbers[0] > 0
	for i := 1; i < len(numbers); i++ {
		prev := numbers[i-1]
		current := numbers[i]
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
	diff := x - y
	if diff < 0 {
		return diff * -1
	}

	return diff
}
