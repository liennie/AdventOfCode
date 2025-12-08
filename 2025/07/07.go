package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
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

	// Part 1
	beams := set.New(start.X)
	next := set.New[int]()
	y := start.Y
	split := 0
loop:
	for {
		next.Clear()
		for beam := range beams {
			switch manifold[space.Point{X: beam, Y: y}] {
			case '^':
				split++
				next.Add(beam-1, beam+1)

			case '.', 'S':
				next.Add(beam)

			default:
				break loop
			}
		}

		beams, next = next, beams
		y++
	}
	log.Part1(split)

	// Part 2
	log.Part2(nil)
}
