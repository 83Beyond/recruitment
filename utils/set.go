package utils

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Len() int {
	return len(s)
}
