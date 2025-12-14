package day06

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

func TestDay06(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, false, 4277556},
		{"input_test.txt", 2, true, 3263827},
		{"input.txt", 1, false, 5381996914800},
		{"input.txt", 2, true, 9627174150897},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			input := helpers.LoadStringList(tc.inputFile)
			answer := CalculateAnswer(input, tc.partTwo)
			assert.Equal(t, tc.expectedAnswer, answer)
		})
	}
}
