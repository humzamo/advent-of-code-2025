package day03

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 03...")

	input := helpers.LoadParsedList[Voltages]("./internal/challenges/day-03/input.txt")

	fmt.Println("The answer to part one is:", CalculateJoltage(input))
	fmt.Println("The answer to part one is:", CalculateJoltagePartTwo(input))

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

func CalculateJoltage(voltages []Voltages) int64 {
	var sum int64

	for _, v := range voltages {
		max := 0
		firstPositionsOfMax := 0
		for i := 0; i < len(v)-1; i++ {
			if v[i] > max {
				max = v[i]
				firstPositionsOfMax = i
			}
		}

		secondMax := 0

		for i := firstPositionsOfMax + 1; i < len(v); i++ {
			if v[i] > secondMax {
				secondMax = v[i]
			}
		}
		joltage := 10*max + secondMax
		sum += int64(joltage)
	}

	return sum
}

func CalculateJoltagePartTwo(voltages []Voltages) int64 {
	sum := &big.Int{}
	iterations := len(voltages[0]) - 12

	for _, v := range voltages {
		max := big.Int{}
		index := 0

		for i := 0; i < iterations; i++ {
			max, index = FindMax(v)
			v = removeDigitFromArray(index, v)
		}
		sum = sum.Add(sum, &max)
	}

	return sum.Int64()
}

// FindMax removes one digit from the voltage to get the max
// It returns the max and the index of the digit to remove
func FindMax(v Voltages) (big.Int, int) {
	index := 0
	currentMax := big.Int{}
	for i := 0; i < len(v); i++ {
		tempArr := removeDigitFromArray(i, v)

		str := ""
		for _, t := range tempArr {
			str += strconv.Itoa(t)
		}

		tempVolt := ParseBigInt(str)
		if tempVolt.Cmp(&currentMax) == 1 {
			currentMax = tempVolt
			index = i
		}
	}

	return currentMax, index
}

func removeDigitFromArray(index int, v Voltages) Voltages {
	tempArr := Voltages{}
	if index == 0 {
		tempArr = append(tempArr, v[1:]...)
	} else if index == len(v)-1 {
		tempArr = append(tempArr, v[:len(v)-1]...)
	} else {
		tempArr = append(tempArr, v[:index]...)
		tempArr = append(tempArr, v[index+1:]...)
	}
	return tempArr
}

func ParseBigInt(str string) big.Int {
	n := new(big.Int)
	n, ok := n.SetString(str, 10)
	if !ok {
		log.Fatal("unable to parse big int")
	}
	return *n
}
