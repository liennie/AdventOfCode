package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type Point3 struct {
	X, Y, Z int
}

func (p Point3) Normalize() Point3 {
	gcd := ints.GCD(ints.GCD(ints.Abs(p.X), ints.Abs(p.Y)), ints.Abs(p.Z))
	return Point3{
		X: p.X / gcd,
		Y: p.Y / gcd,
		Z: p.Z / gcd,
	}
}

func (p Point3) Sub(other Point3) Point3 {
	return Point3{
		X: p.X - other.X,
		Y: p.Y - other.Y,
		Z: p.Z - other.Z,
	}
}

func (p Point3) Add(other Point3) Point3 {
	return Point3{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
	}
}

func (p Point3) Equals(other Point3) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p Point3) ManhattanLen() int {
	return ints.Abs(p.X) + ints.Abs(p.Y) + ints.Abs(p.Z)
}

func (p Point3) Rot90(r Point3) Point3 {
	for i := 0; i < r.X; i++ {
		p.Z, p.Y = p.Y, -p.Z
	}
	for i := 0; i < r.Y; i++ {
		p.X, p.Z = p.Z, -p.X
	}
	for i := 0; i < r.Z; i++ {
		p.Y, p.X = p.X, -p.Y
	}
	return p
}

func (p Point3) Flip() Point3 {
	return Point3{
		X: -p.X,
		Y: -p.Y,
		Z: -p.Z,
	}
}
