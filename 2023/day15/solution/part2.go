package solution

import (
	"sort"
	"strconv"
	"strings"
)

type lens struct {
	size int
	pos int
}

type box struct {
	lenses map[string]lens
	cnt int
}

func (b *box) Remove(label string) {
	delete(b.lenses, label)
}

func (b *box) Add(label string, size int) {
	if l, exist := b.lenses[label]; exist {
		b.lenses[label] = lens{size, l.pos}
		return 
	}
	b.lenses[label] = lens{size, b.cnt}
	b.cnt++
}

func Part2(data string) int {
	strs := parse(data)

	boxes := map[int]*box{}
	for _, s := range strs {
		if isRemove := s[len(s)-1] == '-'; isRemove {
			label := s[:len(s)-1]
			if boxes[hash(label)] == nil {
				continue
			}
			boxes[hash(label)].Remove(label)
			continue
		}

		strs := strings.Split(s, "=")
		label, size := strs[0], toNum(strs[1])
		if boxes[hash(label)] == nil {
			boxes[hash(label)] = &box{lenses: map[string]lens{}}
		}
		boxes[hash(label)].Add(label, size)
	}

	
	sum := 0
	for i, b := range boxes {
		lenses := getOrderedLens(b.lenses)
		for j, l := range lenses {
			sum += (i+1) * (j+1) * l.size
		}
	}
	return sum
}

func toNum(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func getOrderedLens(lenses map[string]lens) []lens {
	var arr []lens

	for _, l := range lenses {
		arr = append(arr, l)
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].pos < arr[j].pos })
	return arr
}

