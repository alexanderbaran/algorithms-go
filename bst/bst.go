package bst

import "fmt"

/* Binary Search Trees (BST) are actual trees compared to the heap data structure so they
need pointers. Allows for fast access O(logn), search O(logn), insertion O(logn),
deletion O(logn).

The tree satisfies the binary search tree property, which states that the key in each
node must be greater than all keys stored in the left sub-tree, and less than all keys
in the right sub-tree.


1. A tree consists of nodes that store unique values.
2. Each node has zero, one, or two child nodes.
3. One of the nodes is designated as the root node that is at the top of the tree
   structure. (This is the “entry point” where all operations start.)
4. Each node has exactly one parent node, except for the root node, which has no parent.
5. Each node’s value is larger than the value of its left child but smaller than the
   value of its right child.
6. Each subtree to the left of a node only contains values that are smaller than the
   node’s value, and each subtree to the right of a node contains only larger values.

To implement:

	- Add and remove
	- Searching
	- Traversals
		- Pre-Order
		- In-Order
		- Post-Order */

type Tree struct {
	Root *Node
	Size int
}

func New() *Tree {
	return &Tree{}
}

type Node struct {
	Key int
	// size  int
	Left  *Node
	Right *Node
}

func (t *Tree) Add(key int) {
	n := &Node{Key: key}
	if t.Root == nil {
		t.Root = n
		t.Size++
		return
	}
	if t.Root.Add(n) {
		t.Size++
	}
}

func (root *Node) Add(n *Node) bool {
	if n.Key < root.Key {
		if root.Left == nil {
			root.Left = n
			return true
		} else {
			return root.Left.Add(n)
		}
	} else if n.Key > root.Key {
		if root.Right == nil {
			root.Right = n
			return true
		} else {
			return root.Right.Add(n)
		}
	}
	return false
}

func (t *Tree) Find(key int) *Node {
	return find(t.Root, key)
}

func find(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if key == n.Key {
		return n
	}
	if key < n.Key {
		return find(n.Left, key)
	}
	return find(n.Right, key)
}

func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Println(n.Key)
		InOrder(n.Right)
	}
}

func PreOrder(n *Node) {
	if n != nil {
		fmt.Println(n.Key)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func PostOrder(n *Node) {
	if n != nil {
		PostOrder(n.Left)
		PostOrder(n.Right)
		fmt.Println(n.Key)
	}
}
