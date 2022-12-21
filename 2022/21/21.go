package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type operation interface {
	do(map[string]*monkey) (int, bool)
	eq(map[string]*monkey, int)
	reset()
}

type constOp int

func (op constOp) do(map[string]*monkey) (int, bool) {
	return int(op), true
}

func (op constOp) eq(_ map[string]*monkey, e int) {
	if e != int(op) {
		evil.Panic("%d != %d", e, int(op))
	}
}

func (op constOp) reset() {}

type mathOp struct {
	a, b string
	m    string

	cacheRes int
	cacheOk  bool
	isCache  bool
}

func (op *mathOp) do(monkeys map[string]*monkey) (res int, ok bool) {
	if op.isCache {
		return op.cacheRes, op.cacheOk
	}
	defer func() {
		op.cacheRes = res
		op.cacheOk = ok
		op.isCache = true
	}()

	a, aok := monkeys[op.a].op.do(monkeys)
	b, bok := monkeys[op.b].op.do(monkeys)
	if !aok || !bok {
		return 0, false
	}

	switch op.m {
	case "+":
		return a + b, true
	case "-":
		return a - b, true
	case "*":
		return a * b, true
	case "/":
		return a / b, true
	default:
		evil.Panic("Invalid operation %q", op.m)
	}
	return 0, false
}

func (op *mathOp) eq(monkeys map[string]*monkey, e int) {
	a, aok := monkeys[op.a].op.do(monkeys)
	b, bok := monkeys[op.b].op.do(monkeys)

	if aok && bok {
		return
	}

	if !aok {
		switch op.m {
		case "+":
			monkeys[op.a].op.eq(monkeys, e-b)
		case "-":
			monkeys[op.a].op.eq(monkeys, e+b)
		case "*":
			monkeys[op.a].op.eq(monkeys, e/b)
		case "/":
			monkeys[op.a].op.eq(monkeys, e*b)
		case "=":
			monkeys[op.a].op.eq(monkeys, b)
		}
	} else {
		switch op.m {
		case "+":
			monkeys[op.b].op.eq(monkeys, e-a)
		case "-":
			monkeys[op.b].op.eq(monkeys, a-e)
		case "*":
			monkeys[op.b].op.eq(monkeys, e/a)
		case "/":
			monkeys[op.b].op.eq(monkeys, a/e)
		case "=":
			monkeys[op.b].op.eq(monkeys, a)
		}
	}
}

func (op *mathOp) reset() {
	op.isCache = false
}

type calcOp struct {
	int
}

func (op *calcOp) do(map[string]*monkey) (int, bool) {
	return 0, false
}

func (op *calcOp) eq(_ map[string]*monkey, e int) {
	op.int = e
}

func (op *calcOp) reset() {}

type monkey struct {
	op operation
}

func parse(filename string) map[string]*monkey {
	res := map[string]*monkey{}
	for line := range load.File(filename) {
		name, op, _ := strings.Cut(line, ": ")

		args := strings.Split(op, " ")
		if len(args) == 1 {
			res[name] = &monkey{
				op: constOp(evil.Atoi(args[0])),
			}
		} else if len(args) == 3 {
			res[name] = &monkey{
				op: &mathOp{a: args[0], m: args[1], b: args[2]},
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
	root, _ := monkeys["root"].op.do(monkeys)
	log.Part1(root)

	// Part 2
	for _, m := range monkeys {
		m.op.reset()
	}

	monkeys["root"].op.(*mathOp).m = "="
	monkeys["humn"].op = &calcOp{}

	monkeys["root"].op.eq(monkeys, 0)
	log.Part2(monkeys["humn"].op.(*calcOp).int)
}
