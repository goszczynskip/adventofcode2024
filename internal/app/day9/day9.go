package day9

import (
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, error) {
	numbers := make([]int, len(input))
	trimmed := strings.ReplaceAll(input, "\n", "")
	for i, char := range trimmed {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}

		numbers[i] = number
	}

	return numbers, nil
}

func Part1(input string) (int, error) {
	numbers, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	buffer := make([]int, len(numbers)*9)

	for i := range buffer {
		buffer[i] = -1
	}

	cursor := 0

	for i, number := range numbers {
		isFile := i%2 == 0

		for j := 0; j < number; j++ {
			if isFile {
				buffer[cursor] = i / 2
			} else {
				buffer[cursor] = -1
			}

			cursor++
		}
	}

	putCursor := 0
	takeCursor := len(buffer)

	for takeCursor > putCursor {
		takeCursor--

		fileId := buffer[takeCursor]
		if fileId == -1 {
			continue
		}

		for {
			targetFileId := buffer[putCursor]
			if targetFileId == -1 {
				buffer[putCursor] = fileId
				buffer[takeCursor] = -1
				break
			}

			putCursor++

			if takeCursor <= putCursor {
				break
			}
		}
	}

	result := 0

	for i, fileId := range buffer {
		if fileId == -1 {
			continue
		}

		result += i * fileId

	}

	return result, nil
}

func fileEndNegative(start int, buffer []int) int {
	endIndex := start - 1

	for firstFile := buffer[start]; endIndex >= 0 && firstFile == buffer[endIndex]; endIndex-- {
	}

	return endIndex + 1
}

func fileEndPositive(start int, buffer []int) int {
	endIndex := start + 1

	for firstFile := buffer[start]; endIndex < len(buffer) && firstFile == buffer[endIndex]; endIndex++ {
	}

	return endIndex
}

func Part2(input string) (int, error) {
	numbers, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	buffer := make([]int, len(numbers)*9)

	for i := range buffer {
		buffer[i] = -1
	}

	cursor := 0

	for i, number := range numbers {
		isFile := i%2 == 0

		for j := 0; j < number; j++ {
			if isFile {
				buffer[cursor] = i / 2
			} else {
				buffer[cursor] = -1
			}

			cursor++
		}
	}

	takeCursor := len(buffer)

	filesMoved := make(map[int]bool)

	for takeCursor > 0 {
		takeCursor--

		fileId := buffer[takeCursor]
		if fileId == -1 {
			continue
		}

		if _, ok := filesMoved[fileId]; ok {
			continue
		}

		fileEnd := fileEndNegative(takeCursor, buffer)
		fileLength := (takeCursor + 1) - fileEnd

		putCursor := 0

		for putCursor < takeCursor {
			targetFileId := buffer[putCursor]

			if targetFileId == -1 {
				slotEnd := fileEndPositive(putCursor, buffer)
				slotLength := slotEnd - putCursor

				if slotLength < fileLength {
					putCursor += slotLength
					continue
				}

				for i, v1 := range buffer[fileEnd : takeCursor+1] {
					buffer[putCursor+i] = v1
					buffer[fileEnd+i] = -1
				}

				filesMoved[fileId] = true

				break
			}

			putCursor++
		}

		takeCursor = fileEnd
	}

	result := 0

	for i, fileId := range buffer {
		if fileId == -1 {
			continue
		}

		result += i * fileId

	}

	return result, nil
}
