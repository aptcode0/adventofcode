package solution

func Part1(data string) int {
	matrix := parse(data)

	sum := 0
	for _, nums := range matrix {
		sum += findNextNum(nums)
	}
	return sum
}

func findNextNum(nums []int) int {
	lasts := []int{nums[len(nums)-1]}
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
		lasts = append(lasts, diffs[len(diffs)-1])
	}

	res := 0
	for i := len(lasts) - 1; i >= 0; i-- {
		res = res + lasts[i]
	}

	return res
}
