package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]

	return toRemove
}
