package main

import (
	"fmt"
	"github.com/dawnpanpan/go-dsa/algs"
	"github.com/dawnpanpan/go-dsa/algs/linkedlist"
)

func main() {
	// sort_test()
	// stack_queue_test()
	// b := "-13" < "-18"
	// fmt.Println(b)
	// pq_test()
	// bst_test()
	// linkedlist_test()
	//graph_test()
	dfs_test()
}

func bst_test() {
	bst := algs.NewBST()
	bst.Put(algs.StringKey("A"), "ada")
	bst.Put(algs.StringKey("C"), "add")
	bst.Put(algs.StringKey("D"), "afd")
	bst.Put(algs.StringKey("B"), "avd")
	bst.Put(algs.StringKey("E"), "aed")
	bst.DeleteMax()
	bst.Inorder()
	fmt.Println(bst.IsBST())
	fmt.Println(bst.Contains(algs.StringKey("C")))
}

func pq_test() {
	var key = algs.GenerateInterfaceSlice(10)
	// key := sort.StringSliceToInterface(k)
	fmt.Println(key...)

	maxPq := algs.NewMaxPQFrom(key)
	fmt.Println(maxPq.Show())

}

func sort_test() {
	for i := 0; i < 10; i++ {
		var slide = []int{-15, 42, 63, 42, 63, 67, 63, 6, 63, 6}
		unsortedSlice := algs.IntSliceToInterface(slide)
		// var unsortedSlice = algs.GenerateInterfaceSlice(10)

		// sort.MergeSort(unsortedSlice)
		// sort.MergeSortBU(unsortedSlice)
		algs.QuickSort(unsortedSlice)
		// sort.QuickSort3Way(unsortedSlice)

		if algs.IsSorted(unsortedSlice) {
			fmt.Println(unsortedSlice...)
		} else {
			fmt.Println("false")
		}
	}
}

func stack_queue_test() {
	tester := algs.NewQueue()
	tester.Test()
}

func linkedlist_test() {
	linkedlist.Demo()
}

func graph_test() {
	// g := algs.NewGraphWithV(5)
	// g.AddEdge(2,3)
	//g := algs.NewGraphWithFile(stdin.NewIn("graph.txt"))
	//g := algs.GraphGeneratorSimple(7, 13)
	//fmt.Println(g)
}

func dfs_test() {
	//g := algs.NewGraphWithFile(stdin.NewIn("graph.txt"))
	g := algs.GraphGeneratorSimple(13, 13)

	s := 0
	dfs := algs.NewDepthFirstSearch(g, s)

	for v := 0; v < g.V(); v++ {
		if dfs.Marked(v) {
			fmt.Printf("%d ", v)
		}
	}

	fmt.Println(g)

}
