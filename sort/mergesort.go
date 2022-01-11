package sort

//Pass a []interface{} by 
//  iargs:=make([]interface{},0)
//  for _,x:=range slice {
// 	   iargs=append(iargs, x)
//  }
func MergeSort(a []interface{}) {
	aux := InitEmptyInterfaceSlice(len(a))
	mergeSort(a, aux, 0, len(a)-1)
}

func mergeSort(a, aux []interface{}, lo int, hi int) {
	if hi <= lo{
		return
	} 

	mid := lo + (hi - lo) / 2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid + 1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a, aux []interface{}, lo int, mid int, hi int) {
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



