package main

import (
	"cmp"
	"slices"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Pair struct {
	a, b space.Point3
	dist int
}

func parse(filename string) []space.Point3 {
	res := []space.Point3{}
	for line := range load.File(filename) {
		coords := evil.Split(line, ",")
		res = append(res, space.Point3{
			X: coords[0],
			Y: coords[1],
			Z: coords[2],
		})
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	boxes := parse(filename)

	// Part 1
	pairs := []Pair{}
	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, Pair{
				a:    boxes[i],
				b:    boxes[j],
				dist: boxes[i].Sub(boxes[j]).LenSquared(),
			})
		}
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.dist, b.dist)
	})

	circuits := map[space.Point3]set.Set[space.Point3]{}
	for _, box := range boxes {
		circuits[box] = set.New(box)
	}

	upto := 1000
	if filename == "test.txt" {
		upto = 10
	}
	for _, pair := range pairs[:upto] {
		if circuits[pair.a].Contains(pair.b) {
			continue
		}

		connected := set.Union(circuits[pair.a], circuits[pair.b])
		for box := range circuits[pair.a] {
			circuits[box] = connected
		}
		for box := range circuits[pair.b] {
			circuits[box] = connected
		}
	}

	sizes := []int{}
	seen := set.New[space.Point3]()
	for _, circuit := range circuits {
		if seen.IntersectsSeq(circuit.All()) {
			continue
		}

		sizes = append(sizes, len(circuit))
		seen.AddSeq(circuit.All())
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)

	log.Part1(ints.Product(sizes[:3]...))

	// Part 2
	circuits = map[space.Point3]set.Set[space.Point3]{}
	for _, box := range boxes {
		circuits[box] = set.New(box)
	}

	cable := 0
	for _, pair := range pairs {
		if circuits[pair.a].Contains(pair.b) {
			continue
		}

		connected := set.Union(circuits[pair.a], circuits[pair.b])
		for box := range circuits[pair.a] {
			circuits[box] = connected
		}
		for box := range circuits[pair.b] {
			circuits[box] = connected
		}

		if len(connected) == len(boxes) {
			cable = pair.a.X * pair.b.X
			break
		}
	}
	log.Part2(cable)
}
