package main

import "fmt"

type treeNode struct{
	data int
	left *treeNode
	right *treeNode
	parent *treeNode
}

type tree struct {
	root *treeNode
}

func (t *tree) search(data int) *treeNode {
	curr, _ := treeSearch(t.root, data, nil)
	return curr
}

func (t *tree) insert(data int) {
	n := &treeNode{
		data:   data,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	if t.root == nil {
		t.root = n
		return
	}
	_, parent := treeSearch(t.root, data, nil)
	n.parent = parent
	if parent.data < n.data {
		parent.right = n
	} else {
		parent.left = n
	}
}

func (t *tree) delete(data int) {
	if t.root == nil {
		return
	}
	curr, parent := treeSearch(t.root, data, nil)
	if curr == nil {
		return
	}
	// leaf node
	if curr.left == nil && curr.right == nil {
		t.leafNodeDelete(curr, parent)
		return
	}
	// node with one child
	if curr.left == nil || curr.right == nil {
		t.oneChildNodeDelete(curr, parent)
		return
	}
	// node with 2 children
	t.twoChildNodeDelete(curr, parent)
}

func (t *tree) leafNodeDelete(curr *treeNode, parent *treeNode) {
	if parent == nil {
		t.root = nil
		return
	}
	if parent.left == curr {
		parent.left = nil
	} else {
		parent.right = nil
	}
}

func (t *tree) oneChildNodeDelete(curr *treeNode, parent *treeNode) {
	grandChildPtr := curr.left
	if curr.right != nil {
		grandChildPtr = curr.right
	}
	if parent == nil {
		t.root = grandChildPtr
		return
	}
	if parent.left == curr {
		parent.left = grandChildPtr
	} else {
		parent.right = grandChildPtr
	}
}

func (t *tree) twoChildNodeDelete(curr *treeNode, parent *treeNode) {
	leftmostNode := dfsLeft(curr.right)
	// swap leftmostNode and curr
	leftNodeOfCurr := curr.left
	rightNodeOfCurr := curr.right

	// move curr to leftmostNode's position and attach any right subtrees
	curr.parent = leftmostNode.parent
	curr.right = leftmostNode.right

	// point parent or root, left and right of curr to leftmostNode
	if parent == nil {
		t.root = leftmostNode
	} else {
		if parent.left == curr {
			parent.left = leftmostNode
		} else {
			parent.right = leftmostNode
		}
	}
	leftmostNode.left = leftNodeOfCurr
	leftmostNode.right = rightNodeOfCurr

	// delete curr
	t.oneChildNodeDelete(curr, curr.parent)
}

// same as traverse
func (t *tree) print() {
	fmt.Printf("print tree \n")
	treePrint(t.root)
	fmt.Printf("\nend print tree \n")
}

func treePrint(node *treeNode) {
	if node == nil {
		return
	}
	treePrint(node.left)
	fmt.Printf(" %v ", node.data)
	treePrint(node.right)
}

func treeSearch(node *treeNode, data int, prev *treeNode) (*treeNode, *treeNode) {
	// not found
	if node == nil {
		return nil, prev
	}
	// found
	if node.data == data {
		return node, prev
	}
	// data is less than curr node's data, scan the left tree
	if data < node.data {
		return treeSearch(node.left, data, node)
	}
	// data is greater than curr node's data, scan the right tree
	return treeSearch(node.right, data, node)
}

func dfsLeft(n *treeNode) *treeNode {
	if n == nil {
		return nil
	}
	// found leaf
	if n.left == nil && n.right == nil {
		return n
	}
	return dfsLeft(n.left)
}

func main() {
	t := &tree{}
	t.insert(5)
	t.insert(3)
	t.insert(1)
	t.insert(2)
	t.print()
	t.delete(2)
	t.print()
	t.delete(3)
	t.print()
	t.delete(5)
	t.print()
	t.delete(1)
	t.print()
}
