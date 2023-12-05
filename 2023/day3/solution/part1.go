package solution

var adjs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

func Part1(data string) int {
	matrix := parse(data)
	sum := 0
	num := 0
	isNumGood := false
	for i, row := range matrix {
		num, isNumGood = 0, false
		for j, ch := range row {
			if !isDigit(ch) {
				if isNumGood {
					sum += num
				}
				isNumGood = false
				num = 0
				continue
			}
			if isGood(matrix, i, j) {
				isNumGood = true
			}
			num = num*10 + int(ch-'0')
		}
		if num > 0 && isNumGood {
			sum += num
		}
	}
	return sum
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isGood(matrix [][]byte, i, j int) bool {
	for _, adj := range adjs {
		if isValid(matrix, i+adj[0], j+adj[1]) &&
			isSpecial(matrix[i+adj[0]][j+adj[1]]) {
			return true
		}
	}
	return false
}

func isValid(matrix [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0])
}

func isSpecial(ch byte) bool {
	return ch != '.' && !isDigit(ch)
}
