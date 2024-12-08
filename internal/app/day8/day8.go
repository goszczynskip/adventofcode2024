package day8

import (
	"fmt"
	"strings"
)

func parseInput(input string) (map[string][][2]int, [2]int) {
	rawLines := strings.Split(input, "\n")

	result := make(map[string][][2]int)
	bounds := [2]int{0, 0}

	for i, line := range rawLines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		bounds[0]++

		rowLength := 0
		for j, char := range line {
			rowLength++

			if string(char) == "." {
				continue
			}

			newEntry := [2]int{i, j}

			if entry, ok := result[string(char)]; ok {
				result[string(char)] = append(entry, newEntry)
			} else {
				arr := make([][2]int, 1)
				arr[0] = newEntry
				result[string(char)] = arr
			}
		}

		if rowLength > bounds[1] {
			bounds[1] = rowLength
		}
	}

	return result, bounds
}

func getKey(a [2]int) string {
	return fmt.Sprintf("%d,%d", a[0], a[1])
}

func findResonantionPlaces(antennas [][2]int, bounds [2]int) map[string][][2]int {
	frequencies := make(map[string][][2]int)

	appendFrequency := func(freq [2]int) {
		if freq[0] < 0 || freq[0] >= bounds[0] || freq[1] < 0 || freq[1] >= bounds[1] {
			return
		}

		key := getKey(freq)
		if place, ok := frequencies[key]; ok {
			frequencies[key] = append(place, freq)
		} else {
			frequencies[key] = [][2]int{freq}
		}
	}

	for _, antennaA := range antennas {
		for _, antennaB := range antennas {
			if antennaA[0] == antennaB[0] && antennaA[1] == antennaB[1] {
				continue
			}

			distance := [2]int{
				antennaA[0] - antennaB[0],
				antennaA[1] - antennaB[1],
			}

			freq1 := [2]int{antennaA[0] + distance[0], antennaA[1] + distance[1]}
			freq2 := [2]int{antennaB[0] - distance[0], antennaB[1] - distance[1]}

			appendFrequency(freq1)
			appendFrequency(freq2)
		}
	}

	return frequencies
}

func findResonantionPlacesWithResonance(antennas [][2]int, bounds [2]int) map[string][][2]int {
	frequencies := make(map[string][][2]int)

	appendFrequency := func(freq [2]int) {
		key := getKey(freq)
		if place, ok := frequencies[key]; ok {
			frequencies[key] = append(place, freq)
		} else {
			frequencies[key] = [][2]int{freq}
		}
	}

	isOutOfBounds := func(freq [2]int) bool {
		return freq[0] < 0 || freq[0] >= bounds[0] || freq[1] < 0 || freq[1] >= bounds[1]
	}

	for _, antennaA := range antennas {
		for _, antennaB := range antennas {
			if antennaA[0] == antennaB[0] && antennaA[1] == antennaB[1] {
				continue
			}

			distance := [2]int{
				antennaA[0] - antennaB[0],
				antennaA[1] - antennaB[1],
			}

			for i := 0; i < bounds[0]+1; i++ {
				freq := [2]int{antennaA[0] + distance[0] * i, antennaA[1] + distance[1] * i}

				if isOutOfBounds(freq) {
					break
				}

				appendFrequency(freq)
			}

			for i := 0; i < bounds[0]+1; i++ {
				freq := [2]int{antennaB[0] - distance[0] * i, antennaB[1] - distance[1] * i}

				if isOutOfBounds(freq) {
					break
				}

				appendFrequency(freq)
			}
		}
	}

	return frequencies
}

func Part1(input string) (int, error) {
	parsedInput, bounds := parseInput(input)
	fmt.Println(bounds)
	uniqueLocations := make(map[string][2]int)
	for _, value := range parsedInput {
		resonationPlaces := findResonantionPlaces(value, bounds)
		for _, places := range resonationPlaces {
			for _, place := range places {
				key := fmt.Sprintf("%d,%d", place[0], place[1])
				if _, ok := uniqueLocations[key]; !ok {
					uniqueLocations[key] = place
				}
			}
		}
	}

	return len(uniqueLocations), nil
}

func Part2(input string) (int, error) {
	parsedInput, bounds := parseInput(input)
	fmt.Println(bounds)
	uniqueLocations := make(map[string][2]int)
	for _, value := range parsedInput {
		resonationPlaces := findResonantionPlacesWithResonance(value, bounds)
		for _, places := range resonationPlaces {
			for _, place := range places {
				key := fmt.Sprintf("%d,%d", place[0], place[1])
				if _, ok := uniqueLocations[key]; !ok {
					uniqueLocations[key] = place
				}
			}
		}
	}

	return len(uniqueLocations), nil
}
