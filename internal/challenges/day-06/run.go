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
	numsList := [][]int{}
	temp := []int{}
	for i := 0; i < len(input[0]); i++ {
		res := ""
		for j := 0; j < len(input); j++ {
			res += string(input[j][i])
		}
		res = strings.TrimSpace(res)
		if res == "" {
			numsList = append(numsList, temp)
			temp = []int{}
			continue
		}

		converted, err := strconv.Atoi(res)
		if err != nil {
			log.Fatal(err)
		}
		temp = append(temp, converted)

		if i == len(input[0])-1 {
			numsList = append(numsList, temp)
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
