package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
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
	max := 0
	for _, inv := range invs {
		max = util.Max(max, util.Sum(inv.calories...))
	}
	log.Part1(max)
}
