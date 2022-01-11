package sort

//Merge Sort Bottom-Up
//Pass a []interface{} by
//  iargs:=make([]interface{},0)
//  for _,x:=range slice {
// 	   iargs=append(iargs, x)
//  }
func MergeSortBU(a []interface{}) {
	aux := InitEmptyInterfaceSlice(len(a))
	N := len(a)
	for sz := 1; sz < N; sz = sz+sz {
		for lo := 0; lo < N-sz; lo += sz+sz {
			mergeBU(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1));
		}
	}
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
