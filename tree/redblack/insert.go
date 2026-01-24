package redblack

func (rbt *RedBlackTreeImpl) insertFixup(node *treeNode) {
	for node.Father.color == red {
		if rbt.isLeftSon(node.Father) {
			uncle := node.Father.Father.Right

			if uncle.color == red {
				node.Father.color, uncle.color = black, black
				node.Father.Father.color = red
				node = node.Father.Father
			} else {
				if rbt.isRightSon(node) {
					node = node.Father
					rbt.leftRotate(node)
				}

				node.Father.color = black
				node.Father.Father.color = red
				rbt.rightRotate(node.Father.Father)
			}
		} else {
			uncle := node.Father.Father.Left

			if uncle.color == red {
				node.Father.color, uncle.color = black, black
				node.Father.Father.color = red
				node = node.Father.Father
			} else {
				if rbt.isLeftSon(node) {
					node = node.Father
					rbt.rightRotate(node)
				}

				node.Father.color = black
				node.Father.Father.color = red
				rbt.leftRotate(node.Father.Father)
			}
		}

	}

	rbt.root.color = black
}

func (rbt *RedBlackTreeImpl) Insert(key int) {
	newNode := rbt.newNode(key)
	y := rbt.nilNode
	x := rbt.root

	for !rbt.isNilNode(x) {
		y = x

		if newNode.Key > x.Key {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	newNode.Father = y

	if rbt.isNilNode(y) {
		rbt.root = newNode
	} else if newNode.Key > newNode.Father.Key {
		y.Right = newNode
	} else {
		y.Left = newNode
	}

	rbt.insertFixup(newNode)
}
