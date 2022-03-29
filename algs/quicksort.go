package algs

import (
	"math/rand"
)

const INSERTION_SORT_CUTOFF = 8

// QuickSort Pass a []interface{} by
//  Call
//  unsortedSlice := IntSliceToInterface(slice []int)
// or
//  unsortedSlice := StringSliceToInterface(slice []string)
func QuickSort(a []interface{}) {
	rand.Shuffle(len(a), func(i, j int) {
		exch(a, i, j)
	})
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []interface{}, lo int, hi int) {
	if lo >= hi {
		return
	}
	// if hi <= lo + INSERTION_SORT_CUTOFF - 1{
	// 	Insertion(a, lo, hi);
	// 	return;
	// }
	j := partition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}

func partition(a []interface{}, lo int, hi int) int {
	i := lo
	j := hi + 1
	for {
		i++
		for ; i <= j && Less(a[i], a[lo]); i++ {
			if i == hi {
				break
			}
		}
		j--
		for ; j >= i && Less(a[lo], a[j]); j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}
		exch(a, i, j)

	}
	exch(a, lo, j)
	return j
}

func exch(a []interface{}, i int, j int) {
	a[i], a[j] = a[j], a[i]
}
