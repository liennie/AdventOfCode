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

func largestJoltage(n int) func(bank []int) int {
	return func(bank []int) int {
		joltage := 0

		k := 0
		for i := range n {
			max := 0
			for j := k; j < len(bank)-n+1+i; j++ {
				b := bank[j]
				if b > max {
					max = b
					k = j + 1
				}
			}

			joltage *= 10
			joltage += max
		}

		return joltage
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	banks := parse(filename)

	// Part 1
	log.Part1(ints.SumFunc(largestJoltage(2), banks...))

	// Part 2
	log.Part2(ints.SumFunc(largestJoltage(12), banks...))
}
