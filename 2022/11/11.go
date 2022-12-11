package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"golang.org/x/exp/slices"
)

type operation struct {
	mask   uint8
	first  int
	op     string
	second int
}

func (op operation) eval(i int) int {
	var a, b int
	if op.mask&1 != 0 {
		a = op.first
	} else {
		a = i
	}
	if op.mask&2 != 0 {
		b = op.second
	} else {
		b = i
	}

	switch op.op {
	case "*":
		return a * b
	case "+":
		return a + b
	}
	evil.Panic("Invalid operation %s", op.op)
	return 0
}

type test struct {
	div   int
	true  int
	false int
}

func (t test) next(i int) int {
	if i%t.div == 0 {
		return t.true
	}
	return t.false
}

type monkey struct {
	items     []int
	operation operation
	test      test
}

func (m monkey) clone() monkey {
	return monkey{
		items:     slices.Clone(m.items),
		operation: m.operation,
		test:      m.test,
	}
}

func cloneMonkeys(monkeys []monkey) []monkey {
	res := make([]monkey, len(monkeys))
	for i, monkey := range monkeys {
		res[i] = monkey.clone()
	}
	return res
}

func parse(filename string) []monkey {
	res := []monkey{}
	var last *monkey
	for line := range load.File(filename) {
		if line == "" {
			continue
		} else if suffix := strings.TrimPrefix(line, "Monkey "); suffix != line {
			id := evil.Atoi(strings.TrimSuffix(suffix, ":"))
			if id != len(res) {
				evil.Panic("Monkeys are out of order")
			}
			res = append(res, monkey{})
			last = &res[id]

		} else if suffix := strings.TrimPrefix(line, "  Starting items: "); suffix != line {
			last.items = evil.Split(suffix, ", ")

		} else if suffix := strings.TrimPrefix(line, "  Operation: new = "); suffix != line {
			args := strings.Split(suffix, " ")
			if len(args) != 3 {
				evil.Panic("Invalid operation len %d", len(args))
			}

			var first, second int
			var mask uint8
			if args[0] != "old" {
				first = evil.Atoi(args[0])
				mask |= 1
			}
			if args[2] != "old" {
				second = evil.Atoi(args[2])
				mask |= 2
			}

			last.operation = operation{
				mask:   mask,
				first:  first,
				op:     args[1],
				second: second,
			}

		} else if suffix := strings.TrimPrefix(line, "  Test: divisible by "); suffix != line {
			last.test.div = evil.Atoi(suffix)

		} else if suffix := strings.TrimPrefix(line, "    If true: throw to monkey "); suffix != line {
			last.test.true = evil.Atoi(suffix)

		} else if suffix := strings.TrimPrefix(line, "    If false: throw to monkey "); suffix != line {
			last.test.false = evil.Atoi(suffix)

		} else {
			evil.Panic("Invalid line %q", line)
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	monkeys := parse(filename)
	clone := cloneMonkeys(monkeys)

	// Part 1
	activity := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.items {
				activity[m]++
				item = monkey.operation.eval(item) / 3
				next := monkey.test.next(item)
				monkeys[next].items = append(monkeys[next].items, item)
			}
			monkeys[m].items = monkey.items[:0]
		}
	}
	slices.SortFunc(activity, func(a, b int) bool { return a > b })
	log.Part1(activity[0] * activity[1])

	// Part 2
	monkeys = clone
	mod := 1
	for _, monkey := range monkeys {
		mod = ints.LCM(mod, monkey.test.div)
	}
	activity = make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.items {
				activity[m]++
				item = monkey.operation.eval(item) % mod
				next := monkey.test.next(item)
				monkeys[next].items = append(monkeys[next].items, item)
			}
			monkeys[m].items = monkey.items[:0]
		}
	}
	slices.SortFunc(activity, func(a, b int) bool { return a > b })
	log.Part2(activity[0] * activity[1])
}
