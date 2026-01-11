package binarysearch

type BinarySearchTree interface {
	Insert(key int, data interface{})
	Search(key int) interface{}
	Delete(key int)
	InOrder()
	PreOrder()
	PosOrder()
}
