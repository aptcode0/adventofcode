package solution

import (
	"strings"
)

func Part2(data string) int {
	inst, graph := parse(data)
	var finalSteps []int
	for node, _ := range graph {
		if strings.HasSuffix(node, "A") {
			steps := findSteps(graph, inst, node, func(node string) bool {
				return strings.HasSuffix(node, "Z")
			})
			finalSteps = append(finalSteps, steps)
		}
	}
	// the input is in a way that target node(one ending with Z) hits in cyclic manner
	// so if 11A reach 11Z in 2 steps then it reach 11Z again after two steps i.e. total 4 steps 
	// if 22A reach 22Z in 3 steps then it will reach 22Z again after 3 steps i.e. total 6 steps
	// So both meet at LCM(2, 3) = 6 (2+2+2(11A 3 cycles), 3+3 (22A 2 cycles))
	return LCM(finalSteps[0], finalSteps[1], finalSteps[2:]...)
}

func LCM(a, b int, nums ...int) int {
	res := a * b / GCD(a, b)

	for _, n := range nums {
		res = LCM(res, n)
	}
	return res
}

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
