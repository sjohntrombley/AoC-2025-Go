package main;

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input string) int {
	position := 50
	zero_count := 0
	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l);
		n, err := strconv.Atoi(l[1:]);
		if err != nil {
			log.Fatal(err);
		}
		n = n % 100;

		switch l[0] {
		case 'L':
			if n > position {
				position += 100 - n;
			} else {
				position -= n;
			}
		case 'R':
			position += n;
			position %= 100;
		default:
			log.Fatal(l[0])
		}

		if position == 0 {
			zero_count += 1;
		}
	}

	return zero_count
}

func part2(input string) int {
	position := 50
	zero_count := 0
	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l);
		n, err := strconv.Atoi(l[1:]);
		if err != nil {
			log.Fatal(err);
		}

		if n >= 100 {
			zero_count += n / 100;
			n = n % 100;
		}

		switch l[0] {
		case 'L':
			if position <= n {
				if position != 0 {
					zero_count += 1;
				}
				position -= n;
				if position < 0 {
					position += 100;
				}
			} else {
				position -= n;
			}
		case 'R':
			position += n;
			if position >= 100 {
				position -= 100;
				zero_count += 1;
			}
		default:
			log.Fatal(l[0])
		}
	}

	return zero_count
}

func main() {
	input_bytes, err := os.ReadFile("input.txt");
	if err != nil {
		log.Fatal(err);
	}
	input := string(input_bytes)

	fmt.Println("Part 1:", part1(input));
	fmt.Println("Part 2:", part2(input));
}
