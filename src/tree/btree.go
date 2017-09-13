package tree

import (
	"fmt"
)

const (
	MaxBTreeKeyNum int = 10
)

type BNode struct {
	KCnt   int
	Keys   []int
	Childs []*BNode
	IsLeaf bool
	parent *BNode
}

func NewBTree() *BNode {
	return &BNode{
		KCnt:   0,
		Keys:   nil,
		Childs: nil,
		IsLeaf: true,
		parent: nil,
	}
}

func (t *BNode) split() (key int, r *BNode) {

	if t == nil || t.KCnt <= MaxBTreeKeyNum {
		return -1, t, nil
	}

	idx := t.KCnt / 2
	left := t
	k := t.Keys[idx]
	lks := t.Keys[:idx]
	rks := t.Keys[idx+1:]
	lcs := t.Childs[:idx]
	rcs := t.Childs[idx:]

	right := &BNode{
		parent: t.parent,
		Keys:   rks,
		Childs: rcs,
		KCnt:   len(lks),
		IsLeaf: t.IsLeaf,
	}
	left.Childs = lcs
	left.Keys = lks
	left.KCnt = len(lks)
	return k, right

}
func (t *BNode) splitChild(i int) *BNode {
	if t == nil {
		return nil
	}

	m := MaxBTreeKeyNum + 1
	if i >= t.KCnt {
		return t
	}
	c := t.Childs[i]
	if c.Keys <= MaxBTreeKeyNum {
		return t
	}
	idx := MaxBTreeKeyNum / 2
	k := c.Keys[idx]

	_, newChild := c.split()

	t.KCnt++
	t.Keys = append(t.Keys, -1)
	t.Childs = append(t.Childs, nil)

	for j := t.KCnt; j > i+1; j-- {
		t.Keys[j] = t.Keys[j-1]
	}
	t.Keys[i+1] = k

	for j := t.KCnt + 1; j > i+1; j-- {
		t.Childs[j] = t.Childs[j-1]
	}
	t.Childs[i+1] = newChild

	for p := t; p != nil && p.KCnt > MaxBTreeKeyNum; p = p.parent {
		pp := p.parent
		if pp == nil {
			key, newc := p.split()
			pnode := &BNode{
				parent: nil,
				Childs: []*BNode{p, newc},
				Keys:   []int{key},
				IsLeaf: false,
				KCnt:   1,
			}
			p.parent = pnode
			newc.parent = pnode
			return t
		}
		for idx = 0; pp.Childs[idx] != p && idx <= p.KCnt; idx++ {
		}

		//can't be hare
		if idx > p.KCnt {
			panic("shouldn't be here!")
			return t
		}
		pp.splitChild(idx)
	}

	return t
}

//@TODO
func (t *BNode) Insert(elem int) *BNode {
	var child *BrNode = nil
	for i := 0; i < t.KCnt; i++ {
		if t.Keys[i] > elem {
			child = t.Childs[i]
		}
		if t.Keys[i] == elem {
			return t
		}
	}
	if child == nil {
		child = t.Childs[t.KCnt]
	}

}
