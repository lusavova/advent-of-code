package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	matrix := [][]string{}

	guardRow := 0
	guardCol := 0

	for _, line := range lines {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	reachableCells := [][]int{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == "^" || matrix[row][col] == ">" || matrix[row][col] == "v" || matrix[row][col] == "<" {
				guardRow = row
				guardCol = col
			} else if matrix[row][col] == "." {
				reachableCells = append(reachableCells, []int{row, col})
			}
		}
	}

	result := 0

	for _, cell := range reachableCells {
		clone := cloneMatrix(matrix)
		clone[cell[0]][cell[1]] = "#"
		if hasLoop(guardRow, guardCol, clone) {
			result++
		}
	}

	fmt.Println(result)
}

func cloneMatrix(original [][]string) [][]string {
	copied := make([][]string, len(original))

	for i := range original {
		copied[i] = make([]string, len(original[i]))
		copy(copied[i], original[i])
	}

	return copied
}

type state struct {
	row       int
	col       int
	direction string
}

func hasLoop(currentRow, currentCol int, matrix [][]string) bool {
	turn := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}
	nextCell := map[string][]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}

	visited := map[state]bool{}
	currentState := state{row: currentRow, col: currentCol, direction: matrix[currentRow][currentCol]}
	for {
		if visited[currentState] {
			return true
		}

		visited[currentState] = true

		direction := currentState.direction
		delta := nextCell[direction]
		nextRow := currentState.row + delta[0]
		nextCol := currentState.col + delta[1]

		if nextRow < 0 || nextRow >= len(matrix) || nextCol < 0 || nextCol >= len(matrix[0]) {
			return false
		}

		if matrix[nextRow][nextCol] == "." {
			temp := matrix[currentRow][currentCol]
			matrix[currentRow][currentCol] = "."
			currentRow = nextRow
			currentCol = nextCol
			matrix[currentRow][currentCol] = temp

		} else if matrix[nextRow][nextCol] == "#" {
			newDirection := turn[matrix[currentRow][currentCol]]
			matrix[currentRow][currentCol] = newDirection
			currentState.direction = newDirection
		}
		currentState.row = currentRow
		currentState.col = currentCol
	}
}
