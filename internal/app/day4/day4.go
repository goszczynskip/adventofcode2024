package day4

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	charsMatrix := make([][]string, len(lines))

	for i, line := range lines {
		chars := strings.Split(line, "")
		if chars[len(chars)-1] == "" {
			chars = chars[:len(chars)-1]
		}
		charsMatrix[i] = chars
	}

	return charsMatrix
}

const XMAS_WORD = "XMAS"

var dirMatrix = [...][2]int{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func getCharFromMatrix(matrix [][]string, i, j int) string {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return ""
	}

	return matrix[i][j]
}

type Location struct {
	i        int
	j        int
	dirIndex int
}

func findWord(matrix [][]string, word string) []Location {
	locations := make([]Location, 0)

	for i, line := range matrix {
		for j := range line {

			correctCount := make([]int, 8)
			for delta, wordChar := range word {
				for dirIndex, dir := range dirMatrix {
					if getCharFromMatrix(
						matrix,
						i+dir[0]*delta,
						j+dir[1]*delta,
					) == string(wordChar) {
						correctCount[dirIndex]++
					}
				}
			}

			for dirIndex, count := range correctCount {
				if count == len(word) {
					locations = append(locations, Location{i, j, dirIndex})
				}
			}
		}
	}

	return locations
}

func Part1(input string) (int, error) {
	charMatrix := parseInput(input)

	locations := findWord(charMatrix, XMAS_WORD)

	return len(locations), nil
}

const MAS_WORD = "MAS"

var transposeMatrix = [...][2]int{
	{-1, 0},
	{0, -1},
	{-1, -1},
}

func Part2(input string) (int, error) {
	charMatrix := parseInput(input)

	fmt.Println(len(charMatrix), len(charMatrix[0]))

	locations := findWord(charMatrix, MAS_WORD)

	uniqueLocations := make(map[string]bool)

	for _, location1 := range locations {
		for _, location2 := range locations {
			if location1.i == location2.i && location1.j == location2.j {
				continue
			}

			for _, transpose := range transposeMatrix {
				transposedDir := [2]int{
          dirMatrix[location2.dirIndex][0] * transpose[0],
          dirMatrix[location2.dirIndex][1] * transpose[1],
        }
				transposedI := location2.i - transposedDir[0]*(len(MAS_WORD)-1)
				transposedJ := location2.j - transposedDir[1]*(len(MAS_WORD)-1)
        transposedDirI := dirMatrix[location2.dirIndex][0]
        if(transpose[0] != 0)  {
          transposedDirI = transposedDirI * transpose[0]
        }

        transposedDirJ := dirMatrix[location2.dirIndex][1]
        if(transpose[1] != 0)  {
          transposedDirJ = transposedDirJ * transpose[1]
        }

				if location1.i == transposedI &&
					location1.j == transposedJ &&
					dirMatrix[location1.dirIndex][0] == transposedDirI &&
					dirMatrix[location1.dirIndex][1] == transposedDirJ {

					stringIIndex := strconv.Itoa(location1.i + dirMatrix[location1.dirIndex][0])
					stringJIndex := strconv.Itoa(location1.j + dirMatrix[location1.dirIndex][1])

					uniqueLocations[stringIIndex+"_"+stringJIndex] = true
				}
			}
		}
	}

	allKeys := slices.Collect(
		maps.Keys(
			uniqueLocations,
		),
	)

	sort.Strings(allKeys)

	return len(uniqueLocations), nil
}
