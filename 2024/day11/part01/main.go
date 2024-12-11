package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	numbersAsStr := strings.Split(string(data), " ")

	multiplier := big.NewInt(2024)
	leftNum := new(big.Int)
	rightNum := new(big.Int)
	temp := new(big.Int)

	for i := 0; i < 25; i++ {
		newNumbers := make([]string, 0, len(numbersAsStr)*2)
		for _, numAsStr := range numbersAsStr {
			if numAsStr == "0" {
				newNumbers = append(newNumbers, "1")
			} else if len(numAsStr)%2 == 0 {
				mid := len(numAsStr) / 2
				leftNum.SetString(numAsStr[:mid], 10)
				rightNum.SetString(numAsStr[mid:], 10)

				newNumbers = append(newNumbers, leftNum.String(), rightNum.String())
			} else {
				temp.SetString(numAsStr, 10)
				temp.Mul(temp, multiplier)
				newNumbers = append(newNumbers, temp.String())
			}
		}
		numbersAsStr = newNumbers

	}

	fmt.Println(len(numbersAsStr))
}
