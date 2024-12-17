package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	registry := []int{}
	program := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Register A: ") {
			number, _ := strconv.Atoi(line[len("Register A: "):])
			registry = append(registry, number)
		} else if strings.HasPrefix(line, "Register B: ") {
			number, _ := strconv.Atoi(line[len("Register B: "):])
			registry = append(registry, number)
		} else if strings.HasPrefix(line, "Register C: ") {
			number, _ := strconv.Atoi(line[len("Register C: "):])
			registry = append(registry, number)
		} else if strings.HasPrefix(line, "Program: ") {
			numbers := strings.Split(line[len("Program: "):], ",")
			for _, n := range numbers {
				number, _ := strconv.Atoi(n)
				program = append(program, number)
			}
		}
	}

	var outputs []int
	instructionPointer := 0
	for instructionPointer < len(program) {
		instruction := program[instructionPointer]

		if instructionPointer+1 >= len(program) {
			break
		}

		operand := program[instructionPointer+1]
		operandValue := getOperandValue(operand, registry)

		switch instruction {
		case 0:
			registry[0] = registry[0] / int(math.Pow(2, float64(operandValue)))
		case 1:
			registry[1] = registry[1] ^ operand // maybe operand value here???
		case 2:
			registry[1] = operandValue % 8
		case 3:
			if registry[0] != 0 {
				instructionPointer = operand
				continue
			}
		case 4:
			registry[1] = registry[1] ^ registry[2]
		case 5:
			result := operandValue % 8
			outputs = append(outputs, result)
		case 6:
			registry[1] = registry[0] / int(math.Pow(2, float64(operandValue)))
		case 7:
			registry[2] = registry[0] / int(math.Pow(2, float64(operandValue)))
		default:
			panic(fmt.Sprintf("Unknown instruction: %d", instruction))
		}
		instructionPointer += 2
	}
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(outputs)), ","), "[]"))

}

func getOperandValue(operand int, registers []int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	}
	panic("operand err")
}
