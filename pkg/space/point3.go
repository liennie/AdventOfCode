package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type Point3 struct {
	X, Y, Z int
}

func (p Point3) Norm() Point3 {
	gcd := ints.GCD(ints.GCD(p.X, p.Y), p.Z)
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

func (p Point3) Mul(other Point3) Point3 {
	return Point3{
		X: p.X * other.X,
		Y: p.Y * other.Y,
		Z: p.Z * other.Z,
	}
}

func (p Point3) Div(other Point3) Point3 {
	return Point3{
		X: p.X / other.X,
		Y: p.Y / other.Y,
		Z: p.Z / other.Z,
	}
}

func (p Point3) Mod(other Point3) Point3 {
	return Point3{
		X: ints.Mod(p.X, other.X),
		Y: ints.Mod(p.Y, other.Y),
		Z: ints.Mod(p.Z, other.Z),
	}
}

func (p Point3) Scale(sc int) Point3 {
	return Point3{
		X: p.X * sc,
		Y: p.Y * sc,
		Z: p.Z * sc,
	}
}

func (p Point3) Abs() Point3 {
	return Point3{
		X: ints.Abs(p.X),
		Y: ints.Abs(p.Y),
		Z: ints.Abs(p.Z),
	}
}

func (p Point3) Area() int {
	return p.X * p.Y * p.Z
}

func (p Point3) Equals(other Point3) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p Point3) ManhattanLen() int {
	return ints.Abs(p.X) + ints.Abs(p.Y) + ints.Abs(p.Z)
}

func (p Point3) LenSquared() int {
	return p.X*p.X + p.Y*p.Y + p.Z*p.Z
}

func (p Point3) Rot90(r Point3) Point3 {
	r = r.Mod(Point3{4, 4, 4})
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
