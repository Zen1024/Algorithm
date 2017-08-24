package tree

import (
	"math/rand"
	"testing"
	"time"
)

func randArr(max, length int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	re := make([]int, length)
	for i := 0; i < length; i++ {
		re[i] = rnd.Intn(max)
	}
	return re
}

func TestPreOrder(t *testing.T) {
	arr := randArr(100, 10)
	tree := NewTree(arr)
	tree.Print()
	tree.PreOrderRec()
	tree.PreOrderNonRec()
	tree.PreOrderNonRecAndStack()
	tree.PostOrderNonRecAndStack()
	tree.LevelOrder()
	// tree.Mirror()
	// tree.Print()
}

func TestUtils(t *testing.T) {
	arr := randArr(100, 10)
	tree := NewTree(arr)
	tree.Print()
	t.Logf("cnt:%d\n", tree.NodeCnt())
	t.Logf("leaf cnt:%d\n", tree.LeafCnt())
	t.Logf("depth:%d\n", tree.Depth())
	tree2 := newRandTree(10, 100)
	t.Logf("tree equals tree_2?%v", tree.Equals(tree2))
	t.Logf("tree equals tree?%v", tree.Equals(tree))
}

func TestLCAandDist(t *testing.T) {
	arr := randArr(100, 10)
	tree := NewTree(arr)
	tree.Print()
	node1 := arr[3]
	node2 := arr[5]
	lca := tree.LCA(node1, node2)
	if lca != nil {
		t.Logf("lca of %d and %d:%d", node1, node2, lca.Element)
	} else {
		t.Fatalf("empty lca!")
	}
	dist := tree.Dist(node1, node2)
	if dist == -1 {
		t.Fatalf("invalid dist!")
	} else {
		t.Logf("dist of %d and %d:%d", node1, node2, dist)
	}

}
