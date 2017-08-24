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
	if elem == nil || q == nil {
		return
	}

	if q.front == nil {
		q.front = &node{
			prev: nil,
			next: nil,
			elem: elem,
		}
		q.end = q.front
	} else {
		nd := &node{
			prev: q.end,
			next: nil,
			elem: elem,
		}
		q.end.next = nd
		q.end = nd
	}
}

func (q *queue) Front() *BinaryNode {
	if q == nil || q.front == nil {
		return nil
	}

	tmp := q.front

	re := tmp.elem
	q.front = tmp.next
	if tmp.next != nil {
		tmp.next.prev = nil
	}

	tmp.next = nil

	return re
}

func (q *queue) Print() {
	fmt.Printf("queue:")
	for p := q.front; p != nil; p = p.next {
		fmt.Printf("%4d", p.elem.Element)
	}
	fmt.Printf("\n")
}

func (q *queue) IsEmpty() bool {
	return q.front == nil
}

func (s *stack) Push(ele *BinaryNode) {
	s.top++
	p := ele
	s.elem = append(s.elem, p)
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

func NewTree(elems []int) *BinaryNode {
	l := len(elems)
	if l == 0 {
		return nil
	}
	var t *BinaryNode = &BinaryNode{
		Element:    elems[0],
		LeftChild:  nil,
		RightChild: nil,
	}
	for i := 1; i < l; i++ {
		t.Insert(elems[i])
	}
	return t
}

func (t *BinaryNode) Insert(ele int) {
	if t == nil {
		t = &BinaryNode{
			Element:    ele,
			LeftChild:  nil,
			RightChild: nil,
		}
		return
	}
	if t.Element == ele {
		return
	}
	if t.Element > ele {
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

func (t *BinaryNode) PreOrderNonRecAndStack() {

	fmt.Printf("\n")
	var p *BinaryNode = t
	for p != nil {
		if p.LeftChild == nil {
			fmt.Printf("%4d ", p.Element)
			p = p.RightChild
		} else {
			pre := p.LeftChild
			//找到前驱
			for pre.RightChild != nil && pre.RightChild != p {
				pre = pre.RightChild
			}
			if pre.RightChild == nil {
				fmt.Printf("%4d ", p.Element)
				pre.RightChild = p
				p = p.LeftChild
			} else {
				pre.RightChild = nil
				p = p.RightChild
			}
		}

	}
}

//中序遍历 -非递归
func (t *BinaryNode) InOrderNonRecAndStack() {
	fmt.Printf("\n")
	var p *BinaryNode = t
	for p != nil {
		if p.LeftChild == nil {
			fmt.Printf("%4d ", p.Element)
			p = p.RightChild
		} else {
			pre := p.LeftChild
			for pre.RightChild != nil && pre.RightChild != p {
				pre = pre.RightChild
			}
			if pre.RightChild == nil {
				pre.RightChild = p
				p = p.LeftChild
			} else {
				fmt.Printf("%4d ", p.Element)
				pre.RightChild = nil
				p = p.RightChild
			}

		}
	}
}

//降序 from-->to
func reversePrint(from, to *BinaryNode) {
	if from == nil || to == nil {
		return
	}

	if from == to {
		fmt.Printf("%4d ", from.Element)
		return
	}
	if from.Element < to.Element {
		reversePrint(from.RightChild, to)
	} else {
		reversePrint(from.LeftChild, to)
	}

	fmt.Printf("%4d", from.Element)
}

func (t *BinaryNode) PostOrderNonRecAndStack() {
	fmt.Printf("\npost order:\n")
	var dummy *BinaryNode = &BinaryNode{
		Element:    -1,
		LeftChild:  t,
		RightChild: nil,
	}
	var p *BinaryNode = dummy
	for p != nil {
		if p.LeftChild == nil {
			p = p.RightChild
		} else {
			pre := p.LeftChild
			for pre.RightChild != nil && pre.RightChild != p {
				pre = pre.RightChild
			}
			if pre.RightChild == nil {
				pre.RightChild = p
				p = p.LeftChild
			} else {
				pre.RightChild = nil
				reversePrint(p.LeftChild, pre)
				p = p.RightChild
			}
		}
	}
	fmt.Printf("\n")
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

func (t *BinaryNode) LevelOrder() {
	fmt.Printf("\n")
	q := newQueue()
	q.Push(t)
	for !q.IsEmpty() {
		p := q.Front()
		if p == nil {
			break
		}
		fmt.Printf("%4d ", p.Element)
		q.Push(p.LeftChild)
		q.Push(p.RightChild)
	}
	fmt.Printf("\n")
}

func (t *BinaryNode) NodeCnt() int {
	if t == nil {
		return 0
	}
	return t.LeftChild.NodeCnt() + t.RightChild.NodeCnt() + 1
}

func (t *BinaryNode) LeafCnt() int {
	if t == nil {
		return 0
	}
	if t.LeftChild == nil && t.RightChild == nil {
		return 1
	}
	return t.LeftChild.LeafCnt() + t.RightChild.LeafCnt()

}

func (t *BinaryNode) Depth() int {
	if t == nil {
		return 0
	}
	l := t.LeftChild.Depth()
	r := t.RightChild.Depth()
	max := l
	if max < r {
		max = r
	}
	return max + 1
}

func (t *BinaryNode) Equals(t2 *BinaryNode) bool {
	if t == nil && t2 == nil {
		return true
	} else {
		if t == nil || t2 == nil {
			return false
		}
	}
	if t.Element != t2.Element {
		return false
	} else {
		return t.LeftChild.Equals(t2.LeftChild) && t.RightChild.Equals(t2.RightChild)
	}
}

func (t *BinaryNode) Mirror() {
	if t == nil {
		return
	}
	t.LeftChild, t.RightChild = t.RightChild, t.LeftChild
	t.LeftChild.Mirror()
	t.RightChild.Mirror()
}

//最近公共祖先
func (t *BinaryNode) LCA(t1, t2 int) *BinaryNode {
	node1, node2 := t.find(t1), t.find(t2)
	if node1 == nil || node2 == nil {
		return nil
	}

	return t.lca(node1, node2)
}
func (t *BinaryNode) lca(t1, t2 *BinaryNode) *BinaryNode {
	if t == nil {
		return nil
	}

	if t.Element == t1.Element || t.Element == t2.Element {
		return t
	}
	l := t.LeftChild.lca(t1, t2)
	r := t.RightChild.lca(t1, t2)
	if l != nil && r != nil {
		return t
	}
	if l != nil {
		return l
	}
	return r
}

func (t *BinaryNode) find(ele int) *BinaryNode {
	if t == nil {
		return nil
	}
	if t.Element == ele {
		return t
	}
	if ele < t.Element {
		return t.LeftChild.find(ele)
	}
	return t.RightChild.find(ele)
}

func (t *BinaryNode) level(d *BinaryNode) int {
	if t == nil || d == nil {
		return -1
	}
	if t.Element == d.Element {
		return 0
	}
	if t == nil {
		return -1
	}
	level := t.LeftChild.level(d)
	if level == -1 {
		level = t.RightChild.level(d)
	}
	if level == -1 {
		return -1
	}
	return level + 1
}

func (t *BinaryNode) Dist(t1, t2 int) int {
	node1, node2 := t.find(t1), t.find(t2)

	if t1 == t2 {
		return 0
	}

	p := t.LCA(t1, t2)
	if p == nil {
		return -1
	}
	l1 := p.level(node1)
	l2 := p.level(node2)
	if l1 == -1 || l2 == -1 {
		return -1
	}
	return l1 + l2
}

//完全二叉树,根据后序遍历 的原理
func (t *BinaryNode) IsCbt() bool {
	q := newQueue()
	q.Push(t)

	for tmp := q.Front(); tmp != nil; {
		l := tmp.LeftChild
		r := tmp.RightChild
		if l == nil {
			l = &BinaryNode{
				Element:    -1,
				LeftChild:  nil,
				RightChild: nil,
			}
		}
		if r == nil {
			r = &BinaryNode{
				Element:    -1,
				LeftChild:  nil,
				RightChild: nil,
			}
		}
		q.Push(l)
		q.Push(r)
	}

	for !q.IsEmpty() {
		tmp := q.Front()
		if tmp.Element != -1 {
			return false
		}
	}
	return true
}
