package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]byte {
	res := [][]byte{}
	for line := range load.File(filename) {
		res = append(res, []byte(line))
	}
	return res
}

type Crossroads struct {
	next map[*Crossroads]int
	all  map[*Crossroads]int
}

func newCrossroads() *Crossroads {
	return &Crossroads{
		next: map[*Crossroads]int{},
		all:  map[*Crossroads]int{},
	}
}

type Graph struct {
	start *Crossroads
	end   *Crossroads
}

func newGraph(forest [][]byte) *Graph {
	start := space.Point{X: -1}
	for x := range forest[0] {
		if forest[0][x] == '.' {
			start.X = x
			break
		}
	}
	evil.Assert(start.X >= 0)

	bottom := len(forest) - 1
	end := space.Point{X: -1, Y: bottom}
	for x := range forest[bottom] {
		if forest[bottom][x] == '.' {
			end.X = x
			break
		}
	}
	evil.Assert(end.X >= 0)

	cross := map[space.Point]*Crossroads{}

	slopes := map[byte]space.Point{
		'>': {X: 1},
		'^': {Y: -1},
		'<': {X: -1},
		'v': {Y: 1},
	}

	var walkFrom func(space.Point) *Crossroads
	walkFrom = func(p space.Point) *Crossroads {
		if c, ok := cross[p]; ok {
			return c
		}
		c := newCrossroads()
		cross[p] = c

		for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
			next := p.Add(dir)
			if next.Y < 0 || next.Y >= len(forest) || next.X < 0 || next.X >= len(forest[next.Y]) || forest[next.Y][next.X] == '#' {
				continue
			}

			l := 0
			prev := p
			cur := p
			slope := false
			crs := false

			for !crs {
				l++
				prev, cur = cur, next

				if cur == start || cur == end {
					crs = true
					break
				}

				for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
					n := cur.Add(dir)
					if n == prev {
						continue
					}

					if n.Y < 0 || n.Y >= len(forest) || n.X < 0 || n.X >= len(forest[n.Y]) || forest[n.Y][n.X] == '#' {
						continue
					}

					if next != cur {
						crs = true
						break
					}

					next = n
					if slopes[forest[cur.Y][cur.X]] == dir.Flip() {
						slope = true
					}
				}
			}

			nc := walkFrom(cur)
			if !slope {
				c.next[nc] = l
			}
			c.all[nc] = l
		}

		return c
	}

	walkFrom(start)

	log.Printf("mapped %d nodes", len(cross))

	return &Graph{
		start: cross[start],
		end:   cross[end],
	}
}

func _longest(start, end *Crossroads, visited set.Set[*Crossroads], slippery bool) (int, bool) {
	if start == end {
		return 0, true
	}

	visited.Add(start)
	defer visited.Remove(start)

	var next map[*Crossroads]int
	if slippery {
		next = start.next
	} else {
		next = start.all
	}

	ml := 0
	for c, cl := range next {
		if visited.Contains(c) {
			continue
		}

		tl, ok := _longest(c, end, visited, slippery)
		if ok {
			ml = max(ml, cl+tl)
		}
	}

	return ml, ml > 0
}

func longest(start, end *Crossroads, slippery bool) int {
	l, ok := _longest(start, end, set.New[*Crossroads](), slippery)
	evil.Assert(ok)
	return l
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	forest := parse(filename)

	// Part 1
	graph := newGraph(forest)
	log.Part1(longest(graph.start, graph.end, true))

	// Part 2
	log.Part2(longest(graph.start, graph.end, false))
}
