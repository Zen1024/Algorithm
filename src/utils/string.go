package utils

// import (
// 	"fmt"
// )

func kmpContains(str, sub string) bool {
	rst := []rune(str)
	rsub := []rune(sub)
	if len(rst) < len(rsub) {
		return false
	}
	lt := len(rst)
	lsub := len(rsub)

	preMap := preFunc(sub)
	q := 0
	for i := 0; i < lt; i++ {
		for q > -1 && rsub[q] != rst[i] {
			q = preMap[q]
		}
		if rsub[q] == rst[i] {
			q++
		}
		if q == lsub-1 {
			return true
		}
		q = preMap[q]
	}

	return false
}

//前缀函数，计算真后缀的最长前缀的长度
// abcabc -> 0:0,1:0,2:0,3:1,4:2,5:3

func preFunc(str string) map[int]int {
	rs := []rune(str)
	l := len(rs)
	k := 0
	re := map[int]int{}
	re[1] = 0

	for i := 1; i < l; i++ {
		for k > 0 && rs[k] != rs[i] {
			k = re[k]
		}
		if rs[k] == rs[i] {
			k++
		}

		re[i] = k
	}
	delete(re, 0)
	return re
}
