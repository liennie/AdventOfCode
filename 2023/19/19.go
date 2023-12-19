package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Part map[string]int

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
			evil.Assert(ok)

			next = w.eval(part)
		}

		if next == "A" {
			for _, v := range part {
				sum += v
			}
		}
	}
	log.Part1(sum)
}
