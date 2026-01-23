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
	Data   interface{}
}

type RedBlackTreeImpl struct {
	Root *treeNode
}
