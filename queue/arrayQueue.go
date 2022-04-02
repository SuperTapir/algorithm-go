package queue

type ArrayQueue struct {
	data  []interface{}
	first int
	last  int
	size  int
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) Enqueue(e interface{}) {
	q.data = append(q.data, e)
	q.size++
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	e := q.data[q.first]
	q.data[q.first] = nil
	q.first++
	q.size--
	return e
}
