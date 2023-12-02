package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part02(lines []string) {
	result := 0
	for _, line := range lines {
		maxValues, err := findMaxValues(line)
		if err != nil {
			return
		}

		power := maxValues[0] * maxValues[1] * maxValues[2]
		result += power
	}

	fmt.Println("Part 02:", result)
}

func findMaxValues(line string) ([]int, error) {
	game := strings.Split(line, ": ")
	subsets := strings.Split(game[1], "; ")

	maxValues := make([]int, 3) // 0: red, 1: blue, 2: green

	for _, subset := range subsets {
		tokens := strings.Split(subset, ", ")

		for _, token := range tokens {
			pair := strings.Split(token, " ")
			color := pair[1]
			numOfItems, err := strconv.Atoi(pair[0])
			if err != nil {
				return nil, err
			}

			index := colorIndex(color)
			if numOfItems > maxValues[index] {
				maxValues[index] = numOfItems
			}
		}
	}

	return maxValues, nil
}

func colorIndex(color string) int {
	switch color {
	case "red":
		return 0
	case "blue":
		return 1
	case "green":
		return 2
	}
	return -1
}
