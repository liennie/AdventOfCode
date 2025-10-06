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
	evil.Assert(match != nil, "line %q does not match regexp %q", line, re.String())

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

func tokens(machine Machine) int {
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
	nf := ints.Frac{
		N: p.X*a.Y - p.Y*a.X,
		D: b.X*a.Y - b.Y*a.X,
	}.Norm()
	if nf.D != 1 {
		return 0
	}
	if nf.N < 0 {
		return 0
	}
	n := nf.N

	// m * Ax + n * Bx == Px
	//
	// m * Ax == Px - n * Bx
	//
	// m == (Px - n * Bx) / Ax
	mf := ints.Frac{
		N: p.X - nf.N*b.X,
		D: a.X,
	}.Norm()
	if mf.D != 1 {
		return 0
	}
	if mf.N < 0 {
		return 0
	}
	m := mf.N

	// m * A + n * B == P
	evil.Assert(a.Scale(m).Add(b.Scale(n)) == p, "incorrect %d, %d, %v, %v", m, n, machine, a.Scale(m).Add(b.Scale(n)))

	return 3*m + n
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	machines := parse(filename)

	// Part 1
	total := 0
	for _, machine := range machines {
		t := tokens(machine)
		total += t
	}
	log.Part1(total)

	// Part 2
	total = 0
	for _, machine := range machines {
		machine.Prize = machine.Prize.Add(space.Point{10000000000000, 10000000000000})
		t := tokens(machine)
		total += t
	}
	log.Part2(total)
}
