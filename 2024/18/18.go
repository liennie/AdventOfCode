package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (res []space.Point) {
	for line := range load.File(filename) {
		n := evil.Split(line, ",")
		evil.Assert(len(n) == 2, "invalid line")
		res = append(res, space.Point{n[0], n[1]})
	}
	return
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	bytes := parse(filename)
	aabb := space.NewAABB(space.Point{0, 0}, space.Point{70, 70})
	if filename != "input.txt" {
		aabb = space.NewAABB(space.Point{0, 0}, space.Point{6, 6})
	}

	// Part 1
	fallen := 1024
	if filename != "input.txt" {
		fallen = 12
	}

	corrupted := set.New(bytes[:fallen]...)
	_, steps, err := path.Shortest(path.GraphFunc[space.Point](func(p space.Point) (edges []path.Edge[space.Point]) {
		for dir := range space.Orthogonal() {
			next := p.Add(dir)
			if aabb.Contains(next) && !corrupted.Contains(next) {
				edges = append(edges, path.Edge[space.Point]{
					Len: 1,
					To:  next,
				})
			}
		}
		return
	}), aabb.Min, path.EndConst(aabb.Max))
	evil.Assert(err == nil, "path not found: ", err)

	log.Part1(steps)
}
