package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tuple struct {
	number string
	depth  int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	numbersAsStr := strings.Split(string(data), " ")

	result := 0
	cache := map[tuple]int{}
	for _, number := range numbersAsStr {
		result += blink(number, 25, cache)
	}

	fmt.Println(result)
}

func blink(number string, depth int, cache map[tuple]int) int {
	if value, ok := cache[tuple{number, depth}]; ok {
		return value
	}
	if number == "" {
		return 0
	}
	if depth == 0 {
		return 1
	}

	leftPart := ""
	rightPart := ""
	if number == "0" {
		leftPart = "1"
	} else if len(number)%2 == 0 {
		leftPart = number[0 : len(number)/2]
		rVal, _ := strconv.Atoi(number[len(number)/2:])
		rightPart = strconv.Itoa(rVal)
	} else {
		val, _ := strconv.Atoi(number)
		leftPart = strconv.Itoa(val * 2024)
	}

	cache[tuple{number, depth}] = blink(leftPart, depth-1, cache) + blink(rightPart, depth-1, cache)
	return cache[tuple{number, depth}]
}
