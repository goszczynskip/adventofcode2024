package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")

	lines = lines[:len(lines)-1]

	positions := make([][]int, len(lines))

	regexDelim := regexp.MustCompile(`\s+`)

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		numbers := regexDelim.Split(line, -1)

		for _, number := range numbers {
			parsedInt, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Failed to parse: %v. Index: %d, line: %v", number, i, line)
				return nil, err
			}

			positions[i] = append(positions[i], parsedInt)
		}

	}

	return positions, nil
}

func Part1(input string) (int, error) {
	positions, err := parseInput(input)

	safeRows := 0

	for _, items := range positions {
		isSafe := checkSafety(items)

		if isSafe {
			safeRows++
		}
	}

	return safeRows, err
}

func checkSafety(items []int) bool {
	delta := 0
	var prevItemB int
	itemB := items[0]
	unsafe := -1

	for j, item := range items[1:] {
		if unsafe > -1 {
			break
		}
		prevItemB = itemB
		itemB = item
		newDelta := prevItemB - itemB

		if newDelta == 0 || newDelta > 0 && delta < 0 || newDelta < 0 && delta > 0 || newDelta > 3 || newDelta < -3 {
			unsafe = j
		}

		delta = newDelta
	}

	return unsafe == -1
}

func removeIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func Part2(input string) (int, error) {
	positions, err := parseInput(input)

	safeRows := 0

	for _, items := range positions {
		isSafe := checkSafety(items)

		if isSafe {
			safeRows++
			continue
		}

		for nToRemove := range items {
      itemsCpy:= removeIndex(items, nToRemove) 

			isSafe = checkSafety(itemsCpy)

			if isSafe {
				safeRows++
				break
			}
		}
	}

	return safeRows, err
}
