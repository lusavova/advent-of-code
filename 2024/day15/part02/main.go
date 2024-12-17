package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
			var matrixRow []string
			for _, token := range tokens {
				switch token {
				case "#":
					matrixRow = append(matrixRow, "#", "#")
				case "O":
					matrixRow = append(matrixRow, "[", "]")
				case ".":
					matrixRow = append(matrixRow, ".", ".")
				case "@":
					matrixRow = append(matrixRow, "@", ".")
					row = len(matrix)
					col = len(matrixRow) - 2
				}
			}
			matrix = append(matrix, matrixRow)
		}
	}

	directions := map[string][]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}
	printMatrix(matrix)

	for _, move := range moves {
		//fmt.Println("MOVE", move)
		offsetRow := directions[move][0]
		offsetCol := directions[move][1]
		newRow := row + offsetRow
		newCol := col + offsetCol

		if matrix[newRow][newCol] == "." {
			matrix[row][col], matrix[newRow][newCol] = matrix[newRow][newCol], matrix[row][col]
			row, col = newRow, newCol
		} else if matrix[newRow][newCol] == "[" || matrix[newRow][newCol] == "]" {
			var boxesToMove [][]int
			visited := make(map[string]bool)

			queue := [][]int{{row, col}}
			visited[fmt.Sprintf("%d,%d", row, col)] = true

			for len(queue) > 0 {
				box := queue[0]
				queue = queue[1:]

				currentRow := box[0]
				currentCol := box[1]
				newRow = currentRow + offsetRow
				newCol = currentCol + offsetCol

				if matrix[newRow][newCol] == "[" || matrix[newRow][newCol] == "]" {
					if matrix[newRow][newCol] == "[" {
						key1 := fmt.Sprintf("%d,%d", newRow, newCol)
						key2 := fmt.Sprintf("%d,%d", newRow, newCol+1)

						if !visited[key1] {
							visited[key1] = true
							boxesToMove = append(boxesToMove, []int{newRow, newCol})
							queue = append(queue, []int{newRow, newCol})
						}
						if !visited[key2] {
							visited[key2] = true
							boxesToMove = append(boxesToMove, []int{newRow, newCol + 1})
							queue = append(queue, []int{newRow, newCol + 1})
						}
					}
					if matrix[newRow][newCol] == "]" {
						key1 := fmt.Sprintf("%d,%d", newRow, newCol-1)
						key2 := fmt.Sprintf("%d,%d", newRow, newCol)

						if !visited[key1] {
							visited[key1] = true
							boxesToMove = append(boxesToMove, []int{newRow, newCol - 1})
							queue = append(queue, []int{newRow, newCol - 1})
						}
						if !visited[key2] {
							visited[key2] = true
							boxesToMove = append(boxesToMove, []int{newRow, newCol})
							queue = append(queue, []int{newRow, newCol})
						}
					}
				}
			}

			if move == ">" {
				sort.Slice(boxesToMove, func(i, j int) bool {
					return boxesToMove[i][1] < boxesToMove[j][1]
				})
			} else if move == "<" {
				sort.Slice(boxesToMove, func(i, j int) bool {
					return boxesToMove[i][1] > boxesToMove[j][1]
				})
			} else if move == "^" {
				sort.Slice(boxesToMove, func(i, j int) bool {
					return boxesToMove[i][0] > boxesToMove[j][0]
				})
			} else if move == "v" {
				sort.Slice(boxesToMove, func(i, j int) bool {
					return boxesToMove[i][0] < boxesToMove[j][0]
				})
			}
			//fmt.Println(boxesToMove)
			if len(boxesToMove) > 0 {
				matrixCopy := copyMatrix(matrix)
				if moveBoxes(boxesToMove, matrixCopy, directions[move]) {
					matrix = copyMatrix(matrixCopy)
					matrix[row][col] = "."
					matrix[row+offsetRow][col+offsetCol] = "@"
					row = row + offsetRow
					col = col + offsetCol
				}
			}

		}
		//printMatrix(matrix)
	}

	result := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == "[" {
				result += row*100 + col
			}
		}
	}

	fmt.Println(result)
}

func moveBoxes(boxes [][]int, matrix [][]string, direction []int) bool {
	offsetRow := direction[0]
	offsetCol := direction[1]
	for i := len(boxes) - 1; i >= 0; i-- {
		row := boxes[i][0]
		col := boxes[i][1]
		nextRow := row + offsetRow
		nextCol := col + offsetCol
		if matrix[nextRow][nextCol] == "#" {
			return false
		}

		matrix[nextRow][nextCol], matrix[row][col] = matrix[row][col], matrix[nextRow][nextCol]
	}
	return true
}

func printMatrix(matrix [][]string) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			fmt.Print(matrix[row][col])
		}
		fmt.Println()
	}
}

func copyMatrix(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}

	return newMatrix
}
