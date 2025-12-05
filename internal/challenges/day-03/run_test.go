package day03

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inputFile      string
	partNumber     int
	partTwo        bool
	expectedAnswer int64
}

func TestDay01(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, false, 357},
		{"input_test.txt", 2, true, 3121910778619},
		{"input.txt", 1, false, 17445},
		{"input.txt", 2, true, 173229689350551},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			input := helpers.LoadParsedList[Voltages](tc.inputFile)
			var answer int64
			if tc.partNumber == 1 {
				answer = int64(CalculateJoltage(input))
			} else {
				str := CalculateJoltagePartTwo(input)
				parsed, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					t.Fail()
				}
				answer = parsed
			}

			assert.Equal(t, tc.expectedAnswer, answer)
		})
	}
}
