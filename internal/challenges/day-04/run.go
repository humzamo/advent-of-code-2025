package day04

import (
	"fmt"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	input := helpers.LoadParsedList[Row]("./internal/challenges/day-04/input.txt")

	fmt.Println("The answer to part one is:", CalculateRolls(input, false))
	fmt.Println("The answer to part two is:", CalculateRolls(input, true))
}

const (
	roll = "@"
)

type Row []string

func (v Row) New() Row {
	return Row{}
}

func (v Row) Convert(str string) Row {
	return strings.Split(str, "")
}

func CalculateRolls(rows []Row, partTwo bool) int {
	sum := 0
	res := 1 // init > 0 to kick off loop

	for res > 0 {
		newRes, newRows := Recursion(rows, partTwo)
		if !partTwo {
			return newRes
		}
		rows = newRows
		res = newRes
		sum += res
	}

	return sum
}

func Recursion(rows []Row, partTwo bool) (int, []Row) {
	res := 0

	for i, row := range rows {
		for j, val := range row {
			if val != roll {
				continue
			}

			count := 0

			// left
			if j > 0 && (row[j-1] == roll) {
				count++
			}
			// right
			if j < len(row)-1 && (row[j+1] == roll) {
				count++
			}
			// above
			if i > 0 && (rows[i-1][j] == roll) {
				count++
			}
			// below
			if i < len(rows)-1 && (rows[i+1][j] == roll) {
				count++
			}
			// top left
			if i > 0 && j > 0 && (rows[i-1][j-1] == roll) {
				count++
			}
			// top right
			if i > 0 && j < len(row)-1 && (rows[i-1][j+1] == roll) {
				count++
			}
			// bottom left
			if i < len(rows)-1 && j > 0 && (rows[i+1][j-1] == roll) {
				count++
			}
			// bottom right
			if i < len(rows)-1 && j < len(row)-1 && (rows[i+1][j+1] == roll) {
				count++
			}
			if count < 4 {
				res++
				if partTwo {
					rows[i][j] = "."
				}
			}
		}
	}

	return res, rows
}
