package linkedlists

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
