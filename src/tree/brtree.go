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

//中序遍历,线索化遍历
func (t *BrNode) InOrder() []*BrNode {
	if t == nil {
		return nil
	}
	re := []*BrNode{}
	p := t
	for p != nil {
		if p.Left == nil {
			re = append(re, p)
			p = p.Right
		} else {
			pre := p.Left
			for pre.Right != nil && pre.Right != p {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = p
				p = p.Left
			} else {
				re = append(re, p)
				p = p.Right
				pre.Right = nil

			}
		}
	}
	return re
}

func (t *BrNode) copy() *BrNode {
	if t == nil {
		return nil
	}
	return &BrNode{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Elem:   t.Elem,
		Color:  t.Color,
	}
}

//@TODO
func (t *BrNode) Copy() *BrNode {
	if t == nil {
		return nil
	}
	re := t.copy()
	re.Left = t.Left.Copy()
	re.Right = t.Right.Copy()
	return re

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

func (t *BrNode) Print(args ...string) {
	if len(args) == 1 {
		fmt.Printf("------[%s]------", args[0])
	}
	fmt.Printf("\nprint b-r tree:\n")
	t.print_level(0)
	fmt.Printf("print end..\n\n")
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
	fmt.Printf("[%s:%d][%p] p[%d:%p],l:[%d:%p],r:[%d:%p],c[%s]\n", name, e, t, p.elem(), p, l.elem(), l, r.elem(), r, c.color())
}

func (c Color) color() string {
	if c == ColorBlack {
		return "b"
	}
	return "r"
}

//get leaves
func (t *BrNode) leaves() []*BrNode {
	var p *BrNode = t
	re := []*BrNode{}
	for p != nil {
		if p.Left == nil {
			re = append(re, p)
			p = p.Right
		} else {
			pre := p.Left
			for pre.Right != nil && pre.Right != p {
				pre = pre.Right
			}
			if pre.Right == nil {
				re = append(re, pre)
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

func (t *BrNode) Valid() bool {
	inseq := t.InOrder()
	li := len(inseq)
	for i := 0; i < li; i++ {
		node := inseq[i]
		if node.Color == ColorRed {
			if (node.Left != nil && node.Left.Color == ColorRed) || (node.Right != nil && node.Right.Color == ColorRed) {
				return false
			}
		}
		leaves := t.leaves()
		ll := len(leaves)
		if ll == 0 {
			return false
		}
		p := inseq[i]
		cnt := p.pathBlackCnt(leaves[0])
		for j := 0; j < ll; j++ {
			c := p.pathBlackCnt(leaves[j])
			if c != cnt {
				return false
			}
		}
	}

	return true

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
	z := elem
	p := z.Parent
	for p != nil && p.Color == ColorRed {
		var uncle *BrNode
		if p == p.Parent.Left {
			uncle = p.Parent.Right
		} else {
			uncle = p.Parent.Left
		}
		if uncle != nil {
			if uncle.Color == ColorRed {
				p.Parent.Color = ColorRed
				p.Color = ColorBlack
				uncle.Color = ColorBlack
				z = p.Parent
				p = z.Parent
				continue
			}
		}
		if p == p.Parent.Left {
			if z == z.Parent.Right {
				p.LeftRotate()
				z = z.Left
				p = z.Parent
				p.Parent.RightRotate()
				p.Color = ColorBlack
				p.Right.Color = ColorRed
				z = p.Left
				p = z.Parent
				continue
			} else {
				p.Parent.RightRotate()
				p.Color = ColorBlack
				p.Right.Color = ColorRed
				continue

			}
		} else if p == p.Parent.Right {
			if z == z.Parent.Left {
				p.RightRotate()
				z = z.Right
				p = z.Parent
				p.Parent.LeftRotate()
				p.Color = ColorBlack
				p.Left.Color = ColorRed
				z = p.Right
				p = z.Parent
				continue
			} else {
				p.Parent.LeftRotate()
				p.Color = ColorBlack
				p.Left.Color = ColorRed
				continue
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
	return re
}

//删除非双孩子的节点，返回节点孩子
func (t *BrNode) delNodeWith01Child(todel *BrNode) (parent, tofix *BrNode) {
	if todel == nil {
		return nil, nil
	}
	if todel.Parent == nil {
		return nil, nil
	}
	if todel.Left == nil && todel.Right == nil {
		if todel.Parent.Left == todel {
			todel.Parent.Left = nil
		} else {
			todel.Parent.Right = nil
		}
		p := todel.Parent
		todel.Parent = nil
		return p, nil
	}
	var newNode *BrNode
	if todel.Left != nil {
		newNode = todel.Left
		todel.Left = nil
	} else {
		newNode = todel.Right
		todel.Right = nil
	}
	newNode.Parent = todel.Parent
	if todel.Parent.Left == todel {
		todel.Parent.Left = newNode
	} else {
		todel.Parent.Right = newNode
	}
	p := todel.Parent
	todel.Parent = nil

	return p, newNode
}

func (t *BrNode) root() *BrNode {
	p := t
	for p.Parent != nil {
		p = p.Parent
	}
	return p
}

func (t *BrNode) delete_fix(tofix, p_tofix *BrNode) *BrNode {
	//被删除的节点为黑色，可推断出其兄弟节点必定不为空，否则违反性质根节点到任意叶子黑高度相同
	// tofix.print("orig-tofix")
	// p_tofix.print("orig-p_tofix")
	// fmt.Printf("\n\n")
	for p_tofix != nil && tofix != t && (tofix == nil || tofix.Color == ColorBlack) {
		// p_tofix.print("p")
		// tofix.print("t")
		// fmt.Printf("\n\n")
		bro := p_tofix.Left
		if tofix == p_tofix.Left {
			bro = p_tofix.Right
		}

		if tofix == p_tofix.Left {
			// fmt.Printf("cond-1")
			//cond -1 convert to other cond
			if bro.Color == ColorRed {
				//t.root().Print("cond-1.0")
				bro.Color = ColorBlack
				p_tofix.Color = ColorRed
				p_tofix.LeftRotate()
				//t.root().Print("cond-1.1")
				continue
			}

			if bro.Color == ColorBlack {
				//cond -2 bro 含有两个黑孩
				if (bro.Left == nil || bro.Left.Color == ColorBlack) && (bro.Right == nil || bro.Right.Color == ColorBlack) {
					//t.root().Print("cond-2.0")
					bro.Color = ColorRed
					tofix = p_tofix
					p_tofix = p_tofix.Parent
					//t.root().Print("cond-2.1")
					continue
				}
				//cond -3 bro-left-red,bro-right-black
				if bro.Left != nil && bro.Left.Color == ColorRed && (bro.Right == nil || bro.Right.Color == ColorBlack) {
					//t.root().Print("cond-3.0")
					bro.Left.Color = ColorBlack
					bro.Color = ColorRed
					bro.RightRotate()
					//t.root().Print("cond-3.1")
					continue
				}

				if bro.Right != nil && bro.Right.Color == ColorRed {
					//t.root().Print("cond-4.0")
					bro.Color = p_tofix.Color
					p_tofix.Color = ColorBlack
					bro.Right.Color = ColorBlack
					p_tofix.LeftRotate()
					tofix = t
					p_tofix = nil
					//t.root().Print("cond-4.1")
					break
				}

			}
		} else {
			// fmt.Printf("cond-2\n")
			if bro.Color == ColorRed {
				//t.root().Print("cond-5.0")
				bro.Color = ColorBlack
				p_tofix.Color = ColorRed
				p_tofix.RightRotate()
				//t.root().Print("cond-5.1")
				continue
			}

			if bro.Color == ColorBlack {
				//cond -2 bro 含有两个黑孩
				if (bro.Left == nil || bro.Left.Color == ColorBlack) && (bro.Right == nil || bro.Right.Color == ColorBlack) {
					//t.root().Print("cond-6.0")
					bro.Color = ColorRed
					tofix = p_tofix
					p_tofix = p_tofix.Parent
					//t.root().Print("cond-6.1")
					continue
				}
				//cond -3 bro-left-red,bro-right-black
				if bro.Left == nil || bro.Left.Color == ColorBlack && (bro.Right == nil || bro.Right.Color == ColorRed) {
					//t.root().Print("cond-7.0")
					bro.Right.Color = ColorBlack
					bro.Color = ColorRed
					bro.LeftRotate()
					//t.root().Print("cond-7.1")
					continue
				}

				if bro.Left != nil && bro.Left.Color == ColorRed {
					//t.root().Print("cond-8.0")
					bro.Color = p_tofix.Color
					p_tofix.Color = ColorBlack
					bro.Left.Color = ColorBlack
					p_tofix.RightRotate()
					//t.root().Print("cond-8.1")
					tofix = t
					p_tofix = nil
					break
				}

			}
		}
	}
	if tofix != nil {
		tofix.Color = ColorBlack
	}

	return t.root()
}

func (t *BrNode) Delete(elem int, s string) *BrNode {
	// fmt.Printf("deleting %d.....\n", elem)
	todel := t.find(elem)
	if todel == nil {
		return t
	}
	// todel.print("todel")
	var p_tofix, tofix *BrNode
	delColor := todel.Color
	if todel.Left == nil || todel.Right == nil {
		// fmt.Printf("cond-1\n")
		p_tofix, tofix = t.delNodeWith01Child(todel)
	} else {
		// fmt.Printf("cond-2\n")
		x := todel
		suc := x.successor()
		x.Elem = suc.Elem
		p_tofix, tofix = t.delNodeWith01Child(suc)
		// suc.print("suc")
		if suc.Color == ColorRed {
			return t
		} else {
			//t.root().Print("orig-2")
			t = t.delete_fix(tofix, p_tofix)
			return t
		}
	}

	if delColor == ColorRed {
		return t
	}
	//t.root().Print("orig-2")
	t = t.delete_fix(tofix, p_tofix)

	return t
}
