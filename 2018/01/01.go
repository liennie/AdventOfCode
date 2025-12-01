package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	nums := evil.SliceAtoi(load.Slice(filename))

	// Part 1
	log.Part1(ints.Sum(nums...))

	// Part 2
	reached := set.New(0)
	sum := 0
	for {
		for _, n := range nums {
			sum += n
			if reached.Contains(sum) {
				log.Part2(sum)
				return
			}
			reached.Add(sum)
		}
	}
}
