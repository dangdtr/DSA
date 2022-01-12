package main

import (
	"fmt"
	sort "go-algorithms/sort"
	queue "go-algorithms/structure/queue"
	pq "go-algorithms/structure/priorityqueue"
	
)

func main() {
	// sort_test()
	// stack_queue_test()
	// b := "-13" < "-18"
	// fmt.Println(b)
	pq_test()
}

func pq_test() {
	var key = sort.GenerateInterfaceSlice(10)
	// key := sort.StringSliceToInterface(k)
	fmt.Println(key...)
	
	maxPq := pq.NewMaxPQFrom(key)
	fmt.Println(maxPq.Show())

}

func sort_test() {
	for i := 0; i < 10; i++ {
		var unsortedSlice = sort.GenerateInterfaceSlice(10)

		// sort.MergeSort(unsortedSlice)
		sort.MergeSortBU(unsortedSlice)
		// sort.QuickSort(unsortedSlice)
		sort.QuickSort3Way(unsortedSlice)

		if sort.IsSorted(unsortedSlice) {
			fmt.Println(unsortedSlice...)
		}else {
			fmt.Println("false")
		}
	}
}

func stack_queue_test() {
	tester := queue.NewQueue()
	tester.Test()
}
