package sort

import (
	"math/rand"
	"testing"
	"time"
)

var testArr []int

func Init() {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	testArr = make([]int, 10)
	for i := 0; i < 10; i++ {
		testArr[i] = rnd.Intn(1000)
		time.Sleep(time.Millisecond)
	}
}

func TestQuick(t *testing.T) {
	Init()
	t.Logf("before:%v", testArr[:10])
	re := QuickSort(testArr)
	t.Logf("after:%v", re[:10])
	if !isSort(re) {
		t.Fail()
	}
}

func TestHeap(t *testing.T) {
	Init()
	t.Logf("before:%v", testArr[:10])
	re := HeapSort(testArr)
	t.Logf("after:%v", re[:10])

	if !isSort(re) {
		t.Fail()
	}
}
