package day3

import (
	"regexp"
	"strconv"
)


func calculate(input string) (int, error) {
	r := regexp.MustCompile(`mul\(((\d+),(\d+))\)`)
	result := 0

	for _, matches := range r.FindAllStringSubmatch(input, -1) {
		v1, v1Err := strconv.Atoi(matches[2])

		if v1Err != nil {
			return 0, v1Err
		}

		v2, v2Err := strconv.Atoi(matches[3])

		if v2Err != nil {
			return 0, v2Err
		}

		result += v1 * v2
	}

	return result, nil
}

func Part1(input string) (int, error) {
  return calculate(input)
}

func Part2(input string) (int, error) {
	dontR := regexp.MustCompile(`don't\(\)`)
	doR := regexp.MustCompile(`do\(\)`)
  slicedInput := ""
	cursor := 0

  for cursor < len(input) {
    donTMatchRaw := dontR.FindStringIndex(input[cursor:])

    var dontMatch int
    if len(donTMatchRaw) == 0 {
      dontMatch = len(input)
    } else {
      dontMatch = donTMatchRaw[0] + cursor
    }


    slicedInput = slicedInput + input[cursor:dontMatch]

    doMatchRaw := doR.FindStringIndex(input[dontMatch:])
    if len(doMatchRaw) == 0 {
      break
    }

    doMatch := doMatchRaw[0] + dontMatch
    cursor = doMatch
  }

	return calculate(slicedInput)
}
