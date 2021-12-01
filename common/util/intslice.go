package util

func Comb(n int) [][]int {
	if n < 0 {
		Panic("Comb(%d)", n)
	}

	if n == 0 {
		return [][]int{{}}
	}

	res := Comb(n - 1)
	l := len(res)
	for i := 0; i < l; i++ {
		p := make([]int, len(res[i]))
		copy(p, res[i])
		p = append(p, n-1)
		res = append(res, p)
	}
	return res
}

func Uniq(ns []int) []int {
	res := []int{}
	s := map[int]bool{}

	for _, n := range ns {
		if !s[n] {
			s[n] = true
			res = append(res, n)
		}
	}

	return res
}

func Contains(ns []int, n int) bool {
	for _, m := range ns {
		if m == n {
			return true
		}
	}
	return false
}
