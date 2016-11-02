package main

/* Binary Search Trees (BST) are actual trees compared to the heap data structure so they
need pointers. Allows for fast access O(logn), search O(logn), insertion O(logn),
deletion O(logn).

The tree satisfies the binary search tree property, which states that the key in each
node must be greater than all keys stored in the left sub-tree, and less than all keys
in the right sub-tree. */

type Node struct {
	key int
	// size  int
	left  *Node
	right *Node
}
