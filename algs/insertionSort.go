package algs

// InsertionSortFunc generic insertion sort with comparator.
func InsertionSortFunc[T any](a []T, less func(a, b T) bool) {
	InsertionFunc(a, 0, len(a), less)
}

func InsertionFunc[T any](a []T, lo, hi int, less func(a, b T) bool) {
	for i := lo; i < hi; i++ {
		for j := i; j > lo; j-- {
			if less(a[j], a[j-1]) {
				exch(a, j, j-1)
			} else {
				break
			}
		}
	}
}
