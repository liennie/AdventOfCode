package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) []int {
	ch := load.File(filename)

	defer func() {
		for range <-ch {
		}
	}()

	return util.SliceAtoi(strings.Split(<-ch, ","))
}

func cost(pos int, crabs []int) int {
	total := 0
	for _, crab := range crabs {
		total += util.Abs(crab - pos)
	}
	return total
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	crabs := parse(filename)

	// Part 1
	min := util.MaxInt
	minPos := util.SliceMin(crabs...)
	maxPos := util.SliceMax(crabs...)
	for pos := minPos; pos <= maxPos; pos++ {
		if c := cost(pos, crabs); c < min {
			min = c
		}
	}
	log.Part1(min)
}
