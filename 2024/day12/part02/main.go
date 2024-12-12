package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type tuple struct {
	row int
	col int
}

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
				directions := map[string][]tuple{}
				area := findAreaAndAddSides(matrix[row][col], row, col, matrix, visited, "", directions)

				final := 0
				for key, val := range directions {
					c := calculateSides(key, val)
					final += c
				}
				fmt.Println(matrix[row][col], final, area, final*area)
				result += final * area
			}
		}
	}

	fmt.Println(result)
}

func calculateSides(direction string, tuples []tuple) int {
	dict := map[int][]int{}
	if direction == "up" || direction == "down" {
		for _, t := range tuples {
			if _, ok := dict[t.row]; !ok {
				dict[t.row] = []int{}
			}
			dict[t.row] = append(dict[t.row], t.col)
		}
	} else {
		for _, t := range tuples {
			if _, ok := dict[t.col]; !ok {
				dict[t.col] = []int{}
			}
			dict[t.col] = append(dict[t.col], t.row)
		}
	}

	cnt := 0
	for _, val := range dict {
		sort.Ints(val)
		if len(val) == 1 {
			cnt++
			continue
		}
		temp := 1
		for i := 1; i < len(val); i++ {
			if val[i] != val[i-1]+1 {
				temp++
			}
		}

		cnt += temp
	}

	return cnt
}

func findAreaAndAddSides(plant string, row, col int, matrix [][]string, visited [][]bool, direction string, directions map[string][]tuple) int {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || matrix[row][col] != plant {
		if _, ok := directions[direction]; !ok {
			directions[direction] = []tuple{}
		}

		directions[direction] = append(directions[direction], tuple{row, col})
		return 0
	}

	if visited[row][col] {
		return 0
	}

	visited[row][col] = true
	return 1 + findAreaAndAddSides(plant, row+1, col, matrix, visited, "down", directions) +
		findAreaAndAddSides(plant, row-1, col, matrix, visited, "up", directions) +
		findAreaAndAddSides(plant, row, col+1, matrix, visited, "right", directions) +
		findAreaAndAddSides(plant, row, col-1, matrix, visited, "left", directions)
}
