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
	end      space.Point
	min, max int
}

func (g Graph) Edges(n Node) []path.Edge[Node] {
	edges := []path.Edge[Node]{}

	for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
		if n.from.Norm().Flip() == dir {
			continue
		}

		if n.from.Norm() == dir {
			npos := n.pos.Add(dir)
			if !g.aabb.Contains(npos) {
				continue
			}

			from := n.from.Add(dir)
			if ints.Abs(from.X) > g.max || ints.Abs(from.Y) > g.max {
				continue
			}

			edges = append(edges, path.Edge[Node]{
				Len: g.blocks[npos.Y][npos.X],
				To: Node{
					from: from,
					pos:  npos,
				},
			})
		} else {
			npos := n.pos.Add(dir.Scale(g.min))
			if !g.aabb.Contains(npos) {
				continue
			}

			cost := 0
			for i := 1; i <= g.min; i++ {
				p := n.pos.Add(dir.Scale(i))
				cost += g.blocks[p.Y][p.X]
			}

			edges = append(edges, path.Edge[Node]{
				Len: cost,
				To: Node{
					from: dir.Scale(g.min),
					pos:  npos,
				},
			})
		}
	}

	return edges
}

func (g Graph) ShortestRemainigDist(n Node) int {
	return n.pos.Sub(g.end).ManhattanLen()
}

var _ path.AStarGraph[Node] = Graph{}

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
			end:    end,
			min:    1,
			max:    3,
		},
		Node{pos: space.Point{X: 0, Y: 0}},
		path.EndFunc[Node](func(n Node) bool { return n.pos == end }),
	)
	evil.Err(err)
	log.Part1(l)

	// Part 2
	_, l, err = path.Shortest(
		Graph{
			blocks: blocks,
			aabb:   aabb,
			end:    end,
			min:    4,
			max:    10,
		},
		Node{pos: space.Point{X: 0, Y: 0}},
		path.EndFunc[Node](func(n Node) bool { return n.pos == end }),
	)
	evil.Err(err)
	log.Part2(l)
}
