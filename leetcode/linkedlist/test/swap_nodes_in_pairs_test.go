package test

import (
	list "algorithm-for-go/base/structure/linkedlist"
	"algorithm-for-go/leetcode/linkedlist"
	"testing"
)

func TestSwapNode(t *testing.T) {
	ll := list.NewLinkedList([]int{1, 2, 3, 4, 5, 6})
	ll.Print()
	linkedlist.SwapNode(ll)
	ll.Print()
}
