package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) set.Set[space.Point] {
	res := set.New[space.Point]()
	y := 0
	for line := range load.File(filename) {
		for x, c := range line {
			if c == '#' {
				res.Add(space.Point{X: x, Y: y})
			}
		}
		y++
	}
	return res
}

func sim(elves set.Set[space.Point], limit int) (set.Set[space.Point], int) {
	dirs := [...][3]space.Point{
		{{Y: -1}, {Y: -1, X: -1}, {Y: -1, X: 1}},
		{{Y: 1}, {Y: 1, X: -1}, {Y: 1, X: 1}},
		{{X: -1}, {X: -1, Y: -1}, {X: -1, Y: 1}},
		{{X: 1}, {X: 1, Y: -1}, {X: 1, Y: 1}},
	}
	dirIndex := 0

	m := 0
	for ; limit < 0 || m < limit; m++ {
		proposed := map[space.Point]space.Point{}
		propCount := map[space.Point]int{}
		for elf := range elves {
			proposed[elf] = elf

			move := elves.Intersects(
				elf.Add(space.Point{-1, -1}),
				elf.Add(space.Point{0, -1}),
				elf.Add(space.Point{1, -1}),
				elf.Add(space.Point{1, 0}),
				elf.Add(space.Point{1, 1}),
				elf.Add(space.Point{0, 1}),
				elf.Add(space.Point{-1, 1}),
				elf.Add(space.Point{-1, 0}),
			)
			if !move {
				continue
			}

			for di := 0; di < 4; di++ {
				dir := dirs[(dirIndex+di)%4]

				move := !elves.Intersects(
					elf.Add(dir[0]),
					elf.Add(dir[1]),
					elf.Add(dir[2]),
				)
				if move {
					to := elf.Add(dir[0])
					proposed[elf] = to
					propCount[to]++
					break
				}
			}
		}

		if len(propCount) == 0 {
			break
		}

		elves = set.New[space.Point]()
		for elf, to := range proposed {
			if propCount[to] > 1 {
				elves.Add(elf)
			} else {
				elves.Add(to)
			}
		}

		dirIndex++
	}

	return elves, m
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	elves := parse(filename)

	// Part 1
	elves10, _ := sim(elves, 10)

	aabb := space.AABB{}
	for elf := range elves10 {
		aabb = aabb.Add(elf)
	}

	count := 0
	for x := aabb.Min.X; x <= aabb.Max.X; x++ {
		for y := aabb.Min.Y; y <= aabb.Max.Y; y++ {
			if !elves10.Contains(space.Point{X: x, Y: y}) {
				count++
			}
		}
	}
	log.Part1(count)

	// Part 2
	_, moves := sim(elves, -1)
	log.Part2(moves + 1)
}
