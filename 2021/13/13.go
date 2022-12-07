package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) ([][]bool, []space.Point) {
	ch := load.File(filename)

	points := map[space.Point]bool{}
	maxX := 0
	maxY := 0
	for line := range ch {
		if line == "" {
			break
		}
		coords := evil.SplitN(line, ",", 2)
		p := space.Point{
			X: coords[0],
			Y: coords[1],
		}
		points[p] = true
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	paper := make([][]bool, maxY+1)
	for y := 0; y <= maxY; y++ {
		paper[y] = make([]bool, maxX+1)
	}
	for p := range points {
		paper[p.Y][p.X] = true
	}

	folds := []space.Point{}
	for line := range ch {
		fold := strings.SplitN(strings.TrimPrefix(line, "fold along "), "=", 2)
		switch fold[0] {
		case "x":
			folds = append(folds, space.Point{X: evil.Atoi(fold[1])})
		case "y":
			folds = append(folds, space.Point{Y: evil.Atoi(fold[1])})
		}
	}

	return paper, folds
}

func fold(paper [][]bool, fold space.Point) [][]bool {
	if fold.X != 0 {
		for y := range paper {
			for x := fold.X + 1; x < len(paper[y]); x++ {
				mx := 2*fold.X - x
				paper[y][mx] = paper[y][mx] || paper[y][x]
			}
			paper[y] = paper[y][:fold.X]
		}
	} else {
		for y := fold.Y + 1; y < len(paper); y++ {
			my := 2*fold.Y - y
			for x := range paper[y] {
				paper[my][x] = paper[my][x] || paper[y][x]
			}
		}
		paper = paper[:fold.Y]
	}
	return paper
}

func countDots(paper [][]bool) int {
	count := 0
	for y := range paper {
		for x := range paper[y] {
			if paper[y][x] {
				count++
			}
		}
	}
	return count
}

func printPaper(paper [][]bool) {
	b := &strings.Builder{}
	for _, line := range paper {
		b.Reset()
		for _, dot := range line {
			if dot {
				b.WriteByte('#')
			} else {
				b.WriteByte(' ')
			}
		}
		log.Part2(b.String())
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	paper, folds := parse(filename)

	// Part 1
	paper = fold(paper, folds[0])
	log.Part1(countDots(paper))

	// Part 2
	for i := 1; i < len(folds); i++ {
		paper = fold(paper, folds[i])
	}
	printPaper(paper)
}
