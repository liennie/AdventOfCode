package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) {
	for line := range load.File(filename) {
		_ = line
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	parse(filename)

	// Part 1
	log.Part1(nil)
}
