package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
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

type Graph struct {
	forest [][]byte
	start  space.Point
	end    space.Point
	max    int
}

type Node struct {
	pos      space.Point
	previous *Node
	len      int
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

	return &Graph{
		forest: forest,
		start:  start,
		end:    end,
	}
}

func (g *Graph) Start() *Node {
	return &Node{
		pos: g.start,
	}
}

func (g *Graph) Edges(n *Node) []path.Edge[*Node] {
	var dirs []space.Point
	switch g.forest[n.pos.Y][n.pos.X] {
	case '>':
		dirs = []space.Point{{X: 1}}
	case '^':
		dirs = []space.Point{{Y: -1}}
	case '<':
		dirs = []space.Point{{X: -1}}
	case 'v':
		dirs = []space.Point{{Y: 1}}
	case '.':
		dirs = []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}}
	}

	next := set.New[space.Point]()
	for _, dir := range dirs {
		e := n.pos.Add(dir)
		x, y := e.X, e.Y
		if x >= 0 && x < len(g.forest[0]) && y >= 0 && y < len(g.forest) && g.forest[y][x] != '#' {
			next.Add(e)
		}
	}

	for p := n.previous; p != nil; p = p.previous {
		next.Remove(p.pos)
	}

	res := []path.Edge[*Node]{}
	for p := range next {
		res = append(res, path.Edge[*Node]{
			Len: 1,
			To: &Node{
				pos:      p,
				previous: n,
				len:      n.len + 1,
			},
		})
	}
	return res
}

func (g *Graph) IsEnd(n *Node) bool {
	if n.pos == g.end {
		g.max = max(g.max, n.len)
	}

	return false
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	forest := parse(filename)

	// Part 1
	graph := newGraph(forest)
	_, _, err := path.Shortest(graph, graph.Start(), graph)
	evil.Assert(err == path.ErrNotFound)
	log.Part1(graph.max)
}
