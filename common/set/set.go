package set

type Set[T comparable] map[T]struct{}

type String = Set[string]

func New[T comparable](items ...T) Set[T] {
	set := Set[T]{}
	set.Add(items...)
	return set
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	if len(b) < len(a) {
		a, b = b, a
	}

	res := Set[T]{}
	for item := range a {
		if b.Contains(item) {
			res.Add(item)
		}
	}
	return res
}
