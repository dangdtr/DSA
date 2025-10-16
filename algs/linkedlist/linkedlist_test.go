package linkedlist

import "testing"

func TestSLinkedList_AddFirst_RemoveFirst_RemoveLast_Size_IsEmpty_Iterator(t *testing.T) {
	ll := NewSLinkedlist[int]()
	if !ll.IsEmpty() || ll.Size() != 0 {
		t.Fatalf("expected empty initially")
	}
	ll.AddFirst(1)
	ll.AddFirst(2)
	ll.AddFirst(3) // list: 3->2->1

	if ll.Size() != 3 || ll.IsEmpty() {
		t.Fatalf("size=%d, empty=%v", ll.Size(), ll.IsEmpty())
	}

	items := ll.Iterator()
	if len(items) != 3 || items[0] != 3 || items[2] != 1 {
		t.Fatalf("iterator=%v, want [3 2 1]", items)
	}

	if v, ok := ll.RemoveFirst(); !ok || v != 3 {
		t.Fatalf("RemoveFirst=%d,%v want 3,true", v, ok)
	}
	ll.RemoveLast() // remove tail (1); list now: 2
	if ll.Size() != 1 {
		t.Fatalf("size after removes=%d, want 1", ll.Size())
	}
	if v, ok := ll.RemoveFirst(); !ok || v != 2 {
		t.Fatalf("RemoveFirst=%d,%v want 2,true", v, ok)
	}
	if !ll.IsEmpty() || ll.Size() != 0 {
		t.Fatalf("should be empty at end")
	}
}


