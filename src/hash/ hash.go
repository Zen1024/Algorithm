package hash

import (
	"math/rand"
)

//散列函数

const (
	goldFactor = 0.6180339887
)

type hFunc func(key, table_len int) int
type node struct {
	key  int
	elem int
	next *node
}

//链表解决冲突
type ch_htable struct {
	nodes   []*node
	hashFnc hFunc
	size    int
}

func divFunc(key, tl int) int {
	//h(k) = k mod m
	return key % tl
}

func mulFunc(key, tl int) int {
	var kg float64 = float64(key) * goldFactor
	//h(k) = m(((k*A)mod1))
	return (int)((float64)(tl) * (kg - (float64)((int64)(kg))))
}

//探查函数

//线性探查

func linerProb(key, i, l int, fnc hFunc) int {
	return fnc(key, l) + i
}

//二次探查
func quadProb(key, i, l, c1, c2 int, fnc hFunc) int {
	return fnc(key, l) + c1*i + c2*i*i
}

//双重散列
func doubleProb(fnc1, fnc2 hFunc, key, i, l int) int {
	return fnc1(key, l) + i*fnc2(key, l)
}

//链表法

func newcht(m int, fnc hFunc) *ch_htable {
	return &ch_htable{
		nodes:   make([]*node, m),
		hashFnc: fnc,
		size:    m,
	}
}

func (t *ch_htable) insert(key int) int {
	slot := t.hashFnc(key, t.size)
	if t.nodes[slot] == nil {
		t.nodes[slot] = &node{
			key:  key,
			next: nil,
		}
		return slot
	} else {
		for nd := t.nodes[slot]; nd.next != nil; nd = nd.next {
			nd.next = &node{
				key:  key,
				next: nil,
			}
		}
	}
	return -1
}

func (t *ch_htable) get(key int) int {
	slot := t.hashFnc(key, t.size)
	if t.nodes[slot] == nil {
		return -1
	}
	for nd := t.nodes[slot]; nd != nil; nd = nd.next {

		if nd.key == key {
			return slot
		}
	}
	return -1
}

func (t *ch_htable) del(key int) int {
	slot := t.hashFnc(key, t.size)
	if t.nodes[slot] == nil {
		return 0
	}
	for nd := t.nodes[slot]; nd.next != nil; nd = nd.next {
		if nd.next != nil {
			if nd.next.key == key {
				nd.next = nd.next.next
				return 0
			}
		}
	}
	return -1
}

//全域散列实现方式之一

const (
	bPrime = 14767
)

type uni_htable struct {
	tl     int
	tables []*node
	//不考虑原始性，这里仅仅保证插入和查找 选取的函数的一致性
	fncFactor map[int][2]int
}

//插入，随机选取函数(((ak+b)%p)%m)
//查找和插入，随机选取的函数要保持一致
func (t *uni_htable) insert(key int) int {
	var a, b int
	a = rand.Intn(bPrime)
	b = rand.Intn(bPrime)
	t.fncFactor[key] = [2]int{a, b}

	slot := ((a*key + b) % bPrime) % t.tl
	if t.tables[slot] != nil {
		for nd := t.tables[slot]; nd.next != nil; nd = nd.next {
			nd.next = &node{
				key:  key,
				next: nil,
			}
		}
	} else {
		t.tables[slot] = &node{
			key:  key,
			next: nil,
		}
	}
	return slot
}

func (t *uni_htable) get(key int) int {
	fct, ok := t.fncFactor[key]
	if !ok {
		return -1
	}
	a, b := fct[0], fct[1]
	slot := ((a*key + b) % bPrime) % t.tl

	for nd := t.tables[slot]; nd != nil; nd = nd.next {
		if nd.key == key {
			return slot
		}
	}

	return -1
}
