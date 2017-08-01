package sort

import (
	"fmt"
)

func isSort(in []int) bool {
	l := len(in)
	if l <= 1 {
		return true
	}
	for i := 0; i < l-1; i++ {
		if in[i] > in[i+1] {
			return false
		}
	}
	return true
}

func quicksort(in []int, l, h int) {
	if l >= h {
		return
	}
	pivot := in[l]
	i := l
	j := h

	for i != j {
		for (in[j] >= pivot) && (i < j) {
			j--
		}
		for (in[i] <= pivot) && (i < j) {
			i++
		}
		if i < j {
			in[i], in[j] = in[j], in[i]
		}
	}
	in[l], in[i] = in[i], in[l]
	quicksort(in, l, i-1)
	quicksort(in, i+1, h)
}

func QuickSort(in []int) []int {
	l := len(in)
	if l <= 1 {
		return in
	}

	tmp := make([]int, l)
	copy(tmp, in[0:])
	quicksort(tmp, 0, l-1)
	return tmp
}

func maxHeap(i, l int, in []int) []int {
	if i >= l {
		return in
	}
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if (left < l) && (in[left] >= in[max]) {
		max = left
	}
	if (right < l) && (in[right] >= in[max]) {
		max = right
	}
	in[i], in[max] = in[max], in[i]
	if i != max {
		maxHeap(max, l, in)
	}

	return in
}

func buildHeap(l int, in []int) {
	i := l / 2
	for ; i >= 0; i-- {
		maxHeap(i, l, in)
	}
}

func HeapSort(in []int) []int {

	l := len(in)
	l2 := l
	buildHeap(l, in)
	for i := l - 1; i >= 0; i-- {
		in[0], in[i] = in[i], in[0]
		l2--
		maxHeap(0, l2, in)
	}

	return in
}

func CountSort(in []int) []int {
	// find max and min
	l := len(in)
	if l == 0 {
		return in
	}

	max := in[0]
	min := in[0]
	for i := 0; i < l; i++ {
		if max < in[i] {
			max = in[i]
		}
		if min > in[i] {
			min = in[i]
		}
	}

	rng := max - min + 1
	tmp := make([]int, rng)
	for i := 0; i < rng; i++ {
		tmp[i] = 0
	}

	for i := 0; i < l; i++ {
		tmp[in[i]-min]++
	}

	re := make([]int, l)
	re_idx := 0
	for i := 0; i < rng; i++ {
		if tmp[i] > 0 {
			for j := 0; j < tmp[i]; j++ {
				re[re_idx] = i + min
				re_idx++
			}
		}
	}
	return re
}

func getNumByIdx(num, idx int) int {
	tmp := num
	for i := 1; i < idx; i++ {
		tmp /= 10
	}
	return tmp % 10
}

func basesortByIdx(in []int, idx int) []int {
	//use count sort
	l := len(in)
	cnt_arr := make([]int, 10)
	cnt_arr_2 := [10][]int{}
	for i := 0; i < 10; i++ {
		cnt_arr[i] = 0
		cnt_arr_2[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		n := getNumByIdx(in[i], idx)
		cnt_arr_2[n][cnt_arr[n]] = in[i]
		cnt_arr[n]++
	}

	j := 0
	re := make([]int, l)
	for i := 0; i < 10; i++ {
		for k := 0; k < cnt_arr[i]; k++ {
			if cnt_arr_2[i][k] > 0 {
				re[j] = cnt_arr_2[i][k]
				j++
			}
		}
	}
	return re
}

func BaseSort(in []int) []int {
	max := in[0]
	l := len(in)
	re := make([]int, l)
	for i := 0; i < l; i++ {
		if max < in[i] {
			max = in[i]
		}
		re[i] = in[i]
	}
	tmp := max
	j := 0
	for ; tmp != 0; j++ {
		tmp /= 10
	}

	for i := 1; i <= j; i++ {
		re = basesortByIdx(re, i)
	}
	return re
}

type Node struct {
	Elem int
	Next *Node
}

func newNode(ele int) *Node {
	return &Node{
		Elem: ele,
		Next: nil,
	}
}

func insertSorted(ele int, node *Node) {
	if node == nil {
		node = &Node{
			Elem: ele,
			Next: nil,
		}
		return
	}
	var tmp *Node = node
	inserted := false
	for tmp.Next != nil {
		if ele <= tmp.Elem {
			newNode := &Node{
				Elem: ele,
				Next: tmp,
			}
			node.Next = newNode
			inserted = true
			break
		}
	}
	if !inserted {
		node.Next = &Node{
			Elem: ele,
			Next: nil,
		}
	}
	return
}

func BucketSort(in []int) []int {
	l := len(in)
	if l <= 1 {
		return in
	}
	max := in[0]
	min := in[0]
	buckets := make([]*Node, l+1)
	for _, ele := range in {
		if max < ele {
			max = ele
		}
		if min > ele {
			min = ele
		}
	}

	width := (max - min) / l
	if width == 0 {
		return in
	}

	for _, ele := range in {
		idx := (ele - min) / width
		insertSorted(ele, buckets[idx])
	}

	re := make([]int, l)
	j := 0
	for i := 0; i <= l; i++ {
		tmp := buckets[i]
		for tmp != nil {
			re[j] = tmp.Elem
			tmp = tmp.Next
			j++
		}
	}
	return re
}
