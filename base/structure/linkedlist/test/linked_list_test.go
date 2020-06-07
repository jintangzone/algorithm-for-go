package test

import (
	"algorithm-for-go/base/structure/linkedlist"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10}
	ll := linkedlist.NewLinkedList(arr)
	ll.Print()
}
