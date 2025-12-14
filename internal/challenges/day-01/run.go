package day01

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	input := helpers.LoadParsedList[instruction]("./internal/challenges/day-01/input.txt")

	fmt.Println("The answer to part one is:", CalculateZeros(input, false))
	fmt.Println("The answer to part two is:", CalculateZeros(input, true))
}

type direction string

const (
	left  direction = "L"
	right direction = "R"
)

type instruction struct {
	direction direction
	count     int
}

func (i instruction) Convert(str string) instruction {
	countStr := str[1:]

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal(err)
	}
	return instruction{
		direction: direction(str[0]),
		count:     count,
	}
}

func (i instruction) New() instruction {
	return instruction{}
}

func CalculateZeros(instructions []instruction, partTwo bool) int {
	zeroCount := 0
	currentPos := 50

	for _, instruction := range instructions {
		distance := instruction.count

		if partTwo {
			// All full rotations pass zero
			fullRotations := int(math.Abs(math.Floor(float64(distance) / 100)))
			zeroCount += fullRotations
		}

		distance = distance % 100

		switch instruction.direction {
		case left:
			delta := currentPos - distance
			// Partial left rotation passing zero
			if partTwo && delta < 0 && currentPos != 0 {
				zeroCount++
			}
			currentPos = mod(delta, 100)
		case right:
			delta := currentPos + distance
			// Partial right rotation passing zero
			if partTwo && delta > 100 {
				zeroCount++
			}
			currentPos = mod(delta, 100)
		}

		// Landing on zero
		if currentPos == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func mod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}
