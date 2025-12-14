package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parse_input(input string) [][]bool {
	var roll_grid [][]bool
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		cur_row := make([]bool, 0, len(line))
		for _, r := range line {
			switch r {
			case '@':
				cur_row = append(cur_row, true)
			case '.':
				cur_row = append(cur_row, false)
			default:
				log.Fatal(line)
			}
		}
		roll_grid = append(roll_grid, cur_row)
	}

	// Add padding to make solving easier
	for i, row := range roll_grid {
		new_row := make([]bool, len(row) + 2)
		copy(new_row[1 : len(new_row) - 1], row)
		roll_grid[i] = new_row
	}
	
	new_roll_grid := make([][]bool, len(roll_grid) + 2)
	new_roll_grid[0] = make([]bool, len(roll_grid[0]))
	copy(new_roll_grid[1 : len(new_roll_grid) - 1], roll_grid)
	new_roll_grid[len(new_roll_grid) - 1] = make([]bool, len(roll_grid[0]))

	return new_roll_grid
}

func part1(roll_grid [][]bool) uint16 {
	var accessable_count uint16
	for r := 1; r < len(roll_grid) - 1; r++ {
		for c := 1; c < len(roll_grid[r]) - 1; c++ {
			if !roll_grid[r][c] {
				continue
			}

			var adjacent_count uint16
			for cr := r - 1; cr < r + 2; cr++ {
				for cc := c - 1; cc < c + 2; cc++ {
					if cr == r && cc == c {
						continue
					}
					if roll_grid[cr][cc] {
						adjacent_count++
					}
				}
			}

			if adjacent_count < 4 {
				accessable_count++
			}
		}
	}

	return accessable_count
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	roll_grid := parse_input(string(input_bytes))
	fmt.Println("Part 1:", part1(roll_grid))
}
