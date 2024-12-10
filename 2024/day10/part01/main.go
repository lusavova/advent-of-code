package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tuple struct {
	row int
	col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var matrix [][]int
	var startPositions []tuple
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []int{})
		for _, token := range line {
			num, _ := strconv.Atoi(string(token))
			matrix[index] = append(matrix[index], num)
			if num == 0 {
				startPositions = append(startPositions, tuple{
					row: index,
					col: len(matrix[index]) - 1,
				})
			}
		}
		index++
	}

	result := 0
	visitedPositions := map[string]bool{}
	for _, startPosition := range startPositions {
		startPositionString := strings.Join([]string{strconv.Itoa(startPosition.row), strconv.Itoa(startPosition.col)}, "")
		result += findUniquePaths(startPosition.row, startPosition.col, matrix, startPositionString, visitedPositions)
	}
	fmt.Println(result)
}

func findUniquePaths(row, col int, matrix [][]int, startPositionStr string, visitedPositions map[string]bool) int {
	if matrix[row][col] == 9 {
		finalPositionStr := strings.Join([]string{strconv.Itoa(row), strconv.Itoa(col)}, "") + startPositionStr
		if _, ok := visitedPositions[finalPositionStr]; !ok {
			visitedPositions[finalPositionStr] = true
			return 1
		}
		return 0
	}

	result := 0
	if row+1 < len(matrix) && matrix[row+1][col] == matrix[row][col]+1 {
		result += findUniquePaths(row+1, col, matrix, startPositionStr, visitedPositions)
	}
	if row-1 >= 0 && matrix[row-1][col] == matrix[row][col]+1 {
		result += findUniquePaths(row-1, col, matrix, startPositionStr, visitedPositions)
	}
	if col+1 < len(matrix[0]) && matrix[row][col+1] == matrix[row][col]+1 {
		result += findUniquePaths(row, col+1, matrix, startPositionStr, visitedPositions)
	}
	if col-1 >= 0 && matrix[row][col-1] == matrix[row][col]+1 {
		result += findUniquePaths(row, col-1, matrix, startPositionStr, visitedPositions)
	}

	return result
}
