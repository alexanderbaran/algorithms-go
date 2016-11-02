package main

// O(n^2)
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}

// O(nlogn)
func mergeSort(a []int) []int {
	// fmt.Println(a)
	if len(a) <= 1 {
		return a
	}
	// left := []int{}
	// right := []int{}
	// for i, v := range a {
	// 	if i < len(a)/2 {
	// 		left = append(left, v)
	// 	} else {
	// 		right = append(right, v)
	// 	}
	// }
	// left = mergeSort(left)
	// right = mergeSort(right)
	left := mergeSort(a[0 : len(a)/2])
	right := mergeSort(a[len(a)/2 : len(a)])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	a := []int{} // Resulting array.

	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			a = append(a, l[0])
			l = l[1:]
		} else {
			a = append(a, r[0])
			r = r[1:]
		}
	}
	for len(l) > 0 {
		a = append(a, l[0])
		l = l[1:]
	}
	for len(r) > 0 {
		a = append(a, r[0])
		r = r[1:]
	}

	// i, j := 0, 0
	// for i < len(l) && j < len(r) {
	// 	if l[i] < r[j] {
	// 		a = append(a, l[i])
	// 		i++
	// 	} else {
	// 		a = append(a, r[j])
	// 		j++
	// 	}
	// }
	// if i != len(l) {
	// 	for ; i < len(l); i++ {
	// 		a = append(a, l[i])
	// 	}
	// }
	// if j != len(r) {
	// 	for ; j < len(r); j++ {
	// 		a = append(a, r[j])
	// 	}
	// }
	return a
}

func heapSort(a []int) []int {
	// 1. buildMaxHeap from unordered array.
	// 2. Find maxElement a[0]
	// 3. Swap elements a[len(a)-1] with a[0]. Now maxElement is at the end of array.
	//    No need to extract.
	// 4. Discard last element from heap, decrementing heapSize.
	// 5. New root may violate max heap property, but children are max heaps.
	//    Run maxHeapify(a, i)
	// 6. Back to step 2 and repeat.
	buildMaxHeap(a)
	heapSize := len(a)
	// result := []int{}
	result := make([]int, heapSize)
	for i := heapSize - 1; i >= 0; i-- {
		// result = append([]int{a[0]}, result...) // Prepend.
		result[i] = a[0]
		a[0], a[len(a)-1] = a[len(a)-1], a[0]
		a = a[0 : len(a)-1]
		maxHeapify(a, 0)
	}
	return result
}

// func duplicate(a []int) []int {
// 	tmp := make([]int, len(a))
// 	copy(tmp, a)
// 	return tmp
// }
