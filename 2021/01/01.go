package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func countIncreases(filename string) int {
	count := 0

	prev := int(^uint(0) >> 1)
	for line := range load.File(filename) {
		n := util.Atoi(line)
		if n > prev {
			count++
		}
		prev = n
	}
	return count
}

func main() {
	defer util.Recover(log.Err)

	// Part 1
	log.Part1(countIncreases("input.txt"))
}
