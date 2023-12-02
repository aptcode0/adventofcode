package solution

func Part2(data string) int64 {
	games := Parse(data)
	var sum int64
	for _, game := range games {
		var power int64 = 1
		for _, cnts := range game.cubes {
			power *= int64(max(cnts))
		}
		sum += power
	}
	return sum
}

func max(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}
