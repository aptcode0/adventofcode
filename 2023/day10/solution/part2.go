package solution

import "strings"

func Part2(data string) int {
	grid := parse(data)
	mainLoop := findMainLoop(grid)
	cols := len(grid[0])
	cnt := 0
	for i, r := range grid {
		for j, _ := range r {
			if mainLoop[i*cols + j] {
				continue
			}
			crossings := findNumOfLeftCrossings(grid, mainLoop, i, j)
			if crossings % 2 != 0 {
				cnt++
			}
		}
	}
	return cnt
}

func findNumOfLeftCrossings(grid [][]byte, mainLoop map[int]bool, i, j int) int {
	cnt := 0
	for k := 0; k < j; k++ {
		if !mainLoop[i * len(grid[0]) + k] { // check pipes on main loop only
			continue
		}
		if crossed := strings.IndexByte("|LJS", grid[i][k]) >= 0; crossed {
			cnt++
		}
	}
	return cnt
}