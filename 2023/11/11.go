package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (set.Set[space.Point], space.AABB) {
	res := set.New[space.Point]()
	aabb := space.AABB{}

	y := 0
	for line := range load.File(filename) {
		for x, c := range line {
			pos := space.Point{X: x, Y: y}
			aabb = aabb.Add(pos)

			if c == '#' {
				res.Add(pos)
			}
		}
		y++
	}

	return res, aabb
}

func expand(galaxies set.Set[space.Point], aabb space.AABB) (set.Set[space.Point], space.AABB) {
	xoff := make([]int, aabb.Max.X+1)
x:
	for x := 0; x <= aabb.Max.X; x++ {
		for y := 0; y <= aabb.Max.Y; y++ {
			if galaxies.Contains(space.Point{X: x, Y: y}) {
				continue x
			}
		}

		for xo := x; xo < len(xoff); xo++ {
			xoff[xo]++
		}
	}

	yoff := make([]int, aabb.Max.Y+1)
y:
	for y := 0; y <= aabb.Max.Y; y++ {
		for x := 0; x <= aabb.Max.X; x++ {
			if galaxies.Contains(space.Point{X: x, Y: y}) {
				continue y
			}
		}

		for yo := y; yo < len(yoff); yo++ {
			yoff[yo]++
		}
	}

	res := set.New[space.Point]()
	for galaxy := range galaxies {
		galaxy = galaxy.Add(space.Point{X: xoff[galaxy.X], Y: yoff[galaxy.Y]})
		res.Add(galaxy)
		aabb = aabb.Add(galaxy)
	}
	return res, aabb
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	galaxies, aabb := parse(filename)

	// Part 1
	galaxies, aabb = expand(galaxies, aabb)
	sum := 0
	for a := range galaxies {
		for b := range galaxies {
			sum += a.Sub(b).ManhattanLen()
		}
	}
	log.Part1(sum / 2)
}
