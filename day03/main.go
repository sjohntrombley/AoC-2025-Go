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

func part2(banks [][]uint8) uint64 {
	var total_joltage uint64
	for _, bank := range banks {
		digits := [12]uint8(bank[len(bank) - 12:])
		for _, joltage := range slices.Backward(bank[:len(bank) - 12]) {
			for i := range digits {
				if joltage < digits[i] {
					break
				}
				digits[i], joltage = joltage, digits[i]
			}
		}
		var joltage uint64
		for _, d := range digits {
			joltage *= 10
			joltage += uint64(d)
		}
		total_joltage += joltage
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
	fmt.Println("Part 2:", part2(banks))
}
