package main

import (
	"sort"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) [][]int {
	res := [][]int{}

	for line := range load.File(filename) {
		res = append(res, util.SliceAtoi(strings.Split(line, "")))
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	heightmap := parse(filename)

	// Part 1
	low := []int{}
	lowPoints := []util.Point{}
	for i := range heightmap {
		for j := range heightmap[i] {
			c := heightmap[i][j]

			if (i == 0 || heightmap[i-1][j] > c) &&
				(j == 0 || heightmap[i][j-1] > c) &&
				(i == len(heightmap)-1 || heightmap[i+1][j] > c) &&
				(j == len(heightmap[i])-1 || heightmap[i][j+1] > c) {
				low = append(low, c)
				lowPoints = append(lowPoints, util.Point{
					X: i,
					Y: j,
				})
			}
		}
	}
	log.Part1(util.Sum(low) + len(low))

	// Part 2
	basins := []int{}
	for _, point := range lowPoints {
		basin := map[util.Point]bool{}
		stack := []util.Point{point}

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			basin[cur] = true

			if cur.X > 0 && heightmap[cur.X-1][cur.Y] < 9 {
				new := util.Point{X: cur.X - 1, Y: cur.Y}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.Y > 0 && heightmap[cur.X][cur.Y-1] < 9 {
				new := util.Point{X: cur.X, Y: cur.Y - 1}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.X < len(heightmap)-1 && heightmap[cur.X+1][cur.Y] < 9 {
				new := util.Point{X: cur.X + 1, Y: cur.Y}
				if !basin[new] {
					stack = append(stack, new)
				}
			}
			if cur.Y < len(heightmap[cur.X])-1 && heightmap[cur.X][cur.Y+1] < 9 {
				new := util.Point{X: cur.X, Y: cur.Y + 1}
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
