package tree

import "cmp"

type Node[E cmp.Ordered] struct {
	data  E
	left  *Node[E]
	right *Node[E]
}

type BinarySearchTree[E cmp.Ordered] struct {
	root *Node[E]
	size int
}

func NewBinarySearchTree[E cmp.Ordered]() *BinarySearchTree[E] {
	return &BinarySearchTree[E]{}
}

func (t *BinarySearchTree[E]) Size() int {
	return t.size
}

func (t *BinarySearchTree[E]) Empty() bool {
	return t.size == 0
}

func (t *BinarySearchTree[E]) Add(value E) {
	root, ok := addRecursive(t.root, value)
	if ok {
		t.size++
	}
	t.root = root
}

func addRecursive[E cmp.Ordered](node *Node[E], value E) (n *Node[E], ok bool) {
	if node == nil {
		return &Node[E]{data: value}, true
	}
	if value < node.data {
		node.left, ok = addRecursive(node.left, value)
	} else if value > node.data {
		node.right, ok = addRecursive(node.right, value)
	}
	// duplicate: do nothing
	return node, ok
}

func (t *BinarySearchTree[E]) Search(value E) bool {
	return searchRecursive(t.root, value)
}

func searchRecursive[E cmp.Ordered](node *Node[E], value E) bool {
	if node == nil {
		return false
	}
	if value < node.data {
		return searchRecursive(node.left, value)
	} else if value > node.data {
		return searchRecursive(node.right, value)
	}
	return true
}

func (t *BinarySearchTree[E]) Del(value E) {
	var deleted bool
	t.root, deleted = delRecursive(t.root, value)
	if deleted {
		t.size--
	}
}

func delRecursive[E cmp.Ordered](node *Node[E], value E) (*Node[E], bool) {
	if node == nil {
		return nil, false
	}
	switch {
	case value < node.data:
		node.left, _ = delRecursive(node.left, value)
	case value > node.data:
		node.right, _ = delRecursive(node.right, value)
	default:
		// node.data == value
		if node.left == nil {
			return node.right, true
		}
		if node.right == nil {
			return node.left, true
		}
		// Two children
		successor := minChild(node.right)
		node.data = successor.data
		node.right, _ = delRecursive(node.right, successor.data)
	}
	return node, true
}

func (t *BinarySearchTree[E]) Min() (E, bool) {
	if t.root == nil {
		var zero E
		return zero, false
	}
	return minChild(t.root).data, true
}

func minChild[E cmp.Ordered](node *Node[E]) *Node[E] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *BinarySearchTree[E]) Max() (E, bool) {
	if t.root == nil {
		var zero E
		return zero, false
	}
	return maxChild(t.root).data, true
}

func maxChild[E cmp.Ordered](node *Node[E]) *Node[E] {
	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *BinarySearchTree[E]) InOrder(visit func(E) bool) {
	inOrder(t.root, visit)
}
func inOrder[E cmp.Ordered](node *Node[E], visit func(E) bool) {
	if node == nil {
		return
	}
	inOrder(node.left, visit)
	if !visit(node.data) {
		return
	}
	inOrder(node.right, visit)
}

func (t *BinarySearchTree[E]) PreOrder(visit func(E) bool) {
	preOrder(t.root, visit)
}
func preOrder[E cmp.Ordered](node *Node[E], visit func(E) bool) {
	if node == nil {
		return
	}
	if !visit(node.data) {
		return
	}
	preOrder(node.left, visit)
	preOrder(node.right, visit)
}

func (t *BinarySearchTree[E]) PostOrder(visit func(E) bool) {
	postOrder(t.root, visit)
}

func postOrder[E cmp.Ordered](node *Node[E], visit func(E) bool) {
	if node == nil {
		return
	}
	postOrder(node.left, visit)
	postOrder(node.right, visit)
	if !visit(node.data) {
		return
	}
}
