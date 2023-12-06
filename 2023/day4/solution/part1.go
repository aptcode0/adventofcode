package solution

func Part1(data string) int {
	cardsData := parse(data)

	sum := 0
	for _, nums := range cardsData {
		points := calculatePoints(nums)
		sum += points
	}

	return sum
}

func calculatePoints(nums NumsData) int {
	seen := map[int]bool{}
	for _, n := range nums.winningNums {
		seen[n] = true
	}
	points := 0
	for _, n := range nums.actualNums {
		if !seen[n] {
			continue
		}
		if points == 0 {
			points = 1
			continue
		}
		points = 2 * points
	}
	return points
}
