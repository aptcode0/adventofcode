package solution

func Part1(data string) int {
	inst, graph := parse(data)
	return findSteps(graph, inst, "AAA", func (node string) bool { return node == "ZZZ"})
}

func findSteps(graph map[string][]string, inst string, start string,  
	isDest func(node string) bool) int {
	curr := start
	steps := 1

	// Just to be safe, kept some limit instead of infinit for loop
	for steps < 1000000 {
		for _, dir := range inst {
			idx := 0
			if dir == 'R' {
				idx = 1
			}
			curr = graph[curr][idx]
			if isDest(curr) {
				return steps
			}
			steps++
		}
	}
	return 0
}
