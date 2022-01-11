package main

import (
	"fmt"
	sort "go-algorithms/sort"
	structure "go-algorithms/structure/queue"
)

func main() {
	sort_test()

}

func sort_test() {
	// var slice = []string{"Gandia", "Arturo", "Alicia"}
	//generate string slice
	// var slide = sort.GenerateStringSlice(10)

	for i := 0; i < 10; i++ {
		// var slice = sort.GenerateInterfaceSlice(10)
		var slice = []int{-15, 42, 63, 64, 95, 67, 63, 6, 63, 6}
		// var slice = []int{64, 6, 6, 63, 67, 95, 63, 42, -15, 63}
		unsortedSlice := make([]interface{}, 0)
		for _, x := range slice {
			unsortedSlice = append(unsortedSlice, x)
		}

		// sort.MergeSort(unsortedSlice)
		// sort.MergeSortBU(unsortedSlice)
		// sort.QuickSort(unsortedSlice)
		sort.QuickSort3Way(unsortedSlice)
		// sort.InsertionSort(unsortedSlice)

		fmt.Println(unsortedSlice...)
		fmt.Println()

	}

}
func stack_queue_test() {
	tester := &structure.Queue{}
	tester.Test()
	fmt.Println(tester.Show())
}
