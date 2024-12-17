package main

import (
	"regexp"

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
	var output []int

	for c.ip < len(c.program) {
		operand := c.program[c.ip+1]

		switch c.program[c.ip] {
		case 0: // adv
			c.a /= ints.Pow(2, c.combo(operand))
			c.ip += 2

		case 1: // bxl
			c.b ^= operand
			c.ip += 2

		case 2: // bst
			c.b = c.combo(operand) % 8
			c.ip += 2

		case 3: // jnz
			if c.a == 0 {
				c.ip += 2
			} else {
				c.ip = operand
			}

		case 4: // bxc
			c.b ^= c.c
			c.ip += 2

		case 5: // out
			output = append(output, c.combo(operand)%8)
			c.ip += 2

		case 6: // bdv
			c.b = c.a / ints.Pow(2, c.combo(operand))
			c.ip += 2

		case 7: // cdv
			c.c = c.a / ints.Pow(2, c.combo(operand))
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

	// Part 1
	log.Part1(evil.Join(computer.run(), ","))
}
