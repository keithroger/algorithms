type stack struct {
	first *node
	N int
}

func (s *stack) push(item int) {
	oldFirst := s.first
	s.first.item = item
	s.first.next = oldFirst
	s.N++
}

func (s *stack) pop() item int{
	item := s.first.item
	s.first = s.first.next
	s.N--
}