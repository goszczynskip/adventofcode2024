package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	printStatements [][]int
	ordering        [][2]int
}

type Node struct {
	next     *Node
	previous *Node
	value    int
}

func parseInput(input string) (Input, error) {
	rawLines := strings.Split(input, "\n")
	result := Input{
		printStatements: make([][]int, 0),
		ordering:        make([][2]int, 0),
	}

	orderingPhase := true
	for i, line := range rawLines {
		if len(line) == 0 {
			orderingPhase = false
			continue
		}

		if orderingPhase {
			orderingParts := strings.Split(line, "|")
			if len(orderingParts) != 2 {
				return Input{}, fmt.Errorf("Invalid ordering line: %s (%d)", line, i)
			}

			int1, err := strconv.Atoi(orderingParts[0])
			if err != nil {
				return Input{}, err
			}

			int2, err := strconv.Atoi(orderingParts[1])
			if err != nil {
				return Input{}, err
			}

			result.ordering = append(result.ordering, [2]int{int1, int2})
		} else {
			printParts := strings.Split(line, ",")
			convertedPrintParts := make([]int, len(printParts))
			for i, part := range printParts {
				convertedPart, err := strconv.Atoi(part)
				if err != nil {
					return Input{}, err
				}
				convertedPrintParts[i] = convertedPart
			}

			result.printStatements = append(result.printStatements, convertedPrintParts)
		}
	}

	return result, nil
}

func buildOrderingGraph(orderings [][2]int) *Node {
	return nil
}

func selectRelevantOrderings(orderings [][2]int, printStatement []int, cache map[int][][2]int) [][2]int {
	result := make([][2]int, 0)
	for _, page := range printStatement {
		if cache[page] != nil {
			result = append(result, cache[page]...)
			continue
		}

		relevantOrderings := make([][2]int, 0)
		for _, ordering := range orderings {
			if ordering[0] == page || ordering[1] == page {
				relevantOrderings = append(relevantOrderings, ordering)
			}
		}
		cache[page] = relevantOrderings

		result = append(result, relevantOrderings...)
	}
	return result
}

func isPageCorrect(page int, after map[int][]int, relevantOrderings [][2]int) bool {
	for _, ordering := range relevantOrderings {
		if ordering[1] == page {
			for _, pageAfter := range after[page] {
				if ordering[0] == pageAfter {
					return false
				}
			}
		}
	}

	return true
}

func isPrintStatementCorrect(printStatement []int, relevantOrderings [][2]int) bool {
	after := make(map[int][]int)

	for i, page := range printStatement {
		after[page] = printStatement[i+1:]
		if !isPageCorrect(page, after, relevantOrderings) {
			return false
		}
	}

	return true
}

func Part1(input string) (int, error) {
	parsedInput, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	relevantOrderingsCache := make(map[int][][2]int)

	correctMiddlePoints := make([]int, 0)

	for _, printStatement := range parsedInput.printStatements {
		relevantOrderings := selectRelevantOrderings(parsedInput.ordering, printStatement, relevantOrderingsCache)
		if isPrintStatementCorrect(printStatement, relevantOrderings) {
			correctMiddlePoints = append(
				correctMiddlePoints,
				printStatement[(len(printStatement)/2):(len(printStatement)/2+1)]...,
			)
		}
	}

	middlePointsSum := 0

	for _, middlePoint := range correctMiddlePoints {
		middlePointsSum += middlePoint
	}

	return middlePointsSum, nil
}

func correctPrintStatement(printStatement []int, relevantOrderings [][2]int) []int {
	finalOrdering := make([]int, len(printStatement))
	for i := 0; i < len(finalOrdering); i++ {
		finalOrdering[i] = -1
	}
	finalOrdering[0] = printStatement[0]

	for i, page := range printStatement[1:] {
		indexToInsert := i
		stopInserting := false
		for k, finalOrder := range finalOrdering {
			if stopInserting {
				break
			}
			for _, ordering := range relevantOrderings {
				if page == ordering[0] && finalOrder == ordering[1] {
					indexToInsert = k
					stopInserting = true
					break

				} else if page == ordering[1] && finalOrder == ordering[0] {
					indexToInsert = k + 1

					break
				}
			}
		}

		for i := len(finalOrdering) - 1; i > indexToInsert; i-- {
			finalOrdering[i] = finalOrdering[i-1]
		}

		finalOrdering[indexToInsert] = page
	}

	return finalOrdering
}

func Part2(input string) (int, error) {
	parsedInput, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	relevantOrderingsCache := make(map[int][][2]int)

	correctMiddlePoints := make([]int, 0)

	for _, printStatement := range parsedInput.printStatements {
		relevantOrderings := selectRelevantOrderings(parsedInput.ordering, printStatement, relevantOrderingsCache)
		correctedPrintStatement := printStatement

		if !isPrintStatementCorrect(printStatement, relevantOrderings) {
			correctedPrintStatement = correctPrintStatement(printStatement, relevantOrderings)

			correctMiddlePoints = append(
				correctMiddlePoints,
				correctedPrintStatement[(len(correctedPrintStatement)/2):(len(correctedPrintStatement)/2+1)]...,
			)
		}

	}

	middlePointsSum := 0

	for _, middlePoint := range correctMiddlePoints {
		middlePointsSum += middlePoint
	}

	return middlePointsSum, nil
}
