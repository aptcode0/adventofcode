package solution

func Part2(data string) int {
	galaxies := parse(data, 1000000)
	return findAllPairDist(galaxies)
}
