package solution

import (
	"strconv"
	"strings"
)

const CardCount = 5
type Hand struct {
	cards []byte
	bid   int
}

var strength = map[byte]int{
	'A':1, 
	'K':2, 
	'Q':3, 
	'J':4, 
	'T':5, 
	'9':6, 
	'8':7, 
	'7':8, 
	'6':9, 
	'5':10, 
	'4':11, 
	'3':12,
	'2':13,
}

var strengthWithJoker = map[byte]int{
	'A':1, 
	'K':2, 
	'Q':3, 
	'T':4, 
	'9':5, 
	'8':6, 
	'7':7, 
	'6':8, 
	'5':9, 
	'4':10, 
	'3':11,
	'2':12,
	'J':13, 
}
type Kind int
const (
	FiveOfKind Kind = iota
	FourOfKind 
	FullHouse
	ThreeOfKind
	TwoPair
	OnePair
	HighCard
)

type Counts struct {
	uniqueLabels int
	maxSeen int
	jokers int
}

func parse(data string) []Hand {
	lines := strings.Split(data, "\n")
	var hands []Hand
	for _, l := range lines {
		parts := strings.Fields(l)
		bid, _ := strconv.Atoi(parts[1])

		hands = append(hands, Hand{[]byte(parts[0]), bid})
	}
	return hands
}
