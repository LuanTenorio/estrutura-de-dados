package binarysearch

type BinarySearchTree interface {
	Insert(key int, data interface{})
	InOrder()
	Search(key int) interface{}
}
