package main

import (
	"fmt"
	"os"
	"strconv"
)

type file struct {
	startIndex int
	size       int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	diskMap := string(data)

	var disk []int
	fileNumber := 0
	var files []file
	for position, item := range diskMap {
		size, _ := strconv.Atoi(string(item))
		if isFile(position) {
			startIndex := len(disk)
			for i := 0; i < size; i++ {
				disk = append(disk, fileNumber)
			}
			files = append(files, file{startIndex: startIndex, size: size})
			fileNumber++
		} else {
			for i := 0; i < size; i++ {
				disk = append(disk, -1)
			}
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		defrag(i, files[i], disk)
	}

	fmt.Println(checksum(disk))
}

func checksum(numbers []int) int64 {
	sum := int64(0)
	for i, num := range numbers {
		if num == -1 {
			continue
		}

		sum += int64(num) * int64(i)
	}
	return sum
}

func defrag(fileNumber int, file file, numbers []int) {
	pointer := 0
	for pointer < file.startIndex {
		// move until it's a dot
		for pointer < file.startIndex && numbers[pointer] != -1 {
			pointer++
		}

		if pointer == file.startIndex {
			break
		}

		cntDots := 0
		startDots := pointer
		for pointer < file.startIndex && numbers[pointer] == -1 {
			cntDots++
			pointer++
		}

		if cntDots >= file.size {
			moveFile(fileNumber, file, numbers, startDots)
			break
		}
	}
}

func moveFile(fileNumber int, file file, numbers []int, startDots int) {
	for i := 0; i < file.size; i++ {
		numbers[startDots+i] = fileNumber
		numbers[file.startIndex+i] = -1
	}
}

func isFile(i int) bool {
	return i%2 == 0
}
