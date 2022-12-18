package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

var rocks = [][]space.Point{
	{
		{0, 0}, {1, 0}, {2, 0}, {3, 0},
	},
	{
		/*   */ {1, 2},
		{0, 1}, {1, 1}, {2, 1},
		/*   */ {1, 0},
	},
	{
		/*           */ {2, 2},
		/*           */ {2, 1},
		{0, 0}, {1, 0}, {2, 0},
	},
	{
		{0, 3},
		{0, 2},
		{0, 1},
		{0, 0},
	},
	{
		{0, 1}, {1, 1},
		{0, 0}, {1, 0},
	},
}

func parse(filename string) []int {
	data := load.Line(filename)
	res := make([]int, 0, len(data))
	for _, j := range data {
		switch j {
		case '<':
			res = append(res, -1)
		case '>':
			res = append(res, 1)
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	jetPattern := parse(filename)

	// Part 1
	top := 0
	tower := set.Set[space.Point]{}
	ji := 0
	ri := 0
	for i := 0; i < 2022; i++ {
		rpos := space.Point{X: 2, Y: top + 4}

	fall:
		for {
			movement := space.Point{X: jetPattern[ji]}
			ji = (ji + 1) % len(jetPattern)
			for _, piece := range rocks[ri] {
				nppos := rpos.Add(piece).Add(movement)
				if tower.Contains(nppos) || nppos.X < 0 || nppos.X >= 7 {
					movement = space.Point{}
					break
				}
			}
			rpos = rpos.Add(movement)

			movement = space.Point{Y: -1}
			for _, piece := range rocks[ri] {
				nppos := rpos.Add(piece).Add(movement)
				if tower.Contains(nppos) || nppos.Y <= 0 {
					break fall
				}
			}
			rpos = rpos.Add(movement)
		}

		for _, piece := range rocks[ri] {
			nppos := rpos.Add(piece)
			top = ints.Max(top, nppos.Y)
			tower.Add(nppos)
		}
		ri = (ri + 1) % len(rocks)
	}
	log.Part1(top)
}
