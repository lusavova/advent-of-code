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

	walls := walls(matrix)
	optimal := bfs(start, end, matrix)
	possibilities := map[int]int{}
	for _, wall := range walls {
		matrix[wall.row][wall.col] = "."
		result := bfs(start, end, matrix)
		if _, ok := possibilities[optimal-result]; !ok {
			possibilities[optimal-result] = 0
		}
		possibilities[optimal-result]++
		matrix[wall.row][wall.col] = "#"
	}

	result := 0
	for picosecond, cnt := range possibilities {
		if picosecond >= 100 {
			result += cnt
		}
	}

	fmt.Println(result)
}

func walls(matrix [][]string) []cell {
	cells := []cell{}
	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[0])-1; col++ {
			if matrix[row][col] == "#" {
				cells = append(cells, cell{row, col})
			}
		}
	}
	return cells
}

func bfs(start, end cell, matrix [][]string) int {
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

	parent := make(map[cell]*cell)

	for len(queue) > 0 {
		steps := len(queue)
		for i := 0; i < steps; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.row == end.row && current.col == end.col {
				return distance
			}

			for _, dir := range directions {
				newRow := current.row + dir[0]
				newCol := current.col + dir[1]

				if !visited[newRow][newCol] && (matrix[newRow][newCol] == "." || matrix[newRow][newCol] == "E") {
					queue = append(queue, cell{row: newRow, col: newCol})
					visited[newRow][newCol] = true
					next := cell{row: newRow, col: newCol}
					parent[next] = &current
				}
			}
		}
		distance++
	}
	return -1
}
