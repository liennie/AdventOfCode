package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Robot struct {
	pos, vel space.Point
}

func parse(filename string) []Robot {
	re := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	var res []Robot
	for line := range load.File(filename) {
		match := re.FindStringSubmatch(line)
		evil.Assert(match != nil, "line '", line, "' does not match regexp '", re.String(), "'")

		res = append(res, Robot{
			pos: space.Point{
				X: evil.Atoi(match[1]),
				Y: evil.Atoi(match[2]),
			},
			vel: space.Point{
				X: evil.Atoi(match[3]),
				Y: evil.Atoi(match[4]),
			},
		})
	}
	return res
}

func printRobots(robots []Robot, max space.Point) {
	cnts := map[space.Point]int{}
	for _, robot := range robots {
		cnts[robot.pos]++
	}

	fmt.Println(strings.Repeat("=", max.X))
	for y := range max.Y {
		for x := range max.X {
			if n, ok := cnts[space.Point{x, y}]; ok {
				if n < 10 {
					fmt.Print(strconv.Itoa(n))
				} else {
					fmt.Print("X")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func medNearest(robots []Robot) float64 {
	dist := make([]float64, len(robots))
	for i, robot := range robots {
		min := math.Inf(1)
		for j, other := range robots {
			if i == j {
				continue
			}

			diff := robot.pos.Sub(other.pos)
			dist := math.Sqrt(float64(diff.X)*float64(diff.X) + float64(diff.Y)*float64(diff.Y))
			if dist < min {
				min = dist
			}
		}
		dist[i] = min
	}

	slices.Sort(dist)
	return dist[len(dist)/2]
}

func moveRobots(robots []Robot, by int, max space.Point) {
	for i := range robots {
		robot := &robots[i]
		robot.pos = robot.pos.Add(robot.vel.Scale(by)).Mod(max)
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	robots := parse(filename)

	max := space.Point{X: 11, Y: 7}
	if filename == "input.txt" {
		max = space.Point{X: 101, Y: 103}
	}

	// Part 1
	futureRobots := slices.Clone(robots)
	moveRobots(futureRobots, 100, max)

	mid := max.Div(space.Point{X: 2, Y: 2})
	f := [4]int{}
	for _, robot := range futureRobots {
		if robot.pos.X == mid.X || robot.pos.Y == mid.Y {
			continue
		}

		i := 0
		if robot.pos.X > mid.X {
			i |= 1
		}
		if robot.pos.Y > mid.Y {
			i |= 2
		}

		f[i]++
	}
	log.Part1(ints.Product(f[:]...))

	// Part 2
	t := 0
	min := math.Inf(1)
	minT := 0
	for range ints.LCM(max.X, max.Y) {
		moveRobots(robots, 1, max)
		t++

		m := medNearest(robots)
		if m < min {
			min = m
			minT = t
		}
	}
	log.Part2(minT)
}
