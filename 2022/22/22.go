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
}
