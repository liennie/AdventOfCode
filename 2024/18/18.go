package main

import (
	"errors"
	"fmt"

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

func graph(aabb space.AABB, corrupted set.Set[space.Point]) path.GraphFunc[space.Point] {
	return func(p space.Point) (edges []path.Edge[space.Point]) {
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
	}
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
	_, steps, err := path.Shortest(graph(aabb, corrupted), aabb.Min, path.EndConst(aabb.Max))
	evil.Assert(err == nil, "path not found: %w", err)

	log.Part1(steps)

	// Part 2
	for _, falling := range bytes[fallen:] {
		corrupted.Add(falling)
		_, _, err := path.Shortest(graph(aabb, corrupted), aabb.Min, path.EndConst(aabb.Max))
		if errors.Is(err, path.ErrNotFound) {
			log.Part2(fmt.Sprintf("%d,%d", falling.X, falling.Y))
			break
		}
	}
}
