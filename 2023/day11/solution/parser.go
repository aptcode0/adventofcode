package solution

import (
	"strings"
)

func parse(data string, size int) [][]int {
	lines := strings.Split(data, "\n")

	cntInRows := make([]int, len(lines))
	cntInCols := make([]int, len(lines[0]))
	for r, l := range lines {
		for c, v := range l {
			if v == '#' {
				cntInRows[r]++
				cntInCols[c]++
			}
		}
	}

	shiftRow := make([]int, len(cntInRows))
	shiftCol := make([]int, len(cntInCols))

	for i, cnt := range cntInRows {
		if cnt == 0 {
			shiftRow[i] += size - 1 // exclude 1 as one empty galaxy row already exists
		}
	}
	for i, cnt := range cntInCols {
		if cnt == 0 {
			shiftCol[i] += size - 1
		}
	}
	for i := 1; i < len(shiftRow); i++ {
		shiftRow[i] += shiftRow[i-1]
	}
	for i := 1; i < len(shiftCol); i++ {
		shiftCol[i] += shiftCol[i-1]
	}

	var galaxiesPos [][]int
	for r, l := range lines {
		var pos [][]int
		for c, v := range l {
			if v == '#' {
				pos = append(pos, []int{r + shiftRow[r], c + shiftCol[c]})
			}
		}

		galaxiesPos = append(galaxiesPos, pos...)
	}
	return galaxiesPos
}
