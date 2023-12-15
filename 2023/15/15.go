package main

import (
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Lens struct {
	label string
	focal int
}

func parse(filename string) []string {
	res := []string{}
	for line := range load.File(filename) {
		res = append(res, strings.Split(line, ",")...)
	}
	return res
}

func hash(s string) int {
	cur := byte(0)
	for _, c := range s {
		evil.Assert(c < 128, "non ASCII character")

		cur += byte(c)
		cur *= 17
	}
	return int(cur)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	seq := parse(filename)

	// Part 1
	sum := 0
	for _, s := range seq {
		sum += hash(s)
	}
	log.Part1(sum)

	// Part 2
	boxes := [256][]Lens{}
	for _, s := range seq {
		if label, ok := strings.CutSuffix(s, "-"); ok {
			i := hash(label)
			boxes[i] = slices.DeleteFunc(boxes[i], func(l Lens) bool { return l.label == label })
		} else if label, n, ok := strings.Cut(s, "="); ok {
			i := hash(label)
			if j := slices.IndexFunc(boxes[i], func(l Lens) bool { return l.label == label }); j >= 0 {
				boxes[i][j].focal = evil.Atoi(n)
			} else {
				boxes[i] = append(boxes[i], Lens{
					label: label,
					focal: evil.Atoi(n),
				})
			}
		} else {
			evil.Panic("invalid step %q", s)
		}
	}

	sum = 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focal
		}
	}
	log.Part2(sum)
}
