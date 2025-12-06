package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Problem struct {
	nums []int
	op   byte
}

func parse(filename string) []Problem {
	raw := [][]string{}
	for line := range load.File(filename) {
		raw = append(raw, strings.Fields(line))
	}

	res := []Problem{}
	for i := range raw[0] {
		p := Problem{}
		for _, r := range raw[:len(raw)-1] {
			p.nums = append(p.nums, evil.Atoi(r[i]))
		}
		p.op = raw[len(raw)-1][i][0]

		evil.Assert(p.op == '*' || p.op == '+', "invalid op %q", p.op)

		res = append(res, p)
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	problems := parse(filename)

	// Part 1
	sum := 0
	for _, p := range problems {
		switch p.op {
		case '+':
			sum += ints.Sum(p.nums...)
		case '*':
			sum += ints.Product(p.nums...)
		}
	}
	log.Part1(sum)

	// Part 2
	log.Part2(nil)
}
