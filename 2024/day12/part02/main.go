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
				_, area := findPerimeter(matrix[row][col], row, col, matrix, visited, "", directions)
				fmt.Println(matrix[row][col])

				final := 0
				for key, val := range directions {
					c := sortSlice(key, val)
					final += c
				}
				fmt.Println(final, area, final*area)
				fmt.Println()
				result += final * area
			}
		}
	}

	fmt.Println(result)
}

func sortSlice(direction string, tuples []tuple) int {
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

func findPerimeter(plant string, row, col int, matrix [][]string, visited [][]bool, direction string, directions map[string][]tuple) (int, int) {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || matrix[row][col] != plant {
		if _, ok := directions[direction]; !ok {
			directions[direction] = []tuple{}
		}

		directions[direction] = append(directions[direction], tuple{row, col})
		return 1, 0
	}

	if visited[row][col] {
		return 0, 0
	}

	visited[row][col] = true
	perimeter1, area1 := findPerimeter(plant, row+1, col, matrix, visited, "down", directions)
	perimeter2, area2 := findPerimeter(plant, row-1, col, matrix, visited, "up", directions)
	perimeter3, area3 := findPerimeter(plant, row, col+1, matrix, visited, "right", directions)
	perimeter4, area4 := findPerimeter(plant, row, col-1, matrix, visited, "left", directions)
	return perimeter1 + perimeter2 + perimeter3 + perimeter4, 1 + area1 + area2 + area3 + area4
}
