package main

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now())
	// a := []int{3, 4, 2, 1, 0, 9, 7, 8, 5, 6}
	// fmt.Println(a[0 : len(a)/2])
	// fmt.Println(a[len(a)/2 : len(a)])

	// n := int(1e4)
	// a := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	a[i] = rand.Intn(n)
	// }
	// sort.InsertionSort(a)
	// fmt.Println(a)

	// a = sort.MergeSort(a)
	// fmt.Println(a)

	// a := []int{1, 2, 7, 3, 10, 3, 4, 5, 14, 0}
	// heap.BuildMaxHeap(a)
	// a = sort.HeapSort(a)
	// fmt.Println(a)

	// root := bst.Node{}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
