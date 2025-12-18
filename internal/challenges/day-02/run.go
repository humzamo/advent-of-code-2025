package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	inputStr := helpers.LoadStringList("./internal/challenges/day-02/input.txt")[0]
	input := ParseInput(inputStr)

	fmt.Println("The answer to part one is:", CalculateInvalidIds(input, false))
	fmt.Println("The answer to part two is:", CalculateInvalidIds(input, true))
}

type Range struct {
	start int
	end   int
}

func ParseInput(input string) []Range {
	ranges := strings.Split(input, ",")
	output := make([]Range, len(ranges))

	for i, r := range ranges {
		splitStr := strings.Split(r, "-")
		start, err := strconv.Atoi(splitStr[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(splitStr[1])
		if err != nil {
			log.Fatal(err)
		}
		output[i] = Range{
			start: start,
			end:   end,
		}
	}
	return output
}

func CalculateInvalidIds(ranges []Range, partTwo bool) int {
	invalidIdSum := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			midpoint := len(str) / 2

			if ok := hasRepeatedSequenceKMP(str, midpoint, partTwo); ok {
				invalidIdSum += i
			}
		}
	}

	return invalidIdSum
}

func hasRepeatedSequenceKMP(s string, maxUnit int, partTwo bool) bool {
	n := len(s)
	lps := make([]int, n)
	half := n / 2

	if !partTwo {
		if n < 2 || n%2 != 0 {
			return false
		}

		if half > maxUnit {
			return false
		}
	}

	for i, j := 1, 0; i < n; {
		if s[i] == s[j] {
			j++
			lps[i] = j
			i++
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	unit := n - lps[n-1]

	if !partTwo {
		p := n - lps[n-1]
		return half%p == 0
	}

	return lps[n-1] > 0 &&
		n%unit == 0 &&
		unit <= maxUnit &&
		n/unit >= 2
}
