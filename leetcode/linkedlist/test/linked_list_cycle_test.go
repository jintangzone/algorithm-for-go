package test

import (
	list "algorithm-for-go/base/structure/linkedlist"
	"algorithm-for-go/leetcode/linkedlist"
	"testing"
)

func TestHasCycle(t *testing.T) {
	ll := list.NewLinkedList([]int{1, 2, 3, 4, 5, 6})
	var tail *list.LinkNode
	var pointer *list.LinkNode
	index := 0
	for p := ll.Head; p != nil; p = p.Next {
		tail = p
		if index == 2 {
			pointer = p
		}
		index++
	}
	tail.Next = pointer
	t.Logf("has cycle: %t", linkedlist.HasCycle(ll))
}
