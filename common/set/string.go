package set

type String map[string]struct{}

func (s String) Add(items ...string) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s String) Remove(items ...string) {
	for _, item := range items {
		delete(s, item)
	}
}
