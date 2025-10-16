package linkedlist

// Node is a singly linked list node that stores a value of type T.
type Node[T any] struct {
	Item T
	Next *Node[T]
}

// NewNode creates a new node with the given item and next pointer.
func NewNode[T any](item T, next *Node[T]) *Node[T] {
	return &Node[T]{Item: item, Next: next}
}

// SLinkedlist is a simple singly linked list that stores values of type T.
type SLinkedlist[T any] struct {
	Head *Node[T]
	n    int // size
}

func NewSLinkedlist[T any]() *SLinkedlist[T] {
	return &SLinkedlist[T]{nil, 0}
}

// AddFirst pushes an item to the front of the list and returns the new head.
func (ll *SLinkedlist[T]) AddFirst(item T) *Node[T] {
	ll.Head = NewNode(item, ll.Head)
	ll.n++
	return ll.Head
}

// RemoveFirst removes the head node and returns its item and a boolean indicating success.
func (ll *SLinkedlist[T]) RemoveFirst() (item T, ok bool) {
	if ll.Head == nil {
		var zero T
		return zero, false
	}
	second := ll.Head.Next
	data := ll.Head.Item
	ll.Head = second
	ll.n--
	return data, true
}

// RemoveLast removes the tail node if it exists.
func (ll *SLinkedlist[T]) RemoveLast() {
	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		ll.n = 0
		return
	}

	secondLast := ll.Head
	for secondLast.Next.Next != nil {
		secondLast = secondLast.Next
	}
	secondLast.Next = nil
	ll.n--
}

func (ll *SLinkedlist[T]) Size() int {
	return ll.n
}

func (ll *SLinkedlist[T]) IsEmpty() bool {
	return ll.n == 0
}

// Iterator returns a snapshot slice of all items in the list from head to tail.
func (ll *SLinkedlist[T]) Iterator() (items []T) {
	for current := ll.Head; current != nil; current = current.Next {
		items = append(items, current.Item)
	}
	return
}
