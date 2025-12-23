package hashtable_test

import (
	"testing"

	"github.com/LuanTenorio/estrutura-de-dados/hashtable"
)

func TestSetAndGet_HappyPath(t *testing.T) {
	ht := hashtable.New(10, hashtable.MultHash)

	s := &hashtable.Student{Name: "Alice", Id: 42}

	if err := ht.Set(s); err != nil {
		return
	}

	got := ht.Get(42)

	if got == nil {
		t.Fatalf("expected student, got nil")
	}

	if got.Name != "Alice" || got.Id != 42 {
		t.Fatalf("unexpected student: %+v", got)
	}
}

func TestGet_NotFound(t *testing.T) {
	ht := hashtable.New(5, hashtable.MultHash)

	if got := ht.Get(1); got != nil {
		t.Fatalf("expected nil for empty table, got %+v", got)
	}
}

func TestSet_TableFullOrConflict(t *testing.T) {
	ht := hashtable.New(1, hashtable.MultHash)

	s1 := &hashtable.Student{Name: "A", Id: 1}
	s2 := &hashtable.Student{Name: "B", Id: 2}

	ht.Set(s1)

	if err := ht.Set(s2); err != nil {
		return
	}

	if got := ht.Get(2); got == nil {
		t.Fatalf("expected student after Set, got nil")
	}
}
