package sort

import (
	"math/rand"
)

const INSERTION_SORT_CUTOFF = 8

//Pass a []interface{} by
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
	if hi <= lo + INSERTION_SORT_CUTOFF - 1{ 
		Insertion(a, lo, hi); 
		return; 
	} 
	j := partition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}

func partition(a []interface{}, lo int, hi int) int {
	i := lo+1
	j := hi
	for {
		for ;Less(a[i], a[lo]);i++ {
			if i == hi {
				break
			}
		}
		for ;Less(a[lo], a[j]);j-- {
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

