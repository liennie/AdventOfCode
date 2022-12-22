package main

import (
	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]byte, []any) {
	ch := load.File(filename)
	defer channel.Drain(ch)

	m := map[space.Point]byte{}
	y := 0
	for line := range ch {
		if line == "" {
			break
		}
		for x, c := range line {
			if c == '.' || c == '#' {
				m[space.Point{X: x, Y: y}] = byte(c)
			}
		}
		y++
	}

	i := []any{}
	a := 0
	for _, c := range <-ch {
		if c >= '0' && c <= '9' {
			a *= 10
			a += int(c - '0')
		} else {
			if a != 0 {
				i = append(i, a)
				a = 0
			}
			i = append(i, byte(c))
		}
	}
	if a != 0 {
		i = append(i, a)
	}

	return m, i
}

func dirVal(d space.Point) int {
	switch d {
	case space.Point{X: 1}:
		return 0
	case space.Point{Y: 1}:
		return 1
	case space.Point{X: -1}:
		return 2
	case space.Point{Y: -1}:
		return 3
	}
	evil.Panic("Invalid dir %v", d)
	return 0
}

type warp struct {
	pos, d space.Point
}

func walk(mp map[space.Point]byte, inst []any, pos, d space.Point, warpMap map[warp]warp) (space.Point, space.Point) {
	for _, i := range inst {
		switch i := i.(type) {
		case int:
			for n := 0; n < i; n++ {
				nextPos := pos.Add(d)
				nextD := d
				t, ok := mp[nextPos]
				if !ok {
					w, ok := warpMap[warp{pos: pos, d: d}]
					if !ok {
						evil.Panic("Oh no")
					}
					nextPos = w.pos
					nextD = w.d
					t = mp[nextPos]
				}

				if t == '.' {
					pos = nextPos
					d = nextD
				} else {
					break
				}
			}

		case byte:
			switch i {
			case 'R':
				d = d.Rot90(1)
			case 'L':
				d = d.Rot90(-1)
			default:
				evil.Panic("Invalid %c", i)
			}

		default:
			evil.Panic("Invalid %T", i)
		}
	}

	return pos, d
}

func addWarps(warpMap map[warp]warp, from, d, to space.Point, rot, cnt int) {
	ff := from
	fd := d.Rot90(1)
	tt := to
	dd := d.Rot90(rot)
	td := dd.Rot90(1)
	for i := 0; i < cnt; i++ {
		warpMap[warp{
			pos: ff,
			d:   d,
		}] = warp{
			pos: tt,
			d:   dd,
		}

		ff = ff.Add(fd)
		tt = tt.Add(td)
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	mp, inst := parse(filename)
	var start space.Point
	for x := 0; x <= 10000; x++ {
		if x == 10000 {
			evil.Panic("Start not found")
		}
		if mp[space.Point{X: x}] == '.' {
			start.X = x
			break
		}
	}

	// Part 1
	warpMap := map[warp]warp{}
	for pos := range mp {
		for _, d := range [...]space.Point{{X: 1}, {Y: 1}, {X: -1}, {Y: -1}} {
			next := pos.Add(d)
			if _, ok := mp[next]; !ok {
				nd := d.Flip()
				for warp := pos; mp[warp] != 0; warp = warp.Add(nd) {
					next = warp
				}
				warpMap[warp{pos: pos, d: d}] = warp{pos: next, d: d}
			}
		}
	}
	pos, d := walk(mp, inst, start, space.Point{X: 1}, warpMap)
	log.Part1((pos.Y+1)*1000 + (pos.X+1)*4 + dirVal(d))

	// Part 2
	if filename == "test.txt" {
		warpMap = map[warp]warp{}
		addWarps(warpMap, space.Point{11, 0}, space.Point{1, 0}, space.Point{15, 11}, 2, 4)
		addWarps(warpMap, space.Point{11, 4}, space.Point{1, 0}, space.Point{15, 8}, 1, 4)
		addWarps(warpMap, space.Point{12, 8}, space.Point{0, -1}, space.Point{11, 7}, -1, 4)
		addWarps(warpMap, space.Point{15, 8}, space.Point{1, 0}, space.Point{11, 3}, 2, 4)

		addWarps(warpMap, space.Point{15, 11}, space.Point{0, 1}, space.Point{0, 4}, -1, 4)
		addWarps(warpMap, space.Point{11, 11}, space.Point{0, 1}, space.Point{0, 7}, 2, 4)
		addWarps(warpMap, space.Point{8, 11}, space.Point{-1, 0}, space.Point{4, 8}, 1, 4)
		addWarps(warpMap, space.Point{7, 7}, space.Point{0, 1}, space.Point{8, 8}, -1, 4)
		addWarps(warpMap, space.Point{3, 7}, space.Point{0, 1}, space.Point{8, 11}, 2, 4)
		addWarps(warpMap, space.Point{0, 7}, space.Point{-1, 0}, space.Point{12, 11}, 1, 4)

		addWarps(warpMap, space.Point{0, 4}, space.Point{0, -1}, space.Point{11, 0}, 2, 4)
		addWarps(warpMap, space.Point{4, 4}, space.Point{0, -1}, space.Point{8, 0}, 1, 4)
		addWarps(warpMap, space.Point{8, 3}, space.Point{-1, 0}, space.Point{7, 4}, -1, 4)
		addWarps(warpMap, space.Point{8, 0}, space.Point{0, -1}, space.Point{3, 4}, 2, 4)
	} else {
		warpMap = map[warp]warp{}
		addWarps(warpMap, space.Point{149, 0}, space.Point{1, 0}, space.Point{99, 149}, 2, 50)
		addWarps(warpMap, space.Point{149, 49}, space.Point{0, 1}, space.Point{99, 99}, 1, 50)
		addWarps(warpMap, space.Point{99, 50}, space.Point{1, 0}, space.Point{100, 49}, -1, 50)
		addWarps(warpMap, space.Point{99, 100}, space.Point{1, 0}, space.Point{149, 49}, 2, 50)

		addWarps(warpMap, space.Point{99, 149}, space.Point{0, 1}, space.Point{49, 199}, 1, 50)
		addWarps(warpMap, space.Point{49, 150}, space.Point{1, 0}, space.Point{50, 149}, -1, 50)

		addWarps(warpMap, space.Point{49, 199}, space.Point{0, 1}, space.Point{149, 0}, 0, 50)
		addWarps(warpMap, space.Point{0, 199}, space.Point{-1, 0}, space.Point{99, 0}, -1, 50)
		addWarps(warpMap, space.Point{0, 149}, space.Point{-1, 0}, space.Point{50, 0}, 2, 50)
		addWarps(warpMap, space.Point{0, 100}, space.Point{0, -1}, space.Point{50, 50}, 1, 50)
		addWarps(warpMap, space.Point{50, 99}, space.Point{-1, 0}, space.Point{49, 100}, -1, 50)
		addWarps(warpMap, space.Point{50, 49}, space.Point{-1, 0}, space.Point{0, 100}, 2, 50)
		addWarps(warpMap, space.Point{50, 0}, space.Point{0, -1}, space.Point{0, 150}, 1, 50)
		addWarps(warpMap, space.Point{100, 0}, space.Point{0, -1}, space.Point{0, 199}, 0, 50)
	}
	pos, d = walk(mp, inst, start, space.Point{X: 1}, warpMap)
	log.Part2((pos.Y+1)*1000 + (pos.X+1)*4 + dirVal(d))
}
