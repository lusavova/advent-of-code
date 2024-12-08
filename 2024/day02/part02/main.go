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
		line := scanner.Text()
		numbers := convertLineToNumbers(line)

		if isSafeRow(numbers) {
			result++
			continue
		}

		if isSafeRowAfterModification(numbers) {
			result++
		}
	}

	fmt.Println(result)
}

func isSafeRowAfterModification(numbers []int) bool {
	isSafe := false
	for i, _ := range numbers {
		newNums := append([]int(nil), numbers[:i]...)
		newNums = append(newNums, numbers[i+1:]...)

		if isSafeRow(newNums) {
			isSafe = true
			break
		}
	}
	return isSafe
}

func convertLineToNumbers(line string) []int {
	tokens := strings.Fields(line)

	var numbers []int
	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		numbers = append(numbers, num)
	}
	return numbers
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
