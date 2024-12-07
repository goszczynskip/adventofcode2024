package app

import (
	"adventofcode/internal/app/day1"
	"adventofcode/internal/app/day2"
	"adventofcode/internal/app/day3"
	"adventofcode/internal/app/day4"
	"adventofcode/internal/app/day5"
	"adventofcode/internal/app/day6"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}

	return string(content), nil
}

type Solution struct {
	part1 func(string) (int, error)
	part2 func(string) (int, error)
}

var solutions = []Solution{
	{day1.Part1, day1.Part2},
	{day2.Part1, day2.Part2},
	{day3.Part1, day3.Part2},
	{day4.Part1, day4.Part2},
	{day5.Part1, day5.Part2},
	{day6.Part1, day6.Part2},
}

// Run executes the main application logic
func Run(day string, test bool) error {
	fmt.Printf("Runs advent of code day: %s \n", day)

	var value int

	dayParts := strings.Split(day, "_")
	dayNumber, err := strconv.Atoi(dayParts[0])
	if err != nil {
		return err
	}

	partNumber, err := strconv.Atoi(dayParts[1])
	if err != nil {
		return err
	}

	var input string

	if test {
		input, err = loadFile(fmt.Sprintf("assets/day%d/test_input.txt", dayNumber))
	} else {
		input, err = loadFile(fmt.Sprintf("assets/day%d/input.txt", dayNumber))
	}

	if err != nil {
		return err
	}

	if partNumber == 1 {
		value, err = solutions[dayNumber-1].part1(input)
	} else {
		value, err = solutions[dayNumber-1].part2(input)
	}

	if err != nil {
		return err
	}

	fmt.Printf("Solution: %d", value)
	fmt.Println("")

	return nil
}
