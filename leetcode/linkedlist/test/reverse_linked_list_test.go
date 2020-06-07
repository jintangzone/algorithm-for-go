package test

import (
	list "algorithm-for-go/base/structure/linkedlist"
	"algorithm-for-go/leetcode/linkedlist"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	ll := list.NewLinkedList([]int{1, 2, 3, 4, 5})
	ll.Print()
	linkedlist.ReverseLinkedList(ll)
	ll.Print()
}
