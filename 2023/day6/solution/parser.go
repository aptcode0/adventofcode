package solution

import (
	"strconv"
	"strings"
)

func parse(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")
	var times []int
	for _, s := range strings.Fields(lines[0])[1:] {
		n, _ := strconv.Atoi(s)
		times = append(times, n)
	}

	var distances []int
	for _, s := range strings.Fields(lines[1])[1:] {
		n, _ := strconv.Atoi(s)
		distances = append(distances, n)
	}

	return times, distances
}

func join(input string) int {
	num, _ := strconv.Atoi(strings.Join(strings.Fields(input)[1:], ""))
	return num
}
