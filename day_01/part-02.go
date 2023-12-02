package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func part02(lines []string) {

	for i, line := range lines {
		line = processLine(line)
		lines[i] = line
	}

	result := sum(lines)
	fmt.Println("Part 02: ", result)
}

func processLine(line string) string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	line = processFirstOccurrence(line, numbers, numbersMap)
	line = processLastOccurrence(line, numbers, numbersMap)

	return line
}

func processFirstOccurrence(line string, numbers []string, numbersMap map[string]int) string {
	minIndex := math.MaxInt32
	minIndexNumber := ""

	for _, number := range numbers {
		index := strings.Index(line, number)
		if index != -1 && index < minIndex {
			minIndex = index
			minIndexNumber = number
		}
	}

	if minIndexNumber != "" {
		line = replaceFirstOccurrence(line, minIndexNumber, strconv.Itoa(numbersMap[minIndexNumber])+minIndexNumber)
	}

	return line
}

func replaceFirstOccurrence(input, old, new string) string {
	firstIndex := strings.Index(input, old)
	if firstIndex == -1 {
		return input
	}

	result := input[:firstIndex] + new + input[firstIndex+len(old):]
	return result
}

func processLastOccurrence(line string, numbers []string, numbersMap map[string]int) string {
	maxIndex := -1
	maxIndexNumber := ""

	for _, number := range numbers {
		index := strings.LastIndex(line, number)
		if index != -1 && index > maxIndex {
			maxIndex = index
			maxIndexNumber = number
		}
	}

	if maxIndexNumber != "" {
		line = replaceLastOccurrence(line, maxIndexNumber, maxIndexNumber+strconv.Itoa(numbersMap[maxIndexNumber]))
	}

	return line
}

func sum(lines []string) int {
	var result int

	for _, v := range lines {
		first, second := extractFirstAndLastDigits(v)
		number, err := strconv.Atoi(first + second)
		if err != nil {
			return 0
		}
		result += number
	}

	return result
}

func extractFirstAndLastDigits(s string) (first, second string) {
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			first = string(s[i])
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			second = string(s[i])
			break
		}
	}

	return first, second
}

func replaceLastOccurrence(input, old, new string) string {
	lastIndex := strings.LastIndex(input, old)
	if lastIndex == -1 {
		return input
	}

	result := input[:lastIndex] + new + input[lastIndex+len(old):]
	return result
}
