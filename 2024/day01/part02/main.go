package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	leftArr := []int{}
	occurences := map[int]int{}
	for _, line := range lines {
		tokens := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(tokens[0])
		leftArr = append(leftArr, leftNum)

		rightNum, _ := strconv.Atoi(tokens[1])
		if _, ok := occurences[rightNum]; !ok {
			occurences[rightNum] = 0
		}
		occurences[rightNum]++
	}

	result := 0
	for _, num := range leftArr {
		if _, ok := occurences[num]; !ok {
			continue
		}
		result += occurences[num] * num
	}

	fmt.Println(result)
}
