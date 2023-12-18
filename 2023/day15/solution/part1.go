package solution

func Part1(data string) int {
	strs := parse(data)

	sum := 0

	for _, s := range strs {
		sum += hash(s)
	}
	return sum
}

func hash(s string) int {
	c := 0
	for _, v := range []byte(s) {
		c += int(v)
		c *= 17
		c %= 256
	}
	return c
}