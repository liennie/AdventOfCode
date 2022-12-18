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

func simulate(jetPattern []int, amt int) int {
	top := 0
	topExtra := 0
	tower := set.Set[space.Point]{}
	ji := 0
	ri := 0

	step := func() {
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

	type cacheKey struct {
		ri int
		ji int
	}
	type cacheValue struct {
		i   int
		top int
	}
	cache := map[cacheKey][]cacheValue{}

	for i := 0; i < amt; i++ {
		if cv, ok := cache[cacheKey{ri: ri, ji: ji}]; ok && len(cv) > 10 {
			id := cv[len(cv)-1].i - cv[len(cv)-2].i
			if i+id < amt {
				td := cv[len(cv)-1].top - cv[len(cv)-2].top
				c := (amt - i) / id
				i += c * id
				topExtra += c * td
			}
		} else {
			cache[cacheKey{ri: ri, ji: ji}] = append(cache[cacheKey{ri: ri, ji: ji}], cacheValue{i: i, top: top})
		}

		if i < amt {
			step()
		}
	}

	return top + topExtra
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	jetPattern := parse(filename)

	// Part 1
	log.Part1(simulate(jetPattern, 2022))
	// Part 2
	log.Part2(simulate(jetPattern, 1000000000000))
}
