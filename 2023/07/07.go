package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

var cardLabels = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

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

func getType(cards [5]int) int {
	cardCnt := make([]int, len(cardLabels))
	for _, c := range cards {
		cardCnt[c]++
	}

	have := make([]int, 6)
	for _, cnt := range cardCnt {
		have[cnt]++
	}

	switch {
	case have[5] == 1:
		return fiveOfAKind
	case have[4] == 1:
		return fourOfAKind
	case have[3] == 1 && have[2] == 1:
		return fullHouse
	case have[3] == 1:
		return threeOfAKind
	case have[2] == 2:
		return twoPair
	case have[2] == 1:
		return onePair
	default:
		return highCard
	}
}

func parse(filename string) []Hand {
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
			typ:   getType(cards),
			bid:   bid,
		})
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	hands := parse(filename)

	// Part 1
	slices.SortFunc(hands, func(a, b Hand) int { return a.compare(b) })
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	log.Part1(sum)
}
