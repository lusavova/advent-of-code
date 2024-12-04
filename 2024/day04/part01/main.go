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

	directions := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			for _, direction := range directions {
				if matrix[row][col] == "X" && isAMatch(row, col, 0, direction[0], direction[1], matrix) {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}

func isAMatch(row, col, index, rowDirection, colDirection int, matrix [][]string) bool {
	if index == len("XMAS") {
		return true
	}

	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != string("XMAS"[index]) {
		return false
	}

	return isAMatch(row+rowDirection, col+colDirection, index+1, rowDirection, colDirection, matrix)
}
