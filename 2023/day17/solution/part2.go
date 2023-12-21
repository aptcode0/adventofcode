package solution

func Part2(data string) int {
	grid := parse(data)
	return bfs(grid, 4, 10)
}
