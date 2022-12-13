package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) ([][]int, space.Point, space.Point) {
	res := [][]int{}
	var start, end space.Point
	y := 0
	for line := range load.File(filename) {
		hm := []int{}
		for x, height := range strings.Split(line, "") {
			switch height {
			case "S":
				height = "a"
				start = space.Point{X: x, Y: y}
			case "E":
				height = "z"
				end = space.Point{X: x, Y: y}
			}
			hm = append(hm, int(height[0]-'a'))
		}
		res = append(res, hm)
		y++
	}
	return res, start, end
}

type graph struct {
	heightmap [][]int
	rev       bool
}

func (g graph) Edges(n space.Point) []path.Edge[space.Point] {
	res := []path.Edge[space.Point]{}
	for _, dir := range [...]space.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		o := n.Add(dir)
		if o.Y < 0 || o.X < 0 || o.Y >= len(g.heightmap) || o.X >= len(g.heightmap[o.Y]) {
			continue
		}

		if (!g.rev && g.heightmap[o.Y][o.X]-g.heightmap[n.Y][n.X] <= 1) ||
			(g.rev && g.heightmap[n.Y][n.X]-g.heightmap[o.Y][o.X] <= 1) {
			res = append(res, path.Edge[space.Point]{Len: 1, To: o})
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	heightmap, start, end := parse(filename)

	// Part 1
	p, err := path.Shortest[space.Point](graph{heightmap: heightmap}, start, path.EndConst(end))
	if err != nil {
		evil.Panic("path error: %w", err)
	}
	log.Part1(len(p) - 1)

	// Part 2
	p, err = path.Shortest[space.Point](graph{heightmap: heightmap, rev: true}, end, path.EndFunc[space.Point](func(n space.Point) bool {
		return heightmap[n.Y][n.X] == 0
	}))
	if err != nil {
		evil.Panic("path error: %w", err)
	}
	log.Part2(len(p) - 1)
}
