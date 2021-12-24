package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type alu struct {
	reg [4]int
}

func (a *alu) String() string {
	return fmt.Sprintf("w: %d, x: %d, y: %d, z: %d", a.reg[0], a.reg[1], a.reg[2], a.reg[3])
}

func (a *alu) store(arg arg, v int) {
	if !arg.reg {
		util.Panic("Not a register")
	}
	a.reg[arg.v] = v
}

func (a *alu) load(arg arg) int {
	if arg.reg {
		return a.reg[arg.v]
	}
	return arg.v
}

func (a *alu) run(instructions []instruction, in *input) {
	for _, inst := range instructions {
		// log.Println(inst)
		switch inst.op {
		case opInp:
			a.store(inst.args[0], in.next())

		case opAdd:
			a.store(inst.args[0], a.load(inst.args[0])+a.load(inst.args[1]))

		case opMul:
			a.store(inst.args[0], a.load(inst.args[0])*a.load(inst.args[1]))

		case opDiv:
			a.store(inst.args[0], a.load(inst.args[0])/a.load(inst.args[1]))

		case opMod:
			a.store(inst.args[0], a.load(inst.args[0])%a.load(inst.args[1]))

		case opEql:
			if a.load(inst.args[0]) == a.load(inst.args[1]) {
				a.store(inst.args[0], 1)
			} else {
				a.store(inst.args[0], 0)
			}

		default:
			util.Panic("Unknown op code %d", inst.op)
		}
		// log.Println(a)
	}
}

func (a *alu) reset() {
	a.reg[0] = 0
	a.reg[1] = 0
	a.reg[2] = 0
	a.reg[3] = 0
}

type input struct {
	in  []int
	pos int
}

func newInput(in ...int) *input {
	return &input{
		in: in,
	}
}

func (i *input) next() int {
	r := i.in[i.pos]
	i.pos++
	return r
}

func (i *input) reset() {
	i.pos = 0
}

type opCode int

const (
	opInp opCode = iota
	opAdd
	opMul
	opDiv
	opMod
	opEql
)

var opStrMap = map[opCode]string{
	opInp: "inp",
	opAdd: "add",
	opMul: "mul",
	opDiv: "div",
	opMod: "mod",
	opEql: "eql",
}

func (op opCode) String() string {
	return opStrMap[op]
}

type arg struct {
	reg bool
	v   int
}

func (a arg) String() string {
	if a.reg {
		return string(byte(a.v) + 'w')
	}
	return strconv.Itoa(a.v)
}

type instruction struct {
	op   opCode
	args []arg
}

func (i instruction) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%s", i.op)
	for _, a := range i.args {
		fmt.Fprintf(b, " %s", a)
	}
	return b.String()
}

func parse(filename string) []instruction {
	instructions := []instruction{}

	for line := range load.File(filename) {
		inst := instruction{}

		p := strings.Split(line, " ")
		switch p[0] {
		case "inp":
			inst.op = opInp
			inst.args = []arg{{reg: true}}
		case "add":
			inst.op = opAdd
			inst.args = []arg{{reg: true}, {reg: false}}
		case "mul":
			inst.op = opMul
			inst.args = []arg{{reg: true}, {reg: false}}
		case "div":
			inst.op = opDiv
			inst.args = []arg{{reg: true}, {reg: false}}
		case "mod":
			inst.op = opMod
			inst.args = []arg{{reg: true}, {reg: false}}
		case "eql":
			inst.op = opEql
			inst.args = []arg{{reg: true}, {reg: false}}
		default:
			util.Panic("Unknown instruction %s", p[0])
		}

		if len(p)-1 != len(inst.args) {
			util.Panic("Invalid number of arguments for %s: %d, expected %d", p[0], len(p)-1, len(inst.args))
		}

		for i, a := range inst.args {
			if len(p[i+1]) == 1 && p[i+1][0] >= 'w' && p[i+1][0] <= 'z' {
				a.reg = true
				a.v = int(p[i+1][0] - 'w')
			} else {
				if a.reg {
					util.Panic("Instruction %s needs a register as arg %d", p[0], i)
				}
				a.v = util.Atoi(p[i+1])
			}
			inst.args[i] = a
		}

		instructions = append(instructions, inst)
	}

	return instructions
}

func split(instructions []instruction) [][]instruction {
	res := [][]instruction{}
	i := 0
	for j, inst := range instructions {
		if inst.op == opInp && j != i {
			res = append(res, instructions[i:j])
			i = j
		}
	}
	res = append(res, instructions[i:])
	return res
}

type best struct {
	maxSeq [14]int
	minSeq [14]int
}

func (b best) better(other best) best {
	for i := range b.maxSeq {
		if other.maxSeq[i] > b.maxSeq[i] {
			b.maxSeq = other.maxSeq
			break
		}
	}
	for i := range b.minSeq {
		if other.minSeq[i] == 0 {
			break
		}
		if other.minSeq[i] < b.minSeq[i] || b.minSeq[i] == 0 {
			b.minSeq = other.minSeq
			break
		}
	}
	return b
}

func (b best) max() int {
	n := 0
	for _, m := range b.maxSeq {
		n *= 10
		n += m
	}
	return n
}

func (b best) min() int {
	n := 0
	for _, m := range b.minSeq {
		n *= 10
		n += m
	}
	return n
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	inst := split(parse(filename))

	states := map[int]best{
		0: {},
	}
	for ii, insts := range inst {
		newStates := map[int]best{}

		for z, seq := range states {
			for i := 1; i <= 9; i++ {
				na := alu{reg: [4]int{0, 0, 0, z}}
				na.run(insts, newInput(i))

				seq.maxSeq[ii] = i
				seq.minSeq[ii] = i
				newStates[na.reg[3]] = newStates[na.reg[3]].better(seq)
			}
		}

		states = newStates
		log.Println(ii+1, len(states))
	}

	max := 0
	min := math.MaxInt
	for z, seq := range states {
		if z != 0 {
			continue
		}

		log.Println(z, seq)

		if seq.max() > max {
			max = seq.max()
		}
		if seq.min() < min {
			min = seq.min()
		}
	}

	log.Part1(max)
	log.Part2(min)
}
