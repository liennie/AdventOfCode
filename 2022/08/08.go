package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]int {
	res := [][]int{}
	for line := range load.File(filename) {
		res = append(res, evil.Split(line, ""))
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	trees := parse(filename)

	// Part 1 && 2
	visible := 0
	max := 0
	directions := []space.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			vis := false
			score := 1
		dir:
			for _, dir := range directions {
				view := 0

				for xx, yy := x+dir.X, y+dir.Y; yy >= 0 && xx >= 0 && yy < len(trees) && xx < len(trees[yy]); xx, yy = xx+dir.X, yy+dir.Y {
					view++
					if trees[yy][xx] >= trees[y][x] {
						score *= view
						continue dir
					}
				}
				vis = true
				score *= view
			}
			if vis {
				visible++
			}
			if score > max {
				max = score
			}
		}
	}
	log.Part1(visible)
	log.Part2(max)
}
