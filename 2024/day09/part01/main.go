package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	diskMap := string(data)

	var disk []int
	fileNumber := 0
	for position, item := range diskMap {
		size, _ := strconv.Atoi(string(item))
		if isFile(position) {
			for i := 0; i < size; i++ {
				disk = append(disk, fileNumber)
			}
			fileNumber++
		} else {
			for i := 0; i < size; i++ {
				disk = append(disk, -1)
			}
		}
	}

	leftPointer := 0
	rightPointer := len(disk) - 1
	for {
		// move until it's a dot
		for leftPointer < len(disk) && disk[leftPointer] != -1 {
			leftPointer++
		}

		// move until it's a number
		for rightPointer >= 0 && disk[rightPointer] == -1 {
			rightPointer--
		}

		if leftPointer >= rightPointer {
			break
		}
		// "swap"
		disk[leftPointer] = disk[rightPointer]
		disk[rightPointer] = -1
	}

	fmt.Println(checksum(disk))
}

func checksum(disk []int) int64 {
	sum := int64(0)
	for i, num := range disk {
		if num == -1 {
			break
		}

		sum += int64(num) * int64(i)
	}

	return sum
}

func isFile(i int) bool {
	return i%2 == 0
}
