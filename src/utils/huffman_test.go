package utils

import (
	"testing"
)

func TestAll(t *testing.T) {
	kv := map[string]int{
		"f": 5,
		"e": 9,
		"c": 12,
		"b": 13,
		"d": 16,
		"a": 45,
	}
	var queue *PQueue = nil
	for k, v := range kv {
		queue = queue.Push(k, v)
	}
	re := HFEncode(kv)
	for k, v := range re {
		t.Logf("%s:%s", k, v)
	}
}
