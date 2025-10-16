package algs

import (
	"cmp"
	"testing"
)

func TestMinQP_Insert_DelMin_Min_Show(t *testing.T) {
	pq := NewMinQP[int](2, cmp.Less[int])
	vals := []int{5, 1, 3, 9, 2}
	for _, v := range vals {
		pq.Insert(v)
	}

	if top := pq.Min(); top != 1 {
		t.Fatalf("Min()=%d, want 1", top)
	}

	order := []int{1, 2, 3, 5, 9}
	for i, exp := range order {
		got := pq.DelMin()
		if got != exp {
			t.Fatalf("DelMin[%d]=%d, want %d", i, got, exp)
		}
	}
}

func TestMinQP_BuildHeapFromSlice(t *testing.T) {
	keys := []int{4, 7, 1, 6}
	pq := NewMinQPFrom(keys, cmp.Less[int])
	order := []int{1, 4, 6, 7}
	for i, exp := range order {
		got := pq.DelMin()
		if got != exp {
			t.Fatalf("DelMin[%d]=%d, want %d", i, got, exp)
		}
	}
}


