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
	var visited [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		matrix = append(matrix, tokens)
		l := len(matrix[0])
		visited = append(visited, make([]bool, l))
	}

	result := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if !visited[row][col] {
				area, perimeter := findAreaAndPerimeter(matrix[row][col], row, col, matrix, visited)
				result += area * perimeter
			}
		}
	}

	fmt.Println(result)
}

func findAreaAndPerimeter(plant string, row, col int, matrix [][]string, visited [][]bool) (int, int) {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || matrix[row][col] != plant {
		return 0, 1
	}

	if visited[row][col] {
		return 0, 0
	}

	visited[row][col] = true

	area1, perimeter1 := findAreaAndPerimeter(plant, row+1, col, matrix, visited)
	area2, perimeter2 := findAreaAndPerimeter(plant, row-1, col, matrix, visited)
	area3, perimeter3 := findAreaAndPerimeter(plant, row, col+1, matrix, visited)
	area4, perimeter4 := findAreaAndPerimeter(plant, row, col-1, matrix, visited)

	return 1 + area1 + area2 + area3 + area4, perimeter1 + perimeter2 + perimeter3 + perimeter4
}
