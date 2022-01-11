package queue

import "fmt"

type Node struct {
	Value interface{}
	Next *Node
}

type Queue struct {
	first *Node
	last *Node	
	size int
}


func (queue *Queue) Enqueue(value interface{}) {
	var newNode Node
	newNode.Value = value

	oldlast := queue.last
	queue.last = &newNode

	if queue.IsEmpty() {
		queue.first = queue.last
	}else {
		oldlast.Next = queue.last
	}

	queue.size++
}

func (queue *Queue) Dequeue() interface{} {
	temp := queue.first.Value
	queue.first = queue.first.Next

	if queue.IsEmpty() {
		queue.last = nil
	}

	return temp
}

func (queue *Queue) Peek(value interface{}) interface{} {
	return queue.first.Value
}

func (queue *Queue) IsEmpty() bool {
	return queue.first == nil
}
func (queue *Queue) QueueSize() int {
	return queue.size;
}


func (queue *Queue) Show() (in []interface{}) {
	current := queue.first

	for current != nil {
		in = append(in, current.Value)
		current = current.Next
	}
	return
}

func (queue *Queue) Test() {
	myQueue := &Queue{}
	myQueue.Enqueue(5)
	myQueue.Enqueue(6)
	myQueue.Enqueue(2)
	myQueue.Enqueue(4)
	myQueue.Enqueue(1)
	myQueue.Dequeue()
	fmt.Println(myQueue.Show()...)
}