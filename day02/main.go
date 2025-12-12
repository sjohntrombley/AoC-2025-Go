package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse_input(input string) [][2]uint64 {
	input = strings.TrimSpace(input)
	var ranges [][2]uint64
	for _, pair_string := range strings.Split(input, ",") {
		dash_split := strings.Split(pair_string, "-");
		if len(dash_split) != 2 {
			log.Fatal(pair_string)
		}
		start, err := strconv.ParseUint(dash_split[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		stop, err := strconv.ParseUint(dash_split[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, [2]uint64{start, stop})
	}
	return ranges
}

func part1(ranges [][2]uint64) uint64 {
	var invalid_sum uint64 = 0
	for _, r := range ranges {
		// TODO: This is very sub-optimal but should be fast enough
		for n := r[0]; n <= r[1]; n++ {
			var digits []uint64 = nil
			m := n
			for m > 0 {
				digits = append(digits, m % 10)
				m /= 10
			}
			l := len(digits)
			if l % 2 == 0 && slices.Equal(digits[:l/2], digits[l/2:]) {
				invalid_sum += n;
			}
		}
	}
	return invalid_sum
}

func part2(ranges [][2]uint64) uint64 {
	var invalid_sum uint64 = 0
	for _, r := range ranges {
		// TODO: Answer is too low
		for n := r[0]; n <= r[1]; n++ {
			var digits []uint64 = nil
			m := n
			for m > 0 {
				digits = append(digits, m % 10)
				m /= 10
			}
			l := len(digits)
			for p := 1; p <= int(math.Ceil(math.Sqrt(float64(l)))); p++ {
				if l % p == 0 {
					group_size := p
					group_count := l / group_size
					if group_count == 1 {
						continue
					}
					first_group := digits[:group_size]
					all_same := true
					for i := 1; i < group_count; i++ {
						cur_group := digits[i * group_size : (i + 1) * group_size]
						if !slices.Equal(first_group, cur_group) {
							all_same = false
							break
						}
					}
					if all_same {
						invalid_sum += n
						break
					}

					if group_size > 1 {
						group_count, group_size = group_size, group_count
						first_group = digits[:group_size]
						all_same = true
						for i := 1; i < group_count; i++ {
							cur_group := digits[i * group_size : (i + 1) * group_size]
							if !slices.Equal(first_group, cur_group) {
								all_same = false
								break
							}
						}
						if all_same {
							invalid_sum += n
							break
						}
					}
				}
			}
		}
	}
	return invalid_sum
}

func main() {
	input_bytes, err := os.ReadFile("input.txt");
	if err != nil {
		log.Fatal(err)
	}
	ranges := parse_input(string(input_bytes))
	fmt.Println("Part 1:", part1(ranges))
	fmt.Println("Part 2:", part2(ranges))
}
