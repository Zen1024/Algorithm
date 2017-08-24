package tree

type AvlNode struct {
	Element int
	Left    *AvlNode
	Right   *AvlNode
	Parent  *AvlNode
}
