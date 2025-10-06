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

func inBounds(octopuses [][]int, i, j int) bool {
	if i < 0 || i >= len(octopuses) {
		return false
	}
	if j < 0 || j >= len(octopuses[i]) {
		return false
	}
	return true
}

func contains(flashes []space.Point, p space.Point) bool {
	for _, flash := range flashes {
		if flash == p {
			return true
		}
	}
	return false
}

func step(octopuses [][]int) int {
	flashes := []space.Point{}
	for i := range octopuses {
		for j := range octopuses[i] {
			if inBounds(octopuses, i, j) {
				octopuses[i][j]++
				if octopuses[i][j] > 9 {
					flashes = append(flashes, space.Point{Y: i, X: j})
				}
			}
		}
	}

	for fi := 0; fi < len(flashes); fi++ {
		f := flashes[fi]
		for i := f.Y - 1; i <= f.Y+1; i++ {
			for j := f.X - 1; j <= f.X+1; j++ {
				if i == f.Y && j == f.X {
					continue
				}
				if !inBounds(octopuses, i, j) {
					continue
				}

				octopuses[i][j]++
				if octopuses[i][j] > 9 && !contains(flashes, space.Point{Y: i, X: j}) {
					flashes = append(flashes, space.Point{Y: i, X: j})
				}
			}
		}
	}

	for _, f := range flashes {
		octopuses[f.Y][f.X] = 0
	}

	return len(flashes)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	octopuses := parse(filename)

	// Part 1
	sync := -1
	total := 0
	for i := 0; i < 100; i++ {
		s := step(octopuses)
		total += s
		if s == 100 && sync == -1 {
			sync = i
		}
	}
	log.Part1(total)

	// Part 2
	for i := 100; sync == -1; i++ {
		s := step(octopuses)
		if s == 100 {
			sync = i + 1
		}
	}
	log.Part2(sync)
}
