package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) [][]int {
	res := [][]int{}
	for line := range load.File(filename) {
		res = append(res, evil.Split(line, ""))
	}
	return res
}

func largestJoltage(bank []int) int {
	maxFirst := 0
	maxPos := -1
	for i, n := range bank[:len(bank)-1] {
		if n > maxFirst {
			maxFirst = n
			maxPos = i
		}
	}

	maxSecond := 0
	for _, n := range bank[maxPos+1:] {
		if n > maxSecond {
			maxSecond = n
		}
	}

	return 10*maxFirst + maxSecond
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	banks := parse(filename)

	// Part 1
	log.Part1(ints.SumFunc(largestJoltage, banks...))
}
