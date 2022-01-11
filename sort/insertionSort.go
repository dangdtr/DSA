package sort

func InsertionSort(a []interface{}) {
	Insertion(a, 0, len(a)); 
}

func Insertion(a []interface{}, lo, hi int) {
	for i := lo; i < hi; i++ {
		for j := i; j > lo; j-- {
			if Less(a[j], a[j-1]) {
				exch(a, j, j-1)
			} else {
				break
			}
		}
	}
}
