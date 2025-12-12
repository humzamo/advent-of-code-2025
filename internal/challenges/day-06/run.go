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
	nums, ops := ParseInput(input)

	fmt.Println("The answer to part one is:", CalculateAnswer(nums, ops))
}

const (
	multiply = "*"
	add      = "+"
)

func ParseInput(input []string) ([][]int, []string) {
	numbersList := [][]int{}
	operations := []string{}

	for i, s := range input {
		split := strings.Fields(s)

		if i == len(input)-1 {
			operations = append(operations, split...)
			break
		}

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

	return numbersList, operations
}

func CalculateAnswer(numbersList [][]int, operations []string) int {
	sum := 0

	transposed := TransposeMatrix(numbersList)

	for i, nums := range transposed {
		op := operations[i]
		res := 1
		if op == multiply {
			for _, n := range nums {
				res = res * n
			}
		}
		if op == add {
			for _, n := range nums {
				res = res + n
			}
			res--
		}
		sum += res
	}

	return sum
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
