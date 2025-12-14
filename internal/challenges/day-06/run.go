package day06

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 06...")

	input := helpers.LoadStringList("./internal/challenges/day-06/input.txt")

	fmt.Println("The answer to part one is:", CalculateAnswer(input, false))
	fmt.Println("The answer to part two is:", CalculateAnswer(input, true))
}

const (
	multiply = "*"
	add      = "+"
)

func CalculateAnswer(input []string, partTwo bool) int {
	sum := 0

	operations := strings.Fields(input[len(input)-1])

	numsList := [][]int{}
	if partTwo {
		numsList = NumbersListPartTwo(input[:len(input)-1])
	} else {
		numsList = NumbersListPartOne(input[:len(input)-1])
	}

	for i, op := range operations {
		res := 1
		switch op {
		case multiply:
			for _, n := range numsList[i] {
				res = res * n
			}
		case add:
			for _, n := range numsList[i] {
				res = res + n
			}
			res--
		}

		sum += res
	}

	return sum
}

func NumbersListPartOne(input []string) [][]int {
	numbersList := [][]int{}

	for _, s := range input {
		split := strings.Fields(s)

		nums := []int{}
		for _, n := range split {
			val, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, val)
		}

		numbersList = append(numbersList, nums)
	}

	return TransposeMatrix(numbersList)
}

func NumbersListPartTwo(input []string) [][]int {
	rowLen := len(input[0])
	numsList := [][]int{}
	section := []int{}

	for i := 0; i < rowLen; i++ {
		// Work out the current number from top to bottom
		res := ""
		for j := 0; j < len(input); j++ {
			res += string(input[j][i])
		}
		res = strings.TrimSpace(res)

		// If there is a blank line (no number)
		// we have all the number for this section.
		// Save them to the list and continue.
		if res == "" {
			numsList = append(numsList, section)
			section = []int{}
			continue
		}

		// If there is a valid number, save to section.
		converted, err := strconv.Atoi(res)
		if err != nil {
			log.Fatal(err)
		}
		section = append(section, converted)

		// For the last column, there is blank line at the end.
		// Save to list here instead.
		if i == rowLen-1 {
			numsList = append(numsList, section)
		}
	}
	return numsList
}

func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
