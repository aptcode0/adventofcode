package solution

type state struct {
	workflow string
	ratings map[byte][]int // min & max of x, m, a, s
}
var categories = []byte{'x', 'm', 'a', 's'}

func Part2(data string) int {
	workflows, _ := parse(data)
	q := []state{{"in", initialRatings()}}
	cnt := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.workflow == "A" {
			cnt += getCount(curr.ratings)	
			continue
		}
		if curr.workflow == "R" {
			continue
		}
		rs := workflows[curr.workflow]

		ratings := curr.ratings
		for i, r := range rs {	
			if  i != len(rs)-1 {
				newR := newRatings(r.cnd, ratings)
				q = append(q, state{r.nxt, newR}) // Split 1...4000 range, one into condition rule pass
				ratings = newRatings(condition{r.cnd.lim, getOpp(r.cnd.comp), r.cnd.part}, ratings) // and keep other rem range in ratings
				continue
			}
			q = append(q, state{r.nxt, ratings})
		}
	}
	return cnt
}

func initialRatings() map[byte][]int {
	ratings := map[byte][]int{}
	for _, v := range categories {
		ratings[v] = []int{1, 4000}
	}
	return ratings
}

func newRatings(cnd condition, ratings map[byte][]int) map[byte][]int {
	newrs := map[byte][]int{}
	for k, r := range ratings {
		newrs[k] = r
	}
	newrs[cnd.part] = newRating(cnd.lim, cnd.comp, ratings[cnd.part])
	return newrs
}

func newRating(lim int, comp string, rng []int) []int {
	newRng := []int{rng[0], rng[1]}
	switch comp {
	case "<":
		newRng[1] = min(rng[1], lim-1)
	case ">":
		newRng[0] = max(rng[0], lim+1)
	case "<=":
		newRng[1] = min(rng[1], lim)	
	case ">=":
		newRng[0] = max(rng[0], lim)	
	}
	return newRng
}

func getCount(ratings map[byte][]int) int {
	cnt := 1
	for _, r := range ratings {
		cnt *= (r[1] - r[0] + 1)
	}
	return cnt
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getOpp(comp string) string {
	if comp == "<" {
		return ">="
	}
	return "<="
}