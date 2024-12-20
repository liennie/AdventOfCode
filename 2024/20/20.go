package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/set"
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

func cheatDests(track map[space.Point]rune, start space.Point, max int) map[space.Point]int {
	res := map[space.Point]int{}
	visited := set.New[space.Point]()
	cur := set.New(start)

	for dist := 1; dist <= max; dist++ {
		visited.AddSeq(cur.All())

		next := set.New[space.Point]()
		for p := range cur {
			for dir := range space.Orthogonal() {
				n := p.Add(dir)
				if visited.Contains(n) {
					continue
				}

				if track[n] == '.' {
					if _, ok := res[n]; !ok {
						res[n] = dist
					}
				}
				next.Add(n)
			}
		}
		cur = next
	}

	return res
}

func cheats(track map[space.Point]rune, path []space.Point, max int) map[int]int {
	dist := map[space.Point]int{}
	for i, p := range path {
		dist[p] = i
	}

	cheats := map[int]int{}
	for _, p := range path {
		for next, d := range cheatDests(track, p, max) {
			if dist[p] > dist[next] {
				continue
			}

			cheat := dist[next] - dist[p] - d
			if cheat > 0 {
				cheats[cheat]++
			}
		}
	}
	return cheats
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	track, start, end := parse(filename)

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

	// Part 1
	cheats100 := 0
	for dist, cnt := range cheats(track, path, 2) {
		if dist >= 100 {
			cheats100 += cnt
		}
	}
	log.Part1(cheats100)

	// Part 2
	cheats100 = 0
	for dist, cnt := range cheats(track, path, 20) {
		if dist >= 100 {
			cheats100 += cnt
		}
	}
	log.Part2(cheats100)
}
