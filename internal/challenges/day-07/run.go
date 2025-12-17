package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	input := helpers.LoadStringList("./internal/challenges/day-07/input_test.txt")

	fmt.Println("The answer to part one is:", CalculateAnswerPartOne(input))
	fmt.Println("The answer to part two is:", CalculateAnswerPartTwo(input))
}

const (
	start   = "S"
	beam    = "1"
	spitter = "^"
	space   = "."
)

func ParseInput(input []string) [][]string {
	grid := make([][]string, len(input))

	for i, s := range input {
		split := strings.Split(s, "")
		grid[i] = split
	}

	// set the starting beam
	for i, s := range grid[0] {
		if s == start {
			grid[1][i] = beam
			break
		}
	}

	return grid
}

func CalculateAnswerPartOne(input []string) int {
	count := 0
	grid := ParseInput(input)

	for i := 2; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i-1][j] == beam {
				if grid[i][j] == spitter {
					grid[i][j-1] = beam
					grid[i][j+1] = beam
					count++
				} else {
					grid[i][j] = beam
				}
			}
		}
	}

	return count
}

func CalculateAnswerPartTwo(input []string) int {
	grid := ParseInput(input)

	for i := 2; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// only care if there is a number above - spaces and splitters are ignored
			if !IsNumber(grid[i-1][j]) {
				continue
			}

			curStr, _ := GetStringAndVal(grid, i, j)
			curAboveStr, curAboveInt := GetStringAndVal(grid, i-1, j)
			curLeftStr, curLeftInt := GetStringAndVal(grid, i, j-1)
			curRightStr, curRightInt := GetStringAndVal(grid, i, j+1)
			curAboveLeftStr, curAboveLeftInt := GetStringAndVal(grid, i-1, j-1)

			// if the current symbol is a splitter,
			// adjust the spaces directly to the left and right
			if curStr == spitter {
				if curLeftStr == space {
					grid[i][j-1] = curAboveStr
				} else {
					grid[i][j-1] = strconv.Itoa(curLeftInt + curAboveInt)
				}

				if curRightStr == space {
					grid[i][j+1] = curAboveStr
				} else {
					grid[i][j+1] = strconv.Itoa(curRightInt + curAboveInt)
				}
				continue
			}

			if curLeftStr != spitter || curAboveLeftStr == space {
				grid[i][j] = curAboveStr
			} else if curStr != space {
				grid[i][j] = strconv.Itoa(curAboveInt + curAboveLeftInt)
			} else {
				grid[i][j] = beam
			}
		}
	}

	count := 0
	for _, str := range grid[len(grid)-1] {
		if str == space {
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		count += num
	}
	return count
}

func IsNumber(s string) bool {
	return s != space && s != spitter
}

func GetStringAndVal(grid [][]string, i, j int) (string, int) {
	str := ""
	val := 0

	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return str, val
	}

	str = grid[i][j]

	var err error
	if IsNumber(str) {
		val, err = strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
	}

	return str, val
}
