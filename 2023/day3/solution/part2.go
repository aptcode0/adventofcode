package solution

import (
	"strconv"
)

func Part2(data string) int {
	matrix := parse(data)
	sum := 0

	for i, row := range matrix {
		for j, ch := range row {
			if ch != '*' {
				continue
			}
			nums := findAdjDigits(matrix, i, j)
			if len(nums) != 2 {
				continue
			}
			sum += nums[0] * nums[1]
		}
	}
	return sum
}

func findAdjDigits(matrix [][]byte, i, j int) []int {
	var nums []int
	visited := initVisited(len(matrix), len(matrix[i]))
	for _, adj := range adjs {
		adjx, adjy := i+adj[0], j+adj[1]

		if !isValid(matrix, adjx, adjy) ||
			!isDigit(matrix[adjx][adjy]) ||
			visited[adjx][adjy] {
			continue
		}
		nums = append(nums, findNumber(matrix, i+adj[0], j+adj[1], visited))
		if len(nums) > 1 {
			return nums
		}
	}
	return nums
}

func initVisited(rows, cols int) [][]bool {
	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}
	return visited
}

func findNumber(matrix [][]byte, i, j int, visited [][]bool) int {
	start, end := findRange(matrix[i], j)
	for col := start; col <= end; col++ {
		visited[i][col] = true
	}
	num, _ := strconv.Atoi(string(matrix[i][start : end+1]))
	return num
}

func findRange(row []byte, curr int) (int, int) {
	start, end := curr, curr
	for start-1 >= 0 && isDigit(row[start-1]) {
		start--
	}
	for end+1 < len(row) && isDigit(row[end+1]) {
		end++
	}
	return start, end
}
