package day07

import (
	"fmt"
	"testing"

	"github.com/humzamo/advent-of-code-2025/internal/helpers"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inputFile      string
	partNumber     int
	partTwo        bool
	expectedAnswer int
}

func TestDay07(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, false, 21},
		{"input_test.txt", 2, true, 40},
		{"input.txt", 1, false, 1678},
		{"input.txt", 2, true, 357525737893560},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			input := helpers.LoadStringList(tc.inputFile)
			answer := 0
			if tc.partTwo {
				answer = CalculateAnswerPartTwo(input, tc.partTwo)
			} else {
				answer = CalculateAnswer(input, tc.partTwo)
			}
			assert.Equal(t, tc.expectedAnswer, answer)
		})
	}
}
