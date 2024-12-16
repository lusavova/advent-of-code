package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var matrix [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		matrix = append(matrix, tokens)
	}

	printMatrix(matrix)
	fmt.Println(bfsWithPriorityQueue(matrix))
}

type cell struct {
	row       int
	col       int
	score     int
	direction int
}

func bfsWithPriorityQueue(matrix [][]string) int {
	startRow := len(matrix) - 2
	startCol := 1

	directions := []struct {
		rowOffset int
		colOffset int
		dir       int
	}{
		{-1, 0, 0},
		{1, 0, 1},
		{0, -1, 2},
		{0, 1, 3},
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, cell{startRow, startCol, 0, 3})

	bestScores := make(map[[3]int]int)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(cell)

		if matrix[current.row][current.col] == "E" {
			return current.score
		}

		state := [3]int{current.row, current.col, current.direction}
		if bestScore, exists := bestScores[state]; exists && current.score >= bestScore {
			continue
		}

		bestScores[state] = current.score

		for _, d := range directions {
			newRow := current.row + d.rowOffset
			newCol := current.col + d.colOffset
			newDirection := d.dir

			if matrix[newRow][newCol] != "#" {
				moveCost := 1
				rotationCost := 0
				if current.direction != newDirection {
					rotationCost = 1000
				}
				newScore := current.score + moveCost + rotationCost

				heap.Push(pq, cell{newRow, newCol, newScore, newDirection})
			}
		}
	}

	return -1
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

type PriorityQueue []cell

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}
func (pq PriorityQueue) Swap(i, j int) {
	temp := pq[i]
	pq[i] = pq[j]
	pq[j] = temp
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(cell))
}
func (pq *PriorityQueue) Pop() interface{} {
	currentQueue := *pq
	length := len(currentQueue)
	topItem := currentQueue[length-1]
	*pq = currentQueue[:length-1]
	return topItem
}
