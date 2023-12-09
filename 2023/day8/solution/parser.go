package solution

import "strings"

func parse(data string) (inst string, graph map[string][]string) {
	lines := strings.Split(data, "\n")

	g := map[string][]string{}
	for _, l := range lines[2:] {  // l e.g. AAA = (BBB, CCC)
		nodes := strings.Split(l, " = ")	// ["AAA", "(BBB, CCC)"]
		nxtNodes := strings.Split(nodes[1], ", ")	// ["(BBB", "CCC)"]
		currNode := nodes[0]
		// Exclude brackets '(' & ')'
		g[currNode] = []string{nxtNodes[0][1:], nxtNodes[1][:len(nxtNodes[1])-1]}
	}
	return lines[0], g
}