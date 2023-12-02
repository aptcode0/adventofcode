package solution

func Part1(data string) int {
	games := Parse(data)

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for _, game := range games {
		valid := true
		for color, cnts := range game.cubes {
			if anyMax(cnts, maxCubes[color]) {
				valid = false
				break
			}
		}
		if valid {
			sum += game.id
		}
	}
	return sum
}

func anyMax(nums []int, t int) bool {
	for _, n := range nums {
		if n > t {
			return true
		}
	}
	return false
}
