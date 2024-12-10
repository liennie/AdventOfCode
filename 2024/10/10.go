package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) map[space.Point]int {
	res := map[space.Point]int{}
	load.Grid(filename, func(p space.Point, r rune) {
		evil.Assert('0' <= r && r <= '9', "invalid char ", string(r))
		res[p] = int(r - '0')
	})
	return res
}

func peaks(topology map[space.Point]int, pos space.Point) set.Set[space.Point] {
	height := topology[pos]

	if height == 9 {
		return set.New(pos)
	}

	sc := set.New[space.Point]()
	for _, dir := range []space.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := pos.Add(dir)
		nextHeight := topology[next]

		if nextHeight == height+1 {
			sc.AddSeq(peaks(topology, next).All())
		}
	}
	return sc
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	topology := parse(filename)

	// Part 1
	sum := 0
	for pos, height := range topology {
		if height == 0 {
			sum += len(peaks(topology, pos))
		}
	}
	log.Part1(sum)
}
