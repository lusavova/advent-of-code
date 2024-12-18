package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const matrixSize = 70
const nBytes = 1024

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	matrix := createMatrix()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		if i == nBytes {
			break
		}
		row, _ := strconv.Atoi(tokens[1])
		col, _ := strconv.Atoi(tokens[0])
		matrix[row][col] = "#"
		i++
	}

	fmt.Println(findShortestPath(matrix))
	printMatrix(matrix)
}

func findShortestPath(matrix [][]string) int {
	targetRow := matrixSize
	targetCol := matrixSize

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	queue := [][]int{}
	queue = append(queue, []int{0, 0})
	cnt := 1
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			current := queue[0]
			queue = queue[1:]

			row := current[0]
			col := current[1]

			if row == targetRow && col == targetCol {
				return cnt
			}

			if matrix[row][col] == "O" || matrix[row][col] == "#" {
				continue
			}
			matrix[row][col] = "O"

			for _, dir := range directions {
				if row+dir[0] >= 0 && row+dir[0] <= matrixSize &&
					col+dir[1] >= 0 && col+dir[1] <= matrixSize {
					queue = append(queue, []int{row + dir[0], col + dir[1]})
				}
			}
		}
		cnt++
	}
	return cnt
}

func createMatrix() [][]string {
	var matrix [][]string
	for row := 0; row <= matrixSize; row++ {
		matrix = append(matrix, make([]string, matrixSize+1))
		for col := 0; col <= matrixSize; col++ {
			matrix[row][col] = "."
		}
	}
	return matrix
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}
