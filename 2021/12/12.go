package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
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

	panic("No start")
}

func countPaths(n *node, visited map[*node]bool) int {
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
		if !visited[con] {
			total += countPaths(con, visited)
		}
	}
	return total
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	start := parse(filename)

	// Part 1
	log.Part1(countPaths(start, nil))
}
