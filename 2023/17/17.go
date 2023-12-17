package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]int {
	res := [][]int{}
	for line := range load.File(filename) {
		res = append(res, evil.Split(line, ""))
	}
	return res
}

type Node struct {
	from space.Point
	pos  space.Point
}

type Graph struct {
	blocks   [][]int
	aabb     space.AABB
	min, max int
}

func (g Graph) Edges(n Node) []path.Edge[Node] {
	edges := []path.Edge[Node]{}

	for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
		if n.from.Normalize().Flip() == dir {
			continue
		}

		cost := 0
		for i := 1; i <= g.max; i++ {
			tdir := dir.Scale(i)
			npos := n.pos.Add(tdir)
			if !g.aabb.Contains(npos) {
				break
			}

			tfrom := n.from.Add(tdir)
			if ints.Abs(tfrom.X) > g.max || ints.Abs(tfrom.Y) > g.max {
				break
			}

			cost += g.blocks[npos.Y][npos.X]
			if i < g.min {
				continue
			}

			var from space.Point
			if n.from.Normalize() == dir {
				from = tfrom
			} else {
				from = tdir
			}

			edges = append(edges, path.Edge[Node]{
				Len: cost,
				To: Node{
					from: from,
					pos:  npos,
				},
			})
		}
	}

	return edges
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	blocks := parse(filename)
	end := space.Point{
		X: len(blocks[0]) - 1,
		Y: len(blocks) - 1,
	}
	aabb := space.NewAABB(space.Point{}, end)

	// Part 1
	_, l, err := path.Shortest(
		Graph{
			blocks: blocks,
			aabb:   aabb,
			min:    1,
			max:    3,
		},
		Node{pos: space.Point{X: 0, Y: 0}},
		path.EndFunc[Node](func(n Node) bool { return n.pos == end }),
	)
	evil.Assert(err == nil, err)
	log.Part1(l)

	// Part 2
	_, l, err = path.Shortest(
		Graph{
			blocks: blocks,
			aabb:   aabb,
			min:    4,
			max:    10,
		},
		Node{pos: space.Point{X: 0, Y: 0}},
		path.EndFunc[Node](func(n Node) bool { return n.pos == end }),
	)
	evil.Assert(err == nil, err)
	log.Part2(l)
}
