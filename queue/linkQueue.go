package queue

type node struct {
	value interface{}
	next  *node
}

type LinkQueue struct {
	first *node
	last  *node
	size  int
}

func (q *LinkQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkQueue) Size() int {
	return q.size
}

func (q *LinkQueue) Enqueue(val interface{}) {
	e := &node{value: val, next: nil}
	if q.IsEmpty() {
		q.first = e
		q.last = e
	} else {
		oldLast := q.last
		q.last = e
		oldLast.next = q.last
	}
	q.size++
}

func (q *LinkQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.first.value
	q.first = q.first.next
	q.size--
	return v
}
