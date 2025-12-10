package main

import (
	"regexp"
	"slices"

	"github.com/liennie/AdventOfCode/pkg/comb"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Machine struct {
	diagram []bool
	buttons [][]int
	joltage []int
}

func parse(filename string) []Machine {
	lineRe := regexp.MustCompile(`^\[([.#]+)\]((?: \([\d,]+\))+) \{([\d,]+)\}$`)
	buttonRe := regexp.MustCompile(`\(([\d,]+)\)`)

	res := []Machine{}
	for line := range load.File(filename) {
		matches := lineRe.FindStringSubmatch(line)
		evil.Assert(matches != nil, "invalid format %q", line)

		m := Machine{}
		for _, d := range matches[1] {
			m.diagram = append(m.diagram, d == '#')
		}

		m.joltage = evil.Split(matches[3], ",")

		buttons := buttonRe.FindAllStringSubmatch(matches[2], -1)
		evil.Assert(buttons != nil, "could not match buttons %q", matches[2])

		for _, b := range buttons {
			m.buttons = append(m.buttons, evil.Split(b[1], ","))
		}

		res = append(res, m)
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	machines := parse(filename)

	// Part 1
	sum := 0
	for _, m := range machines {
		min := len(m.buttons) + 1
		for c := range comb.Comb(m.buttons) {
			if len(c) >= min {
				continue
			}

			diodes := make([]bool, len(m.diagram))
			for _, b := range c {
				for _, i := range b {
					diodes[i] = !diodes[i]
				}
			}
			if slices.Equal(diodes, m.diagram) {
				min = len(c)
			}
		}
		evil.Assert(min <= len(m.buttons), "solution not found")
		sum += min
	}
	log.Part1(sum)

	// Part 2
	log.Part2(nil)
}
