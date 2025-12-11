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

type checklist struct {
	none int
	dac  int
	fft  int
	all  int
}

func (c checklist) add(other checklist) checklist {
	return checklist{
		none: c.none + other.none,
		dac:  c.dac + other.dac,
		fft:  c.fft + other.fft,
		all:  c.all + other.all,
	}
}

func (c checklist) checkDAC() checklist {
	return checklist{
		dac: c.dac + c.none,
		all: c.all + c.fft,
	}
}

func (c checklist) checkFFT() checklist {
	return checklist{
		fft: c.fft + c.none,
		all: c.all + c.dac,
	}
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
	states2 := map[string]checklist{
		"svr": {
			none: 1,
		},
	}
	run = true
	for run {
		run = false
		for n, c := range states2 {
			switch n {
			case "out":
				continue

			case "dac":
				c = c.checkDAC()

			case "fft":
				c = c.checkFFT()
			}
			run = true

			delete(states2, n)
			for _, out := range devices[n].outputs {
				states2[out] = states2[out].add(c)
			}
		}
	}
	log.Part2(states2["out"].all)
}
