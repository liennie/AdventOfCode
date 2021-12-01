package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
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
		n := util.Atoi(line)
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
	defer util.Recover(log.Err)

	const filename = "input.txt"

	// Part 1
	log.Part1(countIncreases(filename, 1))

	// Part 2
	log.Part2(countIncreases(filename, 3))
}
