package util

type Point struct {
	X, Y int
}

func (p Point) Normalize() Point {
	gcd := GCD(Abs(p.X), Abs(p.Y))
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

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}
