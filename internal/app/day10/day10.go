package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int, error) {
	result := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		chars := strings.Split(line, "")

		nextLine := make([]int, len(chars))

		for j, char := range chars {
			number, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}

			nextLine[j] = number
		}
		result = append(result, nextLine)
	}

	return result, nil
}

func findLowestPoints(matrix [][]int) [][2]int {
	lowestPoints := make([][2]int, 0)

	for i, line := range matrix {
		for j, point := range line {
			if point == 0 {
				lowestPoints = append(lowestPoints, [2]int{i, j})
			}
		}
	}

	return lowestPoints
}

func isOutOfBounds(position [2]int, matrix [][]int) bool {
	firstDimLength := len(matrix)
	secondDimLength := len(matrix[0])

	return position[0] < 0 || position[0] >= firstDimLength || position[1] < 0 || position[1] >= secondDimLength
}

func traverse(position [2]int, matrix [][]int, endings *[][2]int, visitedPoints *[][2]int) {
	transformations := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	*visitedPoints = append(*visitedPoints, position)
	currentPointValue := matrix[position[0]][position[1]]

	if currentPointValue == 9 {
		for _, ending := range *endings {
			if ending[0] == position[0] && ending[1] == position[1] {
				return
			}
		}

		*endings = append(*endings, position)
		return
	}

	for _, transformation := range transformations {
		nextPoint := [2]int{position[0] + transformation[0], position[1] + transformation[1]}

		if isOutOfBounds(nextPoint, matrix) {
			// fmt.Printf("Is out of bounds point at: %v\n", nextPoint)
			continue
		}

		nextPointValue := matrix[nextPoint[0]][nextPoint[1]]

		if nextPointValue-currentPointValue != 1 {
			continue
		}

		alreadyVisited := false
		for _, visitedPoint := range *visitedPoints {
			if visitedPoint[0] == nextPoint[0] && visitedPoint[1] == nextPoint[1] {
				// fmt.Printf("Already visited point at: %v\n", nextPoint)
				alreadyVisited = true
			}
		}

		if alreadyVisited {
			continue
		}

		traverse(nextPoint, matrix, endings, visitedPoints)
	}
}

func Part1(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	startingPoints := findLowestPoints(matrix)

	scores := make([]int, 0)

	for _, startingPoint := range startingPoints {
		endingPoints := make([][2]int, 0)
		visitedPoints := make([][2]int, 0)
		traverse(startingPoint, matrix, &endingPoints, &visitedPoints)

		fmt.Printf("Score: %v %d\n", startingPoint, len(endingPoints))
		scores = append(scores, len(endingPoints))
	}

	result := 0
	for _, score := range scores {
		result += score
	}

	return result, nil
}

func traverseWithTrails(position [2]int, matrix [][]int, endings *[][2]int, trails map[string][][2]int, visitedPoints [][2]int) {
	transformations := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

  nextVisitedPoints := make([][2]int, 0)
  nextVisitedPoints = append(nextVisitedPoints, visitedPoints...)
  nextVisitedPoints = append(nextVisitedPoints, position)
    
	currentPointValue := matrix[position[0]][position[1]]

	if currentPointValue == 9 {

    strPath := make([]string, len(nextVisitedPoints))
    for i, point := range nextVisitedPoints {
      strPath[i] = "(" + strconv.Itoa(point[0]) + ":" +  strconv.Itoa(point[1]) + ")"
    }

    strKey := strings.Join(strPath, ",")

    trails[strKey] = nextVisitedPoints

		for _, ending := range *endings {
			if ending[0] == position[0] && ending[1] == position[1] {
				return
			}
		}

		*endings = append(*endings, position)
		return
	}

	for _, transformation := range transformations {
		nextPoint := [2]int{position[0] + transformation[0], position[1] + transformation[1]}

		if isOutOfBounds(nextPoint, matrix) {
			// fmt.Printf("Is out of bounds point at: %v\n", nextPoint)
			continue
		}

		nextPointValue := matrix[nextPoint[0]][nextPoint[1]]

		if nextPointValue-currentPointValue != 1 {
			continue
		}

		alreadyVisited := false
		for _, visitedPoint := range nextVisitedPoints {
			if visitedPoint[0] == nextPoint[0] && visitedPoint[1] == nextPoint[1] {
				// fmt.Printf("Already visited point at: %v\n", nextPoint)
				alreadyVisited = true
			}
		}

		if alreadyVisited {
			continue
		}

		traverseWithTrails(nextPoint, matrix, endings, trails, nextVisitedPoints)
	}
}

func Part2(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	startingPoints := findLowestPoints(matrix)

	scores := make([]int, 0)

	for _, startingPoint := range startingPoints {
		endingPoints := make([][2]int, 0)
		visitedPoints := make([][2]int, 0)
    trails := make(map[string][][2]int)
		traverseWithTrails(startingPoint, matrix, &endingPoints, trails, visitedPoints)

		scores = append(scores, len(trails))
	}

	result := 0
	for _, score := range scores {
		result += score
	}

	return result, nil
}
