package algs

import "testing"

func TestST_PutGetContains_Size_Delete_Keys_IsEmpty(t *testing.T) {
	st := NewST()
	if !st.IsEmpty() || st.Size() != 0 {
		t.Fatalf("expected empty map initially")
	}
	st.Put("a", 1)
	st.Put("b", 2)
	st.Put("a", 3) // overwrite

	if st.Size() != 2 {
		t.Fatalf("size=%d, want 2", st.Size())
	}
	if v := st.Get("a"); v != 3 {
		t.Fatalf("Get(a)=%d, want 3", v)
	}
	if !st.Contains("b") || st.Contains("c") {
		t.Fatalf("Contains mismatch")
	}
	keys := st.Keys()
	if len(keys) != 2 {
		t.Fatalf("Keys length=%d, want 2", len(keys))
	}
	st.Delete("b")
	if st.Contains("b") || st.Size() != 1 {
		t.Fatalf("after delete b, state invalid")
	}
}


