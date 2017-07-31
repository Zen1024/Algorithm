package sort

import (
// "fmt"
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
