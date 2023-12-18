package solution

import "fmt"

func Part2(data string) int {
	grid := parse(data)
	seen := map[string]int{}

	end := 1000000000
	for t := 0; t < end; t++ {
		grid = cycle(grid)
		key := hash(grid)
		if oldT, exist := seen[key]; exist {
			l := t - oldT
			rem := end - t
			times := rem / l
			t += l * times
		} 
		seen[key] = t
	}

	return score(grid)
}

func cycle(grid [][]byte) [][]byte {
	for i := 0; i < 4; i++ {
		rollNorth(grid)
		grid = rotate(grid)
	}
	return grid
}

func hash(grid [][]byte) string {
	return fmt.Sprintf("%v", grid)
}

func score(grid [][]byte) int {
	s := 0
	n := len(grid)
	for i, r := range grid {
		for _, v := range r {
			if v == 'O' {
				s += n - i
			}
		}
	}
	return s
}

func rollNorth(grid [][]byte) {
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			if grid[r][c] != 'O' {
				continue
			}
			for i := r-1; i >= 0; i-- {
				if grid[i][c] != '.' {
					break
				}
				grid[i][c], grid[i+1][c] = 'O', '.'
			}
		}
	}
}

func rotate(grid [][]byte) [][]byte {
	n, m := len(grid), len(grid[0])
	newG := make([][]byte, m)

	for i, _ := range newG {
		newG[i] = make([]byte, n)
	}

	for i, r := range grid {
		for j, v := range r {
			newG[j][n-i-1] = v
		}
	}
	return newG
}