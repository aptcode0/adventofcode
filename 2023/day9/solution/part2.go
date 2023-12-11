package solution

func Part2(data string) int {
	matrix := parse(data)

	sum := 0
	for _, nums := range matrix {
		sum += findPrevNum(nums)
	}
	return sum
}

func findPrevNum(nums []int) int {
	firsts := []int{nums[0]}
	for {
		var diffs []int
		allZero := true
		for i := 0; i < len(nums)-1; i++ {
			diff := nums[i+1] - nums[i]
			if diff != 0 {
				allZero = false
			}
			diffs = append(diffs, diff)
		}
		if allZero {
			break
		}
		nums = diffs
		firsts = append(firsts, diffs[0])
	}

	res := 0
	for i := len(firsts) - 1; i >= 0; i-- {
		res = firsts[i] - res
	}

	return res
}
