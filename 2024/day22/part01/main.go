package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	sum := 0
	for _, num := range nums {
		for range 2000 {
			num = getSecretNum(num)
		}
		sum += num
	}

	fmt.Println(sum)
}

func getSecretNum(num int) int {
	num = prune(mix(num*64, num))
	num = prune(mix(num/32, num))
	num = prune(mix(num*2048, num))
	return num
}

func mix(a, b int) int { return a ^ b }
func prune(a int) int  { return a % 16777216 }
