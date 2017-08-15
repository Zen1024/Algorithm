package tree

import (
	"testing"
)

func TestPreOrder(t *testing.T) {
	tree := newRandTree(10, 100)
	if tree == nil {
		t.Fatalf("empty tree!")
	}
	tree.Print()
	tree.PreOrderRec()
	tree.PreOrderNonRec()
}
