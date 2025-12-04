package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) map[space.Point]rune {
	res := map[space.Point]rune{}
	load.Grid(filename, func(p space.Point, r rune) {
		res[p] = r
	})
	return res
}

func isAccessible(papers map[space.Point]rune, p space.Point) bool {
	adj := 0
	for dir := range space.Neighbors() {
		if papers[p.Add(dir)] == '@' {
			adj++
		}
	}
	return adj < 4
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

		if isAccessible(papers, p) {
			cnt++
		}
	}
	log.Part1(cnt)

	// Part 2
	cnt = 0
	cont := true
	for cont {
		cont = false

		toRemove := set.New[space.Point]()
		for p, r := range papers {
			if r != '@' {
				continue
			}

			if isAccessible(papers, p) {
				toRemove.Add(p)
			}
		}

		if len(toRemove) > 0 {
			cnt += len(toRemove)
			cont = true

			for p := range toRemove {
				papers[p] = '.'
			}
		}
	}
	log.Part2(cnt)
}
