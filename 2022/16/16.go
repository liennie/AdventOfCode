package main

import (
	"regexp"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
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

func maximize(valves map[string]*valve, time int, cur string, closed []string) int {
	max := 0
	for i, c := range closed {
		if c == "" {
			continue
		}
		closed[i] = ""

		dist := dists[distKey{from: cur, to: c}]

		rem := time - dist - 1
		if rem < 0 {
			closed[i] = c
			continue
		}

		gain := (rem * valves[c].rate) + maximize(valves, rem, c, closed)

		if gain > max {
			max = gain
		}

		closed[i] = c
	}
	return max
}

type max2cachekey struct {
	t1, t2 int
	c1, c2 string
	closed string
}

var max2cache = map[max2cachekey]int{}

func maximize2(valves map[string]*valve, time1, time2 int, cur1, cur2 string, closed []string) (max int) {
	if time1 < time2 {
		time1, time2 = time2, time1
		cur1, cur2 = cur2, cur1
	}
	m2key := max2cachekey{
		t1: time1, t2: time2,
		c1: cur1, c2: cur2,
		closed: strings.Join(closed, ""),
	}
	if res, ok := max2cache[m2key]; ok {
		return res
	}
	defer func() {
		max2cache[m2key] = max
	}()

	for i := 0; i < 2; i++ {
		for i, c := range closed {
			if c == "" {
				continue
			}

			closed[i] = ""

			dist := dists[distKey{from: cur1, to: c}]

			rem := time1 - dist - 1
			if rem < 0 {
				closed[i] = c
				continue
			}

			gain := (rem * valves[c].rate) + maximize2(valves, time2, rem, cur2, c, closed)

			if gain > max {
				max = gain
			}

			closed[i] = c
		}

		time1, time2 = time2, time1
		cur1, cur2 = cur2, cur1
	}

	return max
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	valves := parse(filename)
	calculateDistances(valves)

	// Part 1
	closed := []string{}
	for key, valve := range valves {
		if valve.rate > 0 {
			closed = append(closed, key)
		}
	}
	log.Part1(maximize(valves, 30, "AA", closed))

	// Part 2
	log.Part2(maximize2(valves, 26, 26, "AA", "AA", closed))
}
