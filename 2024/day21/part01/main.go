package main

import (
	"fmt"
	"math"
	"strconv"
)

type cell struct {
	row int
	col int
}

type state struct {
	position cell
	path     string
}

type helperResult struct {
	paths []string
	key   string
}

var helperCache = map[string]helperResult{}

func main() {
	matrix := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"", "0", "A"},
	}

	keypad := map[string]cell{
		"7": {0, 0}, "8": {0, 1}, "9": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"1": {2, 0}, "2": {2, 1}, "3": {2, 2},
		"0": {3, 1}, "A": {3, 2},
	}

	matrix2 := [][]string{
		{"", "^", "A"},
		{"<", "v", ">"},
	}

	keypad2 := map[string]cell{
		"A": {0, 2},
		"^": {0, 1},
		">": {1, 2},
		"v": {1, 1},
		"<": {1, 0},
	}

	codes := []string{"140A",
		"143A",
		"349A",
		"582A",
		"964A"}

	totalComplexity := 0
	for _, code := range codes {
		paths := processCode(code, keypad, matrix, "A")
		fmt.Println(len(paths), paths)

		secondPaths := processCodeForRobot(paths, keypad2, matrix2, "A")
		finalPaths := findMinLenPaths(secondPaths)
		fmt.Println(len(finalPaths), finalPaths[0])
		thirdPaths := processCodeForRobot(finalPaths, keypad2, matrix2, "A")
		finalPaths = findMinLenPaths(thirdPaths)
		fmt.Println(len(finalPaths), finalPaths[0])

		sequenceLength := len(finalPaths[0])
		codeValue, _ := strconv.Atoi(code[:len(code)-1])
		totalComplexity += sequenceLength * codeValue

	}

	fmt.Println(totalComplexity)
}

func processCode(code string, keypad map[string]cell, matrix [][]string, startKey string) []string {
	start := keypad[startKey]
	queue := []string{""}

	directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	directionLetter := map[int]string{0: "<", 1: "^", 2: ">", 3: "v"}

	for _, char := range code {
		charStr := string(char)
		paths, nextKey := helper(start, charStr, matrix, directions, directionLetter)
		start = keypad[nextKey]

		temp := []string{}
		for _, existing := range queue {
			for _, p := range paths {
				temp = append(temp, existing+p)
			}
		}
		queue = findMinLenPaths(temp)
	}

	return queue
}

func processCodeForRobot(paths []string, keypad map[string]cell, matrix [][]string, startKey string) []string {
	start := keypad[startKey]
	result := []string{}

	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	directionLetter := map[int]string{0: "<", 1: "v", 2: ">", 3: "^"}

	for _, path := range paths {
		current := start
		currentPaths := []string{""}

		for _, char := range path {
			charStr := string(char)
			robotPaths, nextKey := helper(current, charStr, matrix, directions, directionLetter)
			current = keypad[nextKey]

			temp := []string{}
			for _, existing := range currentPaths {
				for _, p := range robotPaths {
					temp = append(temp, existing+p)
				}
			}
			currentPaths = findMinLenPaths(temp)
		}
		result = append(result, currentPaths...)
	}

	return findMinLenPaths(result)
}

func helper(start cell, target string, matrix [][]string, directions [][]int, directionLetter map[int]string) ([]string, string) {
	cacheKey := fmt.Sprintf("%s:%s", matrix[start.row][start.col], target)
	if result, exists := helperCache[cacheKey]; exists {
		return result.paths, result.key
	}

	queue := []state{{start, ""}}
	visited := map[cell]int{}
	var shortestPaths []string
	shortestLength := -1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position.row < 0 || current.position.row >= len(matrix) ||
			current.position.col < 0 || current.position.col >= len(matrix[0]) {
			continue
		}

		if steps, seen := visited[current.position]; seen && len(current.path) > steps {
			continue
		}
		visited[current.position] = len(current.path)

		if matrix[current.position.row][current.position.col] == target {
			if shortestLength == -1 || len(current.path) == shortestLength {
				shortestLength = len(current.path)
				shortestPaths = append(shortestPaths, current.path+"A")
			} else if len(current.path) > shortestLength {
				break
			}
			continue
		}

		for i, dir := range directions {
			newRow := current.position.row + dir[0]
			newCol := current.position.col + dir[1]

			if newRow >= 0 && newRow < len(matrix) &&
				newCol >= 0 && newCol < len(matrix[0]) &&
				matrix[newRow][newCol] != "" {
				queue = append(queue, state{
					position: cell{newRow, newCol},
					path:     current.path + directionLetter[i],
				})
			}
		}
	}

	helperCache[cacheKey] = helperResult{shortestPaths, target}
	return shortestPaths, target
}

func findMinLenPaths(allPaths []string) []string {
	minLen := math.MaxInt
	for _, p := range allPaths {
		if len(p) < minLen {
			minLen = len(p)
		}
	}

	var minPaths []string
	for _, p := range allPaths {
		if len(p) == minLen {
			minPaths = append(minPaths, p)
		}
	}
	return minPaths
}
