package solution

import (
	"strconv"
	"strings"
)

type Range struct {
	destStart   int
	sourceStart int
	count       int
}

func parse(data string) ([]int, [][]Range) {
	lines := strings.Split(data, "\n\n")
	seeds := findSeedNums(lines[0])
	var mappings [][]Range
	for _, l := range lines[1:] {
		mappings = append(mappings, getMapping(l))
	}

	return seeds, mappings
}

func findSeedNums(input string) []int {
	seedData := strings.Split(input, ": ")[1]
	var seeds []int
	for _, s := range strings.Fields(seedData) {
		seedNum, _ := strconv.Atoi(s)
		seeds = append(seeds, seedNum)
	}
	return seeds
}

func getMapping(input string) []Range {
	var ranges []Range
	for _, s := range strings.Split(input, "\n")[1:] {
		r := strings.Fields(s)
		destStart, _ := strconv.Atoi(r[0])
		sourceStart, _ := strconv.Atoi(r[1])
		count, _ := strconv.Atoi(r[2])
		ranges = append(ranges, Range{destStart, sourceStart, count})
	}

	return ranges
}
