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

	dirs := [...]space.Point3{
		{X: -1}, {X: 1},
		{Y: -1}, {Y: 1},
		{Z: -1}, {Z: 1},
	}

	// Part 1
	cubeSet := set.New(cubes...)
	count := 0
	for cube := range cubeSet {
		for _, dir := range dirs {
			if !cubeSet.Contains(cube.Add(dir)) {
				count++
			}
		}
	}
	log.Part1(count)

	// Part 2
	aabb := space.NewAABB3(cubes[0])
	for _, cube := range cubes {
		aabb = aabb.Add(cube)
	}
	aabb.Min = aabb.Min.Add(space.Point3{-1, -1, -1})
	aabb.Max = aabb.Max.Add(space.Point3{1, 1, 1})

	ext := set.New[space.Point3]()
	flood := set.New(aabb.Min)
	for len(flood) > 0 {
		p, _ := flood.Pop()
		ext.Add(p)
		for _, dir := range dirs {
			np := p.Add(dir)
			if aabb.Contains(np) && !cubeSet.Contains(np) && !ext.Contains(np) {
				flood.Add(np)
			}
		}
	}

	count = 0
	for cube := range cubeSet {
		for _, dir := range dirs {
			if ext.Contains(cube.Add(dir)) {
				count++
			}
		}
	}
	log.Part2(count)
}
