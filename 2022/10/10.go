package main

import (
	"fmt"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type cpu struct {
	x   int
	clk int

	preFunc func(*cpu)
}

func (cpu *cpu) pre() {
	if cpu.preFunc != nil {
		cpu.preFunc(cpu)
	}
}

type instruction interface {
	execute(*cpu)
}

type addx struct {
	amt int
}

func (i addx) execute(cpu *cpu) {
	cpu.pre()
	cpu.clk++

	cpu.pre()
	cpu.clk++
	cpu.x += i.amt
}

type noop struct{}

func (noop) execute(cpu *cpu) {
	cpu.pre()
	cpu.clk++
}

func parse(filename string) []instruction {
	res := []instruction{}
	for line := range load.File(filename) {
		i, args, _ := strings.Cut(line, " ")
		switch i {
		case "addx":
			res = append(res, addx{
				amt: evil.Atoi(args),
			})
		case "noop":
			res = append(res, noop{})
		default:
			evil.Panic("unknown instruction %s", i)
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	instructions := parse(filename)

	// Part 1
	signal := 0
	c := &cpu{
		x:   1,
		clk: 1,

		preFunc: func(cpu *cpu) {
			if (cpu.clk+20)%40 == 0 {
				signal += cpu.x * cpu.clk
			}
		},
	}
	for _, instruction := range instructions {
		instruction.execute(c)
	}
	log.Part1(signal)

	// Part 2
	log.Part2("Drawing:")
	crt := 0
	c = &cpu{
		x:   1,
		clk: 1,

		preFunc: func(cpu *cpu) {
			if ints.Abs(cpu.x-crt) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
			crt++
			if crt == 40 {
				fmt.Println()
				crt = 0
			}
		},
	}
	for _, instruction := range instructions {
		instruction.execute(c)
	}
}
