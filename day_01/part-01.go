package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func part01(lines []string) {
	var result int
	for _, value := range lines {

		var first string
		for i := 0; i < len(value); i++ {
			if unicode.IsDigit(rune(value[i])) {
				first = string(value[i])
				break
			}
		}

		var second string
		for i := len(value) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(value[i])) {
				second = string(value[i])
				break
			}
		}

		number, err := strconv.Atoi(first + second)
		if err != nil {
			return
		}

		result = result + number
	}

	fmt.Println("Part 01: ", result)
}
