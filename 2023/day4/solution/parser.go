package solution

import (
	"strconv"
	"strings"
)

type NumsData struct {
	winningNums []int
	actualNums  []int
}

func parse(data string) []NumsData {
	lines := strings.Split(data, "\n")
	nums := make([]NumsData, len(lines))

	for i, l := range lines {
		onlyNums := strings.Split(l, ": ")[1]
		numStrs := strings.Split(onlyNums, " | ")

		nums[i] = NumsData{getNums(numStrs[0]), getNums(numStrs[1])}
	}
	return nums
}

func getNums(s string) []int {
	var nums []int
	for _, token := range strings.Fields(s) {
		num, _ := strconv.Atoi(token)
		nums = append(nums, num)
	}
	return nums
}
