package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part01(lines []string) {
	result := 0
	for _, line := range lines {
		game := strings.Split(line, ": ")
		gameID, err := strconv.Atoi(strings.Replace(game[0], "Game ", "", 1))
		if err != nil {
			return
		}

		isGamePossible, err := isGamePossible(game)
		if err != nil {
			return
		}
		if isGamePossible {
			result += gameID
		}
	}

	fmt.Println("Part 01: ", result)
}

func isGamePossible(game []string) (bool, error) {
	set := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	isGamePossible := true
	subsets := strings.Split(game[1], "; ")
	for _, subset := range subsets {
		tokens := strings.Split(subset, ", ")
		for _, token := range tokens {
			pair := strings.Split(token, " ")
			color := pair[1]
			numOfItems, err := strconv.Atoi(pair[0])
			if err != nil {
				return false, err
			}
			if set[color] < numOfItems {
				isGamePossible = false
				break
			}
		}
		if !isGamePossible {
			break
		}
	}

	return isGamePossible, nil
}
