package main

import (
	"strconv"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return evil.Fields(load.Line(filename))
}

func blink(stones []int) []int {
	for i := len(stones) - 1; i >= 0; i-- {
		if stones[i] == 0 {
			stones[i] = 1
		} else if s := strconv.Itoa(stones[i]); len(s)%2 == 0 {
			p := ints.Pow(10, len(s)/2)

			// in part 1 order doesn't matter
			stones = append(stones, stones[i]%p)
			stones[i] /= p
		} else {
			stones[i] *= 2024
		}
	}
	return stones
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	stones := parse(filename)

	// Part 1
	blinks := 25
	for range blinks {
		stones = blink(stones)
	}
	log.Part1(len(stones))
}
