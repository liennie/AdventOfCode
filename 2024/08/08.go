package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[rune][]space.Point, space.AABB) {
	res := map[rune][]space.Point{}
	aabb := space.AABB{}
	load.Grid(filename, func(x, y int, r rune) {
		aabb = aabb.Add(space.Point{x, y})
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') {
			res[r] = append(res[r], space.Point{x, y})
		}
	})
	return res, aabb
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	antennas, aabb := parse(filename)

	// Part 1
	antinodes := set.New[space.Point]()
	for _, ant := range antennas {
		for i := range len(ant) - 1 {
			for j := i + 1; j < len(ant); j++ {
				a, b := ant[i], ant[j]
				d := b.Sub(a)

				antiA := a.Sub(d)
				if aabb.Contains(antiA) {
					antinodes.Add(antiA)
				}

				antiB := b.Add(d)
				if aabb.Contains(antiB) {
					antinodes.Add(antiB)
				}
			}
		}
	}
	log.Part1(len(antinodes))

	// Part 2
	antinodes = set.New[space.Point]()
	for _, ant := range antennas {
		for i := range len(ant) - 1 {
			for j := i + 1; j < len(ant); j++ {
				a, b := ant[i], ant[j]
				d := b.Sub(a)

				for antiA := a; aabb.Contains(antiA); antiA = antiA.Sub(d) {
					antinodes.Add(antiA)
				}
				for antiB := b; aabb.Contains(antiB); antiB = antiB.Add(d) {
					antinodes.Add(antiB)
				}
			}
		}
	}
	log.Part2(len(antinodes))
}
