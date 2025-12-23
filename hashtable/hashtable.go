package hashtable

type Hashtable interface {
	Hash(key int) int
	Get(key int) *Student
	Set(student *Student) error
}
