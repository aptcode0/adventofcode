package solution

import (
	"strconv"
	"strings"
)

type condition struct {
	lim int
	comp string
	part byte
}

type rule struct {
	cnd  condition
	nxt string
}

func parse(data string) (map[string][]rule, []map[byte]int) {
	lines := strings.Split(data, "\n")
	workflows, i := getWorkflows(lines)
	ratings := getRatings(lines[i:])
	return workflows, ratings
}

func getWorkflows(lines []string) (map[string][]rule, int) {
	workflows := map[string][]rule{}
	i := 0
	for len(lines[i]) > 0 {
		l := lines[i]
		idx := strings.IndexByte(l, '{')
		
		nm := l[:idx]
		rs := strings.Split(l[idx+1:len(l)-1], ",")

		workflows[nm] = parseRules(rs)
		i++
	}
	return workflows, i+1
}

func getRatings(lines []string) []map[byte]int {
	rules := []map[byte]int{}
	for _, l := range lines {
		parts := strings.Split(l[1:len(l)-1], ",")
		rule := map[byte]int{}

		for _, p := range parts {
			kv := strings.Split(p, "=")
			rule[kv[0][0]] = toNum(kv[1])
		}
		rules = append(rules, rule)
	}
	return rules
}

func parseRules(rs []string) []rule {
	var rules []rule
	for i := 0; i < len(rs)-1; i++ {
		r := rs[i]
		parts := strings.Split(r, ":")
		nxt, cnd := parts[1], parseCondition(parts[0])
		rules = append(rules, rule{cnd, nxt})
	}
	rules = append(rules, rule{nxt:strings.Split(rs[len(rs)-1], ":")[0]}) 
	return rules
}

func parseCondition(s string) condition {
	return condition{part: s[0], comp: s[1:2], lim: toNum(s[2:])}
}

func toNum(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}