package solution

import (
	"strings"
)

func Part2(data string) int {
	lines := strings.Split(data, "\n")
	time, distance := join(lines[0]), join(lines[1])

	return findWays(time, distance)
}
