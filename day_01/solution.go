package main

import (
	"adventofcode/utils"
)

func main() {
	lines, err := utils.GetFileLines("day_01/input.txt")
	if err != nil {
		return
	}
	part01(lines)
	part02(lines)
}
