package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Equation struct {
	target int
	values []int
}

func parse(filename string) []Equation {
	return load.Parse(filename, func(line string) Equation {
		target, values, ok := strings.Cut(line, ":")
		evil.Assert(ok, "Line is missing \":\"")

		eq := Equation{
			target: evil.Atoi(target),
			values: evil.Fields(values),
		}
		evil.Assert(len(eq.values) >= 2, "Equation has too few values")
		return eq
	})
}

func isPossible(target, a, b int, rest ...int) bool {
	if len(rest) == 0 {
		return a*b == target || a+b == target
	}

	if isPossible(target, a*b, rest[0], rest[1:]...) {
		return true
	}
	return isPossible(target, a+b, rest[0], rest[1:]...)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	equations := parse(filename)

	// Part 1
	sum := 0
	for _, eq := range equations {
		if isPossible(eq.target, eq.values[0], eq.values[1], eq.values[2:]...) {
			sum += eq.target
		}
	}
	log.Part1(sum)
}
