package sort

import "math/rand"

// (Recommend) Quick Sort 3 way to duplicate keys. Dont use cut-off!
//  Call
//  unsortedSlice := IntSliceToInterface(slice []int)
// or
//  unsortedSlice := StringSliceToInterface(slice []string)
func QuickSort3Way(a []interface{}) {
	rand.Shuffle(len(a), func(i, j int) {
		exch(a,i,j)
	})
	quickSort3Way(a, 0, len(a)-1)
	IsSorted(a)
}

func quickSort3Way(a []interface{}, lo int, hi int) {

	if hi <= lo {
		return
	}

	lt, gt := partition3Way(a, lo, hi)

	// x[lo..lt-1] < a[lt..gt] < a[gt+1..hi]
	quickSort3Way(a, lo, lt-1)
	quickSort3Way(a, gt+1, hi)

	IsSortedLoHi(a, lo, hi)
}

func partition3Way(a []interface{}, lo int, hi int) (int, int) {

	i := lo + 1
	lt, gt := lo, hi
	v := a[lo]

	// x[lt] === x[lo]
	for i <= gt {
		if Less(a[i], v) {
			exch(a, i, lt)
			lt++
			i++
			continue
		}
		if Less(v, a[i]) {
			exch(a, i, gt)
			gt--
			continue
		} else {
			i++
		}
	}

	return lt, gt
}
