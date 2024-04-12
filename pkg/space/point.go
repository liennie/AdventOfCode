package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type Point struct {
	X, Y int
}

func (p Point) Normalize() Point {
	gcd := ints.GCD(ints.Abs(p.X), ints.Abs(p.Y))
	if gcd == 0 {
		return p
	}
	return Point{
		X: p.X / gcd,
		Y: p.Y / gcd,
	}
}

func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Mul(other Point) Point {
	return Point{
		X: p.X * other.X,
		Y: p.Y * other.Y,
	}
}

func (p Point) Scale(sc int) Point {
	return Point{
		X: p.X * sc,
		Y: p.Y * sc,
	}
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) ManhattanLen() int {
	return ints.Abs(p.X) + ints.Abs(p.Y)
}

func (p Point) Rot90(s int) Point {
	s = ints.Mod(s, 4)
	for i := 0; i < s; i++ {
		p.Y, p.X = p.X, -p.Y
	}
	return p
}

func (p Point) Flip() Point {
	return Point{
		X: -p.X,
		Y: -p.Y,
	}
}
