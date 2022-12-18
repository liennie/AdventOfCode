package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) []space.Point3 {
	res := []space.Point3{}
	for line := range load.File(filename) {
		coords := evil.Split(line, ",")
		res = append(res, space.Point3{
			X: coords[0],
			Y: coords[1],
			Z: coords[2],
		})
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cubes := parse(filename)

	// Part 1
	cubeSet := set.New(cubes...)
	count := 0
	for cube := range cubeSet {
		for _, dir := range [...]space.Point3{
			{X: -1}, {X: 1},
			{Y: -1}, {Y: 1},
			{Z: -1}, {Z: 1},
		} {
			if !cubeSet.Contains(cube.Add(dir)) {
				count++
			}
		}
	}
	log.Part1(count)
}
