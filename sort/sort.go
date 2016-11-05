package sort

import (
	"algorithms-go/bst"
	"algorithms-go/heap"
	"math"
)

// O(n^2)
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}

// O(nlogn)
func MergeSort(a []int) []int {
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
	left := MergeSort(a[0 : len(a)/2])
	right := MergeSort(a[len(a)/2 : len(a)])
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

// O(nlogn)
func HeapSort(a []int) []int {
	// 1. buildMaxHeap from unordered array.
	// 2. Find maxElement a[0]
	// 3. Swap elements a[len(a)-1] with a[0]. Now maxElement is at the end of array.
	//    No need to extract.
	// 4. Discard last element from heap, decrementing heapSize.
	// 5. New root may violate max heap property, but children are max heaps.
	//    Run maxHeapify(a, i)
	// 6. Back to step 2 and repeat.
	heap.BuildMaxHeap(a)
	heapSize := len(a)
	// result := []int{}
	result := make([]int, heapSize)
	for i := heapSize - 1; i >= 0; i-- {
		// result = append([]int{a[0]}, result...) // Prepend.
		result[i] = heap.ExtractMax(&a)
		// a[0], a[len(a)-1] = a[len(a)-1], a[0]
		// a = a[0 : len(a)-1]
		// heap.MaxHeapify(a, 0)
	}
	return result
}

// Best O(nlogn), average O(nlogn), worst O(n^2).
func BSTSort(a []int) []int {
	t := bst.FromSlice(a)
	s := make([]int, len(a))
	i := 0
	bst.InOrderAdd(&s, t.Root, &i)
	return s
}

// O(n+k), k = max
func CountingSort(a []int, max int) {
	// func CountingSort(a []int, max int) []int {
	c := make([]int, max)
	for _, v := range a {
		c[v]++
	}
	// s := make([]int, len(a))
	i := 0
	for j, v := range c {
		if v != 0 {
			for k := 1; k <= v; k++ {
				a[i] = j
				i++
			}
		}
	}
	// return s
}

func RadixSort(a []int, max int) {
	var bin [][]int
	for k := 1; k <= digitLength(max); k++ {
		bin = make([][]int, 10)
		for _, v := range a {
			i := nthDigit(v, k)
			bin[i] = append(bin[i], v)
		}
		i := 0
		for _, s := range bin {
			if s != nil {
				for _, v := range s {
					a[i] = v
					i++
				}
			}
		}
	}
}

func digitLength(x int) int {
	return int(math.Log10(float64(x)) + 1)
}

func nthDigit(number, n int) int {
	return (number / int(math.Pow(10.0, float64(n-1)))) % 10
}

func QuickSort(a []int) []int {
	return nil
}

// func duplicate(a []int) []int {
// 	tmp := make([]int, len(a))
// 	copy(tmp, a)
// 	return tmp
// }
