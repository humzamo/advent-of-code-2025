package day03

import (
	"fmt"
	"log"
	"strconv"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 03...")

	input := helpers.LoadParsedList[Voltages]("./internal/challenges/day-03/input.txt")

	fmt.Println("The answer to part one is:", CalculateJoltage(input, false))
}

type Voltages []int

func (v Voltages) New() Voltages {
	return Voltages{}
}

func (v Voltages) Convert(str string) Voltages {
	output := make(Voltages, len(str))

	for i, s := range str {
		digitStr := string(s)
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			log.Fatal(err)
		}
		output[i] = digit
	}

	return output
}

// digits is a helper map to map digits to their positions
type digits map[int][]int

func CalculateJoltage(voltages []Voltages, partTwo bool) int {
	sum := 0

	for _, v := range voltages {
		fmt.Println(v)
		digits := digits{}
		for i, d := range v {
			digits[d] = append(digits[d], i)
		}
		fmt.Println(digits)
		max := 0
		for k, w := range digits {
			if k > max {
				// the max digit should exclude the final digit in the voltage
				if len(w) == 1 && w[0] == len(v)-1 {
					continue
				}
				max = k
			}
		}
		fmt.Println("max:", max)
		firstPositionsOfMax := digits[max][0]
		secondMax := 0

		for i := firstPositionsOfMax + 1; i < len(v); i++ {
			if v[i] > secondMax {
				secondMax = v[i]
			}
		}
		joltage := 10*max + secondMax
		fmt.Println(joltage)
		sum += joltage
	}

	return sum
}
