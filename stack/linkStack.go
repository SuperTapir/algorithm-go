package stack

type node struct {
	value interface{}
	next  *node
}

type LinkStack struct {
	first *node
	size  int
}

func (s *LinkStack) Size() int {
	return s.size
}

func (s *LinkStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkStack) Push(value interface{}) {
	s.first = &node{value: value, next: s.first}
	s.size++
}

func (s *LinkStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.first.value
	s.first = s.first.next
	s.size--
	return v
}
