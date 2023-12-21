package solution

import (
	"strconv"
)

func Part2(data string) int64 {
	plan := parse(data)

	var x, y int64
	var grid [][]int64
	var b int64

	for _, p := range plan {
		dir := dirs[findDir(p.color)]
		dist := findDist(p.color)

		x, y = x+int64(dir[0])*dist, y+int64(dir[1])*dist

		grid = append(grid, []int64{x, y})
		b += dist
	}

	//Shoelace formula + Pick's theorem
	return findArea(grid) + b / 2 + 1
}

func findDir(color string) int {
	return int(color[len(color)-1] - '0')
}

func findDist(color string) int64 {
	num, _ := strconv.ParseInt(color[1 : len(color)-1], 16, 64)
	return num
}

func findArea(grid [][]int64) int64 {
	var area int64 = 0
	n := len(grid)
	for i := 0; i < n; i++ {
		j := (i+1) % n
		area += grid[i][0]*grid[j][1] - grid[i][1]*grid[j][0]
	}
	return abs(area) / 2
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}