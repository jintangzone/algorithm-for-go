package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
