package solution


func Part1(data string) int {
	workflows, ratings := parse(data)
	sum := 0
	for _, r := range ratings {
		if isAccepted(workflows, "in", r) {
			for _, n := range r {
				sum += n
			}
		}
	}
	return sum
}

func isAccepted(workflows map[string][]rule, start string, rating map[byte]int) bool {
	curr := start
	for curr != "A" && curr != "R" {
		rs := workflows[curr]

		for i, r := range rs {
			if  i == len(rs)-1 {
				curr = r.nxt
				break
			}
			if canPass(r.cnd, rating) {
				curr = r.nxt
				break
			}			
		}
	}
	return curr == "A"
}

func canPass(cnd condition, rating map[byte]int) bool {
	exp, actu := cnd.lim, rating[cnd.part]
	if cnd.comp == "<" {
		return actu < exp
	}
	return actu > exp
}