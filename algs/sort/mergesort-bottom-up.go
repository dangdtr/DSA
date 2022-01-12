package sort

// Merge Sort Bottom-Up.
//  Call
//  unsortedSlice := IntSliceToInterface(slice []int)
//  or
//  unsortedSlice := StringSliceToInterface(slice []string)
func MergeSortBU(a []interface{}) {
	aux := make([]interface{}, len(a), len(a))
	N := len(a)
	for sz := 1; sz < N; sz = sz+sz {
		for lo := 0; lo < N-sz; lo += sz+sz {
			mergeBU(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1));
		}
	}
	IsSorted(a)
}

func min(a, b int) int {
	if a > b {
		return b
	}else {
		return a
	}
}

func mergeBU(a, aux []interface{}, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	i := lo
	j := mid+1;

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else {
			if Less(aux[j], aux[i]) {
				a[k] = aux[j]
				j++
			} else {
				a[k] = aux[i]
				i++
			} 
		} 
	}
}
