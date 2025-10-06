package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

const MinPart = 1
const MaxPart = 4000

type Part map[string]int

type MultiPart map[string]set.Range

func (p MultiPart) amount() int {
	total := 1
	for _, r := range p {
		l := r.Len()
		if l < 0 {
			return 0
		}

		total *= l
	}
	return total
}

type Rule struct {
	cat    string
	op     byte
	thresh int
	next   string
}

func (r Rule) eval(p Part) (string, bool) {
	if r.cat == "" {
		return r.next, true
	}
	switch r.op {
	case '<':
		if p[r.cat] < r.thresh {
			return r.next, true
		}
	case '>':
		if p[r.cat] > r.thresh {
			return r.next, true
		}
	}
	return "", false
}

func (r Rule) multiEval(p MultiPart) (string, MultiPart, MultiPart, bool) {
	if r.cat == "" {
		return r.next, p, nil, true
	}

	nextp := MultiPart{}
	restp := MultiPart{}
	for k, v := range p {
		nextp[k] = v
		restp[k] = v
	}

	switch r.op {
	case '<':
		nextp[r.cat] = nextp[r.cat].Intersection(set.Range{Min: MinPart, Max: r.thresh - 1})
		restp[r.cat] = restp[r.cat].Intersection(set.Range{Min: r.thresh, Max: MaxPart})
		if nextp[r.cat].Len() > 0 {
			return r.next, nextp, restp, true
		}
	case '>':
		nextp[r.cat] = nextp[r.cat].Intersection(set.Range{Min: r.thresh + 1, Max: MaxPart})
		restp[r.cat] = restp[r.cat].Intersection(set.Range{Min: MinPart, Max: r.thresh})
		if nextp[r.cat].Len() > 0 {
			return r.next, nextp, restp, true
		}
	}
	return "", nil, p, false
}

type Workflow struct {
	rules []Rule
}

func (w Workflow) eval(p Part) string {
	for _, rule := range w.rules {
		if next, ok := rule.eval(p); ok {
			return next
		}
	}
	evil.Panic("no rule matched, w: %v, p: %v", w, p)
	return ""
}

type MultiNext struct {
	next string
	part MultiPart
}

func (w Workflow) multiEval(p MultiPart) []MultiNext {
	res := []MultiNext{}
	for _, rule := range w.rules {
		next, nextp, restp, ok := rule.multiEval(p)
		if ok {
			res = append(res, MultiNext{
				next: next,
				part: nextp,
			})
			p = restp
		}
	}
	return res
}

func parseRule(rule string) Rule {
	if cond, next, ok := strings.Cut(rule, ":"); ok {
		if idx := strings.IndexAny(cond, "<>"); idx >= 0 {
			return Rule{
				cat:    cond[:idx],
				op:     cond[idx],
				thresh: evil.Atoi(cond[idx+1:]),
				next:   next,
			}
		}
	}
	return Rule{
		next: rule,
	}
}

func parse(filename string) (map[string]Workflow, []Part) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	workflows := map[string]Workflow{}
	for line := range <-ch {
		name, rules, ok := strings.Cut(line, "{")
		evil.Assert(ok, "invalid format")

		w := Workflow{}
		for _, rule := range strings.Split(strings.TrimSuffix(rules, "}"), ",") {
			w.rules = append(w.rules, parseRule(rule))
		}
		workflows[name] = w
	}

	parts := []Part{}
	for line := range <-ch {
		p := Part{}
		for _, rating := range strings.Split(strings.TrimSuffix(strings.TrimPrefix(line, "{"), "}"), ",") {
			cat, val, ok := strings.Cut(rating, "=")
			evil.Assert(ok, "invalid format")

			p[cat] = evil.Atoi(val)
		}
		parts = append(parts, p)
	}

	return workflows, parts
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	workflows, parts := parse(filename)

	// Part 1
	sum := 0
	for _, part := range parts {
		next := "in"
		for next != "R" && next != "A" {
			w, ok := workflows[next]
			evil.Assert(ok, "missing workflow %s", next)

			next = w.eval(part)
		}

		if next == "A" {
			for _, v := range part {
				sum += v
			}
		}
	}
	log.Part1(sum)

	// Part 2
	accepted := []MultiPart{}
	done := set.String{}
	nexts := []MultiNext{{
		next: "in",
		part: MultiPart{
			"x": set.Range{Min: MinPart, Max: MaxPart},
			"m": set.Range{Min: MinPart, Max: MaxPart},
			"a": set.Range{Min: MinPart, Max: MaxPart},
			"s": set.Range{Min: MinPart, Max: MaxPart},
		},
	}}
	for len(nexts) > 0 {
		next := nexts[0]
		nexts = nexts[1:]

		if next.next == "R" {
			continue
		}
		if next.next == "A" {
			accepted = append(accepted, next.part)
			continue
		}

		evil.Assert(!done.Contains(next.next), "loop detected")

		w, ok := workflows[next.next]
		evil.Assert(ok, "missing workflow %s", next.next)

		nexts = append(nexts, w.multiEval(next.part)...)
		done.Add(next.next)
	}
	sum = 0
	for _, p := range accepted {
		sum += p.amount()
	}
	log.Part2(sum)
}
