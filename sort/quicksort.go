package sort

import (
	"math/rand"
	"time"
)

const INSERTION_SORT_CUTOFF = 8

//Pass a []interface{} by
//  iargs:=make([]interface{},0)
//  for _,x:=range slice {
// 	   iargs=append(iargs, x)
//  }
func QuickSort(a []interface{}) {
	shuffle(a)
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []interface{}, lo int, hi int) {
	if hi <= lo + INSERTION_SORT_CUTOFF - 1{ 
		Insertion(a, lo, hi); 
		return; 
	} 
	j := partition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}

func partition(a []interface{}, lo int, hi int) int {
	i := lo+1
	j := hi
	for {
		for ;Less(a[i], a[lo]);i++ {
			if i == hi {
				break
			}
		}
		for ;Less(a[lo], a[j]);j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}
		exch(a, i, j)
	}
	exch(a, lo, j)
	return j
}

func exch(a []interface{}, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func shuffle(a []interface{}) {
	n := len(a)
	for i := 0; i < n; i++ {
		r := i + uniform(n-i) // between i and n-1
		a[i], a[r] = a[r], a[i]
	}
}
func uniform(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
