package solution

func Part2(data string) int {
	cardsData := parse(data)

	counts := make([]int, len(cardsData))
	for i, _ := range counts {
		counts[i] = 1
	}
	for i, nums := range cardsData {
		matches := calculateMatches(nums)
		updateCounts(counts, matches, i)
	}

	return sum(counts)
}

func calculateMatches(nums NumsData) int {
	seen := map[int]bool{}
	for _, n := range nums.winningNums {
		seen[n] = true
	}
	matches := 0
	for _, n := range nums.actualNums {
		if !seen[n] {
			continue
		}
		matches++
	}
	return matches
}

func updateCounts(counts []int, matches int, currCard int) {
	for i := currCard + 1; i < len(counts) && i <= currCard+matches; i++ {
		counts[i] += counts[currCard]
	}
}

func sum(arr []int) int {
	tot := 0
	for _, n := range arr {
		tot += n
	}
	return tot
}
