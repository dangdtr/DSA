package linkedlist

import (
	"fmt"
)

// generic interface for all nodes
type Node struct {
	Item interface{}
	Next *Node
}

// Create a new Node
func NewNode(item interface{}, next *Node) *Node {
	return &Node{item, next}
}

// single linklist
type SLinkedlist struct {
	Head *Node
	n    int // size
}

func NewSLinkedlist() *SLinkedlist {
	return &SLinkedlist{nil, 0}
}

func (ll *SLinkedlist) AddFirst(item interface{}) *Node {
	ll.Head = NewNode(item, ll.Head)
	ll.n++
	return ll.Head
}

func (ll *SLinkedlist) RemoveFirst() (item interface{}) {
	if ll.Head == nil {
		return nil
	}
	second := ll.Head.Next
	data := ll.Head.Item
	ll.Head = nil
	ll.Head = second
	return data
}

func (ll *SLinkedlist) RemoveLast() {
	if ll.Head == nil {
		return
	}

	SecondLast := ll.Head
	for ; SecondLast.Next.Next != nil; SecondLast = SecondLast.Next {
	}

	SecondLast.Next = nil

}

func (ll *SLinkedlist) Size() int {
	return ll.Size()
}

func (ll *SLinkedlist) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *SLinkedlist) InteratorSlide() (item []interface{}) {
	for current := ll.Head; current != nil; current = current.Next {
		item = append(item, current.Item)
	}
	return
}

func (ll *SLinkedlist) InteratorInt() (item []int) {
	for current := ll.Head; current != nil; current = current.Next {
		item = append(item, current.Item.(int))
	}
	return
}

func (ll *SLinkedlist) InteratorString() (item []string) {
	for current := ll.Head; current != nil; current = current.Next {
		item = append(item, current.Item.(string))
	}
	return
}

func Demo() {
	slist := NewSLinkedlist()
	slist.AddFirst(3)
	slist.AddFirst(4)
	slist.AddFirst(5)
	slist.AddFirst(6)
	slist.AddFirst(7)

	fmt.Println(slist.InteratorSlide())

	slist.RemoveFirst()
	slist.RemoveFirst()
	fmt.Println(slist.InteratorSlide())

}
