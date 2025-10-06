package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
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

func findStart(garden [][]byte) space.Point {
	for y := range garden {
		for x := range garden[y] {
			if garden[y][x] == 'S' {
				garden[y][x] = '.'
				return space.Point{X: x, Y: y}
			}
		}
	}
	evil.Panic("start not found")
	return space.Point{}
}

func walk(garden [][]byte, start set.Set[space.Point], steps int) set.Set[space.Point] {
	plots := start
	for n := 0; n < steps; n++ {
		next := set.New[space.Point]()
		for pos := range plots {
			for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
				n := pos.Add(dir)
				if garden[ints.Mod(n.Y, len(garden))][ints.Mod(n.X, len(garden[0]))] != '#' {
					next.Add(n)
				}
			}
		}
		plots = next
	}
	return plots
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	garden := parse(filename)
	start := findStart(garden)

	// Part 1
	log.Part1(len(walk(garden, set.New(start), 64)))

	// Part 2
	const steps = 26501365

	// all of this just checks that the input behaves nicely
	evil.Assert(len(garden) == len(garden[0]), "not square")
	evil.Assert(len(garden)%2 == 1, "not odd")
	center := space.Point{
		X: len(garden[0]) / 2,
		Y: len(garden) / 2,
	}
	evil.Assert(start == center, "start not in center, start: %v, center: %v", start, center)

	mod := len(garden)

	log.Println("walking 1")
	cycles := []set.Set[space.Point]{}
	cycles = append(cycles, walk(garden, set.New(start), 26501365%mod))

	for p := range cycles[0] {
		evil.Assert(p.X >= 0 && p.X < mod && p.Y >= 0 && p.Y < mod, "out of bounds")
	}

	log.Println("walking 2")
	cycles = append(cycles, walk(garden, cycles[0], mod))
	log.Println("walking 3")
	cycles = append(cycles, walk(garden, cycles[1], mod))
	log.Println("walking 4")
	cycles = append(cycles, walk(garden, cycles[2], mod))

	offsets := [][]space.Point3{
		{
			// 0: center 1
			{Z: 1, X: 0, Y: 0},
			{Z: 2, X: 1, Y: 0},
			{Z: 2, X: 0, Y: -1},
			{Z: 2, X: -1, Y: 0},
			{Z: 2, X: 0, Y: 1},
			{Z: 3, X: 0, Y: 0},
			{Z: 3, X: 2, Y: 0},
			{Z: 3, X: 1, Y: -1},
			{Z: 3, X: 0, Y: -2},
			{Z: 3, X: -1, Y: -1},
			{Z: 3, X: -2, Y: 0},
			{Z: 3, X: -1, Y: 1},
			{Z: 3, X: 0, Y: 2},
			{Z: 3, X: 1, Y: 1},
		},
		{
			// 1: center 2
			{Z: 2, X: 0, Y: 0},
			{Z: 3, X: 1, Y: 0},
			{Z: 3, X: 0, Y: -1},
			{Z: 3, X: -1, Y: 0},
			{Z: 3, X: 0, Y: 1},
		},
		{
			// 2: right corner
			{Z: 1, X: 1, Y: 0},
			{Z: 2, X: 2, Y: 0},
			{Z: 3, X: 3, Y: 0},
		},
		{
			// 3: top corner
			{Z: 1, X: 0, Y: -1},
			{Z: 2, X: 0, Y: -2},
			{Z: 3, X: 0, Y: -3},
		},
		{
			// 4: left corner
			{Z: 1, X: -1, Y: 0},
			{Z: 2, X: -2, Y: 0},
			{Z: 3, X: -3, Y: 0},
		},
		{
			// 5: bottom corner
			{Z: 1, X: 0, Y: 1},
			{Z: 2, X: 0, Y: 2},
			{Z: 3, X: 0, Y: 3},
		},
		{
			// 6: top-right outer edge
			{Z: 1, X: 1, Y: -1},
			{Z: 2, X: 2, Y: -1},
			{Z: 2, X: 1, Y: -2},
			{Z: 3, X: 3, Y: -1},
			{Z: 3, X: 2, Y: -2},
			{Z: 3, X: 1, Y: -3},
		},
		{
			// 7: top-left outer edge
			{Z: 1, X: -1, Y: -1},
			{Z: 2, X: -1, Y: -2},
			{Z: 2, X: -2, Y: -1},
			{Z: 3, X: -1, Y: -3},
			{Z: 3, X: -2, Y: -2},
			{Z: 3, X: -3, Y: -1},
		},
		{
			// 8: bottom-left outer edge
			{Z: 1, X: -1, Y: 1},
			{Z: 2, X: -2, Y: 1},
			{Z: 2, X: -1, Y: 2},
			{Z: 3, X: -3, Y: 1},
			{Z: 3, X: -2, Y: 2},
			{Z: 3, X: -1, Y: 3},
		},
		{
			// 9: bottom-right outer edge
			{Z: 1, X: 1, Y: 1},
			{Z: 2, X: 1, Y: 2},
			{Z: 2, X: 2, Y: 1},
			{Z: 3, X: 1, Y: 3},
			{Z: 3, X: 2, Y: 2},
			{Z: 3, X: 3, Y: 1},
		},
		{
			// 10: top-right inner edge
			{Z: 2, X: 1, Y: -1},
			{Z: 3, X: 2, Y: -1},
			{Z: 3, X: 1, Y: -2},
		},
		{
			// 11: top-left inner edge
			{Z: 2, X: -1, Y: -1},
			{Z: 3, X: -1, Y: -2},
			{Z: 3, X: -2, Y: -1},
		},
		{
			// 12: bottom-left inner edge
			{Z: 2, X: -1, Y: 1},
			{Z: 3, X: -2, Y: 1},
			{Z: 3, X: -1, Y: 2},
		},
		{
			// 13: bottom-right inner edge
			{Z: 2, X: 1, Y: 1},
			{Z: 3, X: 1, Y: 2},
			{Z: 3, X: 2, Y: 1},
		},
	}

	for g, group := range offsets {
		log.Printf("checking group %d", g)
		for y := range garden {
			for x := range garden[y] {
				p := space.Point{X: x, Y: y}
				var want bool
				for i, offset := range group {
					off := space.Point{X: offset.X, Y: offset.Y}.Scale(mod)
					contains := cycles[offset.Z].Contains(p.Add(off))

					if i == 0 {
						want = contains
					} else {
						evil.Assert(contains == want, "misbehaving input, x: %d, y: %d, g: %d, z: %d, off: %v, contains: %v, want: %v ", x, y, g, offset.Z, off, contains, want)
					}
				}
			}
		}
	}

	// input behaves nicely :)
	log.Println("all looks good! let's go")

	groupcnts := make([]int, len(offsets))
	for g, group := range offsets {
		for y := range garden {
			for x := range garden[y] {
				p := space.Point{X: x, Y: y}
				offset := group[0]
				off := space.Point{X: offset.X, Y: offset.Y}.Scale(mod)

				if cycles[offset.Z].Contains(p.Add(off)) {
					groupcnts[g]++
				}
			}
		}
	}

	log.Println(groupcnts)

	reach := steps / mod
	total := 0

	total += groupcnts[0] * reach * reach             // center 1
	total += groupcnts[1] * (reach - 1) * (reach - 1) // center 2
	total += groupcnts[2]                             // right corner
	total += groupcnts[3]                             // top corner
	total += groupcnts[4]                             // left corner
	total += groupcnts[5]                             // bottom corner
	total += groupcnts[6] * reach                     // top-right outer edge
	total += groupcnts[7] * reach                     // top-left outer edge
	total += groupcnts[8] * reach                     // bottom-left outer edge
	total += groupcnts[9] * reach                     // bottom-right outer edge
	total += groupcnts[10] * (reach - 1)              // top-right inner edge
	total += groupcnts[11] * (reach - 1)              // top-left inner edge
	total += groupcnts[12] * (reach - 1)              // bottom-left inner edge
	total += groupcnts[13] * (reach - 1)              // bottom-right inner edge

	log.Part2(total)
}
