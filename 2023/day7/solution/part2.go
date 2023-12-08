package solution

import "fmt"

var upgrade = map[Kind]map[int]Kind{
	FourOfKind:  map[int]Kind{1: FiveOfKind, 4: FiveOfKind},
	FullHouse:   map[int]Kind{2: FiveOfKind, 3: FiveOfKind},
	ThreeOfKind: map[int]Kind{3: FourOfKind, 1: FourOfKind},
	TwoPair:     map[int]Kind{2: FourOfKind, 1: FullHouse},
	OnePair:     map[int]Kind{2: ThreeOfKind, 1: ThreeOfKind},
	HighCard:    map[int]Kind{1: OnePair},
}

func Part2(data string) int {
	hands := parse(data)
	sortHands(hands, FindKindWithJoker, strengthWithJoker)
	return calculateWinnings(hands)
}

func FindKindWithJoker(cards []byte) Kind {
	actualKind := findKind(cards)

	types := map[byte]int{}

	for _, c := range cards {
		types[c]++
	}

	if newKind, exist := upgrade[actualKind][types['J']]; exist {
		return newKind
	}

	return actualKind
}
