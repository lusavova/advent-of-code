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
	coords := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		row, _ := strconv.Atoi(tokens[1])
		col, _ := strconv.Atoi(tokens[0])
		if i == nBytes {
			coords = append(coords, []int{row, col})
		} else {
			matrix[row][col] = "#"
			i++
		}
	}

	for _, coord := range coords {
		matrix[coord[0]][coord[1]] = "#"
		if findShortestPath(matrix) == false {
			fmt.Println(coord)
			printMatrix(matrix)

			return
		}
	}

}

func findShortestPath(matrix [][]string) bool {
	targetRow := matrixSize
	targetCol := matrixSize

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	queue := [][]int{}
	visited := [][]bool{}
	for i := 0; i <= matrixSize; i++ {
		visited = append(visited, make([]bool, matrixSize+1))
	}

	queue = append(queue, []int{0, 0})
	cnt := -1
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			current := queue[0]
			queue = queue[1:]

			row := current[0]
			col := current[1]

			if row == targetRow && col == targetCol {
				return true
			}

			if visited[row][col] || matrix[row][col] == "#" {
				continue
			}
			visited[row][col] = true

			for _, dir := range directions {
				if row+dir[0] >= 0 && row+dir[0] <= matrixSize &&
					col+dir[1] >= 0 && col+dir[1] <= matrixSize {
					queue = append(queue, []int{row + dir[0], col + dir[1]})
				}
			}
		}
		cnt++
	}
	return false
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
