package algs

import (
	"cmp"
	"testing"
)

func TestMaxPQ_Insert_DelMax_Max_Show(t *testing.T) {
	pq := NewMaxPQ[int](2, cmp.Less[int])
	vals := []int{5, 1, 3, 9, 2}
	for _, v := range vals {
		pq.Insert(v)
	}

	if top := pq.Max(); top != 9 {
		t.Fatalf("Max()=%d, want 9", top)
	}

	order := []int{9, 5, 3, 2, 1}
	for i, exp := range order {
		got := pq.DelMax()
		if got != exp {
			t.Fatalf("DelMax[%d]=%d, want %d", i, got, exp)
		}
	}
}

func TestMaxPQ_BuildHeapFromSlice(t *testing.T) {
	keys := []int{4, 7, 1, 6}
	pq := NewMaxPQFrom(keys, cmp.Less[int])
	// After heapify, DelMax should give descending
	order := []int{7, 6, 4, 1}
	for i, exp := range order {
		got := pq.DelMax()
		if got != exp {
			t.Fatalf("DelMax[%d]=%d, want %d", i, got, exp)
		}
	}
}


