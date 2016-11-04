package main

import (
	"algorithms-go/bst"
	"fmt"
	"time"
)

func main() {
	// defer timeTrack(time.Now())
	// a := []int{3, 4, 2, 1, 0, 9, 7, 8, 5, 6}

	// n := int(1e6)
	// a := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	a[i] = rand.Intn(n)
	// }

	// sort.Ints(a)

	// sort.InsertionSort(a)

	// a = sort.MergeSort(a)

	// a = sort.HeapSort(a)

	// fmt.Println(a)

	t := bst.New()
	t.Add(4)
	t.Add(6)
	t.Add(7)
	t.Add(5)
	t.Add(2)
	t.Add(3)
	t.Add(1)
	// fmt.Println(t.Find(6))
	// bst.InOrder(t.Root)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
