package day6

import (
	"fmt"
	"strings"
)

type Vec2 struct {
	y int
	x int
}

type Field struct {
	visited  int
	obstacle bool
	start    bool
}

type Input struct {
	matrix [][]Field
	start  Vec2
	size   Vec2
}

func parseInput(input string) Input {
	rawLines := strings.Split(input, "\n")

	result := Input{
		start:  Vec2{0, 0},
		matrix: make([][]Field, 0),
		size:   Vec2{0, 0},
	}

	for i, line := range rawLines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		result.size.y++

		if result.size.x < len(line) {
			result.size.x = len(line)
		}

		result.matrix = append(result.matrix, make([]Field, len(line)))
		for j, char := range line {
			switch char {
			case '#':
				result.matrix[i][j] = Field{0, true, false}
			case '^':
				result.matrix[i][j] = Field{1, false, true}
				result.start = Vec2{i, j}
			default:
				result.matrix[i][j] = Field{0, false, false}
			}
		}
	}

	return result
}

func printMatrix(matrix [][]Field, blockedPositions []Vec2) {
	for y, row := range matrix {
		for x, field := range row {
			for _, blockedPosition := range blockedPositions {
				if blockedPosition.y == y && blockedPosition.x == x {
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

		nextField := &parsedInput.matrix[nextLocation.y][nextLocation.x]

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

type Vec3 struct {
	y int
	x int
	z int
}

func Part2(input string) (int, error) {
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

	traversedHorizontalAxis := make(map[int][2]*Vec3, 0)
	traversedVerticalAxis := make(map[int][2]*Vec3, 0)

	step := 0

	for {
		step++
		nextLocation := Vec2{
			y: currentLocation.y + direction.y,
			x: currentLocation.x + direction.x,
		}

		if direction.y != 0 {
			if val, ok := traversedVerticalAxis[currentLocation.x*direction.y]; ok {
				if direction.y < 0 && val[1].y > currentLocation.y {
					val[1].y = currentLocation.y
				} else if direction.y > 0 && val[1].y < currentLocation.y {
					val[1].y = currentLocation.y
				}
			} else {
				traversedVerticalAxis[currentLocation.x*direction.y] = [2]*Vec3{
					{currentLocation.y, currentLocation.x, step},
					{currentLocation.y, currentLocation.x, step},
				}
			}
		} else {
			if val, ok := traversedHorizontalAxis[currentLocation.y*direction.x]; ok {
				if direction.x < 0 && val[1].x > currentLocation.x {
					val[1].x = currentLocation.x
				} else if direction.x > 0 && val[1].x < currentLocation.x {
					val[1].x = currentLocation.x
				}
			} else {
				traversedHorizontalAxis[currentLocation.y*direction.x] = [2]*Vec3{
					{currentLocation.y, currentLocation.x, step},
					{currentLocation.y, currentLocation.x, step},
				}
			}
		}

		if nextLocation.y < 0 || nextLocation.y >= parsedInput.size.y || nextLocation.x < 0 || nextLocation.x >= parsedInput.size.x {
			fmt.Println("Out of bounds")
			break
		}

		nextField := &parsedInput.matrix[nextLocation.y][nextLocation.x]

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

	positionsToBlock := make([]Vec2, 0)

	for vKey, vVal := range traversedVerticalAxis {
		var vDirection int
		if vKey < 0 {
			vDirection = -1
		} else {
			vDirection = 1
		}

		distance := (vVal[0].y - vVal[1].y) * vDirection * -1

		for i := 0; i < distance; i++ {
			hKey := (vVal[0].y + i*vDirection) * vDirection * -1

			if hVal, ok := traversedHorizontalAxis[hKey]; ok {
				if hVal[0].z > vVal[0].z {
					continue
				}

				diff := (vKey * vDirection) - hVal[1].x
				if diff < 0 && vDirection < 0 || diff > 0 && vDirection > 0 {
					positionsToBlock = append(positionsToBlock, Vec2{(hKey * vDirection * -1) + vDirection, vKey * vDirection})
				}
			}
		}
	}

	for hKey, hVal := range traversedHorizontalAxis {
		var hDirection int
		if hKey < 0 {
			hDirection = -1
		} else {
			hDirection = 1
		}

		distance := (hVal[0].x - hVal[1].x) * hDirection * -1

		for i := 0; i < distance; i++ {
			vKey := (hVal[0].x + i*hDirection) * hDirection

			if vVal, ok := traversedVerticalAxis[vKey]; ok {
				if vVal[0].z > hVal[0].z {
					continue
				}

				diff := (hKey * hDirection) - vVal[1].y

				if diff < 0 && hDirection > 0 || diff > 0 && hDirection < 0 {
					positionsToBlock = append(positionsToBlock, Vec2{hKey * hDirection, (vKey * hDirection) + hDirection})
				}
			}
		}
	}

	filteredPositionsToBlock := make([]Vec2, 0)

	for _, position := range positionsToBlock {
		skip := false
		for y, line := range parsedInput.matrix {
			for x, field := range line {
				if field.obstacle && position.x == x && position.y == y {
					skip = true
				}
			}
		}

		if !skip {
			filteredPositionsToBlock = append(filteredPositionsToBlock, position)
		}
	}

	fmt.Println("Filtered positions to block:", filteredPositionsToBlock)

	printMatrix(parsedInput.matrix, filteredPositionsToBlock)

	return len(filteredPositionsToBlock), nil
}
