package algs

// MergeSortFunc top-down merge sort with comparator.
func MergeSortFunc[T any](a []T, less func(a, b T) bool) {
	aux := make([]T, len(a))
    mergeSortFunc[T](a, aux, 0, len(a)-1, less)
}

func mergeSortFunc[T any](a, aux []T, lo int, hi int, less func(a, b T) bool) {
	if hi <= lo {
		return
	}

    mid := lo + (hi-lo)/2
    mergeSortFunc[T](a, aux, lo, mid, less)
    mergeSortFunc[T](a, aux, mid+1, hi, less)
    mergeFunc[T](a, aux, lo, mid, hi, less)
}

func mergeFunc[T any](a, aux []T, lo int, mid int, hi int, less func(a, b T) bool) {
	IsSortedLoHiFunc(a, lo, mid, less)
	IsSortedLoHiFunc(a, mid+1, hi, less)

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else {
			if less(aux[j], aux[i]) {
				a[k] = aux[j]
				j++
			} else {
				a[k] = aux[i]
				i++
			}
		}
	}
	IsSortedLoHiFunc(a, lo, hi, less)

}
