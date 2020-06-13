package linkedlist

import "testing"

func TestRemoveElements(t *testing.T) {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := NewLinkedList(nums)
	l := removeElements(head, 6)
	print(l)
}
