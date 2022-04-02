package stack

type ArrayStack struct {
	data []interface{}
	size int
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.data = append(stack.data, e)
	stack.size++
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	stack.size--
	return stack.data[stack.size]
}
