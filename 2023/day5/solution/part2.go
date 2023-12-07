package solution

import (
	"math"
)

// Slow - TODO: improve time complexity, currently slow
func Part2(data string) int {
	seeds, mappings := parse(data)

	ch := make(chan int)
	for i := 0; i < len(seeds); i += 2 {
		go findRes(mappings, seeds[i], seeds[i+1], ch)
	}

	minRes := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		minRes = min(minRes, <-ch)
	}
	return minRes
}

func findRes(mappings [][]Range, seedStart, seedCount int, ch chan int) {
	minRes := math.MaxInt
	for seed := seedStart; seed < seedStart+seedCount; seed++ {
		res := findLastMapping(seed, mappings)
		minRes = min(minRes, res)
	}
	ch <- minRes
}
