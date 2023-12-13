package solution

func Part1(data string) int {
	galaxies := parse(data, 2)
	return findAllPairDist(galaxies)
}

func findAllPairDist(positions [][]int) int {
	sum := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			g1, g2 := positions[i], positions[j]
			dist := abs(g1[0]-g2[0]) + abs(g1[1]-g2[1])
			sum += dist
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
