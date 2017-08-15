package tree

import (
	"fmt"
	"math/rand"
	"time"
)

// 二叉树

type BinaryNode struct {
	// Parent     *BinaryNode
	LeftChild  *BinaryNode
	RightChild *BinaryNode
	Element    int
}

type stack struct {
	top  int
	elem []*BinaryNode
}

type node struct {
	prev *node
	next *node
	elem *BinaryNode
}

type queue struct {
	front *node
	end   *node
}

func newQueue() *queue {
	return &queue{
		front: nil,
		end:   nil,
	}
}

func (q *queue) Push(elem *BinaryNode) {
	if q == nil {
		return
	}

	if q.front == nil {
		q.front = &node{
			prev: nil,
			next: nil,
			elem: elem,
		}
		q.end = q.front
	}
	nd := &node{
		prev: q.end,
		next: nil,
		elem: elem,
	}
	q.front.next = nd
}

func (q *queue) Front() *BinaryNode {
	if q == nil {
		return nil
	}
	if q.front == nil {
		return nil
	}

	re := q.front.elem
	q.front = q.front.next

	return re
}

func (s *stack) Push(ele *BinaryNode) {
	s.top++
	p := ele
	s.elem = append(s.elem, p)
	// fmt.Printf("[pushed]top:%d,ele:%d\n", s.top, len(s.elem))
}

func (s *stack) Pop() *BinaryNode {
	if s.top == -1 {
		return nil
	}
	// fmt.Printf("[poped]top:%d,elem:%d\n", s.top, len(s.elem))
	re := s.elem[s.top]
	s.top--
	if s.top > -1 {
		s.elem = s.elem[:s.top+1]
	} else {
		s.elem = []*BinaryNode{}
	}
	return re
}

func (s *stack) Isempty() bool {
	return s.top == -1
}

func newStack() *stack {
	return &stack{
		top:  -1,
		elem: []*BinaryNode{},
	}
}

//随机构造二叉树
func newRandTree(num, max int) *BinaryNode {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var t *BinaryNode = &BinaryNode{
		Element:    rnd.Intn(max),
		LeftChild:  nil,
		RightChild: nil,
	}
	for i := 1; i < num; i++ {
		ele := rnd.Intn(max)
		t.Insert(ele)
	}
	return t
}

func (t *BinaryNode) Insert(ele int) {

	if t.Element >= ele {
		if t.LeftChild == nil {
			t.LeftChild = &BinaryNode{
				Element:    ele,
				LeftChild:  nil,
				RightChild: nil,
			}
			return
		}
		t.LeftChild.Insert(ele)
	} else {
		if t.RightChild == nil {
			t.RightChild = &BinaryNode{
				Element:    ele,
				LeftChild:  nil,
				RightChild: nil,
			}
			return
		}
		t.RightChild.Insert(ele)
	}

	return
}

//前序遍历，递归版
func (t *BinaryNode) PreOrderRec() {
	if t == nil {
		return
	}
	fmt.Printf("%4d ", t.Element)
	t.LeftChild.PreOrderRec()
	t.RightChild.PreOrderRec()
	return
}

func (t *BinaryNode) PreOrderNonRec() {
	fmt.Printf("\n")
	var p *BinaryNode = t
	s := newStack()
	if t == nil {
		return
	}

	for p != nil || !s.Isempty() {
		for p != nil {
			fmt.Printf("%4d ", p.Element)
			s.Push(p)
			p = p.LeftChild
		}
		if !s.Isempty() {
			tmp := s.Pop()
			p = tmp.RightChild
		}
	}
	return
}

func print_level(level int, t *BinaryNode) {
	if t == nil {
		return
	}
	if (t.LeftChild == nil) && (t.RightChild == nil) {
		for i := 0; i < level; i++ {
			fmt.Printf("__")
		}
		fmt.Printf("%2d\n", t.Element)
		return
	}
	print_level(level+1, t.RightChild)
	for i := 0; i < level; i++ {
		fmt.Printf("__")
	}
	fmt.Printf("%2d\n", t.Element)
	print_level(level+1, t.LeftChild)
}

func (t *BinaryNode) Print() {
	print_level(0, t)
}
