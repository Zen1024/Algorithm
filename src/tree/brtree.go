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
	} else {

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
func (t *BrNode) plrec() (p *BrNode, l *BrNode, r *BrNode, e int, c Color) {
	if t == nil {
		return nil, nil, nil, -1, ColorBlack
	}
	return t.Parent, t.Left, t.Right, t.Elem, t.Color

}

func (t *BrNode) elem() int {
	if t == nil {
		return -1
	}
	return t.Elem
}
func (t *BrNode) print(name string) {
	p, l, r, e, c := t.plrec()
	fmt.Printf("[%s-%d][%p] p[%d:%p],l:[%d:%p],r:[%d:%p],c[%s]\n", name, e, t, p.elem(), p, l.elem(), l, r.elem(), r, c.color())
}

func (c Color) color() string {
	if c == ColorBlack {
		return "b"
	}
	return "r"
}

func (t *BrNode) Valid() bool {
	leaves := t.leaves()
	ll := len(leaves)
	if ll == 0 {
		return false
	}
	cnt := t.pathBlackCnt(leaves[0])

	for i := 0; i < ll; i++ {
		c := t.pathBlackCnt(leaves[i])
		if c != cnt {
			return false
		}
	}
	return true

}

//get leaves
func (t *BrNode) leaves() []*BrNode {
	var p *BrNode = t
	re := []*BrNode{}
	for p != nil {
		if p.Left == nil {
			if p.Right == nil {
				re = append(re, p)
			}
			p = p.Right
		} else {
			pre := p.Left
			for pre.Right != nil && pre.Right != p {
				pre = pre.Right
			}
			if pre.Right == nil {
				if pre.Left == nil {
					re = append(re, pre)
				}
				pre.Right = p
				p = p.Left
			} else {
				pre.Right = nil
				p = p.Right
			}
		}
	}
	return re
}

func (t *BrNode) pathBlackCnt(leaf *BrNode) int {
	if t.Parent != nil || leaf == nil {
		return -1
	}
	var p *BrNode = t
	re := 0

	for p != nil {
		if p.Color == ColorBlack {
			re++
		}
		if p.Elem == leaf.Elem {
			return re
		}
		if p.Elem < leaf.Elem {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	return -1

}

func (t *BrNode) min() *BrNode {
	if t == nil {
		return nil
	}
	if t.Left == nil {
		return t
	}
	p := t

	for p.Left != nil {
		p = p.Left
	}
	return p
}

func (t *BrNode) max() *BrNode {
	if t == nil {
		return nil
	}
	if t.Right == nil {
		return t
	}
	p := t
	for p.Right != nil {
		p = p.Right
	}
	return p
}

//前驱
func (t *BrNode) predecessor() *BrNode {
	if t == nil {
		return nil
	}
	if t.Left != nil {
		return t.Left.max()
	}
	var p, x *BrNode
	p = t.Parent
	x = t
	if p == nil {
		return nil
	}
	for p.Parent != nil && x == p.Left {
		x = p
		p = p.Parent
	}
	return p

}

//后继
func (t *BrNode) successor() *BrNode {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		return t.Right.min()
	}
	var p, x *BrNode
	p = t.Parent
	x = t
	if p == nil {
		return nil
	}
	for p.Parent != nil && x == p.Right {
		x = p
		p = p.Parent
	}
	return p

}

func (t *BrNode) find(elem int) *BrNode {
	var p *BrNode = t
	if p.elem() == elem {
		return p
	}
	for p != nil {
		if p.Elem == elem {
			return p
		}
		if p.Elem > elem {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}
func (t *BrNode) insert_fix(elem *BrNode) *BrNode {
	if elem == nil {
		return t
	}
	if elem.Parent == nil {
		elem.Color = ColorBlack
		return t
	}
	// fmt.Printf("fixing %d...\n", elem.Elem)
	z := elem
	for p := z.Parent; p.Color == ColorRed; {
		var uncle *BrNode
		if p == p.Parent.Left {
			uncle = p.Parent.Right
		} else {
			uncle = p.Parent.Left
		}
		// p.print("p")
		// z.print("z")
		if uncle != nil {
			//cond -1
			// uncle.print("uncle")
			if uncle.Color == ColorRed {
				p.Parent.Color = ColorRed
				p.Color = ColorBlack
				uncle.Color = ColorBlack
				z = p.Parent.Parent
				continue
			}
		}
		// fmt.Printf("[init]tree:----------------\n")
		// t.Print()
		// fmt.Printf("[init]---------------------\n")
		//uncle ==nil or uncle-color ==black //cond-2
		if p == p.Parent.Left {
			if z == z.Parent.Right {
				p.LeftRotate()
				// fmt.Printf("[cond-1]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-1]---------------------\n")
				z.Parent.RightRotate()
				// fmt.Printf("[cond-1.1]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-1.1]---------------------\n")
				z.Color = ColorBlack
				z.Right.Color = ColorRed
			} else {
				p.Parent.RightRotate()
				// fmt.Printf("[cond-1.2]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-1.2]---------------------\n")
				p.Color = ColorBlack
				p.Right.Color = ColorRed

			}
		} else if p == p.Parent.Right {
			if z == z.Parent.Left {
				p.RightRotate()
				// fmt.Printf("[cond-2]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-2]---------------------\n")
				// p.print("p")
				// z.print("z")
				z.Parent.LeftRotate()
				// p.print("p")
				// z.print("z")
				// fmt.Printf("[cond-2.1]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-2.1]---------------------\n")
				z.Color = ColorBlack
				z.Left.Color = ColorRed
			} else {
				p.Parent.LeftRotate()
				// fmt.Printf("[cond-2.2]tree:----------------\n")
				// t.Print()
				// fmt.Printf("[cond-2.2]---------------------\n")
				p.Color = ColorBlack
				p.Left.Color = ColorRed
			}
		}
	}

	// fmt.Printf("fix %d end\n", elem.Elem)
	for t.Parent != nil {
		t = t.Parent
	}
	t.Color = ColorBlack
	return t
}

func (t *BrNode) Insert(elem int) *BrNode {
	if t == nil {
		return nil
	}
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
			return t
		}
	}
	new_node.Parent = ele_parent
	if ele_parent.Elem > elem {
		ele_parent.Left = new_node
	} else {
		ele_parent.Right = new_node
	}
	re := t.insert_fix(new_node)
	//return root
	return re
}

func (t *BrNode) delNodeWith01Child(todel *BrNode) *BrNode {
	if todel == nil {
		return t
	}
	if todel.Parent == nil {
		return nil
	}
	if todel.Left == nil && todel.Right == nil {
		if todel.Parent.Left == todel {
			todel.Parent.Left = nil
		} else {
			todel.Parent.Right = nil
		}
		todel.Parent = nil
		return t
	}
	var new *BrNode
	if todel.Left != nil {
		new = todel.Left
		todel.Left = nil
	} else {
		new = todel.Right
		todel.Right = nil
	}
	new.Parent = todel.Parent
	if todel.Parent.Left == todel {
		todel.Parent.Left = new
	} else {
		todel.Parent.Right = new
	}
	todel.Parent = nil
	return t
}

func (t *BrNode) delete_fix(elem *BrNode) *BrNode { return t }
func (t *BrNode) Delete(elem int) *BrNode {
	todel := t.find(elem)
	if todel == nil {
		return t
	}

	if !(todel.Left != nil && todel.Right != nil) {
		t.delNodeWith01Child(todel)
	} else {

		x := t
		suc := x.successor()

		x.Elem = suc.Elem
		t.delNodeWith01Child(suc)
	}

	return t
}
