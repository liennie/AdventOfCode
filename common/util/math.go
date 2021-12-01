package util

func Mod(a, b int) int {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a*b) / GCD(a, b)
}
