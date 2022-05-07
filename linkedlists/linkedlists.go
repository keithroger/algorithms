package linkedlists

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

type Queue struct {
	first, last *node
	N           int
}

func (q *Queue) IsEmpty() bool {
	return q.N == 0
}

func (q *Queue) Enqueue(item int) {
	oldLast := q.last
	q.last = &node{item, nil}
	if q.N == 0 {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.N++
}

func (q *Queue) Dequeue() int {
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.N--
	return item
}
