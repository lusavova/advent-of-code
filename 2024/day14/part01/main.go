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
		currentRobot := robot{
			position: position{
				row: positionY,
				col: positionX,
			},
			velocity: velocity{
				horizontal: velocityX,
				vertical:   velocityY,
			},
		}
		robots = append(robots, currentRobot)
	}

	var matrix [rows][cols]int
	for i := range robots {
		newRow, newCol := getNextPosition(rows, cols, robots[i].position, robots[i].velocity)
		robots[i].position = position{
			row: newRow,
			col: newCol,
		}
		matrix[robots[i].position.row][robots[i].position.col] += 1
	}

	printMatrix(matrix)

	result1, result2, result3, result4 := getResult(matrix)
	fmt.Println(result1, result2, result3, result4, result1*result2*result3*result4)

}

func getResult(matrix [103][101]int) (int, int, int, int) {
	result1, result2, result3, result4 := 0, 0, 0, 0
	firstRowBounds := rows / 2
	firstColBounds := cols / 2

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if row < firstRowBounds && col < firstColBounds {
				result1 += matrix[row][col]
			} else if row < firstRowBounds && col > firstColBounds {
				result2 += matrix[row][col]
			} else if row > firstRowBounds && col < firstColBounds {
				result3 += matrix[row][col]
			} else if row > firstRowBounds && col > firstColBounds {
				result4 += matrix[row][col]
			}
		}
	}
	return result1, result2, result3, result4
}

func printMatrix(matrix [rows][cols]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(matrix[row][col])
		}
		fmt.Println()
	}
}

func getNextPosition(rows, cols int, position position, velocity velocity) (int, int) {
	newRow := ((position.row+velocity.vertical*100)%rows + rows) % rows
	newCol := ((position.col+velocity.horizontal*100)%cols + cols) % cols
	return newRow, newCol
}
