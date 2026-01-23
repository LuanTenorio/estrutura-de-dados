package redblack

const (
	red uint8 = iota
	black
)

type treeNode struct {
	Left   *treeNode
	Right  *treeNode
	Father *treeNode
	color  uint8
	Key    int
}

type RedBlackTreeImpl struct {
	root    *treeNode
	nilNode *treeNode
}

func NewRBT() RedBlackTree {
	nilNode := &treeNode{color: black}
	return &RedBlackTreeImpl{nilNode: nilNode}
}

func (rbt *RedBlackTreeImpl) newNode(key int) *treeNode {
	return &treeNode{
		color: red,
		Left:  rbt.nilNode,
		Right: rbt.nilNode,
		Key:   key,
	}
}

func (rbt *RedBlackTreeImpl) isNilNode(node *treeNode) bool {
	return node == rbt.nilNode
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

func (rbt *RedBlackTreeImpl) insertFixup(x *treeNode) {

}

func (rbt *RedBlackTreeImpl) insert(key int) {
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
