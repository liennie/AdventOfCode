package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type bitReader struct {
	bits string
	pos  int
}

func (r *bitReader) next(n int) int {
	res, err := strconv.ParseInt(r.bits[r.pos:r.pos+n], 2, 0)
	if err != nil {
		evil.Panic("Invalid bits %q: %w", r.bits[r.pos:r.pos+n], err)
	}
	r.pos += n
	return int(res)
}

var hex2binLut = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func loadFile(filename string) *bitReader {
	ch := load.File(filename)
	defer channel.Drain(ch)

	msg := <-ch

	r := &bitReader{}
	b := &strings.Builder{}

	for _, r := range msg {
		b.WriteString(hex2binLut[r])
	}

	r.bits = b.String()

	return r
}

type packet interface {
	versionSum() int
	value() int
}

type sumPacket struct {
	operatorPacket
}

func (p sumPacket) value() int {
	if len(p.subs) == 0 {
		evil.Panic("No subpackets")
	}

	sum := 0
	for _, sp := range p.subs {
		sum += sp.value()
	}
	return sum
}

func (p sumPacket) String() string {
	b := &strings.Builder{}
	b.WriteRune('(')
	for i, sp := range p.subs {
		if i > 0 {
			b.WriteString(" + ")
		}
		fmt.Fprint(b, sp)
	}
	b.WriteRune(')')
	return b.String()
}

type productPacket struct {
	operatorPacket
}

func (p productPacket) value() int {
	if len(p.subs) == 0 {
		evil.Panic("No subpackets")
	}

	sum := 1
	for _, sp := range p.subs {
		sum *= sp.value()
	}
	return sum
}

func (p productPacket) String() string {
	b := &strings.Builder{}
	b.WriteRune('(')
	for i, sp := range p.subs {
		if i > 0 {
			b.WriteString(" * ")
		}
		fmt.Fprint(b, sp)
	}
	b.WriteRune(')')
	return b.String()
}

type minPacket struct {
	operatorPacket
}

func (p minPacket) value() int {
	if len(p.subs) == 0 {
		evil.Panic("No subpackets")
	}

	min := math.MaxInt
	for _, sp := range p.subs {
		if v := sp.value(); v < min {
			min = v
		}
	}
	return min
}

func (p minPacket) String() string {
	b := &strings.Builder{}
	b.WriteString("Min(")
	for i, sp := range p.subs {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(b, sp)
	}
	b.WriteRune(')')
	return b.String()
}

type maxPacket struct {
	operatorPacket
}

func (p maxPacket) value() int {
	if len(p.subs) == 0 {
		evil.Panic("No subpackets")
	}

	max := math.MinInt
	for _, sp := range p.subs {
		if v := sp.value(); v > max {
			max = v
		}
	}
	return max
}

func (p maxPacket) String() string {
	b := &strings.Builder{}
	b.WriteString("Max(")
	for i, sp := range p.subs {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(b, sp)
	}
	b.WriteRune(')')
	return b.String()
}

type literalValuePacket struct {
	ver int
	val int
}

func (p literalValuePacket) versionSum() int {
	return p.ver
}

func (p literalValuePacket) value() int {
	return p.val
}

func (p literalValuePacket) String() string {
	return strconv.Itoa(p.val)
}

type gtPacket struct {
	operatorPacket
}

func (p gtPacket) value() int {
	if len(p.subs) != 2 {
		evil.Panic("Invalid number of subpackets")
	}

	if p.subs[0].value() > p.subs[1].value() {
		return 1
	}
	return 0
}

func (p gtPacket) String() string {
	return fmt.Sprintf("(%s > %s)", p.subs[0], p.subs[1])
}

type ltPacket struct {
	operatorPacket
}

func (p ltPacket) value() int {
	if len(p.subs) != 2 {
		evil.Panic("Invalid number of subpackets")
	}

	if p.subs[0].value() < p.subs[1].value() {
		return 1
	}
	return 0
}

func (p ltPacket) String() string {
	return fmt.Sprintf("(%s < %s)", p.subs[0], p.subs[1])
}

type eqPacket struct {
	operatorPacket
}

func (p eqPacket) value() int {
	if len(p.subs) != 2 {
		evil.Panic("Invalid number of subpackets")
	}

	if p.subs[0].value() == p.subs[1].value() {
		return 1
	}
	return 0
}

func (p eqPacket) String() string {
	return fmt.Sprintf("(%s = %s)", p.subs[0], p.subs[1])
}

type genericOperatorPacket struct {
	operatorPacket
	id int
}

func (p genericOperatorPacket) impl() packet {
	switch p.id {
	case typeSum:
		return sumPacket{operatorPacket: p.operatorPacket}
	case typeProduct:
		return productPacket{operatorPacket: p.operatorPacket}
	case typeMin:
		return minPacket{operatorPacket: p.operatorPacket}
	case typeMax:
		return maxPacket{operatorPacket: p.operatorPacket}
	case typeGT:
		return gtPacket{operatorPacket: p.operatorPacket}
	case typeLT:
		return ltPacket{operatorPacket: p.operatorPacket}
	case typeEQ:
		return eqPacket{operatorPacket: p.operatorPacket}
	}

	evil.Panic("Invalid operator id %d", p.id)
	return nil
}

type operatorPacket struct {
	ver  int
	subs []packet
}

func (p operatorPacket) versionSum() int {
	sum := p.ver
	for _, sp := range p.subs {
		sum += sp.versionSum()
	}
	return sum
}

const (
	typeSum     = 0
	typeProduct = 1
	typeMin     = 2
	typeMax     = 3
	typeLiteral = 4
	typeGT      = 5
	typeLT      = 6
	typeEQ      = 7
)

func parsePacket(r *bitReader) packet {
	ver := r.next(3)
	id := r.next(3)

	switch id {
	case typeLiteral:
		lit := 0
		for {
			n := r.next(5)

			lit <<= 4
			lit |= n & 0xf

			if n&0x10 == 0 {
				return literalValuePacket{
					ver: ver,
					val: lit,
				}
			}
		}

	default:
		op := genericOperatorPacket{
			operatorPacket: operatorPacket{
				ver: ver,
			},
			id: id,
		}

		lenId := r.next(1)
		if lenId == 0 {
			len := r.next(15)
			to := r.pos + len
			for r.pos < to {
				op.subs = append(op.subs, parsePacket(r))
			}
			if r.pos > to {
				evil.Panic("Len overflow %d > %d", r.pos, to)
			}

		} else {
			subs := r.next(11)
			for i := 0; i < subs; i++ {
				op.subs = append(op.subs, parsePacket(r))
			}
		}

		return op.impl()
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	r := loadFile(filename)
	packet := parsePacket(r)

	log.Printf("%s", packet)

	// Part 1
	log.Part1(packet.versionSum())

	// Part 2
	log.Part2(packet.value())
}
