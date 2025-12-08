package main

import (
	"maps"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]rune, space.Point) {
	res := map[space.Point]rune{}
	start := space.Point{X: -1, Y: -1}
	load.Grid(filename, func(p space.Point, r rune) {
		if r == 'S' {
			start = p
		}

		res[p] = r
	})
	evil.Assert(start.X != -1, "start not found")
	return res, start
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	manifold, start := parse(filename)

	beams := map[int]int{start.X: 1}
	next := map[int]int{}
	y := start.Y
	split := 0
loop:
	for {
		clear(next)
		for beam, cnt := range beams {
			switch manifold[space.Point{X: beam, Y: y}] {
			case '^':
				split++
				next[beam-1] += cnt
				next[beam+1] += cnt

			case '.', 'S':
				next[beam] += cnt

			default:
				break loop
			}
		}

		beams, next = next, beams
		y++
	}
	log.Part1(split)
	log.Part2(ints.SumSeq(maps.Values(beams)))
}
