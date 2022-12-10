package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type cpu struct {
	x   int
	clk int

	preFunc  func(*cpu)
	postFunc func(*cpu)
}

func (cpu *cpu) pre() {
	if cpu.preFunc != nil {
		cpu.preFunc(cpu)
	}
}

func (cpu *cpu) post() {
	if cpu.postFunc != nil {
		cpu.postFunc(cpu)
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
	cpu.post()

	cpu.pre()
	cpu.clk++
	cpu.x += i.amt
	cpu.post()
}

type noop struct{}

func (noop) execute(cpu *cpu) {
	cpu.pre()
	cpu.clk++
	cpu.post()
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
	cpu := &cpu{
		x:   1,
		clk: 1,

		preFunc: func(cpu *cpu) {
			if (cpu.clk+20)%40 == 0 {
				signal += cpu.x * cpu.clk
			}
		},
	}
	for _, instruction := range instructions {
		instruction.execute(cpu)
	}
	log.Part1(signal)
}
