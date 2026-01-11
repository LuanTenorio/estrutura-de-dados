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

func (sbt *BinarySearchTreeImpl) searchNode(key int) *treeNode {
	curNode := sbt.Root

	for curNode != nil && curNode.Key != key {
		if curNode.Key > key {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}

	return curNode
}

func (sbt *BinarySearchTreeImpl) Search(key int) interface{} {
	node := sbt.searchNode(key)

	if node == nil {
		return nil
	}

	return node.Data
}

func (sbt *BinarySearchTreeImpl) deleteLeafNode(node *treeNode, isRoot bool) {
	if isRoot {
		sbt.Root = nil
		return
	}

	if node.Father.Key > node.Key {
		node.Father.Left = nil
	} else {
		node.Father.Right = nil
	}

	node.Father = nil
}

func (sbt *BinarySearchTreeImpl) deleteNodeWithOneNextNode(node *treeNode, isRoot bool) {
	nextNode := node.Left

	if nextNode == nil {
		nextNode = node.Right
	}

	if isRoot {
		sbt.Root = nextNode
		return
	}

	if node.Key > node.Father.Key {
		node.Father.Right = nextNode
	} else {
		node.Father.Left = nextNode
	}

	nextNode.Father = node.Father
}

func (sbt *BinarySearchTreeImpl) deleteNodeWithTwoNextNodes(node *treeNode) {
	successor := node.Right

	for successor.Left != nil {
		successor = successor.Left
	}

	key, data := successor.Key, successor.Data

	sbt.selectDelectionMode(successor, false)

	node.Key, node.Data = key, data
}

func (sbt *BinarySearchTreeImpl) selectDelectionMode(node *treeNode, isRoot bool) {
	if node.Left == nil && node.Right == nil {
		sbt.deleteLeafNode(node, isRoot)
		return
	}

	if node.Left == nil || node.Right == nil {
		sbt.deleteNodeWithOneNextNode(node, isRoot)
		return
	}

	sbt.deleteNodeWithTwoNextNodes(node)
}

func (sbt *BinarySearchTreeImpl) Delete(key int) {
	node := sbt.searchNode(key)

	if node == nil {
		return
	}

	isRoot := sbt.Root == node

	sbt.selectDelectionMode(node, isRoot)
}
