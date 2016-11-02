package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer timeTrack(time.Now())
	// a := []int{3, 4, 2, 1, 0, 9, 7, 8, 5, 6}
	// fmt.Println(a[0 : len(a)/2])
	// fmt.Println(a[len(a)/2 : len(a)])

	n := int(1e7)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}
	// insertionSort(a)
	// fmt.Println(a)

	// a = mergeSort(a)
	// fmt.Println(a)

	// left := []int{0, 3, 6, 8, 9, 10}
	// right := []int{1, 2, 4, 5, 7, 12, 15, 19}
	// fmt.Println(merge(left, right))

	// a := []int{1, 2, 7, 3, 10, 3, 4, 5, 14, 0}
	// buildMaxHeap(a)
	a = heapSort(a)
	// fmt.Println(a)

}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
