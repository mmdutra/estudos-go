package stack

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop
func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]

	return toRemove
}
