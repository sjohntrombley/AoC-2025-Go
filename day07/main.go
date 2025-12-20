package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseInput(input string) ([][]bool, int) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	
	start_col := strings.Index(lines[0], "S")

	// Since every character in the input should be encodable as a single byte in utf8, indexes in a line should
	// directly correspond with indexes in splitters
	splitters := make([][]bool, len(lines) - 1)
	for row_index, line := range lines[1:] {
		splitters[row_index] = make([]bool, len(line))
		for column_index, rune_ := range lines[row_index + 1] {
			if rune_ == '^' {
				splitters[row_index][column_index] = true
			}
		}
	}

	return splitters, start_col
}

func Solve1(splitters [][]bool, start_column int) int {
	row_length := len(splitters[0])
	cur_beams := make([]bool, row_length)
	cur_beams[start_column] = true
	next_beams := make([]bool, row_length)
	var splitter_count int
	for _, splitters_row := range splitters {
		for column_index, has_beam := range cur_beams {
			if has_beam {
				if splitters_row[column_index] {
					splitter_count++
					next_beams[column_index - 1] = true
					next_beams[column_index + 1] = true
				} else {
					next_beams[column_index] = true
				}
			}
		}

		clear(cur_beams)
		cur_beams, next_beams = next_beams, cur_beams
	}

	return splitter_count
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to read input file:", err)
		panic(err)
	}

	splitters, start_column := ParseInput(string(input_bytes))
	fmt.Println("Part 1:", Solve1(splitters, start_column))
}
