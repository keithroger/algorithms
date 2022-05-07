package linked

type Stack struct {
	first *node
	N     int
}

func (s *Stack) Push(item int) {
	oldFirst := s.first
	s.first = &node{item, oldFirst}
	s.N++
}

func (s *Stack) Pop() int {
	item := s.first.item
	s.first = s.first.next
	s.N--

	return item
}
