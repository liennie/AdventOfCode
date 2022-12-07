package main

import (
	"sort"

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

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	heightmap := parse(filename)

	// Part 1
	low := []int{}
	lowPoints := []space.Point{}
	for i := range heightmap {
		for j := range heightmap[i] {
			c := heightmap[i][j]

			if (i == 0 || heightmap[i-1][j] > c) &&
				(j == 0 || heightmap[i][j-1] > c) &&
				(i == len(heightmap)-1 || heightmap[i+1][j] > c) &&
				(j == len(heightmap[i])-1 || heightmap[i][j+1] > c) {
				low = append(low, c)
				lowPoints = append(lowPoints, space.Point{
					X: i,
					Y: j,
				})
			}
		}
	}
	log.Part1(ints.Sum(low...) + len(low))

	// Part 2
	basins := []int{}
	for _, point := range lowPoints {
		basin := map[space.Point]bool{}
		stack := []space.Point{point}

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			basin[cur] = true

			if cur.X > 0 && heightmap[cur.X-1][cur.Y] < 9 {
				new := space.Point{X: cur.X - 1, Y: cur.Y}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.Y > 0 && heightmap[cur.X][cur.Y-1] < 9 {
				new := space.Point{X: cur.X, Y: cur.Y - 1}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.X < len(heightmap)-1 && heightmap[cur.X+1][cur.Y] < 9 {
				new := space.Point{X: cur.X + 1, Y: cur.Y}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.Y < len(heightmap[cur.X])-1 && heightmap[cur.X][cur.Y+1] < 9 {
				new := space.Point{X: cur.X, Y: cur.Y + 1}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
		}

		basins = append(basins, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	log.Part2(basins[0] * basins[1] * basins[2])
}
