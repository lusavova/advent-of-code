package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row int
	col int
}

type velocity struct {
	horizontal int
	vertical   int
}

type robot struct {
	position position
	velocity velocity
}

const rows = 103
const cols = 101

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	outputFile, err := os.Create("matrix_output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	var robots []robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		inputPosition := strings.Split(parts[0][2:], ",")
		positionX, _ := strconv.Atoi(inputPosition[0])
		positionY, _ := strconv.Atoi(inputPosition[1])

		inputVelocity := strings.Split(parts[1][2:], ",")
		velocityX, _ := strconv.Atoi(inputVelocity[0])
		velocityY, _ := strconv.Atoi(inputVelocity[1])
		robot := robot{
			position: position{
				row: positionY,
				col: positionX,
			},
			velocity: velocity{
				horizontal: velocityX,
				vertical:   velocityY,
			},
		}
		robots = append(robots, robot)
	}

	var matrix [rows][cols]int

	for i := 0; i < 10000; i++ {
		fmt.Println("Second", i)
		writeMatrixToFile(writer, matrix, i)
		for i := range robots {
			newRow, newCol := getNextPosition(rows, cols, robots[i].position, robots[i].velocity)
			if matrix[robots[i].position.row][robots[i].position.col] > 0 {
				matrix[robots[i].position.row][robots[i].position.col] -= 1
			}
			robots[i].position = position{
				row: newRow,
				col: newCol,
			}
			matrix[robots[i].position.row][robots[i].position.col] += 1
		}
	}

	fmt.Println("Second", 10000)
	writeMatrixToFile(writer, matrix, 10000)

	result1, result2, result3, result4 := calculateResults(matrix)
	fmt.Println(result1, result2, result3, result4, result1*result2*result3*result4)
}

func writeMatrixToFile(writer *bufio.Writer, matrix [rows][cols]int, second int) {
	writer.WriteString(fmt.Sprintf("Second %d\n", second))

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				writer.WriteString(".")
			} else {
				writer.WriteString(strconv.Itoa(matrix[row][col]))
			}
		}
		writer.WriteString("\n")
	}

	writer.WriteString("\n")
	writer.Flush()
}

func calculateResults(matrix [rows][cols]int) (int, int, int, int) {
	result1, result2, result3, result4 := 0, 0, 0, 0
	firstRowBounds := rows / 2
	firstColBounds := cols / 2

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if row < firstRowBounds && col < firstColBounds {
				result1 += matrix[row][col]
			} else if row < firstRowBounds && col >= firstColBounds {
				result2 += matrix[row][col]
			} else if row >= firstRowBounds && col < firstColBounds {
				result3 += matrix[row][col]
			} else if row >= firstRowBounds && col >= firstColBounds {
				result4 += matrix[row][col]
			}
		}
	}

	return result1, result2, result3, result4
}

func getNextPosition(rows, cols int, position position, velocity velocity) (int, int) {
	newRow := 0
	if position.row+velocity.vertical < 0 {
		newRow = ((position.row + velocity.vertical) + rows) % rows
	} else {
		newRow = (position.row + velocity.vertical) % rows
	}

	newCol := 0
	if position.col+velocity.horizontal < 0 {
		newCol = ((position.col + velocity.horizontal) + cols) % cols
	} else {
		newCol = (position.col + velocity.horizontal) % cols
	}

	return newRow, newCol
}
