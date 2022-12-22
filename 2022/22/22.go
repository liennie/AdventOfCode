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

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	mp, inst := parse(filename)
	var pos space.Point
	for x := 0; x <= 10000; x++ {
		if x == 10000 {
			evil.Panic("Start not found")
		}
		if mp[space.Point{X: x}] == '.' {
			pos.X = x
			break
		}
	}
	d := space.Point{X: 1}

	// Part 1
	for _, i := range inst {
		switch i := i.(type) {
		case int:
			for n := 0; n < i; n++ {
				next := pos.Add(d)
				t, ok := mp[next]
				if !ok {
					nd := d.Flip()
					for warp := pos; mp[warp] != 0; warp = warp.Add(nd) {
						next = warp
					}
					t = mp[next]
				}

				if t == '.' {
					pos = next
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
	log.Part1((pos.Y+1)*1000 + (pos.X+1)*4 + dirVal(d))
}
