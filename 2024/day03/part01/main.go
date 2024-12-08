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

	pattern := `mul\(([0-9]+),([0-9]+)\)`
	regex, _ := regexp.Compile(pattern)

	result := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			firstNum, _ := strconv.Atoi(match[1])
			secondNum, _ := strconv.Atoi(match[2])
			result += firstNum * secondNum
		}
	}

	fmt.Println(result)
}
