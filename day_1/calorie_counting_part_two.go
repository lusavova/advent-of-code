package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumCalories int
	maxThree := []int{0, 0, 0}
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			populateMax(sumCalories, maxThree)
			sumCalories = 0
			continue
		}

		currentCalories, _ := strconv.Atoi(input)
		sumCalories += currentCalories
	}

	fmt.Println(sum(maxThree))
}

func populateMax(num int, max []int) {
	var temp int
	for i := 2; i >= 0; i-- {
		temp = max[i]
		if num >= max[i] {
			max[i] = num
			num = temp
		}
	}
}

func sum(array []int) int {
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}
