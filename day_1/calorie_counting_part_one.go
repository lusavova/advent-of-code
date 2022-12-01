package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	max := 0
	var calories int
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			max = int(math.Max(float64(max), float64(calories)))
			calories = 0
			continue
		}

		numOfCalories, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		calories += numOfCalories
	}

	fmt.Println(max)
}


