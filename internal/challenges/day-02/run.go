package day02

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 02...")

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

			for j := 1; j <= midpoint; j++ {
				if !partTwo {
					if len(str)%2 != 0 {
						continue
					}
					j = midpoint
				}

				if len(str)%j != 0 {
					continue
				}

				splitStr := SplitSubN(str, j)
				slices.Sort(splitStr)
				splitStr = slices.Compact(splitStr)

				if len(splitStr) == 1 {
					invalidIdSum += i
					break
				}
			}
		}
	}

	return invalidIdSum
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
