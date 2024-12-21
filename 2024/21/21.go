package main

import (
	"cmp"
	"maps"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) []string {
	return load.Slice(filename)
}

func parseKeypad(keys string) map[rune]space.Point {
	keys = strings.TrimSpace(keys)
	res := map[rune]space.Point{}
	for y, line := range strings.Split(keys, "\n") {
		line = strings.TrimSpace(line)
		for x, r := range line {
			if r == '-' {
				continue
			}
			res[r] = space.Point{X: x, Y: y}
		}
	}
	return res
}

func arrowLine(dir space.Point) string {
	if dir == (space.Point{X: 0, Y: 0}) {
		return ""
	}

	evil.Assert((dir.X == 0) != (dir.Y == 0), "no diagonal vectors are allowed ", dir)

	switch {
	case dir.X > 0:
		return strings.Repeat(">", dir.X)

	case dir.X < 0:
		return strings.Repeat("<", -dir.X)

	case dir.Y > 0:
		return strings.Repeat("v", dir.Y)

	case dir.Y < 0:
		return strings.Repeat("^", -dir.Y)
	}

	evil.Panic("this should not happen")
	return ""
}

func robotMovement(keypad map[rune]space.Point, sequence string) []string {
	validKeys := set.Collect(maps.Values(keypad))

	res := []string{""}
	cur, ok := keypad['A']
	evil.Assert(ok, "keypad is missing A")

	for _, button := range sequence {
		target, ok := keypad[button]
		evil.Assert(ok, "keypad is missing button ", string(button))

		var arrows []string
		movement := target.Sub(cur)

		if movement == (space.Point{X: 0, Y: 0}) {
			arrows = append(arrows, "")
		} else {
			// this assumes missing button is only in one corner
			if validKeys.Contains(cur.Add(space.Point{X: movement.X})) && movement.X != 0 {
				// horizontal first
				arrows = append(arrows, arrowLine(space.Point{X: movement.X})+arrowLine(space.Point{Y: movement.Y}))
			}
			if validKeys.Contains(cur.Add(space.Point{Y: movement.Y})) && movement.Y != 0 {
				// vertical first
				arrows = append(arrows, arrowLine(space.Point{Y: movement.Y})+arrowLine(space.Point{X: movement.X}))
			}
		}
		cur = target

		next := make([]string, 0, len(res)*len(arrows))
		for _, first := range res {
			for _, second := range arrows {
				next = append(next, first+second+"A")
			}
		}
		res = next
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	codes := parse(filename)

	// Part 1
	numpad := parseKeypad(`
		789
		456
		123
		-0A
	`)
	arrows := parseKeypad(`
		-^A
		<v>
	`)

	keypads := []map[rune]space.Point{
		numpad, arrows, arrows,
	}

	total := 0
	for _, code := range codes {
		sequences := []string{code}
		for _, keypad := range keypads {
			var next []string
			for _, sequence := range sequences {
				next = append(next, robotMovement(keypad, sequence)...)
			}

			slices.SortFunc(next, func(a, b string) int {
				return cmp.Compare(len(a), len(b))
			})
			if i := slices.IndexFunc(next, func(a string) bool { return len(a) > len(next[0]) }); i >= 0 {
				next = next[:i]
			}

			sequences = next
		}

		total += len(sequences[0]) * evil.Atoi(code[:len(code)-1])
	}
	log.Part1(total)
}
