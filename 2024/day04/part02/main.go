package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")

	result := 0
	lines := strings.Split(string(data), "\n")
	matrix := make([][]string, len(lines))
	for row := 0; row < len(lines); row++ {
		tokens := strings.Split(lines[row], "")
		matrix[row] = tokens
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == "A" && isAMatch(row, col, matrix) {
				result++
			}
		}
	}

	fmt.Println(result)
}

func isAMatch(row, col int, matrix [][]string) bool {
	if row-1 >= 0 && row+1 < len(matrix) && col-1 >= 0 && col+1 < len(matrix[0]) {
		topLeft := matrix[row-1][col-1]
		topRight := matrix[row-1][col+1]
		bottomLeft := matrix[row+1][col-1]
		bottomRight := matrix[row+1][col+1]

		return (topLeft == "S" && bottomLeft == "M" && topRight == "S" && bottomRight == "M") ||
			(topLeft == "M" && bottomLeft == "S" && topRight == "M" && bottomRight == "S") ||
			(topLeft == "S" && bottomLeft == "S" && topRight == "M" && bottomRight == "M") ||
			(topLeft == "M" && bottomLeft == "M" && topRight == "S" && bottomRight == "S")
	}
	return false
}
