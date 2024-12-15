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
	initialRow := -1
	initialCol := -1
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
					initialRow = len(matrix) - 1
					initialCol = i
				}
			}
		}

	}

	//printMatrix(matrix)
	fmt.Println()
	for _, move := range moves {
		fmt.Println("MOVE", move)

		initialRow, initialCol = moveBoxes(initialRow, initialCol, move, matrix)
		//printMatrix(matrix)
	}

	printMatrix(matrix)

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

func moveBoxes(row int, col int, move string, matrix [][]string) (int, int) {
	if move == "^" {
		switch matrix[row-1][col] {
		case "#":
			return row, col
		case ".":
			matrix[row][col] = "."
			matrix[row-1][col] = "@"
			return row - 1, col
		case "O":
			i := row - 1
			for matrix[i][col] == "O" {
				i--
			}
			if matrix[i][col] == "." {
				matrix[i][col] = "O"
				matrix[row-1][col] = "@"
				matrix[row][col] = "."
				return row - 1, col
			}
		}
	} else if move == ">" { // DONE
		switch matrix[row][col+1] {
		case "#":
			return row, col
		case ".":
			matrix[row][col] = "."
			matrix[row][col+1] = "@"
			return row, col + 1
		case "O":
			i := col + 1
			for matrix[row][i] == "O" {
				i++
			}
			if matrix[row][i] == "." {
				matrix[row][i] = "O"
				matrix[row][col+1] = "@"
				matrix[row][col] = "."
				return row, col + 1
			}
		}
	} else if move == "v" {
		switch matrix[row+1][col] {
		case "#":
			return row, col
		case ".":
			matrix[row][col] = "."
			matrix[row+1][col] = "@"
			return row + 1, col
		case "O":
			i := row + 1
			for matrix[i][col] == "O" {
				i++
			}
			if matrix[i][col] == "." {
				matrix[i][col] = "O"
				matrix[row+1][col] = "@"
				matrix[row][col] = "."
				return row + 1, col
			}
		}
	} else if move == "<" {
		switch matrix[row][col-1] {
		case "#":
			return row, col
		case ".":
			matrix[row][col] = "."
			matrix[row][col-1] = "@"
			return row, col - 1
		case "O":
			i := col - 1
			for matrix[row][i] == "O" {
				i--
			}
			if matrix[row][i] == "." {
				matrix[row][i] = "O"
				matrix[row][col-1] = "@"
				matrix[row][col] = "."
				return row, col - 1
			}
		}
	}

	return row, col
}

func printMatrix(matrix [][]string) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			fmt.Print(matrix[row][col])
		}
		fmt.Println()
	}
}
