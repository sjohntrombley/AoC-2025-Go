package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func parse_input(input string) [][]uint8 {
	var banks [][]uint8
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		bank := make([]uint8, 0, len(line))
		for _, d := range line {
			bank = append(bank, uint8(d - '0'))
		}
		banks = append(banks, bank)
	}
	return banks
}

func part1(banks [][]uint8) uint16 {
	var total_joltage uint16
	for _, bank := range banks {
		tens := bank[len(bank) - 2]
		ones := bank[len(bank) - 1]
		for _, joltage := range slices.Backward(bank[:len(bank) - 2]) {
			if joltage >= tens {
				if tens > ones {
					ones = tens
				}
				tens = joltage
			}
		}
		total_joltage += uint16(10 * tens + ones)
	}
	return total_joltage
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	banks := parse_input(string(input_bytes))
	fmt.Println("Part 1:", part1(banks))
}
