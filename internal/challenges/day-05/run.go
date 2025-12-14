package day05

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
)

func Run() {
	input := helpers.LoadStringList("./internal/challenges/day-05/input.txt")
	validRange, toTest := ParseInput(input)

	fmt.Println("The answer to part one is:", CountFresh(validRange, toTest))
	fmt.Println("The answer to part two is:", CountValid(validRange))
}

type Range struct {
	Min int
	Max int
}

func ParseInput(input []string) ([]Range, []int) {
	validRanges := []Range{}
	toTest := []int{}

	parseRanges := true
	for _, s := range input {
		if s == "" {
			parseRanges = false
			continue
		}
		if parseRanges {
			start, end := GetStartEnd(s)
			validRanges = append(validRanges,
				Range{
					Min: start,
					Max: end,
				})
		} else {
			val, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			toTest = append(toTest, val)
		}
	}

	compacted := mergeRanges(validRanges)

	return compacted, toTest
}

func GetStartEnd(str string) (int, int) {
	splitStr := strings.Split(str, "-")

	start, err := strconv.Atoi(splitStr[0])
	if err != nil {
		log.Fatal(err)
	}

	end, err := strconv.Atoi(splitStr[1])
	if err != nil {
		log.Fatal(err)
	}

	return start, end
}

func CountFresh(validRanges []Range, toTest []int) int {
	count := 0

	for _, t := range toTest {
		for _, v := range validRanges {
			if t >= v.Min && t <= v.Max {
				count++
				break
			}
		}
	}
	return count
}

func CountValid(ranges []Range) int {
	sum := 0

	for _, r := range ranges {
		sum += r.Max - r.Min + 1
	}

	return sum
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	merged := []Range{}
	current := ranges[0]
	merged = append(merged, current)

	for _, r := range ranges[1:] {
		if r.Min <= current.Max { // Overlapping or touching
			if r.Max > current.Max {
				current.Max = r.Max
				merged[len(merged)-1] = current // update last merged
			}
		} else {
			current = r
			merged = append(merged, current)
		}
	}

	return merged
}
