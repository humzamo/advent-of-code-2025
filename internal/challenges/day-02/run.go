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

	inputStr := helpers.LoadStringList("./internal/challenges/day-02/input_test.txt")[0]
	input := ParseInput(inputStr)

	fmt.Println("The answer to part one is:", CalculateInvalidIds(input))
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

func CalculateInvalidIds(ranges []Range) int {
	invalidIdSum := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			// fmt.Println(str)

			midpoint := len(str) / 2
			// fmt.Println("midpoint:", midpoint)

			for j := 1; j <= midpoint; j++ {
				if len(str)%j != 0 {
					continue
				}
				// fmt.Println(str)
				// fmt.Println("j:", j)
				splitStr := SplitSubN(str, j)
				// fmt.Println("split:", splitStr)
				slices.Sort(splitStr)
				splitStr = slices.Compact(splitStr)
				// fmt.Println("compacted:", splitStr)

				if len(splitStr) == 1 {
					// fmt.Println("MATCH:", i)
					invalidIdSum += i
					break
				}
			}

			/* PART 1
			if len(str)%2 != 0 {
				continue
			}
			midpoint := len(str) / 2
			// fmt.Println("midepoint:", midpoint)
			firstHalf := str[0:midpoint]
			secondHalf := str[midpoint:]
			// fmt.Println("first:", firstHalf)

			// fmt.Println("second:", secondHalf)

			// splitStr := strings.SplitAfterN(str, "", midpoint)
			// fmt.Println(splitStr)

			if firstHalf == secondHalf {
				fmt.Println("match for", str)
				// val, err := strconv.Atoi(firstHalf)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				invalidIdSum += i
			}
			*/
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
