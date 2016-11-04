package bst

import "fmt"

/* Binary Search Trees (BST) are actual trees compared to the heap data structure so they
need pointers. Allows for fast access O(logn), search O(logn), insertion O(logn),
deletion O(logn).

1. A tree consists of nodes that store unique values.
2. Each node has zero, one, or two child nodes.
3. One of the nodes is designated as the root node that is at the top of the tree
   structure. (This is the “entry point” where all operations start.)
4. Each node has exactly one parent node, except for the root node, which has no parent.
5. Each node’s value is larger than the value of its left child but smaller than the
   value of its right child.
6. Each subtree to the left of a node only contains values that are smaller than the
   node’s value, and each subtree to the right of a node contains only larger values.

BST can be used to sort an unsorted array as we can do with a heap data structure.
In practice, the added overhead in time and space for a tree-based sort (particularly
for node allocation) make it inferior to other asymptotically optimal sorts such as
heapsort for static list sorting. On the other hand, it is one of the most efficient
methods of incremental sorting, adding items to a list over time while keeping the
list sorted at all times. */

type Tree struct {
	Root *Node
	Size int
}

func New() *Tree {
	return &Tree{}
}

// O(log n) space in the average case and O(n) in the worst case.
func FromSlice(a []int) *Tree {
	t := &Tree{}
	for i := 0; i < len(a); i++ {
		t.Add(a[i])
	}
	return t
}

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
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
		// } else if n.Key >= root.Key { // Need tree to take duplicates for sort.
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

func minimum(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func maximum(n *Node) *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Println(n.Key)
		InOrder(n.Right)
	}
}

func InOrderAdd(a *[]int, n *Node, i *int) {
	if n != nil {
		InOrderAdd(a, n.Left, i)
		// *a = append(*a, n.Key)
		s := *a
		s[*i] = n.Key
		*i++
		InOrderAdd(a, n.Right, i)
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

func (t *Tree) IsBST() bool {
	if t.Root == nil {
		return true
	}
	return isBST(t.Root, minimum(t.Root).Key, maximum(t.Root).Key)
}

func isBST(n *Node, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Key < min || n.Key > max {
		return false
	}
	return isBST(n.Left, min, n.Key) && isBST(n.Right, n.Key, max)
}
