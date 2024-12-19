package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	patterns := map[string]bool{}
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			tokens := strings.Split(line, ", ")
			for _, token := range tokens {
				patterns[token] = true
			}
		} else if i > 1 {
			lines = append(lines, line)
		}

		i++
	}

	result := 0
	cache := map[string]int{}
	for _, line := range lines {
		result += helper(line, patterns, cache)
	}

	fmt.Println(result)
}

func helper(line string, patterns map[string]bool, cache map[string]int) int {
	if len(line) == 0 {
		return 1
	}

	if _, ok := cache[line]; ok {
		return cache[line]
	}
	result := 0
	for pattern := range patterns {
		if strings.HasPrefix(line, pattern) {
			result += helper(line[len(pattern):], patterns, cache)
		}
	}
	cache[line] = result
	return result
}
