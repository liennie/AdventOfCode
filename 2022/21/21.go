package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type monkey struct {
	op func() int
}

func (m *monkey) yell() int {
	return m.op()
}

func parse(filename string) map[string]*monkey {
	res := map[string]*monkey{}
	for line := range load.File(filename) {
		name, op, _ := strings.Cut(line, ": ")

		args := strings.Split(op, " ")
		if len(args) == 1 {
			n := evil.Atoi(args[0])
			res[name] = &monkey{
				op: func() int {
					return n
				},
			}
		} else if len(args) == 3 {
			var mop func(int, int) int
			switch args[1] {
			case "+":
				mop = func(a, b int) int { return a + b }
			case "-":
				mop = func(a, b int) int { return a - b }
			case "*":
				mop = func(a, b int) int { return a * b }
			case "/":
				mop = func(a, b int) int { return a / b }
			default:
				evil.Panic("Invalid operation %q", op)
			}

			res[name] = &monkey{
				op: func() int {
					return mop(res[args[0]].yell(), res[args[2]].yell())
				},
			}
		} else {
			evil.Panic("Invalid operation %q", op)
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	monkeys := parse(filename)

	// Part 1
	log.Part1(monkeys["root"].yell())
}
