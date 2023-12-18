package solution

func Part1(data string) int {
	grid := parse(data)

	loads := make([]int, len(grid[0]))

	n := len(grid)
	for i, _ := range loads {
		loads[i] = n
	}

	sum := 0

	for i, r := range grid {
		for j, v := range r {
			if v == 'O' {
				sum += loads[j] 	
				loads[j]--
			}
			if v == '#' {
				loads[j] = n - i - 1
			}
		}
	}

	return sum
}