package algs

// IsSortedFunc checks if the slice is sorted according to less.
func IsSortedFunc[T any](a []T, less func(a, b T) bool) bool {
	return IsSortedLoHiFunc(a, 0, len(a)-1, less)
}

// IsSortedLoHiFunc checks if a[lo:hi] is sorted according to less.
func IsSortedLoHiFunc[T any](a []T, lo int, hi int, less func(a, b T) bool) bool {
	for i := lo + 1; i <= hi; i++ {
		if less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}
