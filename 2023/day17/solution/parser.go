package solution

import "strings"

func parse(data string) [][]int {
	lines := strings.Split(data, "\n")
	
	grid := make([][]int, len(lines))
	for i, l := range lines {
		grid[i] = make([]int, len(l))
		for j, v := range l {
			grid[i][j] = int(v - '0') 
		}
	}
	return grid
}