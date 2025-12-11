package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Device struct {
	name    string
	outputs []string
}

func parse(filename string) map[string]Device {
	res := map[string]Device{}
	for line := range load.File(filename) {
		name, out, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "invalid format %q", line)

		outputs := strings.Fields(out)

		res[name] = Device{
			name:    name,
			outputs: outputs,
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	devices := parse(filename)

	// Part 1
	states := map[string]int{
		"you": 1,
	}
	run := true
	for run {
		run = false
		for n, c := range states {
			if n == "out" {
				continue
			}
			run = true

			delete(states, n)
			for _, out := range devices[n].outputs {
				states[out] += c
			}
		}
	}
	log.Part1(states["out"])

	// Part 2
	log.Part2(nil)
}
