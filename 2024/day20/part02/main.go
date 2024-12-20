package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cell struct {
	row int
	col int
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var start cell
	var end cell
	matrix := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, "")
		matrix = append(matrix, elements)
		for i, el := range elements {
			if el == "S" {
				start = cell{row: len(matrix) - 1, col: i}
			} else if el == "E" {
				end = cell{row: len(matrix) - 1, col: i}
			}
		}
	}

	_, paths := bfs(start, end, matrix)
	finalPath := backtrackPath(paths, start, end)
	fmt.Println(calculateCheatsUsingManhattanDistance(finalPath))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateCheatsUsingManhattanDistance(path []cell) int {
	validCheats := 0
	originalPathLength := len(path) - 1

	for i := 0; i < len(path); i++ {
		for j := i + 1; j < len(path); j++ {
			manhattanDistance := abs(path[j].row-path[i].row) + abs(path[j].col-path[i].col)

			if manhattanDistance <= 20 {
				remainingSteps := len(path) - 1 - j
				totalTimeWithCheat := i + manhattanDistance + remainingSteps

				timeSaved := originalPathLength - totalTimeWithCheat

				if timeSaved >= 100 {
					validCheats++
				}
			}
		}
	}

	return validCheats
}

func bfs(start, end cell, matrix [][]string) (int, map[cell]cell) {
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	queue := []cell{}
	queue = append(queue, start)
	visited[start.row][start.col] = true
	distance := 0

	parent := make(map[cell]cell)

	for len(queue) > 0 {
		steps := len(queue)
		for i := 0; i < steps; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.row == end.row && current.col == end.col {
				return distance, parent
			}

			for _, dir := range directions {
				newRow := current.row + dir[0]
				newCol := current.col + dir[1]

				if !visited[newRow][newCol] && (matrix[newRow][newCol] == "." || matrix[newRow][newCol] == "E") {
					queue = append(queue, cell{row: newRow, col: newCol})
					visited[newRow][newCol] = true
					next := cell{row: newRow, col: newCol}
					parent[next] = current
				}
			}
		}
		distance++
	}
	return -1, parent
}

func backtrackPath(parent map[cell]cell, start, end cell) []cell {
	path := []cell{}

	current := end
	for current != start {
		path = append([]cell{current}, path...)
		current = parent[current]
	}
	path = append([]cell{start}, path...)
	return path
}
