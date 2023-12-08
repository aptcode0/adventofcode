package solution

import (
	"sort"
)

func Part1(data string) int {
	hands := parse(data)
	sortHands(hands, findKind, strength)
	return calculateWinnings(hands)
}

func sortHands(hands []Hand, findCardKind func([]byte) Kind, strengthMap map[byte]int) {
	sort.Slice(hands, func(i, j int) bool {
		kindOne := findCardKind(hands[i].cards)
		kindTwo := findCardKind(hands[j].cards)
		if kindOne == kindTwo {
			return CompareByCardStrength(hands[i].cards, hands[j].cards, strengthMap)
		}
		return kindOne > kindTwo
	})
}

func CompareByCardStrength(cards1, cards2 []byte, strengthMap map[byte]int) bool {
	for i := 0; i < CardCount; i++ {
		if cards1[i] != cards2[i] {
			return strengthMap[cards1[i]] >= strengthMap[cards2[i]]
		}
	}
	return true
}

func findKind(cards []byte) Kind {
	types := map[byte]int{}

	for _, c := range cards {
		types[c]++
	}

	uniqueLabels := len(types)
	maxCnt := max(types)

	if uniqueLabels == 1 {
		return FiveOfKind
	}
	if allDifferent := uniqueLabels == CardCount; allDifferent {
		return HighCard
	}

	if onlyOnePair := uniqueLabels == CardCount-1; onlyOnePair {
		return OnePair
	}

	if oneDifferent := maxCnt == CardCount-1; oneDifferent {
		return FourOfKind
	}

	if uniqueLabels == 2 && maxCnt == CardCount-2 {
		return FullHouse
	}

	if maxCnt == CardCount-2 {
		return ThreeOfKind
	}

	return TwoPair
}

func calculateWinnings(hands []Hand) int {
	ans := 0
	for i, h := range hands {
		ans += h.bid * (i + 1)
	}
	return ans
}

func max(m map[byte]int) int {
	mx := 0
	for _, v := range m {
		if v > mx {
			mx = v
		}
	}
	return mx
}
