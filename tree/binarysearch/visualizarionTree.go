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
