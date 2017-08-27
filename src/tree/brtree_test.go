package tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randArr(l, max int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	re := make([]int, l)
	for i := 0; i < l; i++ {
		re[i] = rnd.Intn(max)
	}
	return re
}
func timetooks(begin int64) string {
	delta := time.Now().UnixNano() - begin
	s := delta / int64(1e9)
	ms := (delta % 1e9) / int64(1e6)
	us := (delta % 1e6) / int64(1e3)
	ns := (delta % 1e3) / int64(1)
	if s > int64(0) {
		return fmt.Sprintf("%d.%d.%d.%ds", s, ms, us, ns)
	}
	if ms > int64(0) {
		return fmt.Sprintf("%d.%d.%dms", ms, us, ns)
	}
	if us > int64(0) {
		return fmt.Sprintf("%d.%dus", us, ns)
	}
	if ns > int64(0) {
		return fmt.Sprintf("%dns", ns)
	}
	return ""
}

// func TestRotate(t *testing.T) {
// 	root := &BrNode{
// 		Parent: nil,
// 		Elem:   6,
// 		Color:  ColorBlack,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	x := &BrNode{
// 		Parent: root,
// 		Elem:   2,
// 		Color:  ColorRed,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	root.Left = x
// 	lx := &BrNode{
// 		Parent: x,
// 		Elem:   1,
// 		Color:  ColorBlack,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	x.Left = lx
// 	y := &BrNode{
// 		Parent: x,
// 		Elem:   4,
// 		Color:  ColorBlack,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	x.Right = y
// 	ly := &BrNode{
// 		Parent: y,
// 		Elem:   3,
// 		Color:  ColorRed,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	y.Left = ly
// 	ry := &BrNode{
// 		Parent: y,
// 		Elem:   5,
// 		Color:  ColorRed,
// 		Left:   nil,
// 		Right:  nil,
// 	}
// 	y.Right = ry
// 	root.Print()
// 	root.Left.LeftRotate()
// 	root.Print()
// 	root.Left.RightRotate()
// 	root.Print()
// }

// func TestInsert(t *testing.T) {
// 	max := 10000
// 	inValidCnt := 0
// 	for i := 0; i < 1000; i++ {
// 		arr := randArr(1000, max)
// 		var tree *BrNode = &BrNode{
// 			Parent: nil,
// 			Left:   nil,
// 			Right:  nil,
// 			Elem:   arr[0],
// 			Color:  ColorBlack,
// 		}
// 		la := len(arr)
// 		for i := 1; i < la; i++ {
// 			tree = tree.Insert(arr[i])
// 		}
// 		if !tree.Valid() {
// 			inValidCnt++
// 		}

// 	}
// 	t.Logf("[insert]invalid cnt:%d\n", inValidCnt)
// }

func TestDelete(t *testing.T) {
	arr := randArr(10, 10000)
	la := len(arr)
	tree := &BrNode{
		Elem:   arr[0],
		Color:  ColorBlack,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
	for i := 1; i < la; i++ {
		tree = tree.Insert(arr[i])
	}
	tree.Print()
	cp := tree.Copy()
	cp.Print()
	invalidCnt := 0

	for i := 0; i < la; i++ {
		tree = tree.Delete(arr[i])
		if !tree.Valid() {
			invalidCnt++
			// tree.Print()
		}
	}
	t.Logf("[delete]invalid cnt:%d\n", invalidCnt)
}
