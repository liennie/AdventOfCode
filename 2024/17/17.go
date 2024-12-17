package main

import (
	"regexp"
	"slices"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Computer struct {
	a       int
	b       int
	c       int
	ip      int
	program []int
	verbose bool
}

func (c *Computer) debug(format string, v ...any) {
	if c.verbose {
		log.Printf(format, v...)
	}
}

func (c *Computer) combo(i int) int {
	switch i {
	case 0, 1, 2, 3:
		return i

	case 4:
		return c.a

	case 5:
		return c.b

	case 6:
		return c.c
	}

	evil.Panic("invalid value of a combo operand: %d", i)
	return 0
}

func (c *Computer) run() []int {
	start := *c
	defer func() {
		*c = start
	}()

	var output []int

	for c.ip < len(c.program) {
		operand := c.program[c.ip+1]

		c.debug("A: %d\t B: %d\t C: %d", c.a, c.b, c.c)

		switch c.program[c.ip] {
		case 0: // adv
			combo := c.combo(operand)
			c.debug("0 adv(C%d): c.a = %d / 2^%d", operand, c.a, combo)
			c.a /= ints.Pow(2, combo)
			c.ip += 2

		case 1: // bxl
			c.debug("1 bxl( %d): c.b = %d ~ %d", operand, c.b, operand)
			c.b ^= operand
			c.ip += 2

		case 2: // bst
			combo := c.combo(operand)
			c.debug("2 bst(C%d): c.b = %d %% 8 (%d)", operand, combo, combo%8)
			c.b = combo % 8
			c.ip += 2

		case 3: // jnz
			if c.a == 0 {
				c.debug("3 jnz( %d): --", operand)
				c.ip += 2
			} else {
				c.debug("3 jnz( %d): c.ip = %d\n\n", operand, operand)
				c.ip = operand
			}

		case 4: // bxc
			c.debug("4 bxc(X%d): c.b = %d ~ %d", operand, c.b, c.c)
			c.b ^= c.c
			c.ip += 2

		case 5: // out
			combo := c.combo(operand)
			c.debug("5 out(C%d): %d %% 8 (%d)", operand, combo, combo%8)
			output = append(output, combo%8)
			c.ip += 2

		case 6: // bdv
			combo := c.combo(operand)
			c.debug("6 bdv(C%d): c.b = %d / 2^%d", operand, c.a, combo)
			c.b = c.a / ints.Pow(2, combo)
			c.ip += 2

		case 7: // cdv
			combo := c.combo(operand)
			c.debug("7 cdv(C%d): c.c = %d / 2^%d", operand, c.a, combo)
			c.c = c.a / ints.Pow(2, combo)
			c.ip += 2

		default:
			evil.Panic("invalid instruction opcode: %d", c.program[c.ip])
		}
	}

	return output
}

func parse(filename string) *Computer {
	ch := load.File(filename)
	defer channel.Drain(ch)

	res := &Computer{}
	res.a = evil.Atoi(regexp.MustCompile(`^Register A: (\d+)$`).FindStringSubmatch(<-ch)[1])
	res.b = evil.Atoi(regexp.MustCompile(`^Register B: (\d+)$`).FindStringSubmatch(<-ch)[1])
	res.c = evil.Atoi(regexp.MustCompile(`^Register C: (\d+)$`).FindStringSubmatch(<-ch)[1])
	<-ch //empty line
	res.program = evil.Split(regexp.MustCompile(`^Program: ([0-7,]+)$`).FindStringSubmatch(<-ch)[1], ",")
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	computer := parse(filename)
	// computer.verbose = true

	// Part 1
	log.Part1(evil.Join(computer.run(), ","))

	// Part 2
	// this only works for the real input
	as := []int{0}
	for i := len(computer.program) - 1; i >= 0; i-- {
		target := computer.program[i]
		evil.Assert(len(as) > 0, "as is empty")
		for j := len(as) - 1; j >= 0; j-- {
			a := as[j] * 8
			as = slices.Delete(as, j, j+1)
			for n := range 8 {
				computer.a = a + n
				res := computer.run()
				if len(res) == len(computer.program)-i && res[0] == target {
					as = append(as, a+n)
				}
			}
		}
	}
	log.Part2(ints.Min(as...))
}
