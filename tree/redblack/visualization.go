package redblack

import "fmt"

func (rbt *RedBlackTreeImpl) printInOrder(node *treeNode) {
	if rbt.isNilNode(node) {
		return
	}

	rbt.printInOrder(node.Left)
	fmt.Println(node.Key)
	rbt.printInOrder(node.Right)
}

func (rbt *RedBlackTreeImpl) InOrder() {
	rbt.printInOrder(rbt.root)
}
