package tree

import (
	"fmt"
)

type Color int

const (
	ColorBlack Color = 0
	ColorRed   Color = 1
)

type BrNode struct {
	Color  Color
	Left   *BrNode
	Right  *BrNode
	Parent *BrNode
	Elem   int
}

func (t *BrNode) LeftRotate() {
	if t == nil || (t.Right == nil) {
		return
	}
	x := t
	y := t.Right
	//y上移
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

func (t *BrNode) RightRotate() {
	if t == nil || (t.Left == nil) {
		return
	}
	x := t.Left
	y := t

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

func (t *BrNode) print_level(level int) {
	if t == nil {
		return
	}
	if t.Left == nil && t.Right == nil {
		for i := 0; i < level; i++ {
			fmt.Printf("---")
		}
		color := "b"
		if t.Color == ColorRed {
			color = "r"
		}
		fmt.Printf("%s%2d\n", color, t.Elem)
		return
	}
	t.Right.print_level(level + 1)
	for i := 0; i < level; i++ {
		fmt.Printf("---")
	}
	color := "b"
	if t.Color == ColorRed {
		color = "r"
	}
	fmt.Printf("%s%2d\n", color, t.Elem)
	t.Left.print_level(level + 1)
	return

}

func (t *BrNode) Print() {
	fmt.Printf("print b-r tree:\n")
	t.print_level(0)
	fmt.Printf("print end..\n")
}

func (t *BrNode) get(elem int) *BrNode {
	if t == nil {
		return nil
	}

	p := t
	re := p
	for p != nil {
		re = p
		if p.Elem == elem {
			return p
		}
		if p.Elem < elem {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return re
}

func (t *BrNode) insert_fix(elem *BrNode) {
	p := elem
	if t == elem {
		return
	}
	if p.Parent.Parent == nil {
		return
	}

	for p.Parent.Color == ColorRed {
		if p.Parent.Parent == nil {
			return
		}
		var uncle *BrNode
		if p.Parent == p.Parent.Parent.Left {
			uncle = p.Parent.Parent.Right
		} else {
			uncle = p.Parent.Parent.Left
		}
		if uncle == nil {
			return
		}
		if uncle.Color == ColorRed {
			uncle.Color = ColorBlack
			p.Parent.Color = ColorBlack
			p.Parent.Parent.Color = ColorRed
		}

		if p.Parent == p.Parent.Parent.Left {
			if uncle.Color == ColorBlack {
				if p.Parent.Right == p {
					p.Parent.LeftRotate()
				}
				p.Parent.Color = ColorBlack
				p.Parent.Parent.Color = ColorRed
				p.Parent.Parent.RightRotate()
				p = p.Parent
			}
		} else {
			if uncle.Color == ColorBlack {
				if p.Parent.Right == p {
					p.Parent.LeftRotate()
				}
				p.Parent.Color = ColorBlack
				if p.Parent.Parent == nil {
					fmt.Printf("p:%4d", p.Elem)
				}
				p.Parent.Parent.Color = ColorRed
				p.Parent.Parent.LeftRotate()
				p = p.Parent
			}
		}
	}
}
func (t *BrNode) Insert(elem int) {
	p := t
	new_node := &BrNode{
		Left:  nil,
		Right: nil,
		Elem:  elem,
		Color: ColorRed,
	}
	var ele_parent *BrNode
	for p != nil {
		ele_parent = p
		if p.Elem < elem {
			p = p.Right
		} else if p.Elem > elem {
			p = p.Left
		} else {
			return
		}
	}
	new_node.Parent = ele_parent
	if ele_parent.Elem > elem {
		ele_parent.Left = new_node
	} else {
		ele_parent.Right = new_node
	}
	fmt.Printf("before:%v\n", new_node.Parent)
	t.insert_fix(new_node)
	return
}

func (t *BrNode) delete_fix(elem int) {}
func (t *BrNode) Delete(elem int)     {}
