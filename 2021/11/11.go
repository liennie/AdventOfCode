package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) [][]int {
	res := [][]int{}

	for line := range load.File(filename) {
		res = append(res, util.Split(line, ""))
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

func contains(flashes []util.Point, p util.Point) bool {
	for _, flash := range flashes {
		if flash.Equals(p) {
			return true
		}
	}
	return false
}

func step(octopuses [][]int) int {
	flashes := []util.Point{}
	for i := range octopuses {
		for j := range octopuses[i] {
			if inBounds(octopuses, i, j) {
				octopuses[i][j]++
				if octopuses[i][j] > 9 {
					flashes = append(flashes, util.Point{Y: i, X: j})
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
				if octopuses[i][j] > 9 && !contains(flashes, util.Point{Y: i, X: j}) {
					flashes = append(flashes, util.Point{Y: i, X: j})
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
	defer util.Recover(log.Err)

	const filename = "input.txt"

	octopuses := parse(filename)

	// Part 1
	total := 0
	for i := 0; i < 100; i++ {
		total += step(octopuses)
	}
	log.Part1(total)
}
