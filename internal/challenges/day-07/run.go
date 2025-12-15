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

	fmt.Println("The answer to part one is:", CalculateAnswerPartTwo(input, false))
}

const (
	start   = "S"
	beam    = "|"
	spitter = "^"
	space   = "."
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

func CalculateAnswerPartTwo(input []string, partTwo bool) int {
	grid := ParseInput(input)

	// set the starting beam
	for i, s := range grid[0] {
		if s == start {
			grid[1][i] = "1"
			break
		}
	}

	for i := 2; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// only care if there is a number above - spaces and splitters are ignored
			if IsNumber(grid[i-1][j]) {

				var curAboveInt, curLeftInt, curRightInt, curAboveLeftInt int
				var curAboveStr, curLeftStr, curRightStr, curAboveLeftStr string

				var err error
				curAboveStr = grid[i-1][j]
				if IsNumber(curAboveStr) {
					curAboveInt, err = strconv.Atoi(curAboveStr)
					if err != nil {
						log.Fatal(err)
					}
				}

				if j > 0 {
					curLeftStr = grid[i][j-1]
					if IsNumber(curLeftStr) {
						curLeftInt, err = strconv.Atoi(curLeftStr)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

				if j < len(grid[0])-1 {
					curRightStr = grid[i][j+1]
					if IsNumber(curRightStr) {
						curRightInt, err = strconv.Atoi(curRightStr)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

				if j > 0 {
					curAboveLeftStr = grid[i-1][j-1]
					if IsNumber(curAboveLeftStr) {
						curAboveLeftInt, err = strconv.Atoi(curAboveLeftStr)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

				// if the current space is a splitter
				if grid[i][j] == spitter {
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
				} else {
					if j > 0 && grid[i][j-1] != spitter {
						grid[i][j] = grid[i-1][j]
					} else if grid[i][j] != space {
						grid[i][j] = strconv.Itoa(curAboveInt + curAboveLeftInt)
					} else if j > 0 && grid[i-1][j-1] == space {
						grid[i][j] = curAboveStr
					} else {
						grid[i][j] = "1"
					}
				}
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
