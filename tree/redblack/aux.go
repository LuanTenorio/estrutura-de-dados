package redblack

func (rbt *RedBlackTreeImpl) isNilNode(node *treeNode) bool {
	return node == rbt.nilNode
}

func (rbt *RedBlackTreeImpl) isRightSon(node *treeNode) bool {
	return node != rbt.root && node.Father == node.Father.Father.Right
}

func (rbt *RedBlackTreeImpl) isLeftSon(node *treeNode) bool {
	return node != rbt.root && node.Father == node.Father.Father.Left
}

func (rbt *RedBlackTreeImpl) leftRotate(x *treeNode) {
	y := x.Right
	x.Right = y.Left

	if !rbt.isNilNode(y.Left) {
		y.Left.Father = x
	}

	if !rbt.isNilNode(x.Father) {
		rbt.root = y
	} else if x == x.Father.Left {
		x.Father.Left = y
	} else {
		x.Father.Right = y
	}

	y.Father, y.Left, x.Father = x.Father, x, y
}

func (rbt *RedBlackTreeImpl) rightRotate(x *treeNode) {
	y := x.Left
	x.Left = y.Right

	if !rbt.isNilNode(y.Right) {
		y.Right.Father = x
	}

	if rbt.isNilNode(x.Father) {
		rbt.root = y
	} else if x == x.Father.Right {
		x.Father.Right = y
	} else {
		x.Father.Left = y
	}

	y.Father, y.Right, x.Father = x.Father, x, y
}
