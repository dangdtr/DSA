package main

import (
	"fmt"
	"github.com/dawnpanpan/go-dsa/algs"
	"github.com/dawnpanpan/go-dsa/algs/structure"
	"github.com/dawnpanpan/go-dsa/algs/sort"
)

func main() {
	sort_test()
	// stack_queue_test()
	// b := "-13" < "-18"
	// fmt.Println(b)
	// pq_test()
}

func pq_test() {
	var key = algs.GenerateInterfaceSlice(10)
	// key := sort.StringSliceToInterface(k)
	fmt.Println(key...)
	
	maxPq := structure.NewMaxPQFrom(key)
	fmt.Println(maxPq.Show())

}

func sort_test() {
	for i := 0; i < 10; i++ {
		var slide = []int{-15 ,42 ,63 ,42 ,63 ,67 ,63, 6, 63, 6}
		unsortedSlice := algs.IntSliceToInterface(slide)
		// var unsortedSlice = algs.GenerateInterfaceSlice(10)

		// sort.MergeSort(unsortedSlice)
		// sort.MergeSortBU(unsortedSlice)
		sort.QuickSort(unsortedSlice)
		// sort.QuickSort3Way(unsortedSlice)

		if sort.IsSorted(unsortedSlice) {
			fmt.Println(unsortedSlice...)
		}else {
			fmt.Println("false")
		}
	}
}

func stack_queue_test() {
	tester := structure.NewQueue()
	tester.Test()
}
