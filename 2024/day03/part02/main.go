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
	line := string(data)

	var finalString strings.Builder
	for {
		firstIndexDont := strings.Index(line, "don't()")
		if firstIndexDont == -1 {
			finalString.WriteString(line)
			break
		}
		substring := line[0:firstIndexDont]
		finalString.WriteString(substring)
		line = line[firstIndexDont:]
		firstIndexDo := strings.Index(line, "do()")
		line = line[firstIndexDo:]
	}

	line = finalString.String()

	pattern := `mul\(([0-9]+),([0-9]+)\)`
	re, _ := regexp.Compile(pattern)
	result := 0
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			firstNum, _ := strconv.Atoi(match[1])
			secondNum, _ := strconv.Atoi(match[2])

			result += firstNum * secondNum
		}
	}

	fmt.Println(result)
}
