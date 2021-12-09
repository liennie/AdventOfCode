package main

import (
	"math"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) []int {
	ch := load.File(filename)
	defer util.Drain(ch)
	return util.Split(<-ch, ",")
}

func cost(pos int, crabs []int) int {
	total := 0
	for _, crab := range crabs {
		total += util.Abs(crab - pos)
	}
	return total
}

func cost2(pos int, crabs []int) int {
	total := 0
	for _, crab := range crabs {
		d := util.Abs(crab - pos)
		total += (d * (d + 1)) / 2
	}
	return total
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	crabs := parse(filename)
	minPos := util.SliceMin(crabs...)
	maxPos := util.SliceMax(crabs...)

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
