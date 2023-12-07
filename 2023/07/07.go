package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

var cardLabels1 = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var cardLabels2 = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

const (
	unknown = iota
	highCard
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Hand struct {
	cards [5]int
	typ   int
	bid   int
}

func (h Hand) compare(other Hand) int {
	if c := cmp.Compare(h.typ, other.typ); c != 0 {
		return c
	}
	for i := range h.cards {
		if c := cmp.Compare(h.cards[i], other.cards[i]); c != 0 {
			return c
		}
	}
	return 0
}

func getType(cards [5]int, joker bool) int {
	cardCnt := make([]int, len(cardLabels1))
	for _, c := range cards {
		cardCnt[c]++
	}

	have := make([]int, 6)
	haveWithoutJokers := make([]int, 6)
	for i, cnt := range cardCnt {
		have[cnt]++
		if i > 0 {
			haveWithoutJokers[cnt]++
		}
	}

	switch {
	case have[5] == 1 || (joker && haveWithoutJokers[5-cardCnt[0]] >= 1):
		return fiveOfAKind
	case have[4] == 1 || (joker && haveWithoutJokers[4-cardCnt[0]] >= 1):
		return fourOfAKind
	case (have[3] == 1 && have[2] == 1) || (joker && haveWithoutJokers[2] == 2 && cardCnt[0] == 1): // full house is only worth making from a two pair
		return fullHouse
	case have[3] == 1 || (joker && haveWithoutJokers[3-cardCnt[0]] >= 1):
		return threeOfAKind
	case have[2] == 2: // using jokers to make two pairs is a waste
		return twoPair
	case have[2] == 1 || (joker && haveWithoutJokers[2-cardCnt[0]] >= 1):
		return onePair
	default:
		return highCard
	}
}

func parse(filename string, cardLabels []byte, joker bool) []Hand {
	res := []Hand{}

	for line := range load.File(filename) {
		scards, sbid, ok := strings.Cut(line, " ")
		evil.Assert(ok, "missing bid")
		evil.Assert(len(scards) == 5, "invalid number of cards")
		bid := evil.Atoi(sbid)

		var cards [5]int
	cards:
		for i := range cards {
			for v, c := range cardLabels {
				if scards[i] == c {
					cards[i] = v
					continue cards
				}
			}
			evil.Panic("invalid label %q", scards[i])
		}

		res = append(res, Hand{
			cards: cards,
			typ:   getType(cards, joker),
			bid:   bid,
		})
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	// Part 1
	hands := parse(filename, cardLabels1, false)
	slices.SortFunc(hands, func(a, b Hand) int { return a.compare(b) })
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	log.Part1(sum)

	// Part 1
	hands = parse(filename, cardLabels2, true)
	slices.SortFunc(hands, func(a, b Hand) int { return a.compare(b) })
	sum = 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	log.Part2(sum)
}
