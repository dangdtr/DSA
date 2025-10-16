package algs

// Queue is a generic FIFO queue implemented with a singly linked list.
type Queue[T any] struct {
	first *nodeQ[T]
	last  *nodeQ[T]
	size  int
}

func (queue *Queue[T]) Enqueue(value T) {
	newNode := &nodeQ[T]{Value: value}

	oldlast := queue.last
	queue.last = newNode

	if queue.IsEmpty() {
		queue.first = queue.last
	} else {
		oldlast.Next = queue.last
	}

	queue.size++
}

func (queue *Queue[T]) Dequeue() T {
	temp := queue.first.Value
	queue.first = queue.first.Next

	if queue.IsEmpty() {
		queue.last = nil
	}

	queue.size--
	return temp
}

func (queue *Queue[T]) Peek() T {
	return queue.first.Value
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.first == nil
}
func (queue *Queue[T]) QueueSize() int {
	return queue.size
}

func (queue *Queue[T]) Show() (in []T) {
	current := queue.first

	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// nodeQ is an internal queue node used to avoid leaking implementation details.
type nodeQ[T any] struct {
	Value T
	Next  *nodeQ[T]
}
