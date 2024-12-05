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

	pagesOrder := map[string][]string{}
	updates := [][]string{}
	updates = parseInput(lines, updates, pagesOrder)

	result := 0
	for _, update := range updates {
		if isValidLine(update, pagesOrder) {
			middleElement := update[len(update)/2]
			num, _ := strconv.Atoi(middleElement)
			result += num
		}
	}

	fmt.Println(result)
}

func isValidLine(line []string, pagesOrder map[string][]string) bool {
	for i, current := range line {
		for _, next := range line[i+1:] {
			if !contains(next, pagesOrder[current]) {
				return false
			}
		}
	}
	return true
}

func contains(target string, elements []string) bool {
	for _, el := range elements {
		if target == el {
			return true
		}
	}

	return false
}

func parseInput(lines []string, updates [][]string, pagesOrder map[string][]string) [][]string {
	indexOfNewLine := -1
	for i, line := range lines {
		if line == "" {
			indexOfNewLine = i
			continue
		}
		if indexOfNewLine != -1 {
			row := strings.Split(line, ",")
			updates = append(updates, row)
		} else {
			tokens := strings.Split(line, "|")
			if _, ok := pagesOrder[tokens[0]]; !ok {
				pagesOrder[tokens[0]] = []string{}
			}
			pagesOrder[tokens[0]] = append(pagesOrder[tokens[0]], tokens[1])
		}
	}
	return updates
}
