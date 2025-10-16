package algs

import (
	"testing"
)

func TestStack_PushPopPeek_IsEmpty_Size(t *testing.T) {
	st := NewStack[int]()
	if !st.IsEmpty() || st.StackSize() != 0 {
		// initial state
		t.Fatalf("expected empty stack with size 0")
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)

	if st.StackSize() != 3 {
		t.Fatalf("expected size 3, got %d", st.StackSize())
	}

	if top := st.Peek(); top != 3 {
		t.Fatalf("expected peek 3, got %d", top)
	}

	if v := st.Pop(); v != 3 {
		t.Fatalf("expected pop 3, got %d", v)
	}
	if v := st.Pop(); v != 2 {
		t.Fatalf("expected pop 2, got %d", v)
	}
	if v := st.Pop(); v != 1 {
		t.Fatalf("expected pop 1, got %d", v)
	}

	if !st.IsEmpty() || st.StackSize() != 0 {
		t.Fatalf("expected empty after pops")
	}
}

func TestStack_Iterator_Order(t *testing.T) {
	st := NewStack[int]()
	for i := 1; i <= 5; i++ {
		st.Push(i)
	}
	// Iterator should return items from bottom to top (1..5)
	items := st.Iterator()
	if len(items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(items))
	}
	for i, v := range items {
		exp := i + 1
		if v != exp {
			t.Fatalf("iterator[%d]=%d, want %d", i, v, exp)
		}
	}
}


