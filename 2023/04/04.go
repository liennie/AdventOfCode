package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Card struct {
	have, winning []int
}

func parse(filename string) map[int]Card {
	res := map[int]Card{}

	for line := range load.File(filename) {
		line = strings.TrimPrefix(line, "Card ")

		sid, line, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "missing colon")
		id := evil.Atoi(strings.TrimSpace(sid))

		have, winning, ok := strings.Cut(line, "|")
		evil.Assert(ok, "missing pipe")

		c := Card{}
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

		res[id] = c
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cards := parse(filename)

	// Part 1
	sum := 0
	for _, card := range cards {
		w := 0
	have:
		for _, have := range card.have {
			for _, win := range card.winning {
				if have == win {
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
	}
	log.Part1(sum)
}
