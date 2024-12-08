package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	matrix := [][]string{}
	for row := 0; row < len(lines); row++ {
		tokens := strings.Split(lines[row], "")
		matrix = append(matrix, tokens)
	}

	positions := findPositionsForAntenna(matrix)
	result := findAllAntinodes(positions, matrix)

	for _, val := range matrix {
		fmt.Println(val)
	}

	fmt.Println(result)
}

func findAllAntinodes(positions map[string][]position, matrix [][]string) int {
	antinodesCont := 0
	for _, value := range positions {
		combinations := combinations(value)
		for _, value := range combinations {
			deltaY := value[0].row - value[1].row
			deltaX := value[0].col - value[1].col

			newPositionRow := value[0].row + deltaY
			newPositionCol := value[0].col + deltaX
			antinodesCont = markAntinodeAndGetCount(newPositionRow, newPositionCol, matrix, antinodesCont)

			newPositionRow = value[1].row + deltaY*-1
			newPositionCol = value[1].col + deltaX*-1
			antinodesCont = markAntinodeAndGetCount(newPositionRow, newPositionCol, matrix, antinodesCont)
		}
	}

	return antinodesCont
}
func markAntinodeAndGetCount(newPositionRow, newPositionCol int, matrix [][]string, antinodesCont int) int {
	if newPositionRow >= 0 && newPositionRow < len(matrix) && newPositionCol >= 0 && newPositionCol < len(matrix[0]) {
		if matrix[newPositionRow][newPositionCol] != "#" {
			antinodesCont++
		}
		matrix[newPositionRow][newPositionCol] = "#"
	}
	return antinodesCont
}

func combinations(positions []position) [][]position {
	combinations := [][]position{}
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			pair := []position{positions[i], positions[j]}
			combinations = append(combinations, pair)
		}
	}

	return combinations
}

func findPositionsForAntenna(matrix [][]string) map[string][]position {
	positions := map[string][]position{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			currentChar := matrix[row][col]
			if currentChar == "." {
				continue
			}

			if _, ok := positions[currentChar]; !ok {
				positions[currentChar] = []position{}
			}
			positions[currentChar] = append(positions[currentChar], position{
				row: row,
				col: col,
			})
		}
	}
	return positions
}
