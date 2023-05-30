package generics

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var empty T
		return empty, false
	}

	ind := len(s.vals) - 1
	removed := s.vals[ind]
	s.vals = s.vals[:ind]
	return removed, true
}
