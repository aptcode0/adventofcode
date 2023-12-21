package solution

import "math"


// 0:>, 1:<, 2:V, 3:^
var dirs = [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
var adjs = [][]int{{3, 2}, {2, 3}, {1, 0}, {0, 1}}

type state struct {
	i, j int
	dir int
	steps int
}

func Part1(data string) int {
	grid := parse(data)
	return bfs(grid, 1, 3)
}

func bfs(grid [][]int, minSteps, maxSteps int) int {
	res := math.MaxInt
	r, c := len(grid), len(grid[0])
	q := []state{{0, 0, 0, 0}}
	dists := map[state]int{q[0]: 0}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.i == r-1 && curr.j == c-1 && curr.steps >= minSteps {
			res = min(res, dists[curr])
		}
		var nxtStates []state
		if curr.steps < maxSteps {
			dir := dirs[curr.dir]
			nxt := state{curr.i+dir[0], curr.j+dir[1], curr.dir, curr.steps+1}
			nxtStates = append(nxtStates, nxt)
		} 
		if curr.steps >= minSteps {
			for _, diri := range adjs[curr.dir] {
				dir := dirs[diri]
				nxt := state{curr.i+dir[0], curr.j+dir[1], diri, 1}
				nxtStates = append(nxtStates, nxt)
			}
		}
		
		for _, nxt := range nxtStates {
			if !isValid(nxt.i, nxt.j, r, c) {
				continue
			}
			v := dists[curr] + grid[nxt.i][nxt.j]
			if oldv, exist := dists[nxt]; !exist || oldv > v {
				dists[nxt] = v
				q = append(q, nxt)
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}