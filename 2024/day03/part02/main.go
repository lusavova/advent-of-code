package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input2.txt")
	inputData := string(data)

	var filteredData strings.Builder
	for {
		firstIndexDont := strings.Index(inputData, "don't()")
		if firstIndexDont == -1 {
			filteredData.WriteString(inputData)
			break
		}

		substringUpToFirstDont := inputData[0:firstIndexDont]
		filteredData.WriteString(substringUpToFirstDont)
		inputData = inputData[firstIndexDont:]

		firstIndexDo := strings.Index(inputData, "do()")
		inputData = inputData[firstIndexDo:]
	}

	pattern := `mul\(([0-9]+),([0-9]+)\)`
	regex, _ := regexp.Compile(pattern)
	matches := regex.FindAllStringSubmatch(filteredData.String(), -1)

	result := 0
	for _, match := range matches {
		firstNum, _ := strconv.Atoi(match[1])
		secondNum, _ := strconv.Atoi(match[2])

		result += firstNum * secondNum
	}

	fmt.Println(result)
}
