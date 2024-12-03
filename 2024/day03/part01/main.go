package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")

	result := 0
	lines := strings.Split(string(data), "\n")
	pattern := `mul\(([0-9]+),([0-9]+)\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) >= 3 {
				firstNum, _ := strconv.Atoi(match[1])
				secondNum, _ := strconv.Atoi(match[2])

				result += firstNum * secondNum
			}
		}
	}

	fmt.Println(result)
}
