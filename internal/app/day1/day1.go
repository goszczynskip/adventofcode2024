package day1

import (
	"fmt"
	"math"
	"regexp"
  "sort"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")

  lines = lines[:len(lines)-1]

	firstList := make([]int, len(lines))
	secondList := make([]int, len(lines))

	regexDelim := regexp.MustCompile(`\s+`)

	for i, line := range lines {
    trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		numbers := regexDelim.Split(line, -1)
		parsed1Int, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("Failed to parse 1: %v. Index: %d, line: %v", numbers[0], i, line)
			return firstList, secondList, err
		}

		parsed2Int, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("Failed to parse 2: %v. Index: %d, line: %v", numbers[0], i, line)
			return firstList, secondList, err
		}

		firstList[i] = parsed1Int
		secondList[i] = parsed2Int
	}

	return firstList, secondList, nil
}

func Part1(input string) (int, error) {
	firstList, secondList, err := parseInput(input)

	if err != nil {
		return 0, err
	}

  sort.Ints(firstList) 
  sort.Ints(secondList) 

	diff := 0

	for i, first := range firstList {
		diff += int(math.Abs(float64(first - secondList[i])))
	}

	return diff, nil
}

func Part2(input string) (int, error) {
	firstList, secondList, err := parseInput(input)

	if err != nil {
		return 0, err
	}

  secondMap := make(map[int]int)

  for _, second := range secondList {
    secondMap[second] += 1
  }

	diff := 0

	for _, first := range firstList {
    times := secondMap[first]
    diff += times * first
  }

	return diff, nil
}
