package solution

import (
	"strconv"
	"strings"
)

func parse(data string) [][]int {
	lines := strings.Split(data, "\n")
	
	var in [][]int
	for _, l := range lines {
		var nums []int
		for _, s := range strings.Fields(l) {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		in = append(in, nums)
	}
	return in
}