package day6

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Vec2 struct {
	y int16
	x int16
}

type Field struct {
	directions []Vec2
	visited    int
	obstacle   bool
	start      bool
}

type Input struct {
	matrix [][]*Field
	start  Vec2
	size   Vec2
}

func parseInput(input string) Input {
	rawLines := strings.Split(input, "\n")

	result := Input{
		start:  Vec2{0, 0},
		matrix: make([][]*Field, 0),
		size:   Vec2{0, 0},
	}

	for i, line := range rawLines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		result.size.y++

		if int(result.size.x) < len(line) {
			result.size.x = int16(len(line))
		}

		result.matrix = append(result.matrix, make([]*Field, len(line)))
		for j, char := range line {
			switch char {
			case '#':
				result.matrix[i][j] = &Field{visited: 0, obstacle: true, start: false, directions: make([]Vec2, 0)}
			case '^':
				directions := make([]Vec2, 0)
				directions = append(directions, Vec2{-1, 0})
				result.matrix[i][j] = &Field{
					visited:    1,
					obstacle:   false,
					start:      true,
					directions: directions,
				}
				result.start = Vec2{int16(i), int16(j)}
			default:
				result.matrix[i][j] = &Field{
					visited:    0,
					obstacle:   false,
					start:      false,
					directions: make([]Vec2, 0),
				}
			}
		}
	}

	return result
}

func printMatrixWithBlocked(matrix [][]*Field, blockedPositions []Vec2) {
	for y, row := range matrix {
		for x, field := range row {
			for _, blockedPosition := range blockedPositions {
				if int(blockedPosition.y) == y && int(blockedPosition.x) == x {
					fmt.Print("O")

					goto end
				}
			}
			if field.obstacle {
				fmt.Print("#")
			} else if field.start {
				fmt.Print("^")
			} else {
				fmt.Print(".")
			}
		end:
		}

		fmt.Print("\n")
	}
}

func printMatrix(matrix [][]*Field) {
	for _, row := range matrix {
		for _, field := range row {
			if field.obstacle {
				fmt.Print("#")
			} else if field.start {
				fmt.Print("^")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

func rotateClockwise(direction Vec2) Vec2 {
	return Vec2{direction.x, -direction.y}
}

func Part1(input string) (int, error) {
	parsedInput := parseInput(input)

	fmt.Println("Start:", parsedInput.start)
	fmt.Println("Size:", parsedInput.size)
	fmt.Println("")

	// printMatrix(parsedInput.matrix)

	currentLocation := Vec2{
		y: parsedInput.start.y,
		x: parsedInput.start.x,
	}

	// Top
	direction := Vec2{-1, 0}

	sameLocationCounter := 0

	for {
		nextLocation := Vec2{
			y: currentLocation.y + direction.y,
			x: currentLocation.x + direction.x,
		}

		if nextLocation.y < 0 || nextLocation.y >= parsedInput.size.y || nextLocation.x < 0 || nextLocation.x >= parsedInput.size.x {
			fmt.Println("Out of bounds")
			break
		}

		nextField := parsedInput.matrix[nextLocation.y][nextLocation.x]

		if nextField.obstacle {
			direction = rotateClockwise(direction)
			sameLocationCounter++

			if sameLocationCounter == 4 {
				return 0, fmt.Errorf("Stuck at location x:%d y:%d", currentLocation.x, currentLocation.y)
			}

			continue
		}

		sameLocationCounter = 0
		currentLocation = nextLocation
		nextField.visited = nextField.visited + 1
	}

	distinceLocationsVisited := 0
	for _, row := range parsedInput.matrix {
		for _, field := range row {
			if field.visited > 0 {
				distinceLocationsVisited++
			}
		}
	}

	return distinceLocationsVisited, nil
}

func canGoOut(matrix [][]*Field, start Vec2, size Vec2) bool {
	currentLocation := start
	// Top
	direction := Vec2{-1, 0}

	visitedLocationsMap := make(map[int64]bool)

	for {
		nextLocation := Vec2{
			y: currentLocation.y + direction.y,
			x: currentLocation.x + direction.x,
		}

		if nextLocation.y < 0 || nextLocation.y >= size.y || nextLocation.x < 0 || nextLocation.x >= size.x {
			return true
		}
		key := int64(nextLocation.y) + (int64(nextLocation.x) << 8) + (int64(direction.y+2) << 16) + (int64(direction.x+2) << 24)

		if visitedLocationsMap[key] {
			return false
		}

		visitedLocationsMap[key] = true

		nextField := matrix[nextLocation.y][nextLocation.x]

		if nextField.obstacle {
			direction = rotateClockwise(direction)
			continue
		}

		currentLocation = nextLocation
	}
}

type Vec3 struct {
	y int16
	x int16
	z int16
}

func Part2(input string) (int, error) {
	parsedInput := parseInput(input)

	fmt.Println("Start:", parsedInput.start)
	fmt.Println("Size:", parsedInput.size)
	fmt.Println("")

	printMatrix(parsedInput.matrix)

	currentLocation := Vec2{
		y: parsedInput.start.y,
		x: parsedInput.start.x,
	}

	// Top
	direction := Vec2{-1, 0}

	for {
		nextLocation := Vec2{
			y: currentLocation.y + direction.y,
			x: currentLocation.x + direction.x,
		}

		if nextLocation.y < 0 || nextLocation.y >= parsedInput.size.y || nextLocation.x < 0 || nextLocation.x >= parsedInput.size.x {
			fmt.Println("Out of bounds")
			break
		}

		nextField := parsedInput.matrix[nextLocation.y][nextLocation.x]

		if nextField.obstacle {
			direction = rotateClockwise(direction)

			currentField := parsedInput.matrix[currentLocation.y][currentLocation.x]

			if len(currentField.directions) == 0 {
				currentField.directions = append(currentField.directions, direction)
			} else {
				for _, fieldDirection := range currentField.directions {
					if direction.y == fieldDirection.y && direction.x == fieldDirection.x {
						continue
					}
					currentField.directions = append(currentField.directions, direction)
				}
			}
			continue
		}

		if len(nextField.directions) == 0 {
			nextField.directions = append(nextField.directions, direction)
		} else {
			for _, fieldDirection := range nextField.directions {
				if direction.y == fieldDirection.y && direction.x == fieldDirection.x {
					continue
				}
				nextField.directions = append(nextField.directions, direction)
			}
		}

		currentLocation = nextLocation
		nextField.visited = nextField.visited + 1
	}

	var previousSwapY int
	var previousSwapX int

	blockedLocationsMap := make(map[int64]Vec2, 0)

	for y, row := range parsedInput.matrix {
		for x, field := range row {
			if field.visited > 0 {
				for _, direction := range field.directions {
					if direction.y != 0 {
						previousSwapY = y + int(direction.y)
						previousSwapX = x
					} else {
						previousSwapY = y
						previousSwapX = x + int(direction.x)
					}

					if previousSwapY < 0 || previousSwapY >= int(parsedInput.size.y) || previousSwapX < 0 || previousSwapX >= int(parsedInput.size.x) {
						continue
					}

					if previousSwapX == int(parsedInput.start.x) && previousSwapY == int(parsedInput.start.y) {
						continue
					}

					if parsedInput.matrix[previousSwapY][previousSwapX].obstacle {
						continue
					}

					parsedInput.matrix[previousSwapY][previousSwapX].obstacle = true

					if !canGoOut(parsedInput.matrix, parsedInput.start, parsedInput.size) {
						key := int64(previousSwapY) + (int64(previousSwapX) << 32)
						if _, ok := blockedLocationsMap[key]; !ok {
							blockedLocationsMap[key] = Vec2{y: int16(previousSwapY), x: int16(previousSwapX)}
						}
					}
					parsedInput.matrix[previousSwapY][previousSwapX].obstacle = false
				}
			}
		}
	}

	blockedLocations := slices.Collect(maps.Values(blockedLocationsMap))

	printMatrixWithBlocked(parsedInput.matrix, blockedLocations)

	return len(blockedLocations), nil
}

type Difference struct {
	achar    string
	bchar    string
	position Vec2
}

func compareStrings(a string, b string) []Difference {
	alines := make([]string, 0)
	blines := make([]string, 0)

	for _, line := range strings.Split(a, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		alines = append(alines, line)
	}

	for _, line := range strings.Split(b, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		blines = append(blines, line)
	}

	differences := make([]Difference, 0)
	for i, aline := range alines {
		for j, achar := range aline {
			if achar != rune(blines[i][j]) {
				differences = append(differences, Difference{
					position: Vec2{0, int16(i)},
					achar:    string(achar),
					bchar:    string(blines[i][j]),
				})
			}
		}
	}

	return differences
}

func DebugDay6(validResult string, invalidResult string) {
	comparisons := compareStrings(validResult, invalidResult)
	fmt.Println("Differences:", comparisons)
}
