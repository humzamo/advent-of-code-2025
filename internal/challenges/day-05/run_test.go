package day05

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

func TestDay05(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, false, 3},
		{"input_test.txt", 2, true, 14},
		{"input.txt", 1, false, 652},
		{"input.txt", 2, true, 341753674214273},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			input := helpers.LoadStringList(tc.inputFile)
			validRange, toTest := ParseInput(input)

			answer := 0
			if tc.partTwo {
				answer = CountValid(validRange)
			} else {
				answer = CountFresh(validRange, toTest)
			}

			assert.Equal(t, tc.expectedAnswer, answer)
		})
	}
}
