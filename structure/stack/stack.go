package stack

type Node struct {
	Value interface{}
	Next *Node
}

type Stack struct {
	top *Node
	size int
}

func (stack *Stack) push(value interface{}) {
	newNode := &Node{} // create new node
	
	newNode.Value = value
	newNode.Next = stack.top

	stack.top = newNode
	stack.size++
}

func (stack *Stack) pop() interface{} {
	popValue := stack.top.Value
	if stack.top.Next == nil {
		stack.top = nil
	}else {
		stack.top.Value, stack.top.Next = stack.top.Next.Value, stack.top.Next.Next 
	}
	stack.size--
	return popValue
}

func (stack *Stack) peek() interface{} {	
	return  stack.top.Value
}

func (stack *Stack) stackSize() interface{} {
	return stack.size
}

func (stack *Stack) isEmpty() interface{} {
	return stack.stackSize() == 0
}

func (stack *Stack) show() (in []interface{}) {
	current := stack.top

	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}