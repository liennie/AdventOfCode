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

type Graph struct {
	grid  map[space.Point]rune
	start space.Point
	end   space.Point
}

var _ path.AStarGraph[Node] = Graph{}

func (g Graph) Edges(n Node) (edges []path.Edge[Node]) {
	for dir := range space.Orthogonal() {
		if g.grid[n.pos.Add(dir)] != '.' {
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
}

func (g Graph) Heuristic(n Node) int {
	return n.pos.Sub(g.end).ManhattanLen()
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid, start, end := parse(filename)

	// Part 1
	_, points, err := path.Shortest(
		Graph{grid: grid, start: start, end: end},
		Node{pos: start, dir: space.Point{X: 1}},
		path.EndFunc[Node](func(n Node) bool {
			return n.pos == end
		}),
	)
	evil.Assert(err == nil, "path not found")
	log.Part1(points)

	// Part 2
	paths, points, err := path.AllShortest(
		Graph{grid: grid, start: start, end: end},
		Node{pos: start, dir: space.Point{X: 1}},
		path.EndFunc[Node](func(n Node) bool {
			return n.pos == end
		}),
	)
	evil.Assert(err == nil, "path not found")
	tiles := set.New[space.Point]()
	for _, path := range paths {
		for _, node := range path {
			tiles.Add(node.pos)
		}
	}
	log.Part2(len(tiles))
}
