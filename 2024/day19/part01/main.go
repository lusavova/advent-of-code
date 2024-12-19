package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	cnt := 0
	var patternString string
	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			tokens := strings.Split(line, ", ")
			sort.Slice(tokens, func(i, j int) bool {
				return len(tokens[i]) > len(tokens[j])
			})
			patternString = fmt.Sprintf(`^(%s)*$`, strings.Join(tokens, "|"))
			fmt.Println(patternString)
		} else if i == 1 {
			i++
			continue
		} else {
			pattern := regexp.MustCompile(patternString)
			if pattern.MatchString(line) {
				cnt++
			}
		}

		i++

	}

	fmt.Println(cnt)
}
