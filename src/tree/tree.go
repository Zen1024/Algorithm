package tree

type Tree interface {
	GetRoot() interface{}
	Depth() int
}

type TreeNode interface {
	GetChild() []interface{}
	GetParent() interface{}
}
