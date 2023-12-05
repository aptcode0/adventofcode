package solution

import "strings"

func parse(data string) [][]byte {
	lines := strings.Split(data, "\n")
	matrix := make([][]byte, len(lines))
	for i, l := range lines {
		matrix[i] = make([]byte, len(l))
		for j, ch := range l {
			matrix[i][j] = byte(ch)
		}
	}
	return matrix
}