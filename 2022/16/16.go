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

func maximize(valves map[string]*valve, time int, cur string, closed set.String) int {
	max := 0
	nextClosed := closed.Clone()
	for c := range closed {
		if time == 30 {
			log.Println(c)
		}

		_, dist, err := path.Shortest(graph(valves), cur, path.EndConst(c))
		if err != nil {
			evil.Panic("path error: %w", err)
		}

		rem := time - dist - 1
		if rem < 0 {
			continue
		}

		nextClosed.Remove(c)
		gain := (rem * valves[c].rate) + maximize(valves, rem, c, nextClosed)
		nextClosed.Add(c)

		if gain > max {
			max = gain
		}
	}
	return max
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	valves := parse(filename)

	// Part 1
	closed := set.String{}
	for key, valve := range valves {
		if valve.rate > 0 {
			closed.Add(key)
		}
	}
	log.Part1(maximize(valves, 30, "AA", closed))
}
