package linkedlist

import "fmt"

func NewLinkedList(nums []int) *ListNode {
	head := new(ListNode)
	p := head
	for i := 0; i < len(nums); i++ {
		p.Next = &ListNode{Val: nums[i], Next: nil}
		p = p.Next
	}
	return head.Next
}

func print(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, "-")
	}
	fmt.Println("nil")
}
