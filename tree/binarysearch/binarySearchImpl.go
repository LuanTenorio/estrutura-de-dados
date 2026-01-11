package binarysearch

type treeNode struct {
	Left   *treeNode
	Right  *treeNode
	Father *treeNode
	Key    int
	Data   interface{}
}

type BinarySearchTreeImpl struct {
	Root *treeNode
}

func New() BinarySearchTree {
	return &BinarySearchTreeImpl{
		Root: nil,
	}
}

func createNode(key int, data interface{}) *treeNode {
	return &treeNode{
		Key:  key,
		Data: data,
	}
}

func (bst *BinarySearchTreeImpl) Insert(key int, data interface{}) {
	newNode := createNode(key, data)

	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	var lastNode *treeNode
	curNode := bst.Root

	for curNode != nil {
		lastNode = curNode
		if curNode.Key > key {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}

	newNode.Father = lastNode
	if newNode.Father.Key > key {
		newNode.Father.Left = newNode
	} else {
		newNode.Father.Right = newNode
	}
}
