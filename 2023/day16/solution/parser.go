package solution

import "strings"

func parse(data string) [][]byte {
	lines := strings.Split(data, "\n")
	
	var grid [][]byte
	for _, l := range lines {
		grid = append(grid, []byte(l))
	}
	return grid
}