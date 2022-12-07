package main

import (
	"math"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	ch := load.File(filename)
	defer channel.Drain(ch)
	return evil.Split(<-ch, ",")
}

func cost(pos int, crabs []int) int {
	total := 0
	for _, crab := range crabs {
		total += ints.Abs(crab - pos)
	}
	return total
}

func cost2(pos int, crabs []int) int {
	total := 0
	for _, crab := range crabs {
		d := ints.Abs(crab - pos)
		total += (d * (d + 1)) / 2
	}
	return total
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	crabs := parse(filename)
	minPos := ints.Min(crabs...)
	maxPos := ints.Max(crabs...)

	// Part 1
	min := math.MaxInt
	for pos := minPos; pos <= maxPos; pos++ {
		if c := cost(pos, crabs); c < min {
			min = c
		}
	}
	log.Part1(min)

	// Part 2
	min = math.MaxInt
	for pos := minPos; pos <= maxPos; pos++ {
		if c := cost2(pos, crabs); c < min {
			min = c
		}
	}
	log.Part2(min)
}
