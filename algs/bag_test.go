package algs

import "testing"

func TestBag_AddFirst_Iterator_Size(t *testing.T) {
	b := NewBag[int]()

	b.AddFirst(1)
	b.AddFirst(2)
	b.AddFirst(3)

	items := b.Iterator()
	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}
	// Since AddFirst pushes to front, iteration snapshot is [3,2,1]
	expected := []int{3, 2, 1}
	for i, v := range items {
		if v != expected[i] {
			t.Fatalf("iterator[%d]=%d, want %d", i, v, expected[i])
		}
	}
}


