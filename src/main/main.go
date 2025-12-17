package main

import (
	"fmt"
	"os"
	"strconv"

	day01 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-01"
	day02 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-02"
	day03 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-03"
	day04 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-04"
	day05 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-05"
	day06 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-06"
	day07 "github.com/humzamo/advent-of-code-2025/internal/challenges/day-07"
)

func main() {
	fmt.Print("🎄 Welcome to Advent of Code 2025! 🎄\n\n")

	if len(os.Args) < 2 {
		fmt.Println("Please make sure you enter a day to generate a solution!")
		fmt.Println("Usage: make run day=<day_number>")
		os.Exit(1)
	}

	dayString := os.Args[1]

	dayNumber, err := strconv.Atoi(dayString)
	if err != nil {
		fmt.Printf("invalid day entered! [day=%s]", dayString)
		os.Exit(1)
	}

	if dayNumber > 12 {
		fmt.Println("Sorry there are only 12 days in this challenge!")
		os.Exit(1)
	}

	fmt.Println("Generating solutions for day:", dayNumber)

	switch dayNumber {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	default:
		fmt.Printf("Uh oh, there's no solution for day %s yet... 👀\n", dayString)
		os.Exit(1)
	}
}
