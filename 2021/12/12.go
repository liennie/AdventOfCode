package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type node struct {
	name      string
	connected []*node
}

func (n *node) isBig() bool {
	return n.name == strings.ToUpper(n.name)
}

func parse(filename string) *node {
	nodes := map[string]*node{}

	get := func(name string) *node {
		if n, ok := nodes[name]; ok {
			return n
		}
		n := &node{
			name: name,
		}
		nodes[name] = n
		return n
	}

	for line := range load.File(filename) {
		names := strings.SplitN(line, "-", 2)
		a := get(names[0])
		b := get(names[1])
		a.connected = append(a.connected, b)
		b.connected = append(b.connected, a)
	}

	if start, ok := nodes["start"]; ok {
		return start
	}

	evil.Panic("No start")
	return nil
}

func countPaths(n *node, visited map[*node]bool, canRevisit bool) int {
	if n.name == "end" {
		return 1
	}

	visitedCopy := map[*node]bool{}
	for n, b := range visited {
		visitedCopy[n] = b
	}
	visited = visitedCopy

	if !n.isBig() {
		visited[n] = true
	}

	total := 0
	for _, con := range n.connected {
		vis := visited[con]
		if (!vis || canRevisit) && con.name != "start" {
			total += countPaths(con, visited, (!vis && canRevisit))
		}
	}
	return total
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	start := parse(filename)

	// Part 1
	log.Part1(countPaths(start, nil, false))

	// Part 2
	log.Part2(countPaths(start, nil, true))
}
