package binarysearch

type node struct {
	Left   *node
	Right  *node
	Father *node
	key    int
	data   interface{}
}

type BinarySearchTreeImpl struct {
	Root *node
}

func New() BinarySearchTree {
	return &BinarySearchTreeImpl{
		Root: nil,
	}
}
