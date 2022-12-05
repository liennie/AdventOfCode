package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type assignment struct {
	min, max int
}

func (a assignment) contains(other assignment) bool {
	return a.min <= other.min && a.max >= other.max
}

func (a assignment) overlaps(other assignment) bool {
	return a.max >= other.min && a.min <= other.max
}

func parse(filename string) [][2]assignment {
	res := [][2]assignment{}
	for line := range load.File(filename) {
		a, b, _ := strings.Cut(line, ",")

		aa := util.SplitN(a, "-", 2)
		if len(aa) != 2 || aa[0] > aa[1] {
			util.Panic("Invalid assignment %s", a)
		}

		ba := util.SplitN(b, "-", 2)
		if len(ba) != 2 || ba[0] > ba[1] {
			util.Panic("Invalid assignment %s", b)
		}

		res = append(res, [2]assignment{
			{min: aa[0], max: aa[1]},
			{min: ba[0], max: ba[1]},
		})
	}
	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	assignments := parse(filename)

	// Part 1
	count := 0
	for _, ass := range assignments {
		if ass[0].contains(ass[1]) || ass[1].contains(ass[0]) {
			count++
		}
	}
	log.Part1(count)

	// Part 1
	count = 0
	for _, ass := range assignments {
		if ass[0].overlaps(ass[1]) {
			count++
		}
	}
	log.Part2(count)
}
