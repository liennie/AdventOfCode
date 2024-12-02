package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) [][]int {
	var res [][]int
	for line := range load.File(filename) {
		res = append(res, evil.Fields(line))
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	reports := parse(filename)

	// Part 1
	safe := 0
reports:
	for r, report := range reports {
		evil.Assert(len(report) > 1, "report ", r, " has len 1")

		increasing := report[0] < report[1]
		for i := range len(report) - 1 {
			a, b := report[i], report[i+1]
			if diff := ints.Abs(a - b); 1 > diff || diff > 3 {
				continue reports
			}
			if increasing && a > b {
				continue reports
			}
			if !increasing && a < b {
				continue reports
			}
		}
		safe++
	}
	log.Part1(safe)
}
