package utils

import (
	"fmt"
)

type HFNode struct {
	Left   *HFNode
	Right  *HFNode
	Parent *HFNode
	Key    string
	Value  int
}

//优先级队列
//@TODO 用最小堆实现

type PNode struct {
	Key    string
	Value  int
	Next   *PNode
	Prev   *PNode
	hfNode *HFNode //辅助指针
}

type PQueue struct {
	Front *PNode
	Tail  *PNode
}

func (p *PQueue) pop() *PNode {
	if p == nil {
		return nil
	}
	re := p.Front
	p.Front = re.Next
	if re.Next != nil {
		re.Next.Prev = nil
	}
	re.Next = nil
	return re
}

func (p *PQueue) Pop() (k string, v int) {
	node := p.pop()
	if node == nil {
		return "", -1
	}

	return node.Key, node.Value
}

func (p *PQueue) print() {
	for tmp := p.Front; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%s:%d ", tmp.Key, tmp.Value)
	}
	fmt.Printf("\n")
}

func (p *PQueue) isEmpty() bool {
	return p.Front == nil
}

func (p *PQueue) Push(k string, v int) *PQueue {
	return p.push(&PNode{
		Key:   k,
		Value: v,
	})
}

func (p *PQueue) push(node *PNode) *PQueue {

	if node == nil {
		return p
	}
	v := node.Value

	if p == nil {
		return &PQueue{
			Front: node,
			Tail:  node,
		}
	}
	var q *PNode
	for q = p.Front; q != nil && q.Value < v; q = q.Next {
	}

	if q == nil {
		node.Prev = p.Tail
		p.Tail.Next = node
		p.Tail = node
		return p
	}
	prev := q.Prev
	if prev == nil {
		node.Next = p.Front
		p.Front.Prev = node
		p.Front = node
		return p
	}
	node.Next = q
	q.Prev = node
	prev.Next = node

	return p
}

func newHfTree(keys map[string]int) *HFNode {
	var queue *PQueue = nil
	var re *HFNode

	for k, v := range keys {
		queue = queue.Push(k, v)
	}
	for !queue.isEmpty() {
		// queue.print()
		nl := queue.pop()
		if queue.isEmpty() {
			re = nl.hfNode
			break
		}
		nr := queue.pop()
		kl, vl := nl.Key, nl.Value
		kr, vr := nr.Key, nr.Value

		left := &HFNode{
			Left:  nil,
			Right: nil,
			Key:   kl,
			Value: vl,
		}
		right := &HFNode{
			Left:  nil,
			Right: nil,
			Key:   kr,
			Value: vr,
		}
		if nl.hfNode != nil {
			left = nl.hfNode
		}
		if nr.hfNode != nil {
			right = nr.hfNode
		}

		p := &HFNode{
			Left:   left,
			Right:  right,
			Key:    "",
			Value:  vl + vr,
			Parent: nil,
		}
		left.Parent = p
		right.Parent = p
		// fmt.Printf("\n\n\n")
		// p.print(0)
		queueNode := &PNode{
			Key:    "",
			Value:  vl + vr,
			hfNode: p,
		}
		re = p
		queue.push(queueNode)
	}
	return re
}

func (t *HFNode) leftPrint() (k, s string) {
	p := t
	str := ""
	for p != nil {
		if p.Left != nil {
			str += "0"
			p = p.Left
		} else {
			if p.Right != nil {
				str += "1"
				p = p.Right
			} else {
				k = p.Key
				if p.Parent.Left == p {
					p.Parent.Left = nil
					p.Parent = nil
				} else {
					p.Parent.Right = nil
					p.Parent = nil
				}
				break
			}
		}
	}
	return k, str
}

func (t *HFNode) print(level int) {
	if t == nil {
		return
	}
	t.Right.print(level + 1)
	if t.Left == nil && t.Right == nil {
		for i := 0; i < level; i++ {
			fmt.Printf("----")
		}
		fmt.Printf("%s:%d\n", t.Key, t.Value)
		return
	}
	for i := 0; i < level; i++ {
		fmt.Printf("----")
	}
	fmt.Printf("%s:%d\n", t.Key, t.Value)
	t.Left.print(level + 1)
}

func HFEncode(kv map[string]int) map[string]string {
	tree := newHfTree(kv)
	l := len(kv)
	re := map[string]string{}
	for i := 0; i < l; i++ {
		k, v := tree.leftPrint()
		if k == "" {
			i--
			continue
		}
		re[k] = v
	}
	return re

}
