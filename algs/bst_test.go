package algs

import "testing"

func TestBST_PutGetContains_Size_MinMax_Delete(t *testing.T) {
	b := NewBST[int, string]()

	b.Put(5, "e")
	b.Put(3, "c")
	b.Put(7, "g")
	b.Put(2, "b")
	b.Put(4, "d")
	b.Put(6, "f")
	b.Put(8, "h")

	if b.Size() != 7 || b.IsEmpty() {
		t.Fatalf("expected size 7 and not empty")
	}

	if v, ok := b.Get(4); !ok || v != "d" {
		t.Fatalf("Get(4) = %q,%v; want d,true", v, ok)
	}
	if !b.Contains(6) || b.Contains(9) {
		t.Fatalf("Contains mismatch")
	}

	if min := b.Min(); min != 2 {
		t.Fatalf("Min()=%d, want 2", min)
	}
	if max := b.Max(); max != 8 {
		t.Fatalf("Max()=%d, want 8", max)
	}

	// Delete leaf
	b.Delete(2)
	if b.Contains(2) || b.Size() != 6 {
		t.Fatalf("after delete 2, contains or size incorrect")
	}
	// Delete node with one child
	b.Delete(7) // 7 had children (6,8) -> Hibbard deletion
	if b.Contains(7) || b.Size() != 5 {
		t.Fatalf("after delete 7, contains or size incorrect")
	}
	// DeleteMin and DeleteMax
	b.DeleteMin()
	b.DeleteMax()
	if b.Size() != 3 {
		t.Fatalf("after DeleteMin/Max, size=%d, want 3", b.Size())
	}

	if !b.IsBST() {
		t.Fatalf("tree should satisfy BST invariant")
	}
}


