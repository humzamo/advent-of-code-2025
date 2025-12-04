package day02

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/humzamo/advent-of-code-2025/internal/helpers"
// 	"github.com/stretchr/testify/assert"
// )

// type testCase struct {
// 	inputFile      string
// 	partNumber     int
// 	partTwo        bool
// 	expectedAnswer int
// }

// func TestDay01(t *testing.T) {
// 	cases := []testCase{
// 		{"input_test.txt", 1, false, 3},
// 		{"input_test.txt", 2, true, 6},
// 		{"input.txt", 1, false, 1034},
// 		{"input.txt", 2, true, 6166},
// 	}

// 	for _, tc := range cases {
// 		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
// 			input := helpers.LoadParsedList[instruction](tc.inputFile)
// 			answer := CalculateZeros(input, tc.partTwo)

// 			assert.Equal(t, tc.expectedAnswer, answer)
// 		})
// 	}
// }
