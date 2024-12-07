package main

import (
	"strconv"
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

type Eval struct {
	operators []func(int, int) int
}

func (e Eval) isPossible(target, a, b int, rest ...int) bool {
	if a > target {
		return false
	}

	if len(rest) == 0 {
		for _, op := range e.operators {
			if op(a, b) == target {
				return true
			}
		}
		return false
	}

	for _, op := range e.operators {
		if e.isPossible(target, op(a, b), rest[0], rest[1:]...) {
			return true
		}
	}
	return false
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	equations := parse(filename)

	// Part 1
	e := Eval{
		operators: []func(int, int) int{
			func(a, b int) int { return a * b },
			func(a, b int) int { return a + b },
		},
	}
	sum := 0
	for _, eq := range equations {
		if e.isPossible(eq.target, eq.values[0], eq.values[1], eq.values[2:]...) {
			sum += eq.target
		}
	}
	log.Part1(sum)

	// Part 2
	e = Eval{
		operators: []func(int, int) int{
			func(a, b int) int { return a * b },
			func(a, b int) int { return a + b },
			func(a, b int) int { return evil.Atoi(strconv.Itoa(a) + strconv.Itoa(b)) },
		},
	}
	sum = 0
	for _, eq := range equations {
		if e.isPossible(eq.target, eq.values[0], eq.values[1], eq.values[2:]...) {
			sum += eq.target
		}
	}
	log.Part2(sum)
}
