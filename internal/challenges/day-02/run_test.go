package day02

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

func TestDay02(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, false, 1227775554},
		{"input_test.txt", 2, true, 4174379265},
		{"input.txt", 1, false, 24747430309},
		{"input.txt", 2, true, 30962646823},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			inputStr := helpers.LoadStringList(tc.inputFile)[0]
			input := ParseInput(inputStr)

			answer := CalculateInvalidIds(input, tc.partTwo)

			assert.Equal(t, tc.expectedAnswer, answer)
		})
	}
}
