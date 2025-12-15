package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FreshRange struct {
	start int64
	end   int64
}

func FreshRangeCmp(fr0, fr1 FreshRange) int {
	return int(fr0.start - fr1.start)
}

type FreshRangeSet struct {
	ranges []FreshRange
}

func (frs *FreshRangeSet) Add(fr FreshRange) {
	i, _ := slices.BinarySearchFunc(
		frs.ranges,
		fr,
		// This comparison function ensures that i is either the index that fr
		// should be merged with or the index where fr should be inserted if
		// it should not be merged.
		func(cur_range, new_range FreshRange) int {
			return int(cur_range.end - new_range.start + 1)
		},
	)
	if i == len(frs.ranges) || fr.end+1 < frs.ranges[i].start {
		frs.ranges = slices.Insert(frs.ranges, i, fr)
	} else {
		if fr.start < frs.ranges[i].start {
			frs.ranges[i].start = fr.start
		}
		if fr.end > frs.ranges[i].end {
			frs.ranges[i].end = fr.end
			for i+1 < len(frs.ranges) &&
				frs.ranges[i].end+1 >= frs.ranges[i+1].start {
				if frs.ranges[i+1].end > frs.ranges[i].end {
					frs.ranges[i].end = frs.ranges[i+1].end
				}
				frs.ranges = slices.Delete(frs.ranges, i+1, i+2)
			}
		}
	}
}

func (frs *FreshRangeSet) Contains(ingredient int64) bool {
	i, _ := slices.BinarySearchFunc(
		frs.ranges,
		ingredient,
		func(fr FreshRange, ingredient int64) int {
			return int(fr.end - ingredient)
		},
	)
	return i < len(frs.ranges) && ingredient >= frs.ranges[i].start
}

func parse_input(input string) (FreshRangeSet, []int64) {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) != 2 {
		log.Fatal("Too many blocks in input")
	}

	ranges_string := strings.TrimSpace(blocks[0])
	ingredient_strings := strings.TrimSpace(blocks[1])

	var frs FreshRangeSet
	for line := range strings.Lines(ranges_string) {
		line = strings.TrimSpace(line)
		range_strings := strings.Split(line, "-")
		if len(range_strings) != 2 {
			log.Fatal("Invalid Fresh Range:", line)
		}

		start, err := strconv.ParseInt(range_strings[0], 10, 64)
		if err != nil {
			log.Fatal("Invalid start of range:", line)
		}
		end, err := strconv.ParseInt(range_strings[1], 10, 64)
		if err != nil {
			log.Fatal("Invalid end of range:", line)
		}

		frs.Add(FreshRange{start: start, end: end})
	}

	var ingredients []int64
	for line := range strings.Lines(ingredient_strings) {
		line = strings.TrimSpace(line)
		ingredient, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal("Invalid ingredient:", line)
		}
		ingredients = append(ingredients, ingredient)
	}

	return frs, ingredients
}

func part1(fresh_ranges FreshRangeSet, ingredients []int64) uint16 {
	var fresh_count uint16
	for _, ingredient := range ingredients {
		if fresh_ranges.Contains(ingredient) {
			fresh_count += 1
		}
	}

	return fresh_count
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error Reading Input:", err)
	}
	fresh_ranges, ingredients := parse_input(string(input_bytes))
	fmt.Println("Part 1:", part1(fresh_ranges, ingredients))
}
