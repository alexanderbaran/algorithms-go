package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer timeTrack(time.Now())
	// a := []int{3, 4, 2, 1, 0, 9, 7, 8, 5, 6}

	n := int(1e6)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}

	// sort.Ints(a)

	// sort.InsertionSort(a)

	// a = sort.MergeSort(a)

	// a = sort.HeapSort(a)

	// fmt.Println(a)

	// bst := bst.New()
	// bst.Add(6)
	// bst.Add(4)
	// bst.Add(7)
	// bst.Add(5)
	// bst.Add(5)
	// bst.Add(3)
	// bst.Add(5)
	// spew.Dump(bst)
	// fmt.Println(bst.Find(6))
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
