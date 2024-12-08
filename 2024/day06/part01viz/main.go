package main

import (
	. "aoc2024/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./testInput.txt")
	lines := strings.Split(string(data), "\n")

	matrix := [][]string{}
	visited := [][]bool{}

	guardRow := 0
	guardCol := 0

	for _, line := range lines {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	for range lines {
		visited = append(visited, make([]bool, len(matrix[0])))
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			visited[row][col] = false
			if matrix[row][col] == "^" || matrix[row][col] == ">" || matrix[row][col] == "v" || matrix[row][col] == "<" {
				guardRow = row
				guardCol = col
			}
		}
	}

	ClearScreen()
	fmt.Println(countCells(guardRow, guardCol, matrix, visited))
}

func colorAndPrintMatrix(matrix [][]string) {
	MoveCursor(0, 0)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			color := Color.Gray
			switch matrix[row][col] {
			case ".":
				color = Color.Gray
			case "#":
				color = Color.Blue
			case "^", "v", "<", ">", "*":
				color = Color.Green
			}
			fmt.Printf("%s%s%s ", color, matrix[row][col], Reset)
		}
		fmt.Println()
	}

	Sleep()
}

func countCells(currentRow, currentCol int, matrix [][]string, visited [][]bool) int {
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

	countCells := 1
	visited[currentRow][currentCol] = true

	for currentRow >= 0 && currentRow < len(matrix) &&
		currentCol >= 0 && currentCol < len(matrix[0]) {

		delta := nextCell[matrix[currentRow][currentCol]]
		nextRow := currentRow + delta[0]
		nextCol := currentCol + delta[1]
		needToBreak := true
		if nextRow >= 0 && nextRow < len(matrix) &&
			nextCol >= 0 && nextCol < len(matrix[0]) {

			if matrix[nextRow][nextCol] == "." || matrix[nextRow][nextCol] == "*" {
				if !visited[currentRow][currentCol] {
					countCells++
				}

				temp := matrix[currentRow][currentCol]
				matrix[currentRow][currentCol] = "*"
				visited[currentRow][currentCol] = true

				currentRow = nextRow
				currentCol = nextCol
				matrix[currentRow][currentCol] = temp
			} else if matrix[nextRow][nextCol] == "#" {
				matrix[currentRow][currentCol] = turn[matrix[currentRow][currentCol]]
			}
			needToBreak = false
		}

		if needToBreak {
			break
		}
		colorAndPrintMatrix(matrix)
	}
	return countCells + 1
}
