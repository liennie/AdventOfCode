package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Card struct {
	have, winning []int
	amt           int
}

func parse(filename string) []Card {
	res := []Card{}

	for line := range load.File(filename) {
		_, line, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "missing colon")

		have, winning, ok := strings.Cut(line, "|")
		evil.Assert(ok, "missing pipe")

		c := Card{
			amt: 1,
		}
		for _, h := range strings.Split(have, " ") {
			if h == "" {
				continue
			}
			c.have = append(c.have, evil.Atoi(h))
		}
		for _, w := range strings.Split(winning, " ") {
			if w == "" {
				continue
			}
			c.winning = append(c.winning, evil.Atoi(w))
		}

		res = append(res, c)
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cards := parse(filename)

	// Part 1
	sum := 0
	for i, card := range cards {
		w := 0
		match := 0
	have:
		for _, have := range card.have {
			for _, win := range card.winning {
				if have == win {
					match++

					if w == 0 {
						w = 1
					} else {
						w *= 2
					}

					continue have
				}
			}
		}
		sum += w

		for j := i + 1; j <= i+match; j++ {
			cards[j].amt += card.amt
		}
	}
	log.Part1(sum)
	log.Part2(ints.SumFunc(func(c Card) int { return c.amt }, cards...))
}
