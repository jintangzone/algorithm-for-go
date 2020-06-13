package linkedlist

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 5)
	print(head)
}
