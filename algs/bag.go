package algs

import "github.com/dawnpanpan/go-dsa/algs/linkedlist"

// Bag is a simple unordered collection backed by a singly linked list.
type Bag[T any] struct {
	*linkedlist.SLinkedlist[T]
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{linkedlist.NewSLinkedlist[T]()}
}

// AddFirst inserts an item at the front of the bag.
func (b *Bag[T]) AddFirst(item T) {
	b.SLinkedlist.AddFirst(item)
}

// Iterator returns a snapshot slice of items in insertion order (most-recent first).
func (b *Bag[T]) Iterator() []T {
	return b.SLinkedlist.Iterator()
}
