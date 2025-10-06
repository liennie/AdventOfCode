package main

import (
	"maps"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) []string {
	return load.Slice(filename)
}

type keypad = map[rune]space.Point

func parseKeypad(keys string) keypad {
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

var numpad = parseKeypad(`
	789
	456
	123
	-0A
`)

var arrowpad = parseKeypad(`
	-^A
	<v>
`)

var actionpad = parseKeypad(`A`)

func arrowLine(dir space.Point) string {
	if dir == (space.Point{X: 0, Y: 0}) {
		return ""
	}

	evil.Assert((dir.X == 0) != (dir.Y == 0), "no diagonal vectors are allowed %v", dir)

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

func keypadForKeys(from, to rune) keypad {
	switch from {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return numpad
	case '<', '>', '^', 'v':
		return arrowpad
	case 'A':
		switch to {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return numpad
		case '<', '>', '^', 'v':
			return arrowpad
		case 'A':
			return actionpad
		default:
			evil.Panic("unknown key %c", to)
		}
	default:
		evil.Panic("unknown key %c", from)
	}
	return nil
}

type robotMovementCacheKey struct {
	from, to rune
}

var robotMovementCache = map[robotMovementCacheKey][]string{}

func robotMovement(from, to rune) (res []string) {
	key := robotMovementCacheKey{from, to}
	if res, ok := robotMovementCache[key]; ok {
		return res
	}
	defer func() {
		robotMovementCache[key] = res
	}()

	keypad := keypadForKeys(from, to)

	validKeys := set.Collect(maps.Values(keypad))

	cur, ok := keypad[from]
	evil.Assert(ok, "keypad is missing %c", from)

	target, ok := keypad[to]
	evil.Assert(ok, "keypad is missing %c", to)

	movement := target.Sub(cur)
	if movement == (space.Point{X: 0, Y: 0}) {
		res = append(res, "A")
	} else {
		// this assumes missing button is only in one corner
		if validKeys.Contains(cur.Add(space.Point{X: movement.X})) && movement.X != 0 {
			// horizontal first
			res = append(res, arrowLine(space.Point{X: movement.X})+arrowLine(space.Point{Y: movement.Y})+"A")
		}
		if validKeys.Contains(cur.Add(space.Point{Y: movement.Y})) && movement.Y != 0 {
			// vertical first
			res = append(res, arrowLine(space.Point{Y: movement.Y})+arrowLine(space.Point{X: movement.X})+"A")
		}
	}
	return res
}

type sequenceLenCacheKey struct {
	code      string
	arrowpads int
}

var sequenceLenCache = map[sequenceLenCacheKey]int{}

func sequenceLen(code string, arrowpads int) (res int) {
	key := sequenceLenCacheKey{code, arrowpads}
	if res, ok := sequenceLenCache[key]; ok {
		return res
	}
	defer func() {
		sequenceLenCache[key] = res
	}()

	if arrowpads < 0 {
		return len(code)
	}

	total := 0
	from := 'A'
	for _, to := range code {
		total += ints.MinFunc(func(s string) int {
			return sequenceLen(s, arrowpads-1)
		}, robotMovement(from, to)...)

		from = to
	}
	return total
}

func complexity(code string, arrowpads int) int {
	return sequenceLen(code, arrowpads) * evil.Atoi(code[:len(code)-1])
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	codes := parse(filename)

	// Part 1
	total := 0
	for _, code := range codes {
		total += complexity(code, 2)
	}
	log.Part1(total)

	// Part 2
	total = 0
	for _, code := range codes {
		total += complexity(code, 25)
	}
	log.Part2(total)
}
