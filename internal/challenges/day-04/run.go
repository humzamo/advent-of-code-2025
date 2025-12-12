package day04

import (
	"fmt"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 04...")

	input := helpers.LoadParsedList[Row]("./internal/challenges/day-04/input.txt")

	fmt.Println("The answer to part one is:", CalculateRolls(input))
	// not 1444

	// for _, row := range input {
	// 	output := strings.Join(row, "")
	// 	fmt.Println(output)
	// }
}

const (
	roll      = "@"
	validRoll = "x"
)

type Row []string

func (v Row) New() Row {
	return Row{}
}

func (v Row) Convert(str string) Row {
	return strings.Split(str, "")
}

func CalculateRolls(rows []Row) int {
	res := 0

	for i, row := range rows {
		for j, val := range row {
			if val != roll && val != validRoll {
				continue
			}

			count := 0

			// left
			if j > 0 && (row[j-1] == roll || row[j-1] == validRoll) {
				count++
			}
			// right
			if j < len(row)-1 && (row[j+1] == roll || row[j+1] == validRoll) {
				count++
			}
			// above
			if i > 0 && (rows[i-1][j] == roll || rows[i-1][j] == validRoll) {
				count++
			}
			// below
			if i < len(rows)-1 && (rows[i+1][j] == roll || rows[i+1][j] == validRoll) {
				count++
			}
			// top left
			if i > 0 && j > 0 && (rows[i-1][j-1] == roll || rows[i-1][j-1] == validRoll) {
				count++
			}
			// top right
			if i > 0 && j < len(row)-1 && (rows[i-1][j+1] == roll || rows[i-1][j+1] == validRoll) {
				count++
			}
			// bottom left
			if i < len(rows)-1 && j > 0 && (rows[i+1][j-1] == roll || rows[i+1][j-1] == validRoll) {
				count++
			}
			// bottom right
			if i < len(rows)-1 && j < len(row)-1 && (rows[i+1][j+1] == roll || rows[i+1][j+1] == validRoll) {
				count++
			}
			if count < 4 {
				res++
				rows[i][j] = validRoll
			}
		}
	}

	return res
}
