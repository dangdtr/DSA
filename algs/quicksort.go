package algs

import (
	"math/rand"
)

const InsertionSortCutoff = 8

// QuickSortFunc sorts using quicksort with a provided less comparator, keeping T as any.
func QuickSortFunc[T any](a []T, less func(a, b T) bool) {
	rand.Shuffle(len(a), func(i, j int) {
		exch(a, i, j)
	})
	quickSortFunc(a, 0, len(a)-1, less)
}

func quickSortFunc[T any](a []T, lo int, hi int, less func(a, b T) bool) {
	if lo >= hi {
		return
	}
	if hi <= lo+InsertionSortCutoff-1 {
		InsertionFunc(a, lo, hi, less)
		return
	}
    j := partitionFunc[T](a, lo, hi, less)
    quickSortFunc[T](a, lo, j-1, less)
    quickSortFunc[T](a, j+1, hi, less)
}

func partitionFunc[T any](a []T, lo int, hi int, less func(a, b T) bool) int {
	i := lo
	j := hi + 1
	for {
		i++
		for ; i <= j && less(a[i], a[lo]); i++ {
			if i == hi {
				break
			}
		}
		j--
		for ; j >= i && less(a[lo], a[j]); j-- {
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

func exch[T any](a []T, i int, j int) {
	a[i], a[j] = a[j], a[i]
}
