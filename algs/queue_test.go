package algs

import "testing"

func TestQueue_EnqueueDequeuePeek_IsEmpty_Size(t *testing.T) {
	q := NewQueue[int]()
	if !q.IsEmpty() || q.QueueSize() != 0 {
		t.Fatalf("expected empty queue with size 0")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.QueueSize() != 3 {
		t.Fatalf("expected size 3, got %d", q.QueueSize())
	}

	if front := q.Peek(); front != 1 {
		t.Fatalf("expected peek 1, got %d", front)
	}

	if v := q.Dequeue(); v != 1 {
		t.Fatalf("expected dequeue 1, got %d", v)
	}
	if v := q.Dequeue(); v != 2 {
		t.Fatalf("expected dequeue 2, got %d", v)
	}
	if v := q.Dequeue(); v != 3 {
		t.Fatalf("expected dequeue 3, got %d", v)
	}

	if !q.IsEmpty() || q.QueueSize() != 0 {
		t.Fatalf("expected empty after dequeues")
	}
}

func TestQueue_Show_Order(t *testing.T) {
	q := NewQueue[int]()
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}
	items := q.Show()
	if len(items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(items))
	}
	for i, v := range items {
		exp := i + 1
		if v != exp {
			t.Fatalf("show[%d]=%d, want %d", i, v, exp)
		}
	}
}


