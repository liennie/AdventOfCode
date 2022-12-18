package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"golang.org/x/exp/slices"
)

type inventory struct {
	calories []int
}

func parse(filename string) []inventory {
	res := []inventory{}

	for block := range load.Blocks(filename) {
		inv := inventory{}

		for line := range block {
			inv.calories = append(inv.calories, evil.Atoi(line))
		}

		res = append(res, inv)
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	invs := parse(filename)

	// Part 1
	sums := []int{}
	for _, inv := range invs {
		sums = append(sums, ints.Sum(inv.calories...))
	}
	slices.SortFunc(sums, func(a, b int) bool { return a > b })

	log.Part1(sums[0])

	// Part 2
	log.Part2(ints.Sum(sums[:3]...))
}
