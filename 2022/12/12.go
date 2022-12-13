package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
	"seehuhn.de/go/dijkstra"
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
}

// Edge returns the outgoing edges of the given vertex.
func (g graph) Edges(v space.Point) []space.Point {
	res := []space.Point{}
	for _, dir := range [...]space.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		n := v.Add(dir)
		if n.Y < 0 || n.X < 0 || n.Y >= len(g.heightmap) || n.X >= len(g.heightmap[n.Y]) {
			continue
		}

		if g.heightmap[n.Y][n.X]-g.heightmap[v.Y][v.X] <= 1 {
			res = append(res, n)
		}
	}
	return res
}

func (g graph) Length(e space.Point) int {
	return 1
}

func (g graph) To(e space.Point) space.Point {
	return e
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	heightmap, start, end := parse(filename)

	// Part 1
	path, err := dijkstra.ShortestPath[space.Point, space.Point, int](graph{heightmap: heightmap}, start, end)
	if err != nil {
		evil.Panic("dijkstra error: %w", err)
	}
	log.Part1(len(path))
}
