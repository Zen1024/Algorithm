package utils

import (
	"testing"
)

func TestAll(t *testing.T) {

	var testarr []string = []string{
		"abcabcdeadabca",
		"abc",
		"abab",
	}
	for _, arr := range testarr {
		re := preFunc(arr)
		t.Logf("%v", re)
	}
	t.Logf("%v", kmpContains("abcabcabc", "abc"))
}
