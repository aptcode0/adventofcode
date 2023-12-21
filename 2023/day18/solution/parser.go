package solution

import (
	"strconv"
	"strings"
)

// 0:>, 1:V, 2:<, 3:^
var dirs = [][]int{{0,1}, {1,0}, {0,-1} , {-1,0}}
var chToDir = map[byte]int{
	'R': 0,
	'D': 1,
	'L': 2,
	'U': 3,
}
type record struct {
	dir int
	dist int
	color string
}

func parse(data string) []record {
	lines := strings.Split(data, "\n")
	var plan []record
	for _, l := range lines {
		f := strings.Fields(l)
		plan = append(plan, record{getDir(f[0]), getNum(f[1]), getColor(f[2])})
	}
	return plan
}

func getDir(s string) int {
	return chToDir[s[0]]
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func getColor(s string) string {
	return s[1:len(s)-1] // remove brackets
}