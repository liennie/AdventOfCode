package main

import (
	"fmt"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]byte {
	res := [][]byte{}

	for line := range load.File(filename) {
		res = append(res, []byte(line))
	}

	return res
}

func roll(rocks [][]byte, dir space.Point) {
	aabb := space.NewAABB(space.Point{}, space.Point{X: len(rocks[0]) - 1, Y: len(rocks) - 1})

	var xFrom, xTo, xDir, yFrom, yTo, yDir int
	if dir.X > 0 {
		xFrom, xTo, xDir = aabb.Max.X, aabb.Min.X-1, -1
	} else {
		xFrom, xTo, xDir = aabb.Min.X, aabb.Max.X+1, 1
	}
	if dir.Y > 0 {
		yFrom, yTo, yDir = aabb.Max.Y, aabb.Min.Y-1, -1
	} else {
		yFrom, yTo, yDir = aabb.Min.Y, aabb.Max.Y+1, 1
	}

	for x := xFrom; x != xTo; x += xDir {
		for y := yFrom; y != yTo; y += yDir {
			if rocks[y][x] == 'O' {
				rollTo := space.Point{X: x, Y: y}
				for next := rollTo.Add(dir); aabb.Contains(next) && rocks[next.Y][next.X] == '.'; next = rollTo.Add(dir) {
					rollTo = next
				}
				rocks[rollTo.Y][rollTo.X], rocks[y][x] = rocks[y][x], rocks[rollTo.Y][rollTo.X]
			}
		}
	}
}

func print(rocks [][]byte) {
	for _, line := range rocks {
		log.Printf("%c", line)
	}
}

func northLoad(rocks [][]byte) int {
	sum := 0
	for y := range rocks {
		for x := range rocks[y] {
			if rocks[y][x] == 'O' {
				sum += len(rocks) - y
			}
		}
	}
	return sum
}

func cycle(rocks [][]byte) {
	roll(rocks, space.Point{Y: -1})
	roll(rocks, space.Point{X: -1})
	roll(rocks, space.Point{Y: 1})
	roll(rocks, space.Point{X: 1})
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	rocks := parse(filename)

	// Part 1
	roll(rocks, space.Point{Y: -1})
	log.Part1(northLoad(rocks))

	// Part 2
	maxCycles := 1000000000
	prev := map[string]int{}
	for n := 1; n <= maxCycles; n++ {
		cycle(rocks)
		state := fmt.Sprint(rocks)
		if p, ok := prev[state]; ok {
			cycleLen := n - p
			if (maxCycles-n)%cycleLen == 0 {
				break
			}
		}
		prev[state] = n
	}
	log.Part2(northLoad(rocks))
}
