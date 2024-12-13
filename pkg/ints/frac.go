package ints

type Frac struct {
	N int
	D int
}

func (f Frac) Add(other Frac) Frac {
	f.normSign()
	other.normSign()

	lcm := LCM(f.D, other.D)
	return Frac{
		N: (f.N * (lcm / f.D)) + (other.N * (lcm / other.D)),
		D: lcm,
	}
}

func (f Frac) Sub(other Frac) Frac {
	return f.Add(other.Neg())
}

func (f Frac) Mul(other Frac) Frac {
	return Frac{
		N: f.N * other.N,
		D: f.D * other.D,
	}
}

func (f Frac) Div(other Frac) Frac {
	return f.Mul(other.Inv())
}

func (f Frac) Neg() Frac {
	return Frac{
		N: -f.N,
		D: f.D,
	}
}

func (f Frac) Inv() Frac {
	return Frac{
		N: f.D,
		D: f.N,
	}
}

func (f Frac) Norm() Frac {
	f.normSign()

	gcd := GCD(f.N, f.D)
	if gcd == 0 {
		return f
	}
	return Frac{
		N: f.N / gcd,
		D: f.D / gcd,
	}
}

func (f *Frac) normSign() {
	if f.D < 0 {
		f.N = -f.N
		f.D = -f.D
	}
}
