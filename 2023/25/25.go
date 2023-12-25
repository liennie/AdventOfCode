package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/path"
	"github.com/liennie/AdventOfCode/pkg/set"
)

func parse(filename string) map[string]set.String {
	res := map[string]set.String{}
	for line := range load.File(filename) {
		name, connections, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "missing colon")

		res[name] = set.New(strings.Fields(connections)...)
	}
	for k, s := range res {
		for v := range s {
			if res[v] == nil {
				res[v] = set.New[string]()
			}
			res[v].Add(k)
		}
	}
	return res
}

type conn struct {
	a, b string
	v    int
}

func (c conn) sorted() conn {
	if c.a > c.b {
		return conn{
			a: c.b,
			b: c.a,
		}
	}
	return c
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	connections := parse(filename)

	graph := path.GraphFunc[string](func(n string) []path.Edge[string] {
		res := []path.Edge[string]{}
		for c := range connections[n] {
			res = append(res, path.Edge[string]{
				Len: 1,
				To:  c,
			})
		}
		return res
	})

	// Part 1
	log.Println("looking for paths")
	shortest := map[conn][]string{}
	cnts := map[conn]int{}
	i := 0
	for a := range connections {
		i++
		log.Printf("%d/%d", i, len(connections))

		for b := range connections {
			var s []string
			if ss, ok := shortest[conn{a, b, 0}.sorted()]; ok {
				s = ss
			} else {
				p, _, err := path.Shortest(graph, a, path.EndConst(b))
				evil.Assert(err == nil, "path not found from ", a, " to ", b)

				for i, c := range p {
					shortest[conn{c, b, 0}.sorted()] = p[i:]
				}
				s = p
			}

			for i := 1; i < len(s); i++ {
				cnts[conn{s[i], s[i-1], 0}.sorted()]++
			}
		}
	}

	log.Println("disconnecting")
	conns := []conn{}
	for k, v := range cnts {
		k.v = v
		conns = append(conns, k)
	}
	slices.SortFunc(conns, func(a, b conn) int { return -cmp.Compare(a.v, b.v) })
	for _, c := range conns[:3] {
		connections[c.a].Remove(c.b)
		connections[c.b].Remove(c.a)
	}

	_, _, err := path.Shortest(graph, conns[0].a, path.EndConst(conns[0].b))
	evil.Assert(err == path.ErrNotFound, "disconnected wrong components")

	log.Println("flooding")
	groups := [2]set.String{{}, {}}
	start := [2]string{conns[0].a, conns[0].b}
	for i := 0; i < 2; i++ {
		components := []string{start[i]}
		for len(components) > 0 {
			component := components[0]
			components = components[1:]

			groups[i].Add(component)

			for c := range connections[component] {
				if !groups[i].Contains(c) {
					components = append(components, c)
				}
			}
		}
	}

	log.Part1(len(groups[0]) * len(groups[1]))
}
