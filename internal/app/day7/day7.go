package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	numbers []int64
	result  int64
}

func parseInput(input string) ([]Equation, error) {
	rawLines := strings.Split(input, "\n")

	var equations []Equation

	for i, line := range rawLines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		equation := Equation{result: 0, numbers: make([]int64, 0)}
		eqsplit := strings.Split(line, ": ")
		result, err := strconv.ParseInt(eqsplit[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing line (%d) result: %s. Error: %w", i, line, err)
		}

		equation.result = result

		for j, stringNumber := range strings.Split(eqsplit[1], " ") {
			parsedNumber, err := strconv.ParseInt(stringNumber, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing line (%d) number (%d): %s. Error: %w", i, j, line, err)
			}

			equation.numbers = append(equation.numbers, parsedNumber)
		}

		equations = append(equations, equation)
	}

	return equations, nil
}

func printEquation(equation Equation) {
	stringNumbers := make([]string, len(equation.numbers))
	for i, n := range equation.numbers {
		stringNumbers[i] = strconv.FormatInt(n, 10)
	}
	fmt.Println("Result:", equation.result, "Numbers:", strings.Join(stringNumbers, ", "))
}

func printBranch(branch []*Node) {
	for i, b := range branch {
		fmt.Print(b.number)

		if i >= len(branch)-1 {
			continue
		}

		if b.opAdd == branch[i+1] {
			fmt.Print("+")
		}

		if b.opMul == branch[i+1] {
			fmt.Print("*")
		}

		if b.opCon == branch[i+1] {
			fmt.Print("|")
		}
	}
}

type Node struct {
	opMul  *Node
	opAdd  *Node
	opCon  *Node
	total  int64
	number int64
	equal  bool
}

func buildEquationTree(result int64, numbers []int64, allowedOperators string, acc *Node) (*Node, error) {
	if len(numbers) == 1 {
		return acc, nil
	}

	var addNode *Node
	var mulNode *Node
	var conNode *Node

	prevTotal := numbers[0]
	if acc != nil {
		prevTotal = acc.total
	}

	for _, op := range allowedOperators {
		switch string(op) {
		case "+":
			{

				total := prevTotal + numbers[1]

				currentTree := &Node{total: total, number: numbers[1], equal: result == total}

				addNodeAux, err := buildEquationTree(result, numbers[1:], allowedOperators, currentTree)
				if err != nil {
					return nil, err
				}

				addNode = addNodeAux
				continue
			}
		case "*":
			{

				total := prevTotal * numbers[1]

				currentTree := &Node{total: total, number: numbers[1], equal: result == total}

				mulNodeAux, err := buildEquationTree(result, numbers[1:], allowedOperators, currentTree)
				if err != nil {
					return nil, err
				}

				mulNode = mulNodeAux
				continue
			}
		case "|":
			{
				stringTotal := strconv.FormatInt(prevTotal, 10)
				nextStringValue := strconv.FormatInt(numbers[1], 10)
				concatenated := stringTotal + nextStringValue
				total, err := strconv.ParseInt(concatenated, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("concatenation overflow: %w", err)
				}

				// Verify the concatenation didn't cause overflow
				if total < prevTotal {
					return nil, fmt.Errorf("concatenation result %d is smaller than previous total %d", total, prevTotal)
				}

				currentTree := &Node{total: total, number: numbers[1], equal: result == total}

				conNodeAux, err := buildEquationTree(result, numbers[1:], allowedOperators, currentTree)
				if err != nil {
					return nil, err
				}

				conNode = conNodeAux

				continue
			}
		}
	}

	if acc == nil {
		return &Node{total: prevTotal, number: numbers[0], equal: result == int64(numbers[0]), opAdd: addNode, opMul: mulNode, opCon: conNode}, nil
	} else {
		acc.opAdd = addNode
		acc.opMul = mulNode
		acc.opCon = conNode
		return acc, nil
	}
}

func traverseLeafs(tree *Node, fn func(tree *Node, branch []*Node, containsNew bool), branch *[]*Node, containsNew bool) {
	var branchCopy []*Node

	if branch != nil {
		branchCopy = make([]*Node, len(*branch))
		copy(branchCopy, *branch)
	} else {
		branchCopy = make([]*Node, 0)
	}

	branchCopy = append(branchCopy, tree)

	if tree.opAdd == nil && tree.opMul == nil && tree.opCon == nil {
		fn(tree, branchCopy, containsNew)
		return
	}

	if tree.opAdd != nil {
		traverseLeafs(tree.opAdd, fn, &branchCopy, containsNew)
	}
	if tree.opMul != nil {
		traverseLeafs(tree.opMul, fn, &branchCopy, containsNew)
	}
	if tree.opCon != nil {
		traverseLeafs(tree.opCon, fn, &branchCopy, true)
	}
}

func Part1(input string) (int, error) {
	equations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var result int64 = 0
	equationCount := 0
	for _, eq := range equations {
		eqTree, err := buildEquationTree(int64(eq.result), eq.numbers, "+*", nil)
		if err != nil {
			return 0, err
		}

		canBeSolved := false
		traverseLeafs(
			eqTree,
			func(tree *Node, branch []*Node, containsNew bool) {
				err := verifyEquation(eq, branch)

				if err != nil {
					// fmt.Printf("Traverse error: %v", err)
				} else {
					canBeSolved = true
				}
			},
			nil,
			false,
		)

		if canBeSolved {
			nextResult := result + eq.result
			if nextResult-eq.result == result {
				result = nextResult
			} else {
				fmt.Println("Overflow")
			}
			equationCount++
		}
	}

	fmt.Println(result)
	fmt.Println(equationCount)

	return 0, nil
}

func verifyEquation(equation Equation, branch []*Node) error {
	branchLen := len(branch)
	var total int64 = branch[0].number
	for i, b := range branch {
		if i >= branchLen-1 {
			continue
		}

		if b.opAdd == branch[i+1] {
			total += branch[i+1].number
			continue
		}
		if b.opMul == branch[i+1] {
			total *= branch[i+1].number
			continue
		}
		if b.opCon == branch[i+1] {
			stringTotal := strconv.FormatInt(total, 10)
			nextStringValue := strconv.FormatInt(branch[i+1].number, 10)
			concatenated := stringTotal + nextStringValue
			totalAux, err := strconv.ParseInt(concatenated, 10, 64)
			if err != nil {
				return fmt.Errorf("concatenation overflow: %w", err)
			}

			// Verify the concatenation didn't cause overflow
			if totalAux < total {
				return fmt.Errorf("concatenation result %d is smaller than previous total %d", totalAux, total)
			}

			// fmt.Printf("Concatenation: %s + %s = %s (as number: %d)\n",
			// 	stringTotal, nextStringValue, concatenated, totalAux)
			total = totalAux
			continue
		}
	}

	if total != equation.result {
		return fmt.Errorf("Equation not solved correctly. Expected: %d, got: %d", equation.result, total)
	}

	return nil
}

func Part2(input string) (int, error) {
	equations, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var result int64 = 0
	equationCount := 0
	for _, eq := range equations {
		eqTree, err := buildEquationTree(int64(eq.result), eq.numbers, "+*|", nil)
		if err != nil {
			return 0, err
		}

		canBeSolved := false
		traverseLeafs(
			eqTree,
			func(tree *Node, branch []*Node, containsNew bool) {
        if tree.equal {
					canBeSolved = true
				}
			},
			nil,
			false,
		)

		if canBeSolved {
			nextResult := result + eq.result
			if nextResult-eq.result == result {
				result = nextResult
			} else {
				fmt.Println("Overflow")
			}
			equationCount++
		}
	}

	fmt.Println(result)
	fmt.Println(equationCount)

	return 0, nil
}
