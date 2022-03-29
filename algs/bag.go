package algs

import "github.com/dawnpanpan/go-dsa/algs/linkedlist"

type Bag struct {
	*linkedlist.SLinkedlist
}

func NewBag() *Bag {
	return &Bag{linkedlist.NewSLinkedlist()}
}
