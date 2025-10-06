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
	OpCONST = ""
	OpOR    = "OR"
	OpAND   = "AND"
	OpXOR   = "XOR"
)

type Gate struct {
	in1, in2 string
	op       string
	out      string
	value    bool
	valueSet bool
}

func (g *Gate) get(gates Gates) (res bool) {
	defer func() {
		g.value = res
		g.valueSet = true
	}()

	if g.valueSet {
		return g.value
	}

	if !g.valueSet && g.op == OpCONST {
		evil.Panic("empty gate")
	}

	if _, ok := gates[g.in1]; !ok {
		evil.Panic("invalid gate %s -> %s", g.in1, g.out)
	}
	in1 := gates[g.in1].get(gates)

	if _, ok := gates[g.in2]; !ok {
		evil.Panic("invalid gate %s -> %s", g.in2, g.out)
	}
	in2 := gates[g.in2].get(gates)

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

func (g *Gate) reset() {
	if g.op != OpCONST {
		g.valueSet = false
	}
}

type Gates map[string]*Gate

func (g Gates) bits() uint8 {
	for i := uint8(0); ; i++ {
		if _, ok := g[fmt.Sprintf("z%02d", i)]; !ok {
			return i
		}
	}
}

func (g Gates) reset() {
	for _, gate := range g {
		gate.reset()
	}
}

func (g Gates) get(n string) uint64 {
	number := uint64(0)
	for i := 0; ; i++ {
		wire := fmt.Sprintf("%s%02d", n, i)
		if _, ok := g[wire]; !ok {
			break
		}
		number |= b2i(g[wire].get(g)) << i
	}
	return number
}

func (g Gates) run() uint64 {
	return g.get("z")
}

func (g Gates) set(n string, v uint64) {
	for i := 0; ; i++ {
		wire := fmt.Sprintf("%s%02d", n, i)
		if _, ok := g[wire]; !ok {
			break
		}

		gate := g[wire]
		if gate.op != OpCONST {
			evil.Panic("trying to set non-const gate %s %s %s -> %s", gate.in1, gate.op, gate.in2, gate.out)
		}

		gate.value = (v & 1) != 0
		v >>= 1
	}
}

func (g Gates) add(x, y uint64) uint64 {
	g.set("x", x)
	g.set("y", y)
	g.reset()
	return g.run()
}

func parse(filename string) Gates {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	gates := Gates{}

	for line := range <-ch {
		out, value, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "invalid line format %q", line)

		gates[out] = &Gate{
			op:       OpCONST,
			value:    boolValue(value),
			valueSet: true,
		}
	}

	for line := range <-ch {
		fields := strings.Fields(line)
		evil.Assert(len(fields) == 5, "invalid line format %q", line)
		evil.Assert(fields[3] == "->", "invalid line format %q", line)

		in1, op, in2, out := fields[0], fields[1], fields[2], fields[4]
		evil.Assert(op == OpOR || op == OpAND || op == OpXOR, "invalid operation %s", op)

		gates[out] = &Gate{
			in1: in1,
			in2: in2,
			op:  op,
			out: out,
		}
	}

	return gates
}

func b2i(v bool) uint64 {
	if v {
		return 1
	}
	return 0
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	gates := parse(filename)

	// Part 1
	log.Part1(gates.run())

	// Part 2
	// solved by hand
}
