package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type object uint8

const (
	air object = iota
	rock
	sand
)

func parse(filename string) map[space.Point]object {
	res := map[space.Point]object{}
	for line := range load.File(filename) {
		path := []space.Point{}
		for _, point := range strings.Split(line, " -> ") {
			coords := evil.Split(point, ",")
			path = append(path, space.Point{X: coords[0], Y: coords[1]})
		}

		for i := 1; i < len(path); i++ {
			v := path[i].Sub(path[i-1]).Norm()

			for p := path[i-1]; !p.Equals(path[i]); p = p.Add(v) {
				res[p] = rock
			}
			res[path[i]] = rock
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cave := parse(filename)

	// Part 1
	bottom := math.MinInt
	for p := range cave {
		bottom = ints.Max(bottom, p.Y)
	}

	count := 0
	start := space.Point{X: 500, Y: 0}
sand:
	for {
		p := start
		for {
			if np := p.Add(space.Point{Y: 1}); cave[np] == air {
				p = np
			} else if np := p.Add(space.Point{X: -1, Y: 1}); cave[np] == air {
				p = np
			} else if np := p.Add(space.Point{X: 1, Y: 1}); cave[np] == air {
				p = np
			} else {
				cave[p] = sand
				count++
				break
			}

			if p.Y > bottom {
				break sand
			}
		}
	}
	log.Part1(count)

	// Part 2
	floor := bottom + 2
	for x := 500 - floor; x <= 500+floor; x++ {
		cave[space.Point{X: x, Y: floor}] = rock
	}

	for cave[start] == air {
		p := start
		for {
			if np := p.Add(space.Point{Y: 1}); cave[np] == air {
				p = np
			} else if np := p.Add(space.Point{X: -1, Y: 1}); cave[np] == air {
				p = np
			} else if np := p.Add(space.Point{X: 1, Y: 1}); cave[np] == air {
				p = np
			} else {
				cave[p] = sand
				count++
				break
			}
		}
	}
	log.Part1(count)
}
