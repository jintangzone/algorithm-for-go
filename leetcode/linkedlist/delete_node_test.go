package linkedlist

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	head := NewLinkedList([]int{5, 1, 9})
	deleteNode(head)
	print(head)
}
