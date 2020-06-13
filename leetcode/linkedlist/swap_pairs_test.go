package linkedlist

import "testing"

func TestSwapPairs(t *testing.T) {
	head := NewLinkedList([]int{1, 2, 3, 4})
	ll := swapPairs(head)
	print(ll)

}
