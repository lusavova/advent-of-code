package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	for _, startPosition := range startPositions {
		result += findAllPaths(startPosition.row, startPosition.col, matrix)
	}

	fmt.Println(result)
}

func findAllPaths(row, col int, matrix [][]int) int {
	if matrix[row][col] == 9 {
		return 1
	}

	result := 0
	if row+1 < len(matrix) && matrix[row+1][col] == matrix[row][col]+1 {
		result += findAllPaths(row+1, col, matrix)
	}
	if row-1 >= 0 && matrix[row-1][col] == matrix[row][col]+1 {
		result += findAllPaths(row-1, col, matrix)
	}
	if col+1 < len(matrix[0]) && matrix[row][col+1] == matrix[row][col]+1 {
		result += findAllPaths(row, col+1, matrix)
	}
	if col-1 >= 0 && matrix[row][col-1] == matrix[row][col]+1 {
		result += findAllPaths(row, col-1, matrix)
	}

	return result
}
