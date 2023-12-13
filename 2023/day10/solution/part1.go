package solution

import "strings"

// Sequence - UP, DOWN, LEFT, RIGHT
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var allowedPipe = []string{"S|JL", "S|F7", "S-J7", "S-LF"}
var nxtValidPipe = []string{"|7F", "|JL", "-LF", "-7J"}

func Part1(data string) int {
	grid := parse(data)
	
	return len(findMainLoop(grid)) / 2
}

func findMainLoop(grid [][]byte) map[int]bool {
	cols := len(grid[0])

	start := findStart(grid)
	seen := map[int]bool{}
	q := [][]int{start} // enqueue start

	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:] //dequeue
		seen[i*cols+j] = true
		for k, dir := range dirs {
			if isValid(grid, i, j, k, dir, seen) {
				q = append(q, []int{i + dir[0], j + dir[1]}) // enqueue
			}
		}
	}
	return seen
}

func findStart(grid [][]byte) []int {
	for i, r := range grid {
		for j, c := range r {
			if c == 'S' {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func isValid(grid [][]byte, i, j, k int, dir []int, seen map[int]bool) bool {
	nxti, nxtj := i+dir[0], j+dir[1]
	if nxti < 0 || nxtj < 0 || nxti >= len(grid) || nxtj >= len(grid[0]) ||
		seen[nxti*len(grid[0])+nxtj] {
		return false
	}
	if strings.IndexByte(allowedPipe[k], grid[i][j]) < 0 {
		return false
	}
	return strings.IndexByte(nxtValidPipe[k], grid[nxti][nxtj]) >= 0
}
