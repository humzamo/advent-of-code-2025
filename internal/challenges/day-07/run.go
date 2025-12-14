package day07

import (
	"fmt"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	input := helpers.LoadStringList("./internal/challenges/day-07/input.txt")

	fmt.Println("The answer to part one is:", CalculateAnswer(input, false))
}

const (
	start   = "S"
	beam    = "|"
	spitter = "^"
)

func ParseInput(input []string) [][]string {
	output := make([][]string, len(input))

	for i, s := range input {
		split := strings.Split(s, "")
		output[i] = split
	}

	return output
}

func CalculateAnswer(input []string, partTwo bool) int {
	count := 0
	grid := ParseInput(input)

	// set the starting beam
	for i, s := range grid[0] {
		if s == start {
			grid[1][i] = beam
			break
		}
	}

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
	// for i := 0; i < len(grid); i++ {
	// 	fmt.Println(grid[i])
	// }
	return count
}
