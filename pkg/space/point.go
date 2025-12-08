package space

import (
	"iter"

	"github.com/liennie/AdventOfCode/pkg/ints"
)

type Point struct {
	X, Y int
}

func (p Point) Norm() Point {
	gcd := ints.GCD(p.X, p.Y)
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

func (p Point) Div(other Point) Point {
	return Point{
		X: p.X / other.X,
		Y: p.Y / other.Y,
	}
}

func (p Point) Mod(other Point) Point {
	return Point{
		X: ints.Mod(p.X, other.X),
		Y: ints.Mod(p.Y, other.Y),
	}
}

func (p Point) Scale(sc int) Point {
	return Point{
		X: p.X * sc,
		Y: p.Y * sc,
	}
}

func (p Point) ManhattanLen() int {
	return ints.Abs(p.X) + ints.Abs(p.Y)
}

func (p Point) LenSquared() int {
	return p.X*p.X + p.Y*p.Y
}

// Rot90 rotates clockwise if positive x points right
// and positive y points down.
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

func Orthogonal() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, d := range [...]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			if !yield(d) {
				return
			}
		}
	}
}

func Neighbors() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, d := range [...]Point{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}} {
			if !yield(d) {
				return
			}
		}
	}
}
