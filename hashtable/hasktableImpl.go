package hashtable

import "errors"

type HashFunc func(key int, tableSize int) int

type Student struct {
	Name       string
	Id         int
	N1, N2, N3 float32
}

type HashtableImpl struct {
	TABLE_SIZE, Amount int
	Itens              []*Student
	hashFunc           HashFunc
}

func New(TABLE_SIZE int, hash HashFunc) Hashtable {
	return &HashtableImpl{
		TABLE_SIZE: TABLE_SIZE,
		Itens:      make([]*Student, TABLE_SIZE),
		hashFunc:   hash,
	}
}

func DivHash(key int, TABLE_SIZE int) int {
	return (key & 0x7fffffff) % TABLE_SIZE
}

func MultHash(key int, TABLE_SIZE int) int {
	A := 0.6180339887
	value := float64(key) * A
	value = value - float64(int(value))

	return int(value) * TABLE_SIZE
}

func (h *HashtableImpl) Hash(key int) int {
	return h.hashFunc(key, h.TABLE_SIZE)
}

func (h *HashtableImpl) Set(student *Student) error {
	if h.Amount == h.TABLE_SIZE {
		return errors.New("The table is full")
	}

	index := h.Hash(student.Id)

	if h.Itens[index] == nil {
		return errors.New("Conflict when inserting element")
	}

	h.Itens[index] = student
	h.Amount++

	return nil
}

func (h *HashtableImpl) Get(key int) *Student {
	index := h.Hash(key)

	return h.Itens[index]
}
