package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type button struct {
	buttonType string
	x          int
	y          int
	cost       int
}

type move struct {
	buttons []button
	priceX  int64
	priceY  int64
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	buttonPattern := `Button (A|B): X\+(\d+), Y\+(\d+)`
	prizePattern := `Prize: X=(\d+), Y=(\d+)`

	moves := []move{}
	var currentMove move

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button") {
			regex := regexp.MustCompile(buttonPattern)
			matches := regex.FindStringSubmatch(line)
			if len(matches) > 0 {
				firstNum, _ := strconv.Atoi(matches[2])
				secondNum, _ := strconv.Atoi(matches[3])
				button := button{
					buttonType: matches[1],
					x:          firstNum,
					y:          secondNum,
				}
				if button.buttonType == "A" {
					button.cost = 3
				} else {
					button.cost = 1
				}
				currentMove.buttons = append(currentMove.buttons, button)
			}
		} else if strings.HasPrefix(line, "Prize") {
			regex := regexp.MustCompile(prizePattern)
			matches := regex.FindStringSubmatch(line)
			if len(matches) > 0 {
				x, _ := strconv.ParseInt(matches[1], 10, 64)
				y, _ := strconv.ParseInt(matches[2], 10, 64)
				currentMove.priceX = x + 10000000000000
				currentMove.priceY = y + 10000000000000
			}
		} else {
			if len(currentMove.buttons) > 0 {
				moves = append(moves, currentMove)
				currentMove = move{buttons: []button{}}
			}
		}
	}

	if len(currentMove.buttons) > 0 {
		moves = append(moves, currentMove)
		currentMove = move{buttons: []button{}}
	}

	result := int64(0)
	for _, m := range moves {
		a := m.buttons[0]
		a1 := float64(a.x)
		a2 := float64(a.y)

		b := m.buttons[1]
		b1 := float64(b.x)
		b2 := float64(b.y)

		c1 := float64(m.priceX)
		c2 := float64(m.priceY)

		// cramer's rule
		// https://www.youtube.com/watch?v=vXqlIOX2itM
		temp := a1*b2 - b1*a2
		if temp == 0 {
			continue
		}

		tempX := c1*b2 - c2*b1
		tempY := a1*c2 - a2*c1

		x := tempX / temp
		y := tempY / temp

		if x != float64(int64(x)) || y != float64(int64(y)) {
			continue
		}

		result += 3*int64(x) + int64(y)
	}

	fmt.Println(result)
}
