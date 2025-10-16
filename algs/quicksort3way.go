package algs

import "math/rand"

// QuickSort3WayFunc (Recommend) Quick Sort 3 way to duplicate keys. Dont use cut-off!
func QuickSort3WayFunc[T any](a []T, less func(a, b T) bool) {
	rand.Shuffle(len(a), func(i, j int) {
		exch(a, i, j)
	})
	quickSort3WayFunc[T](a, 0, len(a)-1, less)
	IsSortedFunc[T](a, less)
}

func quickSort3WayFunc[T any](a []T, lo int, hi int, less func(a, b T) bool) {

	if hi <= lo {
		return
	}

	lt, gt := partition3WayFunc[T](a, lo, hi, less)

	// x[lo..lt-1] < a[lt..gt] < a[gt+1..hi]
	quickSort3WayFunc[T](a, lo, lt-1, less)
	quickSort3WayFunc[T](a, gt+1, hi, less)

	IsSortedLoHiFunc[T](a, lo, hi, less)
}

func partition3WayFunc[T any](a []T, lo int, hi int, less func(a, b T) bool) (int, int) {

	i := lo + 1
	lt, gt := lo, hi
	v := a[lo]

	// x[lt] === x[lo]
	for i <= gt {
		if less(a[i], v) {
			exch(a, i, lt)
			lt++
			i++
			continue
		}
		if less(v, a[i]) {
			exch(a, i, gt)
			gt--
			continue
		} else {
			i++
		}
	}

	return lt, gt
}
