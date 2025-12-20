package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func ParseInput1(input string) []Problem {
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

func Solve(problems []Problem) uint64 {
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

func ParseInput2(input string) []Problem {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	var problem_start_columns []int
	for i, r := range lines[len(lines) - 1] {
		if !unicode.IsSpace(r) {
			problem_start_columns = append(problem_start_columns, i)
		}
	}

	problems := make([]Problem, 0, len(problem_start_columns))
	for problem_i, rune_lb := range problem_start_columns {
		// rune_ub is exclusive
		var rune_ub int
		if problem_i + 1 == len(problem_start_columns) {
			rune_ub = len(lines[0])
		} else {
			// Assuming operators appear in the first column for a problem and
			// problems are separated by a column containing only spaces
			rune_ub = problem_start_columns[problem_i + 1] - 1
		}

		number_count := rune_ub - rune_lb
		number_bytes := make([][]byte, number_count)
		// Assuming every rune is a single byte
		for number_i := range number_count {
			byte_i := rune_lb + number_i
			for _, line := range lines[: len(lines) - 1] {
				if !unicode.IsSpace(rune(line[byte_i])) {
					number_bytes[number_i] = append(
						number_bytes[number_i],
						line[byte_i],
					)
				}
			}
		}

		numbers := make([]uint64, len(number_bytes))
		for i, bytes := range number_bytes {
			var err error
			numbers[i], err = strconv.ParseUint(string(bytes), 10, 64)
			if err != nil {
				log.Panic(problem_i, err)
			}
		}

		var operator Operator
		switch rune(lines[len(lines) - 1][rune_lb]) {
		case '+':
			operator = Add
		case '*':
			operator = Multiply
		default:
			log.Fatal("Invalid Operator:", rune(lines[len(lines) - 1][rune_lb]))
		}

		problems = append(
			problems,
			Problem{
				Numbers: numbers,
				Operator: operator,
			},
		)
	}

	return problems
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error Reading Input:", err)
	}

	problems1 := ParseInput1(string(input_bytes))
	fmt.Println("Part 1:", Solve(problems1))
	problems2 := ParseInput2(string(input_bytes))
	fmt.Println("Part 2:", Solve(problems2))
}
