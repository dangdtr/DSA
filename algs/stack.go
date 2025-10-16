package algs

// Stack is a generic LIFO stack implemented with a singly linked list node.
type Stack[T any] struct {
	top  *nodeS[T]
	size int
}

func (stack *Stack[T]) Push(value T) {
	newNode := &nodeS[T]{Value: value, Next: stack.top}
	stack.top = newNode
	stack.size++
}

func (stack *Stack[T]) Pop() T {
	value := stack.top.Value
	stack.top = stack.top.Next
	stack.size--
	return value
}

func (stack *Stack[T]) Peek() T {
	return stack.top.Value
}

func (stack *Stack[T]) StackSize() int {
	return stack.size
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack[T]) Show() (in []T) {
	current := stack.top

	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Iterator() (items []T) {
	var temp []T
	for current := stack.top; current != nil; current = current.Next {
		temp = append(temp, current.Value)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		items = append(items, temp[i])
	}
	return
}

// nodeS is an internal stack node type.
type nodeS[T any] struct {
	Value T
	Next  *nodeS[T]
}
