package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type stack struct {
	crates []string
}

func move(count int, from, to *stack) {
	for i := 0; i < count; i++ {
		last := len(from.crates) - 1
		to.crates = append(to.crates, from.crates[last])
		from.crates = from.crates[:last]
	}
}

type step struct {
	count int
	from  int
	to    int
}

func parse(filename string) ([]stack, []step) {
	ch := load.File(filename)

	stacks := []stack{}
	for line := range ch {
		if line == "" {
			break
		}

		for i := 0; i*4 < len(line); i++ {
			if len(stacks) <= i {
				stacks = append(stacks, stack{})
			}

			pos := i * 4
			if line[pos] == '[' && line[pos+2] == ']' {
				stacks[i].crates = append(stacks[i].crates, string(line[pos+1]))
			}
		}
	}
	for _, stack := range stacks {
		for i, j := 0, len(stack.crates)-1; i < j; i, j = i+1, j-1 {
			stack.crates[i], stack.crates[j] = stack.crates[j], stack.crates[i]
		}
	}

	steps := []step{}
	for line := range ch {
		count, line, _ := strings.Cut(line, " from ")
		count = strings.TrimPrefix(count, "move ")
		from, to, _ := strings.Cut(line, " to ")

		steps = append(steps, step{
			count: util.Atoi(count),
			from:  util.Atoi(from) - 1,
			to:    util.Atoi(to) - 1,
		})
	}

	return stacks, steps
}

func top(stacks []stack) string {
	res := &strings.Builder{}
	for _, stack := range stacks {
		res.WriteString(stack.crates[len(stack.crates)-1])
	}
	return res.String()
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	stacks, steps := parse(filename)

	// Part 1
	for _, step := range steps {
		move(step.count, &stacks[step.from], &stacks[step.to])
	}
	log.Part1(top(stacks))
}
