package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) map[space.Point]rune {
	res := map[space.Point]rune{}
	load.Grid(filename, func(p space.Point, r rune) {
		res[p] = r
	})
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	papers := parse(filename)

	// Part 1
	cnt := 0
	for p, r := range papers {
		if r != '@' {
			continue
		}

		adj := 0
		for dir := range space.Neighbors() {
			if papers[p.Add(dir)] == '@' {
				adj++
			}
		}
		if adj < 4 {
			cnt++
		}
	}
	log.Part1(cnt)
}
