package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var matrix [][]string
	var moves []string
	row := -1
	col := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

		} else if strings.HasPrefix(line, "<") ||
			strings.HasPrefix(line, ">") ||
			strings.HasPrefix(line, "^") ||
			strings.HasPrefix(line, "v") {
			moves = append(moves, strings.Split(line, "")...)
		} else {
			tokens := strings.Split(line, "")
			matrix = append(matrix, tokens)

			for i, token := range tokens {
				if token == "@" {
					row = len(matrix) - 1
					col = i
				}
			}
		}

	}

	directions := map[string][]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}
	for _, move := range moves {
		offsetRow := directions[move][0]
		offsetCol := directions[move][1]
		newRow := row + offsetRow
		newCol := col + offsetCol

		if matrix[newRow][newCol] == "." {
			matrix[row][col], matrix[newRow][newCol] = matrix[newRow][newCol], matrix[row][col]
			row, col = newRow, newCol
		} else if matrix[newRow][newCol] == "O" {
			var boxesToMove [][]int
			for matrix[newRow][newCol] == "O" {
				boxesToMove = append(boxesToMove, []int{newRow, newCol})
				newRow += offsetRow
				newCol += offsetCol
			}
			if matrix[newRow][newCol] == "." {
				for i := len(boxesToMove) - 1; i >= 0; i-- {
					currentRow := boxesToMove[i][0]
					currentCol := boxesToMove[i][1]
					matrix[newRow][newCol], matrix[currentRow][currentCol] = matrix[currentRow][currentCol], matrix[newRow][newCol]
					newRow, newCol = currentRow, currentCol
				}
				matrix[newRow][newCol], matrix[row][col] = matrix[row][col], matrix[newRow][newCol]
				row, col = newRow, newCol
			}
		}
	}

	result := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == "O" {
				result += row*100 + col
			}
		}
	}

	fmt.Println(result)
}

func printMatrix(matrix [][]string) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			fmt.Print(matrix[row][col])
		}
		fmt.Println()
	}
}
