package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func sum(i []int) int {
	sum := 0
	for _, n := range i {
		sum += n
	}
	return sum
}

func countIncreases(filename string, window int) int {
	count := 0

	buf := make([]int, window)
	i := 0
	prev := int(^uint(0) >> 1)
	for line := range load.File(filename) {
		n := ints.Atoi(line)
		buf[i%window] = n

		if i >= (window - 1) {
			s := sum(buf)

			if s > prev {
				count++
			}
			prev = s
		}

		i++
	}
	return count
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	// Part 1
	log.Part1(countIncreases(filename, 1))

	// Part 2
	log.Part2(countIncreases(filename, 3))
}
