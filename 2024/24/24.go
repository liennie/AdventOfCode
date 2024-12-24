package main

import (
	"fmt"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func boolValue(value string) bool {
	switch value {
	case "0":
		return false
	case "1":
		return true
	}
	evil.Panic("invalid bool value %s", value)
	return false
}

const (
	OpOR  = "OR"
	OpAND = "AND"
	OpXOR = "XOR"
)

type Gate struct {
	in1, in2 string
	op       string
	out      string
}

func (g Gate) get(wires map[string]bool, gates map[string]Gate) (res bool) {
	defer func() {
		wires[g.out] = res
	}()

	var in1, in2, ok bool
	if in1, ok = wires[g.in1]; !ok {
		if _, ok := gates[g.in1]; !ok {
			evil.Panic("invalid gate %s -> %s", g.in1, g.out)
		}
		in1 = gates[g.in1].get(wires, gates)
	}
	if in2, ok = wires[g.in2]; !ok {
		if _, ok := gates[g.in2]; !ok {
			evil.Panic("invalid gate %s -> %s", g.in2, g.out)
		}
		in2 = gates[g.in2].get(wires, gates)
	}

	switch g.op {
	case OpOR:
		return in1 || in2

	case OpAND:
		return in1 && in2

	case OpXOR:
		return in1 != in2
	}
	evil.Panic("invalid operation %s", g.op)
	return false
}

func parse(filename string) (map[string]bool, map[string]Gate) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	wires := map[string]bool{}
	for line := range <-ch {
		wire, value, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "invalid line format ", line)

		wires[wire] = boolValue(value)
	}

	gates := map[string]Gate{}
	for line := range <-ch {
		fields := strings.Fields(line)
		evil.Assert(len(fields) == 5, "invalid line format ", line)

		in1, op, in2, out := fields[0], fields[1], fields[2], fields[4]
		evil.Assert(op == OpOR || op == OpAND || op == OpXOR, "invalid operation ", op)

		gates[out] = Gate{
			in1: in1,
			in2: in2,
			op:  op,
			out: out,
		}
	}

	return wires, gates
}

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	wires, gates := parse(filename)

	// Part 1
	number := 0
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if _, ok := gates[wire]; !ok {
			break
		}

		number |= b2i(gates[wire].get(wires, gates)) << i
	}
	log.Part1(number)
}
