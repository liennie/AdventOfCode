package main

import (
	"slices"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) (left, right []int) {
	for line := range load.File(filename) {
		ns := evil.Fields(line)
		evil.Assert(len(ns) == 2, "len is not 2: ", len(ns))

		left = append(left, ns[0])
		right = append(right, ns[1])
	}

	return
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	left, right := parse(filename)
	evil.Assert(len(left) == len(right), "slices have different len: ", len(left), " != ", len(right))

	// Part 1
	sortedLeft := slices.Sorted(slices.Values(left))
	sortedRight := slices.Sorted(slices.Values(right))

	dist := 0
	for i := range sortedLeft {
		dist += ints.Abs(sortedLeft[i] - sortedRight[i])
	}

	log.Part1(dist)
}
