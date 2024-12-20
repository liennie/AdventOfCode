package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]rune, space.Point, space.Point) {
	res := map[space.Point]rune{}
	var start, end space.Point
	load.Grid(filename, func(p space.Point, r rune) {
		if r == 'S' {
			start = p
			r = '.'
		} else if r == 'E' {
			end = p
			r = '.'
		}

		res[p] = r
	})
	return res, start, end
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	track, start, end := parse(filename)

	// Part 1
	path, _, err := path.Shortest(path.GraphFunc[space.Point](func(p space.Point) (edges []path.Edge[space.Point]) {
		for dir := range space.Orthogonal() {
			next := p.Add(dir)
			if track[next] == '.' {
				edges = append(edges, path.Edge[space.Point]{
					Len: 1,
					To:  next,
				})
			}
		}
		return
	}), start, path.EndConst(end))
	evil.Assert(err == nil, "no path")

	dist := map[space.Point]int{}
	for i, p := range path {
		dist[p] = i
	}

	cheats100 := 0
	for _, p := range path {
		for dir := range space.Orthogonal() {
			if track[p.Add(dir)] != '#' {
				continue
			}
			next := p.Add(dir.Scale(2))
			if track[next] != '.' {
				continue
			}
			if dist[p] > dist[next] {
				continue
			}

			cheat := dist[next] - dist[p] - 2
			if cheat >= 100 {
				cheats100++
			}
		}
	}
	log.Part1(cheats100)
}
