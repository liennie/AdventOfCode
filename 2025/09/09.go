package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) []space.Point {
	res := []space.Point{}
	for line := range load.File(filename) {
		coords := evil.Split(line, ",")
		res = append(res, space.Point{
			X: coords[0],
			Y: coords[1],
		})
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	tiles := parse(filename)

	// Part 1
	max := 0
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			a := tiles[i]
			b := tiles[j]
			area := b.Sub(a).Abs().Add(space.Point{X: 1, Y: 1}).Area()

			if area > max {
				max = area
			}
		}
	}
	log.Part1(max)

	// Part 2
	max = 0
	for i := 0; i < len(tiles)-1; i++ {
	rectangles:
		for j := i + 1; j < len(tiles); j++ {
			aabb := space.NewAABB(tiles[i], tiles[j])
			area := aabb.Size().Area()

			aabb = aabb.Expand(-1)

			for k, tile := range tiles {
				next := tiles[(k+1)%len(tiles)]
				line := space.NewAABB(tile, next)
				if aabb.Overlaps(line) {
					continue rectangles
				}
			}

			if area > max {
				max = area
			}
		}
	}
	log.Part2(max)
}
