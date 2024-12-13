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
	priceX  int
	priceY  int
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
				x, _ := strconv.Atoi(matches[1])
				y, _ := strconv.Atoi(matches[2])
				currentMove.priceX = x
				currentMove.priceY = y
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

	result := 0
	for _, m := range moves {
		firstButton := m.buttons[0]
		firstButtonX := firstButton.x
		firstButtonY := firstButton.y

		secondButton := m.buttons[1]
		secondButtonX := secondButton.x
		secondButtonY := secondButton.y

		for y := 1; y < 100; y++ {
			if (m.priceX-secondButtonX*y)%firstButtonX == 0 {
				x := (m.priceX - secondButtonX*y) / firstButtonX
				if x*firstButtonY+y*secondButtonY == m.priceY {
					result += x*3 + y*1
					break
				}
			}

		}
	}

	fmt.Println(result)
}
