package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

type Computers map[string]set.Set[string]

func parse(filename string) Computers {
	res := Computers{}
	for line := range load.File(filename) {
		left, right, ok := strings.Cut(line, "-")
		evil.Assert(ok, "invalid line format")

		if _, ok := res[left]; !ok {
			res[left] = set.New[string]()
		}
		res[left].Add(right)

		if _, ok := res[right]; !ok {
			res[right] = set.New[string]()
		}
		res[right].Add(left)
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	computers := parse(filename)

	// Part 1
	total := 0
	for computer, connections := range computers {
		for left := range connections.All() {
			for right := range connections.All() {
				if left == right {
					continue
				}

				if computers[left].Contains(right) {
					if strings.HasPrefix(computer, "t") || strings.HasPrefix(left, "t") || strings.HasPrefix(right, "t") {
						total++
					}
				}
			}
		}
	}
	total /= 6 // we count each triplet six times
	log.Part1(total)
}
