package main

import (
	"slices"
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

func parse1(filename string) []Problem {
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

func parse2(filename string) []Problem {
	lines := load.Slice(filename)
	maxLine := ints.MaxSeqFunc(func(s string) int { return len(s) }, slices.Values(lines))
	for i := range lines {
		lines[i] = lines[i] + strings.Repeat(" ", maxLine-len(lines[i]))
	}

	res := []Problem{}
	p := Problem{}
	for i := range lines[0] {
		n := 0
		op := byte(0)
		empty := true

		for _, line := range lines[:len(lines)-1] {
			if line[i] == ' ' {
				continue
			}

			evil.Assert('0' <= line[i] && line[i] <= '9', "invalid char %q", line[i])
			n *= 10
			n += int(line[i] - '0')
			empty = false
		}

		if c := lines[len(lines)-1][i]; c != ' ' {
			op = c
		}

		if !empty {
			p.nums = append(p.nums, n)
			if op != 0 {
				p.op = op
			}
		} else {
			evil.Assert(p.op == '*' || p.op == '+', "invalid op %q", p.op)
			res = append(res, p)
			p = Problem{}
		}
	}
	if p.op != 0 {
		res = append(res, p)
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	// Part 1
	problems := parse1(filename)
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
	problems = parse2(filename)
	sum = 0
	for _, p := range problems {
		switch p.op {
		case '+':
			sum += ints.Sum(p.nums...)
		case '*':
			sum += ints.Product(p.nums...)
		}
	}
	log.Part2(sum)
}
