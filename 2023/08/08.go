package main

import (
	"regexp"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

const (
	Left = iota
	Right
)

type Node struct {
	left, right string
}

func parse(filename string) ([]int, map[string]Node) {
	ch := load.File(filename)
	sdirs := <-ch
	dirs := make([]int, len(sdirs))
	for i, c := range sdirs {
		switch c {
		case 'L':
			dirs[i] = Left
		case 'R':
			dirs[i] = Right
		default:
			evil.Panic("invalid dir %q", c)
		}
	}
	<-ch // skip empty line

	nodes := map[string]Node{}
	nodeRe := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	for line := range ch {
		n := nodeRe.FindStringSubmatch(line)
		evil.Assert(len(n) == 4, "%q did not match", line)
		nodes[n[1]] = Node{
			left:  n[2],
			right: n[3],
		}
	}

	return dirs, nodes
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	dirs, nodes := parse(filename)

	// Part 1
	if _, ok := nodes["AAA"]; ok {
		cur := "AAA"
		cnt := 0
		dirIdx := 0
		for cur != "ZZZ" {
			dir := dirs[dirIdx]
			dirIdx++
			if dirIdx == len(dirs) {
				dirIdx = 0
			}

			node, ok := nodes[cur]
			evil.Assert(ok, "missing node %s", cur)

			switch dir {
			case Left:
				cur = node.left
			case Right:
				cur = node.right
			default:
				evil.Panic("invalid dir %d", dir)
			}
			cnt++
		}
		log.Part1(cnt)
	} else {
		log.Part1("Missing node AAA, skipping part 1")
	}

	// Part 2
	starts := []string{}
	for n := range nodes {
		if strings.HasSuffix(n, "A") {
			starts = append(starts, n)
		}
	}
	total := 1
	for _, start := range starts {
		cur := start
		cnt := 0
		dirIdx := 0
		for !strings.HasSuffix(cur, "Z") {
			dir := dirs[dirIdx]
			dirIdx++
			if dirIdx == len(dirs) {
				dirIdx = 0
			}

			node, ok := nodes[cur]
			evil.Assert(ok, "missing node %s", cur)

			switch dir {
			case Left:
				cur = node.left
			case Right:
				cur = node.right
			default:
				evil.Panic("invalid dir %d", dir)
			}
			cnt++
		}

		total = ints.LCM(total, cnt)
	}
	log.Part2(total)
}
