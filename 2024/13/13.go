package main

import (
	"regexp"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Machine struct {
	A     space.Point
	B     space.Point
	Prize space.Point
}

func parseLine(re regexp.Regexp, line string) space.Point {
	match := re.FindStringSubmatch(line)
	evil.Assert(match != nil, "line '", line, "' does not match regexp '", re.String(), "'")

	return space.Point{
		X: evil.Atoi(match[1]),
		Y: evil.Atoi(match[2]),
	}
}

func parse(filename string) []Machine {
	buttonA := regexp.MustCompile(`^Button A: X\+(\d+), Y\+(\d+)$`)
	buttonB := regexp.MustCompile(`^Button B: X\+(\d+), Y\+(\d+)$`)
	prize := regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

	var res []Machine
	for block := range load.Blocks(filename) {
		res = append(res, Machine{
			A:     parseLine(*buttonA, <-block),
			B:     parseLine(*buttonB, <-block),
			Prize: parseLine(*prize, <-block),
		})

		channel.Drain(block)
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	machines := parse(filename)

	// Part 1
	total := 0
	for _, machine := range machines {
		if na, nb := machine.A.Norm(), machine.B.Norm(); na == nb {
			evil.Panic("special case a == b, %+v", machine)
		}

		a, b, p := machine.A, machine.B, machine.Prize

		// m * A + n * B == P
		//
		// m * Ax + n * Bx == Px   // *  Ay
		// m * Ay + n * By == Py   // * -Ax
		//
		//   m * Ax * Ay + n * Bx * Ay ==   Px * Ay
		// - m * Ax * Ay - n * By * Ax == - Py * Ax
		//
		// n * Bx * Ay - n * By * Ax == Px * Ay - Py * Ax
		//
		// n * (Bx * Ay - By * Ax) == Px * Ay - Py * Ax
		//
		// n == (Px * Ay - Py * Ax) / (Bx * Ay - By * Ax)
		n := ints.Frac{
			N: p.X*a.Y - p.Y*a.X,
			D: b.X*a.Y - b.Y*a.X,
		}.Norm()
		if n.D != 1 {
			continue
		}

		// m * Ax + n * Bx == Px
		//
		// m * Ax == Px - n * Bx
		//
		// m == (Px - n * Bx) / Ax
		m := ints.Frac{
			N: p.X - n.N*b.X,
			D: a.X,
		}.Norm()
		if n.D != 1 {
			continue
		}

		total += 3*m.N + n.N
	}
	log.Part1(total)
}
