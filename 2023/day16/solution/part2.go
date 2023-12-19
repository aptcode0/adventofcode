package solution

func Part2(data string) int {
	grid := parse(data)

	var maxRes int
	// > & <
	for i := 0; i < len(grid); i++ {
		maxRes = max(maxRes, castRay(grid, i, 0, 0))
		maxRes = max(maxRes, castRay(grid, i, len(grid[0])-1, 1))
	}
	// V & ^
	for i := 0; i < len(grid[0]); i++ {
		maxRes = max(maxRes, castRay(grid, 0, i, 2))
		maxRes = max(maxRes, castRay(grid, len(grid)-1, i, 3))
	}
	return maxRes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}