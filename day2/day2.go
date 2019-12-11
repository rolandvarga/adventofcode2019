package main

import (
	"errors"
	"fmt"
)

var (
	masterInput = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 2, 23, 6, 27, 2, 6, 27, 31, 2, 13, 31, 35, 1, 10, 35, 39, 2, 39, 13, 43, 1, 43, 13, 47, 1, 6, 47, 51, 1, 10, 51, 55, 2, 55, 6, 59, 1, 5, 59, 63, 2, 9, 63, 67, 1, 6, 67, 71, 2, 9, 71, 75, 1, 6, 75, 79, 2, 79, 13, 83, 1, 83, 10, 87, 1, 13, 87, 91, 1, 91, 10, 95, 2, 9, 95, 99, 1, 5, 99, 103, 2, 10, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0}
	dayTwoInput = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 2, 23, 6, 27, 2, 6, 27, 31, 2, 13, 31, 35, 1, 10, 35, 39, 2, 39, 13, 43, 1, 43, 13, 47, 1, 6, 47, 51, 1, 10, 51, 55, 2, 55, 6, 59, 1, 5, 59, 63, 2, 9, 63, 67, 1, 6, 67, 71, 2, 9, 71, 75, 1, 6, 75, 79, 2, 79, 13, 83, 1, 83, 10, 87, 1, 13, 87, 91, 1, 91, 10, 95, 2, 9, 95, 99, 1, 5, 99, 103, 2, 10, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0}
)

func PartOne() int {
	dayTwoInput[1] = 12
	dayTwoInput[2] = 2
	for i := 0; i <= len(dayTwoInput); i = i + 4 {
		opcode := dayTwoInput[i]
		if opcode == 99 {
			break
		}
		a := dayTwoInput[i+1]
		b := dayTwoInput[i+2]
		c := dayTwoInput[i+3]
		switch opcode {
		case 1:
			dayTwoInput[c] = dayTwoInput[a] + dayTwoInput[b]
		case 2:
			dayTwoInput[c] = dayTwoInput[a] * dayTwoInput[b]
		}
	}
	return dayTwoInput[0]
}

func runPair(input []int) error {
	for i := 0; i <= len(input); i = i + 4 {
		paramsA := input[i+1]
		paramsB := input[i+2]
		paramsC := input[i+3]

		if paramsA >= len(input) || paramsB >= len(input) || paramsC >= len(input) {
			return errors.New("index out of range")
		}

		opcode := input[i]
		switch opcode {
		case 1:
			input[paramsC] = input[paramsA] + input[paramsB]
		case 2:
			input[paramsC] = input[paramsA] * input[paramsB]
		case 99:
			return nil
		}
	}
	return nil
}

func PartTwo() int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copyInput := make([]int, len(masterInput))
			copy(copyInput, masterInput)

			copyInput[1] = noun
			copyInput[2] = verb

			err := runPair(copyInput)
			if err != nil {
				continue
			}

			if copyInput[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func main() {
	fmt.Printf("PartOne solution: %d\n", PartOne())
	fmt.Printf("PartTwo solution: %d\n", PartTwo())
}
