package solution

import (
	"strconv"
	"strings"
)

func parse(data string) ([]string, [][]int) {
	lines := strings.Split(data, "\n")

	records := make([]string, len(lines))
	damagedCnts := make([][]int, len(lines))

	for i, l := range lines {
		fields := strings.Fields(l)
		records[i] = fields[0]

		damagedCnts[i] = parseInts(fields[1])
	}
	return records, damagedCnts
}

func parseInts(str string) []int {
	var nums []int
	for _, s := range strings.Split(str, ",") {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	return nums
}