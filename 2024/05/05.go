package main

import (
	"slices"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

type Rule struct {
	before, after int
}

func parse(filename string) (set.Set[Rule], [][]int) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	rules := set.New[Rule]()
	for line := range <-ch {
		rule := evil.SplitN(line, "|", 2)
		rules.Add(Rule{rule[0], rule[1]})
	}

	var updates [][]int
	for line := range <-ch {
		updates = append(updates, evil.Split(line, ","))
	}

	return rules, updates
}

func ordered(update []int, rules set.Set[Rule]) bool {
	for i := range len(update) - 1 {
		for j := i + 1; j < len(update); j++ {
			if rules.Contains(Rule{update[j], update[i]}) {
				return false
			}
		}
	}
	return true
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	rules, updates := parse(filename)

	// Part 1
	sum := 0
	for _, update := range updates {
		if ordered(update, rules) {
			sum += update[len(update)/2]
		}
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for _, update := range updates {
		if !ordered(update, rules) {
			slices.SortFunc(update, func(a, b int) int {
				switch {
				case rules.Contains(Rule{a, b}):
					return -1

				case rules.Contains(Rule{b, a}):
					return 1

				default:
					return 0
				}
			})

			sum += update[len(update)/2]
		}
	}
	log.Part2(sum)
}
