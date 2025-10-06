package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Pipe struct {
	t           rune
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
				t: c,
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
	mask := 0
	for i, dir := range [4]space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
		for _, conn := range res[start.Add(dir)].connections {
			if conn == start {
				sPipe.connections[connIdx] = start.Add(dir)
				connIdx++
				mask |= 1 << i
				break
			}
		}
	}
	evil.Assert(connIdx == 2, "start is not part of a loop")

	var ok bool
	sPipe.t, ok = map[int]rune{
		0b0011: 'L',
		0b0101: '-',
		0b1001: 'F',
		0b0110: 'J',
		0b1010: '|',
		0b1100: '7',
	}[mask]
	evil.Assert(ok, "invalid start shape")

	res[start] = sPipe

	return res, start
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	pipes, start := parse(filename)
	sPipe, ok := pipes[start]
	evil.Assert(ok, "missing starting pipe")

	loop := set.New(
		start.Scale(3).Add(space.Point{X: 1, Y: 1}),
		start.Scale(3).Add(space.Point{X: 1, Y: 1}).Add(connMap[sPipe.t][0]),
		start.Scale(3).Add(space.Point{X: 1, Y: 1}).Add(connMap[sPipe.t][1]),
	)
	minPipe := start
	maxPipe := start

	// Part 1
	prev := start
	cur := sPipe.next(sPipe.connections[0])
	total := 1
	for cur != start {
		nPipe, ok := pipes[cur]

		// Part 2 setup
		loop.Add(
			cur.Scale(3).Add(space.Point{X: 1, Y: 1}),
			cur.Scale(3).Add(space.Point{X: 1, Y: 1}).Add(connMap[nPipe.t][0]),
			cur.Scale(3).Add(space.Point{X: 1, Y: 1}).Add(connMap[nPipe.t][1]),
		)
		minPipe = space.Point{X: min(minPipe.X, cur.X), Y: min(minPipe.Y, cur.Y)}
		maxPipe = space.Point{X: max(maxPipe.X, cur.X), Y: max(maxPipe.Y, cur.Y)}

		// Part 1
		evil.Assert(ok, "missing pipe at %v", cur)
		prev, cur = cur, nPipe.next(prev)
		total++
	}
	log.Part1(total / 2)

	// Part 2
	minBigPipe := minPipe.Scale(3)
	maxBigPipe := maxPipe.Scale(3).Add(space.Point{X: 2, Y: 2})

	outside := loop.Clone()
	flood := set.New(minBigPipe)
	for len(flood) > 0 {
		p, _ := flood.Pop()
		outside.Add(p)

		for _, dir := range [4]space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
			next := p.Add(dir)
			if outside.Contains(next) {
				continue
			}
			if next.X < minBigPipe.X || next.Y < minBigPipe.Y || next.X > maxBigPipe.X || next.Y > maxBigPipe.Y {
				continue
			}

			flood.Add(next)
		}
	}

	inside := 0
	for x := minPipe.X; x <= maxPipe.X; x++ {
		for y := minPipe.Y; y <= maxPipe.Y; y++ {
			if !outside.Contains(space.Point{X: x, Y: y}.Scale(3).Add(space.Point{X: 1, Y: 1})) {
				inside++
			}
		}
	}
	log.Part2(inside)
}
