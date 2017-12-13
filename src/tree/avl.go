package tree

import (
	"fmt"
)

type AvlNode struct {
	Element int
	Left    *AvlNode
	Right   *AvlNode
	Parent  *AvlNode
}

func (root *AvlNode) GetH() int {
	if root == nil {
		return 0
	}
	lh := root.Left.geth() + 1
	rh := root.Right.geth() + 1
	if lh > rh {
		return lh
	}
	return rh
}

func (root *AvlNode) Print() int {
	root.printLevel(0)
}

func (root *AvlNode) printLevel(level int) {
	if root == nil {
		return
	}
	if root.Right == nil {
		for i := 0; i < level; i++ {
			fmt.Printf("----")
		}
		fmt.Printf("%4d\n", root.Element)
		return
	}
	root.Right.printLevel(level + 1)
	for i := 0; i < level; i++ {
		fmt.Printf("----")
	}
	fmt.Printf("%4d\n", root.Element)
	root.Left.printLevel(level + 1)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func (node *AvlNode) isvalid() bool {
	if node == nil {
		return true
	}
	if abs(node.Left.GetH(), node.Right.GetH()) > 1 {
		return false
	}
	return true
}

//线索化，遍历
func (root *AvlNode) IsValid() bool {
	var p *AvlNode = root

	for p != nil {
		if p.Left == nil {
			if !p.isvalid() {
				return false
			}
			p = p.Right
		} else {
			pre := p.Left
			for pre.Right != nil && pre.Right != p {
				pre = pre.Right
			}
			if pre.Right == nil {
				if !p.isvalid() {
					return false
				}
				pre.Right = p
				p = p.Left
			} else {
				pre.Right = nil
				p = p.Right
			}
		}

	}
}

func (root *AvlNode) LeftRotate() {
	if root == nil || (root.Right == nil) {
		return
	}
	x := root
	y := root.Right

	y.Parent = x.Parent
	y.Parent = x.Parent
	if x.Parent != nil {
		if x.Parent.Left == x {
			x.Parent.Left = y
		} else {
			x.Parent.Right = y
		}
	}

	//x下移
	x.Parent = y

	//y左->x右
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Left = x
}

func (root *AvlNode) RightRotate() {
	if root == nil || (root.Right == nil) {
		return
	}
	x := root.Left
	y := root
	//x上移
	x.Parent = y.Parent
	if y.Parent != nil {
		if y.Parent.Left == y {
			y.Parent.Left = x
		} else {
			y.Parent.Right = x
		}
	}

	//y下移
	y.Parent = x

	//x右-->y左
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Right = y
}

func (root *AvlNode) Insert(elem int) *AvlNode {

}

func (root *AvlNode) Delete(elem int) *AvlNode {

}
