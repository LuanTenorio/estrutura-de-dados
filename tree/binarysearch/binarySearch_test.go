package binarysearch_test

import (
	"testing"

	"github.com/LuanTenorio/estrutura-de-dados/tree/binarysearch"
)

var fixtureNumbers = []int{15, 6, 18, 3, 7, 2, 4, 14, 9, 17, 20}

func buildTree(t *testing.T) binarysearch.BinarySearchTree {
	t.Helper()

	bst := binarysearch.New()
	if bst == nil {
		t.Fatalf("Erro ao criar árvore")
	}

	for _, key := range fixtureNumbers {
		bst.Insert(key, key)
	}

	return bst
}

func TestInsertAndSearch(t *testing.T) {
	bst := buildTree(t)

	for _, key := range fixtureNumbers {
		value, ok := bst.Search(key).(int)
		if !ok || value != key {
			t.Fatalf("Wrong node: key = %d searched = %v", key, value)
		}
	}

	if bst.Search(999) != nil {
		t.Fatalf("Esperava nil para chave inexistente")
	}
}

func TestDeleteLeafNode(t *testing.T) {
	bst := buildTree(t)

	bst.Delete(2)
	if bst.Search(2) != nil {
		t.Fatalf("Folha não removida corretamente")
	}

	if value, ok := bst.Search(3).(int); !ok || value != 3 {
		t.Fatalf("Árvore corrompida após remover folha")
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	bst := buildTree(t)

	bst.Delete(14)
	if bst.Search(14) != nil {
		t.Fatalf("Nó com um filho não removido corretamente")
	}

	if value, ok := bst.Search(9).(int); !ok || value != 9 {
		t.Fatalf("Filho não reencadeado após remover nó com um filho")
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	bst := buildTree(t)

	bst.Delete(6)
	if bst.Search(6) != nil {
		t.Fatalf("Nó com dois filhos não removido corretamente")
	}

	if value, ok := bst.Search(3).(int); !ok || value != 3 {
		t.Fatalf("Subárvore esquerda corrompida após remoção")
	}

	if value, ok := bst.Search(7).(int); !ok || value != 7 {
		t.Fatalf("Subárvore direita corrompida após remoção")
	}
}

func TestDeleteRootWithTwoChildren(t *testing.T) {
	bst := buildTree(t)

	bst.Delete(15)
	if bst.Search(15) != nil {
		t.Fatalf("Root não removida corretamente")
	}

	if value, ok := bst.Search(17).(int); !ok || value != 17 {
		t.Fatalf("Sucessor não promovido corretamente após remover root")
	}
}
