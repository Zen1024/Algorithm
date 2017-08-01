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

type SortFnc func(in []int) []int

var allFunc map[string]SortFnc = map[string]SortFnc{
	"QuickSort":  QuickSort,
	"HeapSort":   HeapSort,
	"CountSort":  CountSort,
	"BaseSort":   BaseSort,
	"BucketSort": BucketSort,
}

func TestAll(t *testing.T) {
	for name, fnc := range allFunc {
		Init()
		re := fnc(testArr)
		if !isSort(re) {
			t.Fatalf("func:%s failed!", name)
		}
	}
}
