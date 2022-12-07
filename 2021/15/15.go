package main

import (
	"math"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]int {
	res := [][]int{}

	for line := range load.File(filename) {
		res = append(res, ints.Split(line, ""))
	}

	return res
}

func smallestRisk(risk [][]int) int {
	totalRisk := make([][]int, len(risk))
	for i := range risk {
		totalRisk[i] = make([]int, len(risk[i]))
		for j := range totalRisk[i] {
			totalRisk[i][j] = math.MaxInt
		}
	}
	totalRisk[0][0] = 0

	points := map[space.Point]bool{
		{X: 0, Y: 0}: true,
	}
	for len(points) > 0 {
		mr := math.MaxInt
		p := space.Point{X: len(risk[0]), Y: len(risk)}
		for pp := range points {
			if totalRisk[pp.Y][pp.X] < mr {
				mr = totalRisk[pp.Y][pp.X]
				p = pp
			}
		}
		delete(points, p)

		cur := totalRisk[p.Y][p.X]

		for _, dir := range []space.Point{{Y: -1}, {Y: 1}, {X: -1}, {X: 1}} {
			n := p.Add(dir)
			if n.Y >= 0 && n.X >= 0 && n.Y < len(risk) && n.X < len(risk[n.Y]) &&
				cur+risk[n.Y][n.X] < totalRisk[n.Y][n.X] {
				points[n] = true
				totalRisk[n.Y][n.X] = cur + risk[n.Y][n.X]
			}
		}
	}

	end := space.Point{Y: len(risk) - 1}
	end.X = len(risk[end.Y]) - 1

	return totalRisk[end.Y][end.X]
}

func expand(risk [][]int) [][]int {
	full := make([][]int, len(risk)*5)
	for i := range full {
		ri := i % len(risk)
		imul := i / len(risk)

		full[i] = make([]int, len(risk[ri])*5)
		for j := range full[i] {
			rj := j % len(risk[ri])
			jmul := j / len(risk[ri])

			full[i][j] = (((risk[ri][rj] + (imul + jmul)) - 1) % 9) + 1
		}
	}

	return full
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	risk := parse(filename)

	// Part 1
	log.Part1(smallestRisk(risk))

	// Part 2
	log.Part2(smallestRisk(expand(risk)))
}
