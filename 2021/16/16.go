package main

import (
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type bitReader struct {
	bits string
	pos  int
}

func (r *bitReader) next(n int) int {
	res, err := strconv.ParseInt(r.bits[r.pos:r.pos+n], 2, 0)
	if err != nil {
		util.Panic("Invalid bits %q: %w", r.bits[r.pos:r.pos+n], err)
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
	defer util.Drain(ch)

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
}

type literalValuePacket struct {
	ver   int
	value int
}

func (p literalValuePacket) versionSum() int {
	return p.ver
}

type operatorPacket struct {
	ver  int
	id   int
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
	typeLiteral = 4
)

func parsePacket(r *bitReader) packet {
	ver := r.next(3)
	_ = ver

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
					ver:   ver,
					value: lit,
				}
			}
		}

	default:
		op := operatorPacket{
			ver: ver,
			id:  id,
		}

		lenId := r.next(1)
		if lenId == 0 {
			len := r.next(15)
			to := r.pos + len
			for r.pos < to {
				op.subs = append(op.subs, parsePacket(r))
			}
			if r.pos > to {
				util.Panic("Len overflow %d > %d", r.pos, to)
			}

		} else {
			subs := r.next(11)
			for i := 0; i < subs; i++ {
				op.subs = append(op.subs, parsePacket(r))
			}
		}

		return op
	}
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	r := loadFile(filename)
	packet := parsePacket(r)

	// Part 1
	log.Part1(packet.versionSum())
}
