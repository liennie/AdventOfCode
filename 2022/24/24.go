package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Blizzard struct {
	Start space.Point
	Dir   space.Point
}

type Valley struct {
	Start     space.Point
	End       space.Point
	InnerAABB space.AABB
	Blizzards []Blizzard
}

func (v Valley) Graph() *ValleyGraph {
	g := &ValleyGraph{
		End:   v.End,
		Empty: set.New[space.Point3](),
	}

	size := v.InnerAABB.Size()
	lcm := ints.LCM(size.X, size.Y)
	for z := range lcm {
		blizz := set.New[space.Point]()
		for _, b := range v.Blizzards {
			blizz.Add(v.InnerAABB.Wrap(b.Start.Add(b.Dir.Scale(z))))
		}
		for p := range v.InnerAABB.All() {
			if !blizz.Contains(p) {
				g.Empty.Add(space.Point3{X: p.X, Y: p.Y, Z: z})
			}
		}
		g.Empty.Add(
			space.Point3{X: v.Start.X, Y: v.Start.Y, Z: z},
			space.Point3{X: v.End.X, Y: v.End.Y, Z: z},
		)
	}
	g.WrapZ = lcm

	return g
}

type ValleyGraph struct {
	End   space.Point
	WrapZ int
	Empty set.Set[space.Point3]
}

func (g *ValleyGraph) Edges(p space.Point3) (edges []path.Edge[space.Point3]) {
	// defer func() {
	// 	log.Printf("edges for p %v: %v", p, edges)
	// }()

	for _, dir := range []space.Point{
		{X: 1},  // right
		{Y: -1}, // up
		{X: -1}, // left
		{Y: 1},  // down
		{},      // wait
	} {
		next := p.Add(space.Point3{X: dir.X, Y: dir.Y, Z: 1})
		if next.Z >= g.WrapZ {
			next.Z = 0
		}

		if g.Empty.Contains(next) {
			edges = append(edges, path.Edge[space.Point3]{Len: 1, To: next})
		}
	}
	return edges
}

func (g *ValleyGraph) ShortestRemainigDist(p space.Point3) int {
	return space.Point{X: p.X, Y: p.Y}.Sub(g.End).ManhattanLen()
}

func (g *ValleyGraph) IsEnd(p space.Point3) bool {
	return space.Point{X: p.X, Y: p.Y} == g.End
}

var _ path.AStarGraph[space.Point3] = &ValleyGraph{}
var _ path.End[space.Point3] = &ValleyGraph{}

func parse(filename string) Valley {
	res := Valley{
		Start: space.Point{X: 1, Y: 0},
	}
	load.Grid(filename, func(p space.Point, r rune) {
		if p == res.Start {
			evil.Assert(r == '.', "start not found at %v", res.Start)
			return
		}
		if p.Y == 0 || p.X == 0 {
			evil.Assert(r == '#', "non-wall found on %v: %c", p, r)
			return
		}
		if res.End.X != 0 && res.End.Y != 0 && p == res.End {
			evil.Assert(r == '.', "end not found at %v", res.End)
			return
		}
		if res.End.X != 0 && p.X == res.End.X+1 {
			evil.Assert(r == '#', "non-wall found on %v: %c", p, r)
			return
		}
		if res.End.Y != 0 && p.Y == res.End.Y {
			evil.Assert(r == '#', "non-wall found on %v: %c", p, r)
			return
		}
		if res.InnerAABB.Contains(p) {
			evil.Assert(r != '#', "wall found on %v", p)
		}

		switch r {
		case '>':
			res.Blizzards = append(res.Blizzards, Blizzard{Start: p, Dir: space.Point{X: 1}})
		case '^':
			res.Blizzards = append(res.Blizzards, Blizzard{Start: p, Dir: space.Point{Y: -1}})
		case '<':
			res.Blizzards = append(res.Blizzards, Blizzard{Start: p, Dir: space.Point{X: -1}})
		case 'v':
			res.Blizzards = append(res.Blizzards, Blizzard{Start: p, Dir: space.Point{Y: 1}})
		case '.':
			// empty space
		case '#':
			if res.InnerAABB.Contains(space.Point{X: p.X - 1, Y: p.Y}) {
				// end of line
				res.End.X = p.X - 1
			} else {
				// end of map
				res.End.Y = p.Y
			}
		default:
			evil.Panic("invalid char %c", r)
		}
		if r != '#' && p.Y != res.End.Y {
			res.InnerAABB = res.InnerAABB.Add(p)
		}
	})
	evil.Assert(res.End == res.InnerAABB.Max.Add(space.Point{Y: 1}), "weird shape of map")
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid := parse(filename)

	// Part 1
	graph := grid.Graph()

	_, dist, err := path.Shortest(graph, space.Point3{X: grid.Start.X, Y: grid.Start.Y, Z: 0}, graph)
	evil.Err(err)

	log.Part1(dist)

	// Part 2
	// back to start
	graph.End = grid.Start
	_, dist2, err := path.Shortest(graph, space.Point3{X: grid.End.X, Y: grid.End.Y, Z: ints.Mod(dist, graph.WrapZ)}, graph)
	evil.Err(err)

	log.Print("back to start: ", dist2)

	// and back to end again
	graph.End = grid.End
	_, dist3, err := path.Shortest(graph, space.Point3{X: grid.Start.X, Y: grid.Start.Y, Z: ints.Mod(dist+dist2, graph.WrapZ)}, graph)
	evil.Err(err)

	log.Print("back to end: ", dist3)

	log.Part2(dist + dist2 + dist3)
}
