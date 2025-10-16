package algs

// MergeSortBUFunc bottom-up merge sort with comparator.
func MergeSortBUFunc[T any](a []T, less func(a, b T) bool) {
	aux := make([]T, len(a))
	N := len(a)
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
            mergeBUFunc[T](a, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1), less)
		}
	}
    IsSortedFunc[T](a, less)
}

// use builtin generic min

func mergeBUFunc[T any](a, aux []T, lo int, mid int, hi int, less func(a, b T) bool) {
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
}
