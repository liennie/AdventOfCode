package main

import (
	"regexp"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/set"
)

type valve struct {
	rate    int
	leadsTo []string
}

var inputRe = regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

func parse(filename string) map[string]*valve {
	valves := map[string]*valve{}
	for line := range load.File(filename) {
		match := inputRe.FindStringSubmatch(line)
		if match == nil {
			evil.Panic("Line %q does not match", line)
		}

		valves[match[1]] = &valve{
			rate:    evil.Atoi(match[2]),
			leadsTo: strings.Split(match[3], ", "),
		}
	}
	return valves
}

func graph(valves map[string]*valve) path.Graph[string] {
	return path.GraphFunc[string](func(n string) []path.Edge[string] {
		res := []path.Edge[string]{}
		for _, l := range valves[n].leadsTo {
			res = append(res, path.Edge[string]{
				Len: 1,
				To:  l,
			})
		}
		return res
	})
}

type distKey struct {
	from, to string
}

var dists = map[distKey]int{}

func calculateDistances(valves map[string]*valve) {
	g := graph(valves)

	for from := range valves {
		for to := range valves {
			_, dist, err := path.Shortest(g, from, path.EndConst(to))
			if err != nil {
				evil.Panic("path error: %w", err)
			}

			dists[distKey{from: from, to: to}] = dist
		}
	}
}

func maximize(valves map[string]*valve, time int, cur string, closed set.String) int {
	max := 0
	nextClosed := closed.Clone()
	for c := range closed {
		nextClosed.Remove(c)

		dist := dists[distKey{from: cur, to: c}]

		rem := time - dist - 1
		if rem < 0 {
			continue
		}

		gain := (rem * valves[c].rate) + maximize(valves, rem, c, nextClosed)

		if gain > max {
			max = gain
		}

		nextClosed.Add(c)
	}
	return max
}

func maximize2(valves map[string]*valve, time1, time2 int, cur1, cur2 string, closed set.String) int {
	max := 0
	nextClosed := closed.Clone()
	for c1 := range closed {
		nextClosed.Remove(c1)

		if time1 == 26 {
			log.Println(c1)
		}

		dist1 := dists[distKey{from: cur1, to: c1}]

		rem1 := time1 - dist1 - 1
		if rem1 < 0 {
			continue
		}

		gain := (rem1 * valves[c1].rate) + maximize2(valves, rem1, time2, c1, cur2, nextClosed)

		if gain > max {
			max = gain
		}

		for c2 := range closed {
			if c2 == c1 {
				continue
			}

			if time1 == 26 && time2 == 26 {
				log.Println(c1, c2)
			}

			nextClosed.Remove(c2)

			dist2 := dists[distKey{from: cur2, to: c2}]

			rem2 := time2 - dist2 - 1
			if rem2 < 0 {
				continue
			}

			gain := (rem1 * valves[c1].rate) + (rem2 * valves[c2].rate) + maximize2(valves, rem1, rem2, c1, c2, nextClosed)

			if gain > max {
				max = gain
			}

			nextClosed.Add(c2)
		}

		nextClosed.Add(c1)
	}
	return max
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	valves := parse(filename)
	calculateDistances(valves)

	// Part 1
	closed := set.String{}
	for key, valve := range valves {
		if valve.rate > 0 {
			closed.Add(key)
		}
	}
	log.Part1(maximize(valves, 30, "AA", closed))

	// Part 2
	log.Part2(maximize2(valves, 26, 26, "AA", "AA", closed))
}
