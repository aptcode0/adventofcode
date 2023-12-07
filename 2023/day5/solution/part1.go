package solution

import (
	"math"
)

func Part1(data string) int {
	seeds, mappings := parse(data)

	minRes := math.MaxInt
	for _, seed := range seeds {
		res := findLastMapping(seed, mappings)
		minRes = min(minRes, res)
	}
	return minRes
}

func findLastMapping(seed int, mappings [][]Range) int {
	currVal := seed
	for _, ranges := range mappings {
		r := findRange(currVal, ranges)
		if r != nil {
			currVal = r.destStart + (currVal - r.sourceStart)
		}
	}

	return currVal
}

func findRange(curr int, ranges []Range) *Range {
	for _, r := range ranges {
		if r.sourceStart <= curr && curr < r.sourceStart+r.count {
			return &r
		}
	}
	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
