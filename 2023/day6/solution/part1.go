package solution

import (
	"fmt"
)

func Part1(data string) int {
	times, distances := parse(data)
	ans := 1
	fmt.Println(times, distances)
	for i := 0; i < len(times); i++ {
		ans *= findWays(times[i], distances[i])
	}
	return ans
}

func findWays(time, dist int) int {
	l, r := 0, time

	for l < r {
		speed := l + (r-l)/2
		actualDist := calcDistance(time, speed)
		if actualDist > dist {
			r = speed
		} else {
			l = speed + 1
		}
	}

	return time - 2*r + 1
}

func calcDistance(totalTime int, speed int) int {
	return (totalTime - speed) * speed
}
