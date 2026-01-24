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
