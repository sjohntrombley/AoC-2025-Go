package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operator uint8
const (
	Add Operator = iota
	Multiply
)

type Problem struct{
	Numbers []uint64
	Operator Operator
}

func ParseInput(input string) []Problem {
	var problem_string_grid [][]string
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		var row_strings []string
		for element := range strings.SplitSeq(line, " ") {
			if element != "" {
				row_strings = append(row_strings, element)
			}
		}

		problem_string_grid = append(problem_string_grid, row_strings)
	}

	problems := make([]Problem, 0, len(problem_string_grid[0]))
	for j := range problem_string_grid[0] {
		var numbers []uint64
		for i := range len(problem_string_grid) - 1 {
			new_number, err := strconv.ParseUint(problem_string_grid[i][j], 10, 64)
			if err != nil {
				log.Fatal("Error parsing number in problem:", err)
			}
			numbers = append(numbers, new_number)
		}

		var operator Operator
		switch problem_string_grid[len(problem_string_grid) - 1][j] {
		case "+":
			operator = Add
		case "*":
			operator = Multiply
		default:
			log.Fatal("Invalid Operator Error:", problem_string_grid[len(problem_string_grid) - 1][j])
		}

		problems = append(problems, Problem{Numbers: numbers, Operator: operator})
	}

	return problems
}

func Part1(problems []Problem) uint64 {
	var grandTotal uint64
	for _, problem := range problems {
		var solution uint64
		switch problem.Operator {
		case Add:
			solution = 0
			for _, number := range problem.Numbers {
				solution += number
			}
		case Multiply:
			solution = 1
			for _, number := range problem.Numbers {
				solution *= number
			}
		default:
			log.Fatalf("Invalid Problem.Operator value: %#v", problem)
		}

		grandTotal += solution
	}

	return grandTotal
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error Reading Input:", err)
	}

	problems := ParseInput(string(input_bytes))

	fmt.Println("Part 1:", Part1(problems))
}
