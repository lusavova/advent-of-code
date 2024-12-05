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
	invalidLines := [][]string{}
	for _, update := range updates {
		if !isValidLine(update, pagesOrder) {
			invalidLines = append(invalidLines, update)
		}
	}

	for _, invalidUpdate := range invalidLines {
		orderedUpdate := getValidLine(invalidUpdate, pagesOrder)

		middleElement := orderedUpdate[len(orderedUpdate)/2]
		num, _ := strconv.Atoi(middleElement)
		result += num
	}

	fmt.Println(result)
}

func getValidLine(line []string, pagesOrder map[string][]string) []string {
	graph := map[string][]string{}

	countDependencies := map[string]int{}
	for _, element := range line {
		countDependencies[element] = 0
	}

	for key, value := range pagesOrder {
		if contains(key, line) {
			for _, element := range value {
				if contains(element, line) {
					graph[key] = append(graph[key], element)
					countDependencies[element]++
				}
			}
		}
	}

	orderedLine := []string{}
	queue := []string{}
	for node, numOfDependencies := range countDependencies {
		if numOfDependencies == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		orderedLine = append(orderedLine, current)

		for _, neighbor := range graph[current] {
			countDependencies[neighbor]--
			if countDependencies[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return orderedLine
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
