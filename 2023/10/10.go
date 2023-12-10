package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Pipe struct {
	connections [2]space.Point
}

func (p Pipe) next(from space.Point) space.Point {
	switch from {
	case p.connections[0]:
		return p.connections[1]
	case p.connections[1]:
		return p.connections[0]
	}
	evil.Panic("cannot reach pipe %v from %v", p, from)
	return space.Point{}
}

var connMap = map[rune][2]space.Point{
	'|': {
		{Y: 1},
		{Y: -1},
	},
	'-': {
		{X: 1},
		{X: -1},
	},
	'L': {
		{X: 1},
		{Y: -1},
	},
	'J': {
		{X: -1},
		{Y: -1},
	},
	'7': {
		{X: -1},
		{Y: 1},
	},
	'F': {
		{X: 1},
		{Y: 1},
	},
}

func parse(filename string) (map[space.Point]Pipe, space.Point) {
	res := map[space.Point]Pipe{}
	var start space.Point

	y := 0
	for line := range load.File(filename) {
		for x, c := range line {
			pos := space.Point{X: x, Y: y}

			if c == '.' {
				continue
			}
			if c == 'S' {
				start = pos
				continue
			}

			m, ok := connMap[c]
			evil.Assert(ok, "invalid connection")

			res[pos] = Pipe{
				connections: [2]space.Point{
					pos.Add(m[0]),
					pos.Add(m[1]),
				},
			}
		}

		y++
	}

	sPipe := Pipe{}
	connIdx := 0
	for _, dir := range [4]space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
		for _, conn := range res[start.Add(dir)].connections {
			if conn == start {
				sPipe.connections[connIdx] = start.Add(dir)
				connIdx++
				break
			}
		}
	}
	evil.Assert(connIdx == 2, "start is not part of a loop")
	res[start] = sPipe

	return res, start
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	pipes, start := parse(filename)
	sPipe, ok := pipes[start]
	evil.Assert(ok, "missing starting pipe")

	// Part 1
	prev := start
	cur := sPipe.next(sPipe.connections[0])
	total := 1
	for cur != start {
		nPipe, ok := pipes[cur]
		evil.Assert(ok, "missing pipe at ", cur)
		prev, cur = cur, nPipe.next(prev)
		total++
	}
	log.Part1(total / 2)
}
