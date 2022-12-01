package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
	"golang.org/x/exp/slices"
)

type inventory struct {
	calories []int
}

func parse(filename string) []inventory {
	res := []inventory{}
	add := true
	var last *inventory

	for line := range load.File(filename) {
		if line == "" {
			add = true
			continue
		}
		if add {
			add = false
			res = append(res, inventory{})
			last = &res[len(res)-1]
		}

		last.calories = append(last.calories, util.Atoi(line))
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	invs := parse(filename)

	// Part 1
	sums := []int{}
	for _, inv := range invs {
		sums = append(sums, util.Sum(inv.calories...))
	}
	slices.SortFunc(sums, func(a, b int) bool { return a > b })

	log.Part1(sums[0])

	// Part 2
	log.Part2(util.Sum(sums[:3]...))
}
