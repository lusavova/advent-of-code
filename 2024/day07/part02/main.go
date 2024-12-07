package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	var result int64
	for _, line := range lines {
		tokens := strings.Split(line, ": ")

		sum, _ := strconv.Atoi(tokens[0])
		targetSum := int64(sum)
		numbersAsStrings := strings.Split(tokens[1], " ")
		numbers := []int64{}
		for _, num := range numbersAsStrings {
			temp, _ := strconv.Atoi(num)
			numbers = append(numbers, int64(temp))
		}

		if matchTarget(numbers, numbers[0], 1, targetSum) {
			result += targetSum
		}
	}

	fmt.Println(result)
}

func matchTarget(numbers []int64, runningSum int64, index int, target int64) bool {
	if index == len(numbers) {
		return runningSum == target
	}

	currentStr := strconv.FormatInt(runningSum, 10) + strconv.FormatInt(numbers[index], 10)
	value, _ := strconv.ParseInt(currentStr, 10, 64)
	return matchTarget(numbers, runningSum+numbers[index], index+1, target) ||
		matchTarget(numbers, runningSum*numbers[index], index+1, target) ||
		matchTarget(numbers, value, index+1, target)
}
