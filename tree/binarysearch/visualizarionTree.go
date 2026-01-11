package binarysearch

import "fmt"

func printInOrder(node *treeNode) {
	if node == nil {
		return
	}

	printInOrder(node.Left)
	fmt.Printf(" %d", node.Key)
	printInOrder(node.Right)
}

func (bst *BinarySearchTreeImpl) InOrder() {
	printInOrder(bst.Root)
}

func printPreOrder(node *treeNode) {
	if node == nil {
		return
	}

	fmt.Printf(" %d", node.Key)
	printPreOrder(node.Left)
	printPreOrder(node.Right)
}

func (bst *BinarySearchTreeImpl) PreOrder() {
	printPreOrder(bst.Root)
}

func printPosOrder(node *treeNode) {
	if node == nil {
		return
	}

	printPosOrder(node.Left)
	printPosOrder(node.Right)
	fmt.Printf(" %d", node.Key)
}

func (bst *BinarySearchTreeImpl) PosOrder() {
	printPosOrder(bst.Root)
}
