package solution

import "fmt"


// dir {0, 1} '/'={-1, 0}, '\'={1, 0}, '|'={-1, 0},{1, 0}, '-'=> {0,1}
// dir {0, -1} '/'={1, 0}, '\'={-1, 0}, '|'={-1, 0},{1, 0}, '-'=> {0,-1}
// dir {1, 0} '/'={0, -1}, '\'={0, 1}, '|'={1,0}, '-'=> {0, -1},{0, 1}
// dir {-1, 0} '/'={0, 1}, '\'={0, -1}, '|'={-1,0}, '-'=> {0, -1},{0, 1}

// 0:>, 1:<, 2:V, 3:^

var dirs = [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
var newMirrDir = map[byte][]int{
	'/': []int{3, 2, 1, 0},
	'\\': []int{2, 3, 0, 1},
} 

func Part1(data string) int {
	grid := parse(data)
	return castRay(grid, 0, 0, 0)
}

func castRay(grid [][]byte, x, y, diri int) int {
	visited := map[string]bool{} // x+y
	seen := map[string]bool{} 	// x+y+dir
	// in queue x, y, dirIdx
	var q [][]int
	for _, d := range newDir(grid[x][y], diri) {
		q = append(q, []int{x, y, d}) 
	}
	
	for len(q) > 0 {
		x, y, diri = q[0][0], q[0][1], q[0][2]
		q = q[1:]
		k := fmt.Sprintf("%d|%d|%d", x, y, diri)
		if seen[k] {
			continue
		}
		seen[k] = true
		visited[fmt.Sprintf("%d|%d", x, y)] = true
		
		newX, newY := x + dirs[diri][0], y + dirs[diri][1]
		if !isValid(newX, newY, len(grid), len(grid[0])) {
			continue
		}

		for _, d := range newDir(grid[newX][newY], diri) {
			q = append(q, []int{newX, newY, d}) 
		}
		
	}
	return len(visited)
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols 
}

func newDir(tile byte, diri int) []int {
	switch tile {
	case '.': 
		return []int{diri}
	case '/':
		return []int{newMirrDir[tile][diri]}
	case '\\':
		return []int{newMirrDir[tile][diri]}
	case '-':
		if diri == 0 || diri == 1 {
			return []int{diri}
		}
		return []int{0, 1}
	case '|':
		if diri == 2 || diri == 3 {
			return []int{diri}
		}
		return []int{2, 3}
	}
	return []int{}
}
