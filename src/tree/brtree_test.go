package tree

import (
	"testing"
)

func TestRotate(t *testing.T) {
	root := &BrNode{
		Parent: nil,
		Elem:   6,
		Color:  ColorBlack,
		Left:   nil,
		Right:  nil,
	}
	x := &BrNode{
		Parent: root,
		Elem:   2,
		Color:  ColorRed,
		Left:   nil,
		Right:  nil,
	}
	root.Left = x
	lx := &BrNode{
		Parent: x,
		Elem:   1,
		Color:  ColorBlack,
		Left:   nil,
		Right:  nil,
	}
	x.Left = lx
	y := &BrNode{
		Parent: x,
		Elem:   4,
		Color:  ColorBlack,
		Left:   nil,
		Right:  nil,
	}
	x.Right = y
	ly := &BrNode{
		Parent: y,
		Elem:   3,
		Color:  ColorRed,
		Left:   nil,
		Right:  nil,
	}
	y.Left = ly
	ry := &BrNode{
		Parent: y,
		Elem:   5,
		Color:  ColorRed,
		Left:   nil,
		Right:  nil,
	}
	y.Right = ry
	root.Print()
	root.Left.LeftRotate()
	root.Print()
	root.Left.RightRotate()
	root.Print()
	root.Insert(10)
	root.Print()
	root.Insert(40)
	root.Print()
}
