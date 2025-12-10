package main

import (
	"math"
	"regexp"
	"slices"

	"github.com/liennie/AdventOfCode/pkg/comb"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/seq"
	"github.com/liennie/AdventOfCode/pkg/set"
)

type Machine struct {
	diagram []bool
	buttons [][]int
	joltage []int
}

func parse(filename string) []Machine {
	lineRe := regexp.MustCompile(`^\[([.#]+)\]((?: \([\d,]+\))+) \{([\d,]+)\}$`)
	buttonRe := regexp.MustCompile(`\(([\d,]+)\)`)

	res := []Machine{}
	for line := range load.File(filename) {
		matches := lineRe.FindStringSubmatch(line)
		evil.Assert(matches != nil, "invalid format %q", line)

		m := Machine{}
		for _, d := range matches[1] {
			m.diagram = append(m.diagram, d == '#')
		}

		m.joltage = evil.Split(matches[3], ",")

		buttons := buttonRe.FindAllStringSubmatch(matches[2], -1)
		evil.Assert(buttons != nil, "could not match buttons %q", matches[2])

		for _, b := range buttons {
			m.buttons = append(m.buttons, evil.Split(b[1], ","))
		}

		res = append(res, m)
	}
	return res
}

func addJoltage(joltage []int, button []int, n int) {
	for _, i := range button {
		joltage[i] += n
	}
}

func minJoltagePresses(buttons [][]int, target []int) int {
	target = slices.Clone(target)
	buttons = slices.Clone(buttons)

	max := ints.Sum(target...) + 1

	type group struct {
		buttons [][]int
		idx     int

		common []int
	}

	deadEnd := func(group group, target []int) bool {
		for _, u := range group.common {
			if target[u] < target[group.idx] {
				return true
			}
		}
		return false
	}

	groups := []group{}
	for len(buttons) > 0 {
		n := make([]int, len(target))
		for _, b := range buttons {
			addJoltage(n, b, 1)
		}
		min := math.MaxInt
		idx := -1
		for j := range n {
			if n[j] > 0 {
				ch := comb.Choose(n[j]+target[j]-1, n[j]-1)
				if ch < min {
					min = ch
					idx = j
				}
			}
		}
		evil.Assert(idx != -1, "min not found")

		cur := [][]int{}
		buttons = slices.DeleteFunc(buttons, func(b []int) bool {
			if slices.Contains(b, idx) {
				cur = append(cur, b)
				return true
			}
			return false
		})
		evil.Assert(len(cur) > 0, "no buttons to press")

		groups = append(groups, group{
			buttons: cur,
			idx:     idx,

			common: slices.Collect(set.Intersection(slices.Collect(seq.Map(slices.Values(cur), func(s []int) set.Set[int] {
				return set.New(s...)
			}))...).All()),
		})
	}

	var minPresses func([]group, []int, int) int
	minPresses = func(groups []group, target []int, indent int) (joltage int) {
		if ints.Sum(target...) == 0 {
			return 0
		}
		if len(groups) == 0 {
			return max
		}
		group := groups[0]

		if deadEnd(group, target) {
			return max
		}

		var press func([][]int, int) int
		press = func(buttons [][]int, m int) int {
			button := buttons[0]

			if len(buttons) == 1 {
				addJoltage(target, button, -m)
				if ints.Min(target...) < 0 {
					addJoltage(target, button, m)
					return max
				}

				p := m + minPresses(groups[1:], target, indent+1)
				addJoltage(target, button, m)
				return p
			}

			min := max
			for n := 0; n <= m; n++ {
				addJoltage(target, button, -n)
				p := n + press(buttons[1:], m-n)
				if p < min {
					min = p
				}
				addJoltage(target, button, n)
			}
			return min
		}
		return press(group.buttons, target[group.idx])
	}

	return minPresses(groups, target, 0)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	machines := parse(filename)

	// Part 1
	sum := 0
	for _, m := range machines {
		min := len(m.buttons) + 1
		for c := range comb.Comb(m.buttons) {
			if len(c) >= min {
				continue
			}

			diodes := make([]bool, len(m.diagram))
			for _, b := range c {
				for _, i := range b {
					diodes[i] = !diodes[i]
				}
			}
			if slices.Equal(diodes, m.diagram) {
				min = len(c)
			}
		}
		evil.Assert(min <= len(m.buttons), "solution not found")
		sum += min
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for i, m := range machines {
		presses := minJoltagePresses(m.buttons, m.joltage)
		log.Printf("%d %v %v == %v", i, m.buttons, m.joltage, presses)
		sum += presses
	}
	log.Part2(sum)
}
