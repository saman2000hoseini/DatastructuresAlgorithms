package ds

import "fmt"

type BSTNode struct {
	data                int
	left, right, parent *BSTNode
}

type BST struct {
	length uint
	root   *BSTNode
}

func newBSTNode(data int) *BSTNode {
	return &BSTNode{data, nil, nil, nil}
}

func NewBST() *BST {
	return &BST{0, nil}
}

func InOrder(root *BSTNode) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Print(string(root.data) + " ")
	InOrder(root.right)
}

func searchNode(data int, root *BSTNode) *BSTNode {

	if root.data > data {
		return searchNode(data, root.left)
	} else {
		return searchNode(data, root.right)
	}
}

func (myBst *BST) FindNode(data int) *BSTNode {
	if myBst.root == nil || myBst.root.data == data {
		return myBst.root
	}
	return searchNode(data, myBst.root)
}

func insertBstNode(root *BSTNode, data int) {
	if root.data > data {
		if root.left == nil {
			root.left = newBSTNode(data)
			root.left.parent = root
			return
		} else {
			insertBstNode(root.left, data)
		}
	} else {
		if root.right == nil {
			root.right = newBSTNode(data)
			root.right.parent = root
			return
		} else {
			insertBstNode(root.right, data)
		}
	}
}

func (myBst *BST) InsertNode(data int) {
	if myBst.root == nil {
		myBst.root = newBSTNode(data)
		return
	}
	insertBstNode(myBst.root, data)
}

func minBstNode(root *BSTNode) *BSTNode {
	if root != nil && root.left != nil {
		return minBstNode(root.left)
	}
	return root
}

func deleteBstNode(root *BSTNode, data int) {
	if root == nil {
		return
	}
	if root.data == data {
		if root.right == nil {
			root = root.left
			return
		} else if root.left == nil {
			root = root.right
			return
		}
		root.data = minBstNode(root.right).data
		deleteBstNode(root.right, data)
		return
	} else if root.data > data {
		deleteBstNode(root.left, data)
	} else {
		deleteBstNode(root.right, data)
	}
}

func (myBst *BST) DeleteNode(data int) {
	if myBst.root == nil {
		return
	}
	deleteBstNode(myBst.root, data)
}
