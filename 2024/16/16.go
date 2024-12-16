package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]rune, space.Point, space.Point) {
	var start, end space.Point
	res := map[space.Point]rune{}
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

type Node struct {
	pos space.Point
	dir space.Point
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid, start, end := parse(filename)

	// Part 1
	_, points, err := path.Shortest(
		path.GraphFunc[Node](func(n Node) (edges []path.Edge[Node]) {
			for dir := range space.Orthogonal() {
				if grid[n.pos.Add(dir)] != '.' {
					continue
				}

				p := 1
				if dir == n.dir.Flip() {
					p += 2000
				} else if dir != n.dir {
					p += 1000
				}

				edges = append(edges, path.Edge[Node]{
					Len: p,
					To: Node{
						pos: n.pos.Add(dir),
						dir: dir,
					},
				})
			}
			return
		}),
		Node{pos: start, dir: space.Point{X: 1}},
		path.EndFunc[Node](func(n Node) bool {
			return n.pos == end
		}),
	)
	evil.Assert(err == nil, "path not found")
	log.Part1(points)
}
