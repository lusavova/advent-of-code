package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	leftArr := []int{}
	rightArr := []int{}
	for _, line := range lines {
		tokens := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(tokens[0])
		rightNum, _ := strconv.Atoi(tokens[1])

		leftArr = append(leftArr, leftNum)
		rightArr = append(rightArr, rightNum)
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	result := 0
	for i := 0; i < len(leftArr); i++ {
		n := leftArr[i] - rightArr[i]
		if n < 0 {
			n *= -1
		}

		result += n
	}

	fmt.Println(result)
}
