package linkedlist

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	head = reverseList(head)
	print(head)
}
