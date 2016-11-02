package main

// https://en.wikipedia.org/wiki/Binary_heap
// http://www.cse.hut.fi/en/research/SVG/TRAKLA2/tutorials/heap_tutorial/index.html

/* Heap does not need pointers. Can be represented by an array. Good for extracting min
and max values from a min-heap and max-heap data structure.

Heaps where the parent key is greater than or equal to (≥) the child keys are called
max-heaps; those where it is less than or equal to (≤) are called min-heaps. */

// O(logn)
func maxHeapify(a []int, i int) {
	heapSize := len(a)
	left := 2*i + 1 // Index starting at 0, not 1.
	right := 2*i + 2
	largest := i
	if left < heapSize && a[left] > a[largest] {
		largest = left
	}
	if right < heapSize && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest)
	}
}

// O(n)
func buildMaxHeap(a []int) {
	heapSize := len(a)
	for i := (heapSize / 2) - 1; i >= 0; i-- {
		maxHeapify(a, i)
	}
}
