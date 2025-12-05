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
	fmt.Println("The answer to part one is:", CalculateJoltage(input, true))
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

func CalculateJoltage(voltages []Voltages, partTwo bool) int64 {
	var sum int64

	k := 2
	if partTwo {
		k = 12
	}

	for _, v := range voltages {
		res := PickLargestK(v, k)
		joltage := digitSliceToInt64(res)
		sum += joltage
	}

	return sum
}

// Given a slice of digits and a target length K,
// choose the lexicographically largest possible subsequence of length K.
func PickLargestK(v Voltages, k int) Voltages {
	n := len(v)
	result := make(Voltages, 0, k)

	start := 0
	for pick := 0; pick < k; pick++ {
		// the latest index we are allowed to choose from
		// while still having enough digits left
		end := n - (k - pick)

		// find the largest digit in v[start:end+1]
		bestDigit := -1
		bestIndex := start
		for i := start; i <= end; i++ {
			if v[i] > bestDigit {
				bestDigit = v[i]
				bestIndex = i
			}
		}

		// choose it
		result = append(result, bestDigit)

		// next search begins after the chosen digit
		start = bestIndex + 1
	}

	return result
}

func digitSliceToInt64(res Voltages) int64 {
	str := ""
	for _, t := range res {
		str += strconv.Itoa(t)
	}

	joltage, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return joltage
}
