package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Pulse struct {
	from, to string
	high     bool
}

type Module interface {
	connectInput(string)
	connectOutput(string)

	pulse(string, bool) []Pulse
}

type FlipFlop struct {
	name  string
	state bool
	out   []string
}

func (f *FlipFlop) connectInput(string) {}

func (f *FlipFlop) connectOutput(out string) {
	f.out = append(f.out, out)
}

func (f *FlipFlop) pulse(from string, high bool) []Pulse {
	if high {
		return nil
	}

	f.state = !f.state

	res := []Pulse{}
	for _, out := range f.out {
		res = append(res, Pulse{
			from: f.name,
			to:   out,
			high: f.state,
		})
	}
	return res
}

type Conjunction struct {
	name string
	in   map[string]bool
	out  []string
}

func (c *Conjunction) connectInput(in string) {
	if c.in == nil {
		c.in = map[string]bool{}
	}
	c.in[in] = false
}

func (c *Conjunction) connectOutput(out string) {
	c.out = append(c.out, out)
}

func (c *Conjunction) pulse(from string, high bool) []Pulse {
	c.in[from] = high

	oh := false
	for _, h := range c.in {
		if !h {
			oh = true
		}
	}

	res := []Pulse{}
	for _, out := range c.out {
		res = append(res, Pulse{
			from: c.name,
			to:   out,
			high: oh,
		})
	}
	return res
}

type Broadcast struct {
	name string
	out  []string
}

func (b *Broadcast) connectInput(string) {}

func (b *Broadcast) connectOutput(out string) {
	b.out = append(b.out, out)
}

func (b *Broadcast) pulse(from string, high bool) []Pulse {
	res := []Pulse{}
	for _, out := range b.out {
		res = append(res, Pulse{
			from: b.name,
			to:   out,
			high: high,
		})
	}
	return res
}

type Sink struct{}

func (*Sink) connectInput(string)        {}
func (*Sink) connectOutput(string)       {}
func (*Sink) pulse(string, bool) []Pulse { return nil }

func parse(filename string) map[string]Module {
	type conn struct {
		from string
		to   []string
	}
	conns := []conn{}
	for line := range load.File(filename) {
		from, tos, ok := strings.Cut(line, "->")
		evil.Assert(ok, "missing arrow")

		c := conn{
			from: strings.TrimSpace(from),
		}
		for _, to := range strings.Split(tos, ",") {
			c.to = append(c.to, strings.TrimSpace(to))
		}
		conns = append(conns, c)
	}

	modules := map[string]Module{}
	for _, c := range conns {
		if name, ok := strings.CutPrefix(c.from, "%"); ok {
			modules[name] = &FlipFlop{name: name}
		} else if name, ok := strings.CutPrefix(c.from, "&"); ok {
			modules[name] = &Conjunction{name: name}
		} else {
			modules[name] = &Broadcast{name: name}
		}
	}

	for _, c := range conns {
		from := strings.TrimLeft(c.from, "%&")
		for _, to := range c.to {
			to = strings.TrimLeft(to, "%&")
			modules[from].connectOutput(to)
			if m, ok := modules[to]; ok {
				m.connectInput(from)
			} else {
				modules[to] = &Sink{}
			}
		}
	}

	return modules
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	modules := parse(filename)

	// Part 1
	total := map[bool]int{}
	for n := 0; n < 1000; n++ {
		pulses := []Pulse{{to: "broadcaster", high: false}}

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]

			total[pulse.high]++

			pulses = append(pulses, modules[pulse.to].pulse(pulse.from, pulse.high)...)
		}
	}
	log.Part1(total[false] * total[true])

	// Part 2
	lcm := 1
	// "input.lcm" was created manually by converting the input into dot
	// and analyzing the rendered graph
	for line := range load.File("input.lcm") {
		lcm = ints.LCM(lcm, evil.Atoi(line))
	}
	log.Part2(lcm)
}
