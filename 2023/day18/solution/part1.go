package solution

type point struct {
	x, y int
}
func Part1(data string) int {
	plan := parse(data)

	grid := map[point]bool{}
	x, y := 0, 0
	grid[point{x, y}] = true

	for _, p := range plan {
		for i := 0; i < p.dist; i++ {
			d := dirs[p.dir]
			grid[point{x+d[0], y+d[1]}] = true
			x, y = x+d[0], y+d[1]
		} 
	}
	expand(grid, point{1, 1})
	return len(grid)
}

func expand(grid map[point]bool, curr point) {
	grid[curr] = true
	for _, d := range dirs {
		p := point{curr.x+d[0], curr.y+d[1]}
		if !grid[p] {
			expand(grid, p)
		}
	}
}
